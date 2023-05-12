package repository

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository/util"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

// dbInsertUser represents user data which should be added into the database
type dbInsertUser struct {
	Username       string `db:"username"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	Fullname       string `db:"fullname"`
	Status         string `db:"status"`
	Bio            string `db:"bio"`
}

func newDBInsertUser(user *domain.User) *dbInsertUser {
	return &dbInsertUser{
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname,
		Status:         user.Status,
		Bio:            user.Bio,
	}
}

// dbReturnUser represents user data which is returned from the database
type dbReturnUser struct {
	Username       string    `db:"username"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
	Fullname       string    `db:"fullname"`
	Status         string    `db:"status"`
	Bio            string    `db:"bio"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func newDomainUser(user *dbReturnUser) *domain.User {
	return &domain.User{
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname,
		Status:         user.Status,
		Bio:            user.Bio,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

// dbReturnUsers represents slice of dbReturnUser type returned from the database
type dbReturnUsers []*dbReturnUser

func newDomainUsers(dbUsers dbReturnUsers) domain.Users {
	users := make(domain.Users, len(dbUsers))
	for i, user := range dbUsers {
		users[i] = newDomainUser(user)
	}
	return users
}

type userRepository struct {
	db *sqlx.DB
	ds *goqu.SelectDataset // dataSet is used to specify a table's name
}

func newUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
		ds: goqu.From("users"),
	}
}

func (r *userRepository) CreateUser(c context.Context, user *domain.User) error {
	query := util.CreateEntityQuery(r.ds, newDBInsertUser(user))
	return util.BaseExecDB(c, r.db, query, "CreateUser")
}

func (r *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	query := util.GetEntityByFieldQuery(r.ds, "username", username)
	data, err := util.GetDB(c, r.db, &dbReturnUser{}, query, "GetUserByEmail")
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
	query := util.GetEntityByFieldQuery(r.ds, "email", email)
	data, err := util.GetDB(c, r.db, &dbReturnUser{}, query, "GetUserByEmail")
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
	query := util.SearchEntitiesByFieldQuery(r.ds, "username", username, params)
	data, err := util.SelectDB(c, r.db, &dbReturnUsers{}, query, "SearchUsersByUsername")
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
	query := util.UpdateEntityFieldQuery(r.ds, "username", username, "hashed_password", hashedPassword)
	return util.ExecCheckDB(c, r.db, query, "UpdateUserHashedPassword")
}

func (r *userRepository) UpdateUserFullname(c context.Context, username, fullname string) error {
	query := util.UpdateEntityFieldQuery(r.ds, "username", username, "fullname", fullname)
	return util.ExecCheckDB(c, r.db, query, "UpdateUserFullname")
}

func (r *userRepository) UpdateUserStatus(c context.Context, username, status string) error {
	query := util.UpdateEntityFieldQuery(r.ds, "username", username, "status", status)
	return util.ExecCheckDB(c, r.db, query, "UpdateUserStatus")
}

func (r *userRepository) UpdateUserBio(c context.Context, username, bio string) error {
	query := util.UpdateEntityFieldQuery(r.ds, "username", username, "bio", bio)
	return util.ExecCheckDB(c, r.db, query, "UpdateUserBio")
}

func (r *userRepository) DeleteUser(c context.Context, username string) error {
	query := util.DeleteEntityQuery(r.ds, "username", username)
	return util.ExecCheckDB(c, r.db, query, "DeleteUser")
}
