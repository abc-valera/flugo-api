package infrastructure

import (
	"errors"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/service"
	"github.com/golang-jwt/jwt"
)

const secretKey = "12345678901234567890123456789012"

type jwtToken struct {
	accessDuration  time.Duration
	refreshDuration time.Duration
}

func newTokenMaker(accessDuration, refreshDuration time.Duration) service.TokenMaker {
	return &jwtToken{
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}
}

func (s *jwtToken) createToken(username string, isRefresh bool, duration time.Duration) (string, *domain.Payload, error) {
	payload := domain.NewPayload(username, isRefresh, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil, domain.NewInternalError("jwtToken.createToken", err)
	}
	return token, payload, nil
}

func (s *jwtToken) CreateAccessToken(username string) (string, *domain.Payload, error) {
	return s.createToken(username, false, s.accessDuration)
}

func (s *jwtToken) CreateRefreshToken(username string) (string, *domain.Payload, error) {
	return s.createToken(username, true, s.refreshDuration)
}

func (s *jwtToken) VerifyToken(token string) (*domain.Payload, error) {
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
