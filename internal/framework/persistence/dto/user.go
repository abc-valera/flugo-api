package dto

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
)

// InsertUser represents user data which should be added into the database
type InsertUser struct {
	Username       string `db:"username"`
	Email          string `db:"email"`
	HashedPassword string `db:"hashed_password"`
	Fullname       string `db:"fullname"`
	Status         string `db:"status"`
	Bio            string `db:"bio"`
}

func NewInsertUser(user *domain.User) *InsertUser {
	return &InsertUser{
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname,
		Status:         user.Status,
		Bio:            user.Bio,
	}
}

// ReturnUser represents user data which is returned from the database
type ReturnUser struct {
	Username       string    `db:"username"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
	Fullname       string    `db:"fullname"`
	Status         string    `db:"status"`
	Bio            string    `db:"bio"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func NewDomainUser(user *ReturnUser) *domain.User {
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

// ReturnUsers represents slice of ReturnUser type returned from the database
type ReturnUsers []*ReturnUser

func NewDomainUsers(dbUsers ReturnUsers) domain.Users {
	users := make(domain.Users, len(dbUsers))
	for i, user := range dbUsers {
		users[i] = NewDomainUser(user)
	}
	return users
}
