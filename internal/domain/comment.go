package domain

import "time"

type Comment struct {
	ID        int
	Username  string
	JokeID    int
	Text      string
	CreatedAt time.Time
}

func NewComment(username, text string, jokeID int) *Comment {
	return &Comment{
		Username:  username,
		JokeID:    jokeID,
		Text:      text,
		CreatedAt: time.Now(),
	}
}

type Comments []*Comment
