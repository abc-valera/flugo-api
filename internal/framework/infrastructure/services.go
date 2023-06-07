package infrastructure

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/domain/service"
)

func NewServices(
	accessDuration, refreshDuration time.Duration,
	senderAddress, senderPassword string,
) *service.Services {
	return &service.Services{
		Logger:        newLogger(),
		PasswordMaker: newPasswordMaker(),
		TokenMaker:    newTokenMaker(accessDuration, refreshDuration),
		EmailSender:   newEmailMaker(senderAddress, senderPassword),
	}
}
