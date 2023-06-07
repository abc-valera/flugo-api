package api

import (
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	Code    domain.Code `json:"code"`
	Message string      `json:"message"`
}

func errorHandler(c *fiber.Ctx, err error) error {
	if e, ok := err.(*domain.Error); ok {
		if e.Code == domain.CodeInvalidArgument ||
			e.Code == domain.CodeNotFound ||
			e.Code == domain.CodeAlreadyExists {
			return c.Status(fiber.StatusBadRequest).JSON(&errorResponse{
				Code:    e.Code,
				Message: e.Msg,
			})
		}

		if e.Code == domain.CodePermissionDenied ||
			e.Code == domain.CodeUnauthenticated {
			return c.Status(fiber.StatusUnauthorized).JSON(&errorResponse{
				Code:    e.Code,
				Message: e.Msg,
			})
		}

		if e.Code == domain.CodeInternal {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	if e, ok := err.(*fiber.Error); ok {
		return c.Status(e.Code).JSON(&errorResponse{
			Message: e.Message,
		})
	}

	return c.SendStatus(fiber.StatusInternalServerError)
}
