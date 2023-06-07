package persistence

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/framework/persistence/orm"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

// dbInsertLike represents like data which should be added into the database
type dbInsertLike struct {
	Username string `db:"username"`
	JokeID   int    `db:"joke_id"`
}

func newDBInsertLike(like *domain.Like) *dbInsertLike {
	return &dbInsertLike{
		Username: like.Username,
		JokeID:   like.JokeID,
	}
}

// dbReturnLike represents like data which is returned from the database
type dbReturnLike struct {
	Username  string    `db:"username"`
	JokeID    int       `db:"joke_id"`
	CreatedAt time.Time `db:"created_at"`
}

func newDomainLike(like *dbReturnLike) *domain.Like {
	return &domain.Like{
		Username:  like.Username,
		JokeID:    like.JokeID,
		CreatedAt: like.CreatedAt,
	}
}

// dbReturnLikes represents slice of dbReturnLike type returned from the database
type dbReturnLikes []*dbReturnLike

func newDomainLikes(dbLikes dbReturnLikes) domain.Likes {
	likes := make(domain.Likes, len(dbLikes))
	for i, like := range dbLikes {
		likes[i] = newDomainLike(like)
	}
	return likes
}

type likeRepository struct {
	q  orm.Queriers
	ds *goqu.SelectDataset
}

func newLikeRepository(db *sqlx.DB) repository.LikeRepository {
	return &likeRepository{
		q:  orm.NewQueriers(db),
		ds: goqu.From("likes"),
	}
}

func (r *likeRepository) CreateLike(c context.Context, like *domain.Like) error {
	query := orm.CreateEntityQuery(r.ds, newDBInsertLike(like))
	return r.q.Exec(c, query, "CreateLike")
}

func (r *likeRepository) CalcLikesOfJoke(c context.Context, jokeID int) (int, error) {
	var likes_count int
	query, _, _ := r.ds.Select(
		goqu.COUNT("*").As("likes_count")).
		Where(goqu.C("joke_id").Eq(jokeID)).
		ToSQL()
	err := r.q.Get(c, &likes_count, query, "CalcLikesOfJoke")
	return likes_count, err
}

func (r *likeRepository) GetJokesUserLiked(c context.Context, username string, params *domain.SelectParams) (domain.Jokes, error) {
	jokes := &dbReturnJokes{}
	query, _, _ := r.ds.
		Select("jokes.id", "jokes.username", "jokes.title", "jokes.text", "jokes.explanation", "jokes.created_at", "jokes.updated_at").
		InnerJoin(
			goqu.T("jokes"),
			goqu.On(goqu.Ex{"joke_id": goqu.I("jokes.id")})).
		Where(goqu.I("likes.username").Eq(username)).
		Order(orm.OrderedExpression(params)).
		Limit(params.Limit).
		Offset(params.Offset).
		ToSQL()
	err := r.q.Select(c, jokes, query, "GetJokesUserLiked")
	return newDomainJokes(*jokes), err
}

func (r *likeRepository) GetUsersWhoLikedJoke(c context.Context, jokeID int, params *domain.SelectParams) (domain.Users, error) {
	users := &dbReturnUsers{}
	query, _, _ := r.ds.
		Select("users.username", "users.email", "users.hashed_password", "users.fullname", "users.status", "users.bio", "users.created_at", "users.updated_at").
		InnerJoin(
			goqu.T("users"),
			goqu.On(goqu.Ex{"likes.username": goqu.I("users.username")})).
		Where(goqu.C("joke_id").Eq(jokeID)).
		Order(orm.OrderedExpression(params)).
		Limit(params.Limit).
		Offset(params.Offset).
		ToSQL()
	err := r.q.Select(c, users, query, "GetUsersWhoLikedJoke")
	return newDomainUsers(*users), err
}

func (r *likeRepository) DeleteLike(c context.Context, username string, jokeID int) error {
	query, _, _ := r.ds.
		Where(goqu.Ex{
			"username": goqu.Op{"eq": username},
			"joke_id":  goqu.Op{"eq": jokeID},
		}).
		Delete().
		ToSQL()
	err := r.q.CheckExec(c, query, "DeleteLike")
	if err != nil {
		return err
	}
	return nil
}
