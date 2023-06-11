package dto

import (
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
)

// InsertJoke represents joke data which should be added into the database
type InsertJoke struct {
	Username    string `db:"username"`
	Title       string `db:"title"`
	Text        string `db:"text"`
	Explanation string `db:"explanation"`
}

func NewInsertJoke(joke *domain.Joke) *InsertJoke {
	return &InsertJoke{
		Username:    joke.Username,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
	}
}

// ReturnJoke represents joke data which is returned from the database
type ReturnJoke struct {
	ID          int       `db:"id"`
	Username    string    `db:"username"`
	Title       string    `db:"title"`
	Text        string    `db:"text"`
	Explanation string    `db:"explanation"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewDomainJoke(joke *ReturnJoke) *domain.Joke {
	return &domain.Joke{
		ID:          joke.ID,
		Username:    joke.Username,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
		CreatedAt:   joke.CreatedAt,
		UpdatedAt:   joke.UpdatedAt,
	}
}

// ReturnJokes represents slice of ReturnJoke type returned from the database
type ReturnJokes []*ReturnJoke

func NewDomainJokes(dbJokes ReturnJokes) domain.Jokes {
	jokes := make(domain.Jokes, len(dbJokes))
	for i, joke := range dbJokes {
		jokes[i] = NewDomainJoke(joke)
	}
	return jokes
}
