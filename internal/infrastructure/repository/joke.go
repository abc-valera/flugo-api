package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository/util"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

// dbInsertJoke represents joke data which should be added into the database
type dbInsertJoke struct {
	Username    string `db:"username"`
	Title       string `db:"title"`
	Text        string `db:"text"`
	Explanation string `db:"explanation"`
}

func newDBInsertJoke(joke *domain.Joke) *dbInsertJoke {
	return &dbInsertJoke{
		Username:    joke.Username,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
	}
}

// dbReturnJoke represents joke data which is returned from the database
type dbReturnJoke struct {
	ID          int       `db:"id"`
	Username    string    `db:"username"`
	Title       string    `db:"title"`
	Text        string    `db:"text"`
	Explanation string    `db:"explanation"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func newDomainJoke(joke *dbReturnJoke) *domain.Joke {
	return &domain.Joke{
		ID:          joke.ID,
		Username:    joke.Username,
		Title:       joke.Title,
		Text:        joke.Text,
		Explanation: joke.Explanation,
		CreatedAt:   joke.CreatedAt,
		UpdatedAt:   joke.UpdatedAt,
	}
}

// dbReturnJokes represents slice of dbReturnJoke type returned from the database
type dbReturnJokes []*dbReturnJoke

func newDomainJokes(dbJokes dbReturnJokes) domain.Jokes {
	jokes := make(domain.Jokes, len(dbJokes))
	for i, joke := range dbJokes {
		jokes[i] = newDomainJoke(joke)
	}
	return jokes
}

type jokeRepository struct {
	db *sqlx.DB
	ds *goqu.SelectDataset
}

func newJokeRepository(db *sqlx.DB) repository.JokeRepository {
	return &jokeRepository{
		db: db,
		ds: goqu.From("jokes"),
	}
}

func (r *jokeRepository) CreateJoke(c context.Context, joke *domain.Joke) error {
	query := util.CreateEntityQuery(r.ds, newDBInsertJoke(joke))
	return util.BaseExecDB(c, r.db, query, "CreateJoke")
}

func (r *jokeRepository) GetJokeByID(c context.Context, id int) (*domain.Joke, error) {
	query := util.GetEntityByFieldQuery(r.ds, "id", fmt.Sprint(id))
	data, err := util.GetDB(c, r.db, &dbReturnJoke{}, query, "GetJokeByID")
	if err != nil {
		return nil, err
	}

	joke, ok := data.(*dbReturnJoke)
	if !ok {
		return nil, domain.NewInternalError("GetJokeByID: type assertation failed", nil)
	}
	return newDomainJoke(joke), nil
}

func (r *jokeRepository) GetJokes(c context.Context, params *domain.SelectParams) (domain.Jokes, error) {
	query := util.GetEntitiesQuery(r.ds, params)
	data, err := util.SelectDB(c, r.db, &dbReturnJokes{}, query, "GetJokes")
	if err != nil {
		return domain.Jokes{}, err
	}

	jokes, ok := data.(*dbReturnJokes)
	if !ok {
		return domain.Jokes{}, domain.NewInternalError("GetJokes: type assertation failed", nil)
	}
	return newDomainJokes(*jokes), nil
}

func (r *jokeRepository) GetJokesByAuthor(c context.Context, username string, params *domain.SelectParams) (domain.Jokes, error) {
	query := util.GetEntitiesByFieldQuery(r.ds, "username", username, params)
	data, err := util.SelectDB(c, r.db, &dbReturnJokes{}, query, "GetJokesByAuthor")
	if err != nil {
		return domain.Jokes{}, err
	}

	jokes, ok := data.(*dbReturnJokes)
	if !ok {
		return domain.Jokes{}, domain.NewInternalError("GetJokesByAuthor: type assertation failed", nil)
	}
	return newDomainJokes(*jokes), nil
}

func (r *jokeRepository) UpdateJokeExplanation(c context.Context, id int, explanation string) error {
	query := util.UpdateEntityFieldQuery(r.ds, "id", fmt.Sprint(id), "explanation", explanation)
	return util.ExecCheckDB(c, r.db, query, "UpdateJokeExplanation")
}

func (r *jokeRepository) DeleteJoke(c context.Context, id int) error {
	query := util.DeleteEntityQuery(r.ds, "id", fmt.Sprint(id))
	return util.ExecCheckDB(c, r.db, query, "DeleteUser")
}

// TODO: implement function
func (r *jokeRepository) DeleteJokesByAuthor(c context.Context, username string) error {
	return nil
}
