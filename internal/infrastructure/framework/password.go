package framework

import (
	"github.com/abc-valera/flugo-api/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type PasswordFramework interface {
	// HashPassword returns hash of the provided password.
	//
	// Returned codes:
	// 	- Internal
	HashPassword(password string) (string, error)

	// CheckPassword checks if provided password matches provided hash.
	// If matches returns nil.
	//
	// Returned codes:
	//  - InvalidArgument
	// 	- Internal
	CheckPassword(password, hashedPassword string) error
}

type bcryptPassword struct{}

func newPasswordFramework() PasswordFramework {
	return &bcryptPassword{}
}

func (f *bcryptPassword) HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.NewInternalError("bcryptPassword.hashPassword", err)
	}
	return string(hashPassword), err
}

func (f *bcryptPassword) CheckPassword(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return domain.NewErrWithMsg(domain.CodeInvalidArgument, "Provided wrong password")
		}
		return domain.NewInternalError("bcryptPassword.CheckPassword", err)
	}
	return nil
}
