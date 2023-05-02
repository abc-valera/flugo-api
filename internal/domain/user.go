package domain

import "time"

type User struct {
	Username       string
	Email          string
	HashedPassword string
	Fullname       string
	Status         string
	Bio            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewUser(username, email, hashedPassword, fullname, status, bio string) *User {
	return &User{
		Username:       username,
		Email:          email,
		HashedPassword: hashedPassword,
		Fullname:       fullname,
		Status:         status,
		Bio:            bio,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

type Users []*User
