package domain

import (
	"time"
)

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

type Users []*User
