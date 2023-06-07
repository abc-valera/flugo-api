package persistence

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/framework/persistence/orm"
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
	q  orm.Queriers
	ds *goqu.SelectDataset // dataSet is used to specify a table's name
}

func newUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		q:  orm.NewQueriers(db),
		ds: goqu.From("users"),
	}
}

func (r *userRepository) WithTx(txFunc func() error) error {
	txEcec, err := r.q.StartTx()
	if err != nil {
		return domain.NewInternalError("ususerRepo.WithTx", err)
	}
	r.q.ExecQuerier = txEcec

	dbExec, err := r.q.PerformTx(txFunc)
	if err != nil {
		return err
	}
	r.q.ExecQuerier = dbExec

	return nil
}

func (r *userRepository) CreateUser(c context.Context, user *domain.User) error {
	query := orm.CreateEntityQuery(r.ds, newDBInsertUser(user))
	return r.q.Exec(c, query, "CreateUser")
}

func (r *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	user := &dbReturnUser{}
	query := orm.GetEntityByFieldQuery(r.ds, "username", username)
	err := r.q.Get(c, user, query, "GetUserByUsername")
	if err != nil {
		return nil, err
	}
	return newDomainUser(user), nil
}

func (r *userRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	user := &dbReturnUser{}
	query := orm.GetEntityByFieldQuery(r.ds, "email", email)
	err := r.q.Get(c, user, query, "GetUserByEmail")
	if err != nil {
		return nil, err
	}
	return newDomainUser(user), nil
}

func (r *userRepository) SearchUsersByUsername(c context.Context, username string, params *domain.SelectParams) (domain.Users, error) {
	users := &dbReturnUsers{}
	query := orm.SearchEntitiesByFieldQuery(r.ds, "username", username, params)
	err := r.q.Select(c, users, query, "SearchUsersByUsername")
	if err != nil {
		return domain.Users{}, err
	}
	return newDomainUsers(*users), nil
}

func (r *userRepository) UpdateUserHashedPassword(c context.Context, username, hashedPassword string) error {
	query := orm.UpdateEntityFieldQuery(r.ds, "username", username, "hashed_password", hashedPassword)
	return r.q.CheckExec(c, query, "UpdateUserHashedPassword")
}

func (r *userRepository) UpdateUserFullname(c context.Context, username, fullname string) error {
	query := orm.UpdateEntityFieldQuery(r.ds, "username", username, "fullname", fullname)
	return r.q.CheckExec(c, query, "UpdateUserFullname")
}

func (r *userRepository) UpdateUserStatus(c context.Context, username, status string) error {
	query := orm.UpdateEntityFieldQuery(r.ds, "username", username, "status", status)
	return r.q.CheckExec(c, query, "UpdateUserStatus")
}

func (r *userRepository) UpdateUserBio(c context.Context, username, bio string) error {
	query := orm.UpdateEntityFieldQuery(r.ds, "username", username, "bio", bio)
	return r.q.CheckExec(c, query, "UpdateUserBio")
}

func (r *userRepository) DeleteUser(c context.Context, username string) error {
	query := orm.DeleteEntityQuery(r.ds, "username", username)
	return r.q.CheckExec(c, query, "DeleteUser")
}
