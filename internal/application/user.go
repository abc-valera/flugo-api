package application

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
)

type UserService interface {
	// UpdateUserPassword checks if provided oldPassword's hash matches user's current PasswordHash.
	// If mathces it creates hash of the provided newPassword and updates user entity with new hash.
	//
	// Returned codes:
	//  - NotFound (if no user with such email)
	//  - InvalidArgument (if provided wrong oldPassword)
	//  - Internal
	UpdateUserPassword(c context.Context, username, oldPassword, newPassword string) error

	// DeleteUser checks if provided password's hash matches user's with provided username PasswordHash.
	// If mathces it deletes user entry for ever.
	//
	// Returned codes:
	//  - NotFound (if no user with such email)
	//  - InvalidArgument (if provided wrong oldPassword)
	//  - Internal
	DeleteUser(c context.Context, username, password string) error
}

type userService struct {
	userRepo    domain.UserRepository
	passwordPkg domain.PasswordPackage
}

func newUserService(userRepo domain.UserRepository, passwordPkg domain.PasswordPackage) UserService {
	return &userService{
		userRepo:    userRepo,
		passwordPkg: passwordPkg,
	}
}

func (s *userService) UpdateUserPassword(c context.Context, username, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}
	if err = s.passwordPkg.CheckPassword(oldPassword, user.HashedPassword); err != nil {
		return err
	}

	hashedPassword, err := s.passwordPkg.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.userRepo.UpdateUserHashedPassword(c, username, hashedPassword)
}

func (s *userService) DeleteUser(c context.Context, username, password string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}
	if err = s.passwordPkg.CheckPassword(password, user.HashedPassword); err != nil {
		return err
	}
	return s.userRepo.DeleteUser(c, username)
}
