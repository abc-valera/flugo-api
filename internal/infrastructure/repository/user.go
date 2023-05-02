package repository

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
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

type userRepository struct {
	db *sqlx.DB
	ds *goqu.SelectDataset // dataSet is used to specify a table's name
}

func newUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
		ds: goqu.From("users"),
	}
}

func (r *userRepository) CreateUser(c context.Context, user *domain.User) error {
	query := createEntityQuery(r.ds, newDBInsertUser(user))
	return baseExecDB(c, r.db, query, "CreateUser")
}

func (r *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	query := getEntityByFieldQuery(r.ds, "username", username)
	data, err := getDB(c, r.db, &dbReturnUser{}, query, "GetUserByEmail")
	if err != nil {
		return nil, err
	}
	user, ok := data.(*dbReturnUser)
	if !ok {
		return nil, domain.NewInternalError("GetUserByUsername: type assertation failed", nil)
	}
	return newDomainUser(user), nil
}

func (r *userRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	query := getEntityByFieldQuery(r.ds, "email", email)
	data, err := getDB(c, r.db, &dbReturnUser{}, query, "GetUserByEmail")
	if err != nil {
		return nil, err
	}

	user, ok := data.(*dbReturnUser)
	if !ok {
		return nil, domain.NewInternalError("GetUserByEmail: type assertation failed", nil)
	}
	return newDomainUser(user), nil
}

func (r *userRepository) SearchUsersByUsername(c context.Context, username string, params *domain.SelectParams) (domain.Users, error) {
	query := searchEntitiesByFieldQuery(r.ds, "username", username, params)
	data, err := selectDB(c, r.db, &dbReturnUsers{}, query, "SearchUsersByUsername")
	if err != nil {
		return domain.Users{}, err
	}
	users, ok := data.(*dbReturnUsers)
	if !ok {
		return domain.Users{}, domain.NewInternalError("SearchUsersByUsername: type assertation failed", nil)
	}
	return newDomainUsers(*users), nil
}

func (r *userRepository) UpdateUserHashedPassword(c context.Context, username, hashedPassword string) error {
	query := updateEntityFieldQuery(r.ds, "username", username, "hashed_password", hashedPassword)
	return execCheckDB(c, r.db, query, "UpdateUserHashedPassword")
}

func (r *userRepository) UpdateUserFullname(c context.Context, username, fullname string) error {
	query := updateEntityFieldQuery(r.ds, "username", username, "fullname", fullname)
	return execCheckDB(c, r.db, query, "UpdateUserFullname")
}

func (r *userRepository) UpdateUserStatus(c context.Context, username, status string) error {
	query := updateEntityFieldQuery(r.ds, "username", username, "status", status)
	return execCheckDB(c, r.db, query, "UpdateUserStatus")
}

func (r *userRepository) UpdateUserBio(c context.Context, username, bio string) error {
	query := updateEntityFieldQuery(r.ds, "username", username, "bio", bio)
	return execCheckDB(c, r.db, query, "UpdateUserBio")
}

func (r *userRepository) DeleteUser(c context.Context, username string) error {
	query := deleteEntityQuery(r.ds, "username", username)
	return execCheckDB(c, r.db, query, "DeleteUser")
}
