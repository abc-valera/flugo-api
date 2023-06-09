package repository

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
)

type LikeRepository interface {
	BaseRepository

	// CreateLike creates new like entity in the database.
	// Returns error if specified username already likes specified joke.
	//
	// Returned codes:
	//  - AlreadyExists
	//  - Internal
	CreateLike(c context.Context, like *domain.Like) error

	// CalcLikesOfJoke returns number of users who liked specified joke.
	// Returns error if joke doesn't exist.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	CalcLikesOfJoke(c context.Context, jokeID int) (int, error)

	// GetJokesUserLiked returns liked jokes of a user from the database.
	// Returns error if there is no user with such username.
	// Returns empty joke slice if none liked jokes found.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetJokesUserLiked(c context.Context, username string, params *domain.SelectParams) (domain.Jokes, error)

	// GetUsersWhoLikesJoke returns users who liked specified joke from the database.
	// Returns error if there is no joke with such id.
	// Returns empty user slice if none users liked.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetUsersWhoLikedJoke(c context.Context, jokeID int, params *domain.SelectParams) (domain.Users, error)

	// DeleteLike deletes user's like to a specified joke.
	// Returns error if user doesn't like specified joke.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	DeleteLike(c context.Context, username string, jokeID int) error
}
