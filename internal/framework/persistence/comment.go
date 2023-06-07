package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/framework/persistence/orm"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

// dbInsertComment represents comment data which should be added into the database
type dbInsertComment struct {
	Username string `db:"username"`
	JokeID   int    `db:"joke_id"`
	Text     string `db:"text"`
}

func newDBInsertComment(comment *domain.Comment) *dbInsertComment {
	return &dbInsertComment{
		Username: comment.Username,
		JokeID:   comment.JokeID,
		Text:     comment.Text,
	}
}

// dbReturnComment represents comment data which is returned from the database
type dbReturnComment struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	JokeID    int       `db:"joke_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func newDomainComment(comment *dbReturnComment) *domain.Comment {
	return &domain.Comment{
		ID:        comment.ID,
		Username:  comment.Username,
		JokeID:    comment.JokeID,
		Text:      comment.Text,
		CreatedAt: comment.CreatedAt,
	}
}

// dbReturnComments represents slice of dbReturnComment type returned from the database
type dbReturnComments []*dbReturnComment

func newDomainComments(dbComments dbReturnComments) domain.Comments {
	comments := make(domain.Comments, len(dbComments))
	for i, comment := range dbComments {
		comments[i] = newDomainComment(comment)
	}
	return comments
}

type commentRepository struct {
	q  orm.Queriers
	ds *goqu.SelectDataset
}

func newCommentRepository(db *sqlx.DB) repository.CommentRepository {
	return &commentRepository{
		q:  orm.NewQueriers(db),
		ds: goqu.From("comments"),
	}
}

func (r *commentRepository) CreateComment(c context.Context, comment *domain.Comment) error {
	query := orm.CreateEntityQuery(r.ds, newDBInsertComment(comment))
	return r.q.Exec(c, query, "CreateComment")
}

func (r *commentRepository) GetComment(c context.Context, id int) (*domain.Comment, error) {
	comment := &dbReturnComment{}
	query := orm.GetEntityByFieldQuery(r.ds, "id", fmt.Sprint(id))
	err := r.q.Get(c, comment, query, "GetComment")
	if err != nil {
		return nil, err
	}
	return newDomainComment(comment), nil
}

func (r *commentRepository) GetCommentsOfUser(c context.Context, username string, params *domain.SelectParams) (domain.Comments, error) {
	comments := &dbReturnComments{}
	query := orm.GetEntitiesByFieldQuery(r.ds, "username", username, params)
	err := r.q.Select(c, comments, query, "GetCommentsOfUser")
	if err != nil {
		return domain.Comments{}, err
	}
	return newDomainComments(*comments), nil
}

func (r *commentRepository) GetCommentsOfJoke(c context.Context, jokeID int, params *domain.SelectParams) (domain.Comments, error) {
	comments := &dbReturnComments{}
	query := orm.GetEntitiesByFieldQuery(r.ds, "joke_id", fmt.Sprint(jokeID), params)
	err := r.q.Select(c, comments, query, "GetCommentsOfJoke")
	if err != nil {
		return domain.Comments{}, err
	}
	return newDomainComments(*comments), nil
}

func (r *commentRepository) DeleteComment(c context.Context, id int) error {
	query := orm.DeleteEntityQuery(r.ds, "id", fmt.Sprint(id))
	return r.q.CheckExec(c, query, "DeleteComment")
}
