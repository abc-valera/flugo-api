package domain

import "time"

type Joke struct {
	ID          int
	Username    string
	Title       string
	Text        string
	Explanation string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewJoke(username, title, text, explanation string) *Joke {
	return &Joke{
		Username:    username,
		Title:       title,
		Text:        text,
		Explanation: explanation,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type Jokes []*Joke
