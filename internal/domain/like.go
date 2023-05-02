package domain

import "time"

type Like struct {
	Username  string
	JokeID    int
	CreatedAt time.Time
}

func NewLike(username string, jokeID int) *Like {
	return &Like{
		Username:  username,
		JokeID:    jokeID,
		CreatedAt: time.Now(),
	}
}

type Likes []*Like
