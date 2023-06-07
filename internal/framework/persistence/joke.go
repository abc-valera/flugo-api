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
	q  orm.Queriers
	ds *goqu.SelectDataset
}

func newJokeRepository(db *sqlx.DB) repository.JokeRepository {
	return &jokeRepository{
		q:  orm.NewQueriers(db),
		ds: goqu.From("jokes"),
	}
}

func (r *jokeRepository) CreateJoke(c context.Context, joke *domain.Joke) error {
	query := orm.CreateEntityQuery(r.ds, newDBInsertJoke(joke))
	return r.q.Exec(c, query, "CreateJoke")
}

func (r *jokeRepository) GetJokeByID(c context.Context, id int) (*domain.Joke, error) {
	joke := &dbReturnJoke{}
	query := orm.GetEntityByFieldQuery(r.ds, "id", fmt.Sprint(id))
	err := r.q.Get(c, joke, query, "GetJokeByID")
	if err != nil {
		return nil, err
	}
	return newDomainJoke(joke), nil
}

func (r *jokeRepository) GetJokes(c context.Context, params *domain.SelectParams) (domain.Jokes, error) {
	jokes := &dbReturnJokes{}
	query := orm.GetEntitiesQuery(r.ds, params)
	err := r.q.Select(c, jokes, query, "GetJokes")
	if err != nil {
		return domain.Jokes{}, err
	}
	return newDomainJokes(*jokes), nil
}

func (r *jokeRepository) GetJokesByAuthor(c context.Context, username string, params *domain.SelectParams) (domain.Jokes, error) {
	jokes := &dbReturnJokes{}
	query := orm.GetEntitiesByFieldQuery(r.ds, "username", username, params)
	err := r.q.Select(c, jokes, query, "GetJokesByAuthor")
	if err != nil {
		return domain.Jokes{}, err
	}
	return newDomainJokes(*jokes), nil
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
