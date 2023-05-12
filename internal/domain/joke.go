package domain

import (
	"time"
)

type Joke struct {
	ID          int
	Username    string
	Title       string
	Text        string
	Explanation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Jokes []*Joke
