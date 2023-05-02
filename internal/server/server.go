package server

import (
	"github.com/abc-valera/flugo-api/internal/delivery/api"
	"github.com/abc-valera/flugo-api/internal/infrastructure/framework"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
	"github.com/abc-valera/flugo-api/internal/service"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func RunServer() error {
	// init config
	c, err := LoadConfig(".")
	if err != nil {
		return err
	}

	// init db connection
	conn, err := sqlx.Open(c.DatabaseDriver, c.DatabaseUrl)
	if err != nil {
		return err
	}
	err = conn.Ping()
	if err != nil {
		return err
	}

	// Init layers
	repos := repository.NewRepositories(conn)
	frameworks := framework.NewFrameworks(c.AccessTokenDuration, c.RefreshTokenDuration)
	services := service.NewServices(repos, frameworks)

	// Init app
	return api.RunAPI(c.PORT, frameworks, repos, services)
}
