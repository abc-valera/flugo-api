package middleware

import (
	"strings"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

const (
	authHeaderKey  = "authorization"
	authTypeBearer = "bearer"
	AuthPayloadKey = "auth_payload"
)

func NewAuthMiddleware(tokenFr service.TokenMaker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(authHeaderKey)
		if len(authHeader) == 0 {
			return domain.NewErrWithMsg(domain.CodeUnauthenticated,
				"Authorization is not provided (should be under 'authorization' header key)")
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			return domain.NewErrWithMsg(domain.CodeUnauthenticated,
				"Invalid authorization (must be 'bearer <token>')")
		}

		authType := strings.ToLower(fields[0])
		if authType != authTypeBearer {
			return domain.NewErrWithMsg(domain.CodeUnauthenticated,
				"Invalid authorization type (only bearer supported)")
		}

		accessToken := fields[1]
		payload, err := tokenFr.VerifyToken(accessToken)
		if err != nil {
			return err
		}
		if payload.IsRefresh {
			return domain.NewErrWithMsg(domain.CodeUnauthenticated,
				"Provided refresh token")
		}
		c.Locals(AuthPayloadKey, payload)

		return c.Next()
	}
}
