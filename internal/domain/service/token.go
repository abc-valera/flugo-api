package service

import "github.com/abc-valera/flugo-api/internal/domain"

type TokenMaker interface {
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
	//  - Unauthenticated (if provided outdated token)
	//  - InvalidArgument (if provided wrong token)
	//  - Internal
	VerifyToken(token string) (*domain.Payload, error)
}
