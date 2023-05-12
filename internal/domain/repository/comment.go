package repository

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
)

type CommentRepository interface {
	// CreateComment creates new comment entity in the database.
	//
	// Returned codes:
	//  - Internal
	CreateComment(c context.Context, comment *domain.Comment) error

	// GetComment returns comment with such ID from the database.
	// Returns error if there is no comments with such ID.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetComment(c context.Context, id int) (*domain.Comment, error)

	// GetCommentsOfUser returns comments by specified user.
	// Returns error if there is no user with such username.
	// Returns empty comment slice if none comments found.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetCommentsOfUser(c context.Context, username string, params *domain.SelectParams) (domain.Comments, error)

	// GetCommentsOfJoke returns comments of specified joke from the database.
	// Returns error if there is no joke with such id.
	// Returns empty comment slice if none comments found.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetCommentsOfJoke(c context.Context, jokeID int, params *domain.SelectParams) (domain.Comments, error)

	// DeleteComment deletes user's comment to a specified joke.
	// Returns error if user didn't comment specified joke.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	DeleteComment(c context.Context, id int) error
}
