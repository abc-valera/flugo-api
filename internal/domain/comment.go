package domain

import (
	"time"
)

type Comment struct {
	ID        int
	Username  string
	JokeID    int
	Text      string
	CreatedAt time.Time
}

type Comments []*Comment
