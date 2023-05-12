package domain

import (
	"time"
)

type Like struct {
	Username  string
	JokeID    int
	CreatedAt time.Time
}

type Likes []*Like
