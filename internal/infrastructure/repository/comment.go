package repository

import (
	"context"
	"fmt"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
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

type commentRepository struct {
	db *sqlx.DB
	ds *goqu.SelectDataset
}

func newCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentRepository{
		db: db,
		ds: goqu.From("comments"),
	}
}

func (r *commentRepository) CreateComment(c context.Context, comment *domain.Comment) error {
	query := createEntityQuery(r.ds, newDBInsertComment(comment))
	return baseExecDB(c, r.db, query, "CreateComment")
}

func (r *commentRepository) GetComment(c context.Context, id int) (*domain.Comment, error) {
	query := getEntityByFieldQuery(r.ds, "id", fmt.Sprint(id))
	data, err := getDB(c, r.db, &dbReturnComment{}, query, "GetComment")
	if err != nil {
		return nil, err
	}

	comment, ok := data.(*dbReturnComment)
	if !ok {
		return nil, domain.NewInternalError("commentRepository.GetComment: type assertation failed", nil)
	}
	return newDomainComment(comment), nil
}

func (r *commentRepository) GetCommentsOfUser(c context.Context, username string, params *domain.SelectParams) (domain.Comments, error) {
	query := getEntitiesByFieldQuery(r.ds, "username", username, params)
	data, err := selectDB(c, r.db, &dbReturnComments{}, query, "GetCommentsOfUser")
	if err != nil {
		return domain.Comments{}, err
	}

	comments, ok := data.(*dbReturnComments)
	if !ok {
		return domain.Comments{}, domain.NewInternalError("GetCommentsOfUser: type assertation failed", nil)
	}
	return newDomainComments(*comments), nil
}

func (r *commentRepository) GetCommentsOfJoke(c context.Context, jokeID int, params *domain.SelectParams) (domain.Comments, error) {
	query := getEntitiesByFieldQuery(r.ds, "joke_id", fmt.Sprint(jokeID), params)
	data, err := selectDB(c, r.db, &dbReturnComments{}, query, "GetCommentsOfJoke")
	if err != nil {
		return domain.Comments{}, err
	}

	comments, ok := data.(*dbReturnComments)
	if !ok {
		return domain.Comments{}, domain.NewInternalError("GetCommentsOfJoke: type assertation failed", nil)
	}
	return newDomainComments(*comments), nil
}

func (r *commentRepository) DeleteComment(c context.Context, id int) error {
	query := deleteEntityQuery(r.ds, "id", fmt.Sprint(id))
	return execCheckDB(c, r.db, query, "DeleteComment")
}
