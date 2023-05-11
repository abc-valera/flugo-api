package domain

import (
	"context"
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

type JokeRepository interface {
	// CreateJoke creates new joke entity in the database.
	// Returns error if specified username already has a joke with such name.
	//
	// Returned codes:
	//  - AlreadyExists
	//  - Internal
	CreateJoke(c context.Context, joke *Joke) error

	// GetJokeByID returns joke with such ID from the database.
	// Returns error if there is no jokes with such ID.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetJokeByID(c context.Context, id int) (*Joke, error)

	// GetJokes returns jokes from the database.
	// Returns empty joke slice if none found.
	//
	// Returned codes:
	//  - Internal
	GetJokes(c context.Context, params *SelectParams) (Jokes, error)

	// GetJokesByAuthor returns jokes by specified username from the database.
	// Returns error if there is no user with such username.
	// Returns empty joke slice if none jokes found.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetJokesByAuthor(c context.Context, username string, params *SelectParams) (Jokes, error)

	// UpdateJokeExplanation updates joke's explanation.
	// Returns error if joke with such ID doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	UpdateJokeExplanation(c context.Context, id int, explanation string) error

	// DeleteJoke deletes joke with specified ID.
	// Returns error if joke with such ID doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	DeleteJoke(c context.Context, id int) error

	// DeleteJokesByAuthor deletes all jokes by specified username.
	// Returns error if user with such username doesn't exist.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	DeleteJokesByAuthor(c context.Context, username string) error
}
