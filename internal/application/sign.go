package application

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/domain/service"
)

// TODO: replace domain.User for another type?

type SignUsecase interface {
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

	// SignRefresh exchages given refresh token for the access token for the same user.
	//
	// Returned codes:
	//  - Unauthenticated (if provided outdated token)
	//  - InvalidArgument (if provided wrong token)
	//  - Internal
	SignRefresh(c context.Context, refreshToken string) (string, error)
}

type signUsecase struct {
	userRepo      repository.UserRepository
	passwordMaker service.PasswordMaker
	tokenMaker    service.TokenMaker
	emailSender   service.EmailSender
	msgBroker     service.MessagingBroker
}

func newSignUsecase(
	userRepo repository.UserRepository,
	passwordMaker service.PasswordMaker,
	tokenMaker service.TokenMaker,
	emailSender service.EmailSender,
	msgBroker service.MessagingBroker,
) SignUsecase {
	return &signUsecase{
		userRepo:      userRepo,
		passwordMaker: passwordMaker,
		tokenMaker:    tokenMaker,
		emailSender:   emailSender,
		msgBroker:     msgBroker,
	}
}

func (s *signUsecase) SignUp(c context.Context, user *domain.User, password string) error {
	hashedPassword, err := s.passwordMaker.HashPassword(password)
	if err != nil {
		return err
	}
	user.HashedPassword = hashedPassword

	txFunc := func() error {
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
		}
		return s.msgBroker.SendVerifyEmailTask(c, user.Email)
	}

	return s.userRepo.WithTx(txFunc)
}

func (s *signUsecase) SignIn(c context.Context, email, password string) (*domain.User, string, string, error) {
	user, err := s.userRepo.GetUserByEmail(c, email)
	if err != nil {
		return nil, "", "", err
	}

	if err = s.passwordMaker.CheckPassword(password, user.HashedPassword); err != nil {
		return nil, "", "", err
	}

	access, _, err := s.tokenMaker.CreateAccessToken(user.Username)
	if err != nil {
		return nil, "", "", err
	}
	refresh, _, err := s.tokenMaker.CreateRefreshToken(user.Username)
	if err != nil {
		return nil, "", "", err
	}

	return user, access, refresh, nil
}

func (s *signUsecase) SignRefresh(c context.Context, refreshToken string) (string, error) {
	payload, err := s.tokenMaker.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, _, err := s.tokenMaker.CreateAccessToken(payload.Username)
	if err != nil {
		return "", err
	}

	return accessToken, err
}
