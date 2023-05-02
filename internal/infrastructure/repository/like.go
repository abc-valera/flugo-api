package repository

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type LikeRepository interface {
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

type likeRepository struct {
	db *sqlx.DB
	ds *goqu.SelectDataset
}

func newLikeRepository(db *sqlx.DB) LikeRepository {
	return &likeRepository{
		db: db,
		ds: goqu.From("likes"),
	}
}

func (r *likeRepository) CreateLike(c context.Context, like *domain.Like) error {
	query := createEntityQuery(r.ds, newDBInsertLike(like))
	return baseExecDB(c, r.db, query, "CreateLike")
}

func (r *likeRepository) CalcLikesOfJoke(c context.Context, jokeID int) (int, error) {
	var likes_count int
	query, _, _ := r.ds.Select(
		goqu.COUNT("*").As("likes_count")).
		Where(goqu.C("joke_id").Eq(jokeID)).
		ToSQL()
	err := r.db.GetContext(c, &likes_count, query)
	return likes_count, handlePGErr(err, " CalcLikesOfJoke")
}

func (r *likeRepository) GetJokesUserLiked(c context.Context, username string, params *domain.SelectParams) (domain.Jokes, error) {
	jokes := &dbReturnJokes{}
	query, _, _ := r.ds.
		Select("jokes.id", "jokes.username", "jokes.title", "jokes.text", "jokes.explanation", "jokes.created_at", "jokes.updated_at").
		InnerJoin(
			goqu.T("jokes"),
			goqu.On(goqu.Ex{"joke_id": goqu.I("jokes.id")})).
		Where(goqu.I("likes.username").Eq(username)).
		Order(orderedExpression(params)).
		Limit(params.Limit).
		Offset(params.Offset).
		ToSQL()
	err := r.db.SelectContext(c, jokes, query)
	return newDomainJokes(*jokes), handlePGErr(err, "GetJokesUserLiked")
}

func (r *likeRepository) GetUsersWhoLikedJoke(c context.Context, jokeID int, params *domain.SelectParams) (domain.Users, error) {
	users := &dbReturnUsers{}
	query, _, _ := r.ds.
		Select("users.username", "users.email", "users.hashed_password", "users.fullname", "users.status", "users.bio", "users.created_at", "users.updated_at").
		InnerJoin(
			goqu.T("users"),
			goqu.On(goqu.Ex{"likes.username": goqu.I("users.username")})).
		Where(goqu.C("joke_id").Eq(jokeID)).
		Order(orderedExpression(params)).
		Limit(params.Limit).
		Offset(params.Offset).
		ToSQL()
	err := r.db.SelectContext(c, users, query)
	return newDomainUsers(*users), handlePGErr(err, "GetUsersWhoLikedJoke")
}

// TODO: move to ORM?
func (r *likeRepository) DeleteLike(c context.Context, username string, jokeID int) error {
	query, _, _ := r.ds.
		Where(goqu.Ex{
			"username": goqu.Op{"eq": username},
			"joke_id":  goqu.Op{"eq": jokeID},
		}).
		Delete().
		ToSQL()
	res, err := r.db.ExecContext(c, query)
	if err != nil {
		return domain.NewInternalError("DeleteLike", err)
	}
	affected, err := res.RowsAffected()
	if err != nil || affected == 0 {
		return &domain.Error{Code: domain.CodeNotFound}
	}
	return nil
}
