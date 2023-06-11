package persistence

import (
	"context"
	"fmt"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/framework/persistence/dto"
	"github.com/abc-valera/flugo-api/internal/framework/persistence/orm"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

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

func (r *commentRepository) WithTx(txFunc func() error) error {
	return r.q.WithTx(txFunc)
}

func (r *commentRepository) CreateComment(c context.Context, comment *domain.Comment) error {
	query := orm.CreateEntityQuery(r.ds, dto.NewInsertComment(comment))
	return r.q.Exec(c, query, "CreateComment")
}

func (r *commentRepository) GetComment(c context.Context, id int) (*domain.Comment, error) {
	comment := &dto.ReturnComment{}
	query := orm.GetEntityByFieldQuery(r.ds, "id", fmt.Sprint(id))
	err := r.q.Get(c, comment, query, "GetComment")
	if err != nil {
		return nil, err
	}
	return dto.NewDomainComment(comment), nil
}

func (r *commentRepository) GetCommentsOfUser(c context.Context, username string, params *domain.SelectParams) (domain.Comments, error) {
	comments := &dto.ReturnComments{}
	query := orm.GetEntitiesByFieldQuery(r.ds, "username", username, params)
	err := r.q.Select(c, comments, query, "GetCommentsOfUser")
	if err != nil {
		return domain.Comments{}, err
	}
	return dto.NewDomainComments(*comments), nil
}

func (r *commentRepository) GetCommentsOfJoke(c context.Context, jokeID int, params *domain.SelectParams) (domain.Comments, error) {
	comments := &dto.ReturnComments{}
	query := orm.GetEntitiesByFieldQuery(r.ds, "joke_id", fmt.Sprint(jokeID), params)
	err := r.q.Select(c, comments, query, "GetCommentsOfJoke")
	if err != nil {
		return domain.Comments{}, err
	}
	return dto.NewDomainComments(*comments), nil
}

func (r *commentRepository) DeleteComment(c context.Context, id int) error {
	query := orm.DeleteEntityQuery(r.ds, "id", fmt.Sprint(id))
	return r.q.CheckExec(c, query, "DeleteComment")
}
