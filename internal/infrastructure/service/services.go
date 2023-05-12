package service

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/application/service"
)

func NewServices(
	accessDuration, refreshDuration time.Duration,
	senderAddress, senderPassword string,
) *service.Services {
	return &service.Services{
		PasswordService: newPasswordService(),
		TokenService:    newTokenService(accessDuration, refreshDuration),
		EmailService:    newEmailService(senderAddress, senderPassword),
	}
}
