package pkg

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
)

func NewPackages(
	accessDuration, refreshDuration time.Duration,
	senderAddress, senderPassword string,
) *domain.Packages {
	return &domain.Packages{
		PasswordPkg: newPasswordPackage(),
		TokenPkg:    newTokenPackage(accessDuration, refreshDuration),
		EmailPkg:    newEmailPackage(senderAddress, senderPassword),
	}
}
