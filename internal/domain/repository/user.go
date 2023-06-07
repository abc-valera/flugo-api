package repository

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
)

type UserRepository interface {
	BaseRepository

	// CreateUser creates new user entity in the database.
	// Returns error if user with same username or email already exists.
	//
	// Returned codes:
	//  - AlreadyExists
	//  - Internal
	CreateUser(c context.Context, user *domain.User) error

	// GetUserByUsername returns user entity with such email from the database.
	// Returns error if user with such username doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetUserByUsername(c context.Context, username string) (*domain.User, error)

	// GetUserByEmail returns user entity with such email from the database.
	// Returns error if user with such email doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetUserByEmail(c context.Context, email string) (*domain.User, error)

	// SearchUsersByUsername returns users whose usernames follow the pattern '*<username>*'.
	// Supports ordering by 'orderBy' with specified 'order' (ASC or DESC).
	// Returns empty users slice if none found.
	//
	// Returned codes:
	//  - Internal
	SearchUsersByUsername(c context.Context, username string, params *domain.SelectParams) (domain.Users, error)

	// UpdateUserHashedPassword updates user's hashedPassword.
	// Returns error if user with such username doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	UpdateUserHashedPassword(c context.Context, username, hashedPassword string) error

	// UpdateUserFullname updates user's fullname.
	// Returns error if user with such username doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	UpdateUserFullname(c context.Context, username, fullname string) error

	// UpdateUserStatus updates user's status.
	// Returns error if user with such username doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	UpdateUserStatus(c context.Context, username, status string) error

	// UpdateUserBio updates user's bio.
	// Returns error if user with such username doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	UpdateUserBio(c context.Context, username, bio string) error

	// UpdateUserBio deletes user with provided username.
	// Returns error if user with such username doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	DeleteUser(c context.Context, username string) error
}
