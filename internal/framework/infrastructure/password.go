package infrastructure

import (
	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/service"
	"golang.org/x/crypto/bcrypt"
)

type bcryptPassword struct{}

func newPasswordMaker() service.PasswordMaker {
	return &bcryptPassword{}
}

func (s *bcryptPassword) HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", domain.NewInternalError("bcryptPassword.hashPassword", err)
	}
	return string(hashPassword), err
}

func (s *bcryptPassword) CheckPassword(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return domain.NewErrWithMsg(domain.CodeInvalidArgument, "Provided wrong password")
		}
		return domain.NewInternalError("bcryptPassword.CheckPassword", err)
	}
	return nil
}
