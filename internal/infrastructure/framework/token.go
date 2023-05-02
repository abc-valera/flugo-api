package framework

import (
	"errors"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/golang-jwt/jwt"
)

type TokenFramework interface {
	// CreateAccessToken creates access token with given username.
	//
	// Returned codes:
	//  - Internal
	CreateAccessToken(username string) (string, *domain.Payload, error)

	// CreateRefreshToken creates refresh token with given username.
	//
	// Returned codes:
	//  - Internal
	CreateRefreshToken(username string) (string, *domain.Payload, error)

	// VerifyToken verifies given token and, if it's correct,
	// returns its payload.
	//
	// Returned codes:
	//  - CodeUnauthenticated (if provided wrong token)
	//  - Internal
	VerifyToken(token string) (*domain.Payload, error)
}

const secretKey = "12345678901234567890123456789012"

type jwtToken struct {
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func newTokenFramework(accessDuration, refreshDuration time.Duration) TokenFramework {
	return &jwtToken{
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}
}

func (f *jwtToken) createToken(username string, isRefresh bool, duration time.Duration) (string, *domain.Payload, error) {
	payload := domain.NewPayload(username, isRefresh, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil, domain.NewInternalError("jwtToken.createToken", err)
	}
	return token, payload, nil
}

func (f *jwtToken) CreateAccessToken(username string) (string, *domain.Payload, error) {
	return f.createToken(username, false, f.accessDuration)
}

func (f *jwtToken) CreateRefreshToken(username string) (string, *domain.Payload, error) {
	return f.createToken(username, true, f.refreshDuration)
}

func (f *jwtToken) VerifyToken(token string) (*domain.Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.NewErrWithMsg(domain.CodeUnauthenticated, "Provided wrong signing method")
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &domain.Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, domain.ErrExpiredToken) {
			return nil, domain.ErrExpiredToken
		}
		return nil, domain.NewErrWithMsg(domain.CodeUnauthenticated, "Provided invalid token")
	}

	payload, ok := jwtToken.Claims.(*domain.Payload)
	if !ok {
		return nil, domain.NewErrWithMsg(domain.CodeUnauthenticated, "Provided invalid token")
	}
	return payload, nil
}
