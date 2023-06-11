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

type jokeRepository struct {
	q  orm.Queriers
	ds *goqu.SelectDataset
}

func newJokeRepository(db *sqlx.DB) repository.JokeRepository {
	return &jokeRepository{
		q:  orm.NewQueriers(db),
		ds: goqu.From("jokes"),
	}
}

func (r *jokeRepository) WithTx(txFunc func() error) error {
	return r.q.WithTx(txFunc)
}

func (r *jokeRepository) CreateJoke(c context.Context, joke *domain.Joke) error {
	query := orm.CreateEntityQuery(r.ds, dto.NewInsertJoke(joke))
	return r.q.Exec(c, query, "CreateJoke")
}

func (r *jokeRepository) GetJokeByID(c context.Context, id int) (*domain.Joke, error) {
	joke := &dto.ReturnJoke{}
	query := orm.GetEntityByFieldQuery(r.ds, "id", fmt.Sprint(id))
	err := r.q.Get(c, joke, query, "GetJokeByID")
	if err != nil {
		return nil, err
	}
	return dto.NewDomainJoke(joke), nil
}

func (r *jokeRepository) GetJokes(c context.Context, params *domain.SelectParams) (domain.Jokes, error) {
	jokes := &dto.ReturnJokes{}
	query := orm.GetEntitiesQuery(r.ds, params)
	err := r.q.Select(c, jokes, query, "GetJokes")
	if err != nil {
		return domain.Jokes{}, err
	}
	return dto.NewDomainJokes(*jokes), nil
}

func (r *jokeRepository) GetJokesByAuthor(c context.Context, username string, params *domain.SelectParams) (domain.Jokes, error) {
	jokes := &dto.ReturnJokes{}
	query := orm.GetEntitiesByFieldQuery(r.ds, "username", username, params)
	err := r.q.Select(c, jokes, query, "GetJokesByAuthor")
	if err != nil {
		return domain.Jokes{}, err
	}
	return dto.NewDomainJokes(*jokes), nil
}

func (r *jokeRepository) UpdateJokeExplanation(c context.Context, id int, explanation string) error {
	query := orm.UpdateEntityFieldQuery(r.ds, "id", fmt.Sprint(id), "explanation", explanation)
	return r.q.CheckExec(c, query, "UpdateJokeExplanation")
}

func (r *jokeRepository) DeleteJoke(c context.Context, id int) error {
	query := orm.DeleteEntityQuery(r.ds, "id", fmt.Sprint(id))
	return r.q.CheckExec(c, query, "DeleteUser")
}

// TODO: implement function
func (r *jokeRepository) DeleteJokesByAuthor(c context.Context, username string) error {
	return nil
}
