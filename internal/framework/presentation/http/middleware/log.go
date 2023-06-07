package middleware

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

func NewLoggerMiddleware(log service.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now()

		errHandler := c.App().ErrorHandler

		err := c.Next()
		if err != nil {
			if err := errHandler(c, err); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		code := c.Response().StatusCode()
		errCode := domain.Code("")
		if err != nil {
			errCode = domain.ErrorCode(err)
		}

		if errCode == "" {
			log.Info("REQUEST",
				"protocol", "http",
				"method", c.Method(),
				"path", c.Path(),
				"http_code", code,
				"duration(ms)", time.Since(startTime).Milliseconds(),
			)
		} else if errCode != domain.CodeInternal {
			log.Warn("REQUEST",
				"protocol", "http",
				"method", c.Method(),
				"path", c.Path(),
				"http_code", code,
				"err", err.Error(),
				"err_code", errCode,
				"err_msg", domain.ErrorMessage(err),
				"duration(ms)", time.Since(startTime).Milliseconds(),
			)
		} else {
			log.Error("REQUEST",
				"protocol", "http",
				"method", c.Method(),
				"path", c.Path(),
				"http_code", code,
				"err", err.Error(),
				"err_code", errCode,
				"err_msg", domain.ErrorMessage(err),
				"duration(ms)", time.Since(startTime).Milliseconds(),
			)
		}

		return err
	}
}
