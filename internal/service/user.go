package service

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/infrastructure/framework"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
)

type UserService interface {
	// SignUp performs user sign-up: it creates new user entity with unique username and email.
	// SignUp also creates hash of the password provided by user
	// (the original password is not stored due to security issues).
	//
	// Returned codes:
	//  - AlreadyExists (if user with such username/email already exists)
	//  - Internal
	SignUp(c context.Context, user *domain.User, password string) error

	// SignIn performs user sign-in: it checks if user with provided email exists,
	// then creates hash of the provided password and compares it to the stored in programm hash.
	// The SignIn returns user, accessToken for accessing most endpoints and refreshToken to renew accessToken.
	//
	// Returned codes:
	//  - NotFound (if no user with such email)
	//  - InvalidArgument (if provided wrong password)
	//  - Internal
	SignIn(c context.Context, email, password string) (*domain.User, string, string, error)

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
	userRepo   repository.UserRepository
	passwordFr framework.PasswordFramework
	tokenFr    framework.TokenFramework
}

func newUsserService(repos *repository.Repositories, frs *framework.Frameworks) UserService {
	return &userService{
		userRepo:   repos.UserRepository,
		passwordFr: frs.PasswordFramework,
		tokenFr:    frs.TokenFramework,
	}
}

func (s *userService) SignUp(c context.Context, user *domain.User, password string) error {
	hashedPassword, err := s.passwordFr.HashPassword(password)
	if err != nil {
		return err
	}

	user.HashedPassword = hashedPassword
	if err = s.userRepo.CreateUser(c, user); err != nil {
		if domain.ErrorCode(err) == domain.CodeAlreadyExists {
			_, e := s.userRepo.GetUserByUsername(c, user.Username)
			if e == nil {
				return domain.NewErrWithMsg(domain.CodeAlreadyExists, "User with such username already exists")
			}
			_, e = s.userRepo.GetUserByEmail(c, user.Email)
			if e == nil {
				return domain.NewErrWithMsg(domain.CodeAlreadyExists, "User with such email already exists")
			}
		}
		return err
	}
	return nil
}

func (s *userService) SignIn(c context.Context, email, password string) (*domain.User, string, string, error) {
	user, err := s.userRepo.GetUserByEmail(c, email)
	if err != nil {
		return nil, "", "", err
	}

	if err = s.passwordFr.CheckPassword(password, user.HashedPassword); err != nil {
		return nil, "", "", err
	}

	access, _, err := s.tokenFr.CreateAccessToken(user.Username)
	if err != nil {
		return nil, "", "", err
	}
	refresh, _, err := s.tokenFr.CreateRefreshToken(user.Username)
	if err != nil {
		return nil, "", "", err
	}

	return user, access, refresh, nil
}

func (s *userService) UpdateUserPassword(c context.Context, username, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetUserByUsername(c, username)
	if err != nil {
		return err
	}
	if err = s.passwordFr.CheckPassword(oldPassword, user.HashedPassword); err != nil {
		return err
	}

	hashedPassword, err := s.passwordFr.HashPassword(newPassword)
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
	if err = s.passwordFr.CheckPassword(password, user.HashedPassword); err != nil {
		return err
	}
	return s.userRepo.DeleteUser(c, username)
}
