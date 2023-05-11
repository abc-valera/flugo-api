package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository/util"
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
	db *sqlx.DB
	ds *goqu.SelectDataset
}

func newCommentRepository(db *sqlx.DB) domain.CommentRepository {
	return &commentRepository{
		db: db,
		ds: goqu.From("comments"),
	}
}

func (r *commentRepository) CreateComment(c context.Context, comment *domain.Comment) error {
	query := util.CreateEntityQuery(r.ds, newDBInsertComment(comment))
	return util.BaseExecDB(c, r.db, query, "CreateComment")
}

func (r *commentRepository) GetComment(c context.Context, id int) (*domain.Comment, error) {
	query := util.GetEntityByFieldQuery(r.ds, "id", fmt.Sprint(id))
	data, err := util.GetDB(c, r.db, &dbReturnComment{}, query, "GetComment")
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
	query := util.GetEntitiesByFieldQuery(r.ds, "username", username, params)
	data, err := util.SelectDB(c, r.db, &dbReturnComments{}, query, "GetCommentsOfUser")
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
	query := util.GetEntitiesByFieldQuery(r.ds, "joke_id", fmt.Sprint(jokeID), params)
	data, err := util.SelectDB(c, r.db, &dbReturnComments{}, query, "GetCommentsOfJoke")
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
	query := util.DeleteEntityQuery(r.ds, "id", fmt.Sprint(id))
	return util.ExecCheckDB(c, r.db, query, "DeleteComment")
}
