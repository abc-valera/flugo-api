package application

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
)

// TODO: replace domain.User for another type?

type SignService interface {
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
}

type signService struct {
	userRepo    domain.UserRepository
	passwordPkg domain.PasswordPackage
	tokenPkg    domain.TokenPackage
	emailPkg    domain.EmailPackage
}

func newSignService(uR domain.UserRepository,
	pPkg domain.PasswordPackage,
	tPkg domain.TokenPackage,
	ePkg domain.EmailPackage,
) SignService {
	return &signService{
		userRepo:    uR,
		passwordPkg: pPkg,
		tokenPkg:    tPkg,
		emailPkg:    ePkg,
	}
}

func (s *signService) SignUp(c context.Context, user *domain.User, password string) error {
	hashedPassword, err := s.passwordPkg.HashPassword(password)
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

func (s *signService) SignIn(c context.Context, email, password string) (*domain.User, string, string, error) {
	user, err := s.userRepo.GetUserByEmail(c, email)
	if err != nil {
		return nil, "", "", err
	}

	if err = s.passwordPkg.CheckPassword(password, user.HashedPassword); err != nil {
		return nil, "", "", err
	}

	access, _, err := s.tokenPkg.CreateAccessToken(user.Username)
	if err != nil {
		return nil, "", "", err
	}
	refresh, _, err := s.tokenPkg.CreateRefreshToken(user.Username)
	if err != nil {
		return nil, "", "", err
	}

	return user, access, refresh, nil
}
