package pkg

import (
	"errors"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/golang-jwt/jwt"
)

const secretKey = "12345678901234567890123456789012"

type jwtToken struct {
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func newTokenPackage(accessDuration, refreshDuration time.Duration) domain.TokenPackage {
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
