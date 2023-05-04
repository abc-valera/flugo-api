package server

import (
	"log"

	"github.com/abc-valera/flugo-api/internal/delivery/api"
	"github.com/abc-valera/flugo-api/internal/infrastructure/framework"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
	"github.com/abc-valera/flugo-api/internal/service"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunServer() error {
	// init config
	c, err := LoadConfig(".")
	if err != nil {
		return err
	}

	// init db connection
	db, err := sqlx.Open(c.DatabaseDriver, c.DatabaseUrl)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	// init migrations
	m, err := migrate.New("file://migration", c.DatabaseUrl)
	if err != nil {
		return err
	}
	// !!! Migrate down for testing only
	if err := m.Down(); err != nil {
		// return err
		log.Println(err)
	}
	if err := m.Up(); err != nil {
		// return err
		log.Println(err)
	}

	// Init layers
	repos := repository.NewRepositories(db)
	frameworks := framework.NewFrameworks(
		c.AccessTokenDuration, c.RefreshTokenDuration,
		c.EmailSenderAddress, c.EmailSenderPassword)
	services := service.NewServices(repos, frameworks)

	// Init app
	return api.RunAPI(c.PORT, frameworks, repos, services)
}
