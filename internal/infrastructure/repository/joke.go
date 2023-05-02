package repository

import (
	"context"
	"fmt"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type JokeRepository interface {
	// CreateJoke creates new joke entity in the database.
	// Returns error if specified username already has a joke with such name.
	//
	// Returned codes:
	//  - AlreadyExists
	//  - Internal
	CreateJoke(c context.Context, joke *domain.Joke) error

	// GetJokeByID returns joke with such ID from the database.
	// Returns error if there is no jokes with such ID.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetJokeByID(c context.Context, id int) (*domain.Joke, error)

	// GetJokes returns jokes from the database.
	// Returns empty joke slice if none found.
	//
	// Returned codes:
	//  - Internal
	GetJokes(c context.Context, params *domain.SelectParams) (domain.Jokes, error)

	// GetJokesByAuthor returns jokes by specified username from the database.
	// Returns error if there is no user with such username.
	// Returns empty joke slice if none jokes found.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	GetJokesByAuthor(c context.Context, username string, params *domain.SelectParams) (domain.Jokes, error)

	// UpdateJokeExplanation updates joke's explanation.
	// Returns error if joke with such ID doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	UpdateJokeExplanation(c context.Context, id int, explanation string) error

	// DeleteJoke deletes joke with specified ID.
	// Returns error if joke with such ID doesn't exists.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	DeleteJoke(c context.Context, id int) error

	// DeleteJokesByAuthor deletes all jokes by specified username.
	// Returns error if user with such username doesn't exist.
	//
	// Returned codes:
	//  - NotFound
	//  - Internal
	DeleteJokesByAuthor(c context.Context, username string) error
}

type jokeRepository struct {
	db *sqlx.DB
	ds *goqu.SelectDataset
}

func newJokeRepository(db *sqlx.DB) JokeRepository {
	return &jokeRepository{
		db: db,
		ds: goqu.From("jokes"),
	}
}

func (r *jokeRepository) CreateJoke(c context.Context, joke *domain.Joke) error {
	query := createEntityQuery(r.ds, newDBInsertJoke(joke))
	return baseExecDB(c, r.db, query, "CreateJoke")
}

func (r *jokeRepository) GetJokeByID(c context.Context, id int) (*domain.Joke, error) {
	query := getEntityByFieldQuery(r.ds, "id", fmt.Sprint(id))
	data, err := getDB(c, r.db, &dbReturnJoke{}, query, "GetJokeByID")
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
	query := getEntitiesQuery(r.ds, params)
	data, err := selectDB(c, r.db, &dbReturnJokes{}, query, "GetJokes")
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
	query := getEntitiesByFieldQuery(r.ds, "username", username, params)
	data, err := selectDB(c, r.db, &dbReturnJokes{}, query, "GetJokesByAuthor")
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
	query := updateEntityFieldQuery(r.ds, "id", fmt.Sprint(id), "explanation", explanation)
	return execCheckDB(c, r.db, query, "UpdateJokeExplanation")
}

func (r *jokeRepository) DeleteJoke(c context.Context, id int) error {
	query := deleteEntityQuery(r.ds, "id", fmt.Sprint(id))
	return execCheckDB(c, r.db, query, "DeleteUser")
}

// TODO: implement function
func (r *jokeRepository) DeleteJokesByAuthor(c context.Context, username string) error {
	return nil
}
