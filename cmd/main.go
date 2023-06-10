package main

import (
	"time"

	"github.com/charmbracelet/log"

	"github.com/abc-valera/flugo-api/internal/application"
	"github.com/abc-valera/flugo-api/internal/framework/infrastructure"
	"github.com/abc-valera/flugo-api/internal/framework/messaging/redis"
	"github.com/abc-valera/flugo-api/internal/framework/persistence"
	"github.com/abc-valera/flugo-api/internal/framework/presentation/http/api"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Contains all configuration variables
type Config struct {
	PORT                 string        `mapstructure:"HTTP_PORT"`
	DatabaseDriver       string        `mapstructure:"DATABASE_DRIVER"`
	DatabaseUrl          string        `mapstructure:"DATABASE_URL"`
	RedisPort            string        `mapstructure:"REDIS_PORT"`
	RedisUser            string        `mapstructure:"REDIS_USER"`
	RedisPass            string        `mapstructure:"REDIS_PASS"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(configPath string) (Config, error) {
	viper.SetConfigFile(configPath)

	// Override variables from file with the environmet variables
	viper.AutomaticEnv()

	config := Config{}
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

func main() {
	// init config
	c, err := LoadConfig(".env")
	if err != nil {
		log.Fatal(err)
	}

	// init db connection
	db, err := sqlx.Open(c.DatabaseDriver, c.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Initialized database connection")

	// init migrations
	m, err := migrate.New("file://migration", c.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	// !!! Migrate down for testing only
	// if err := m.Down(); err != nil {
	// 	// return err
	// 	log.Info(err)
	// }
	if err := m.Up(); err != nil {
		// return err
		log.Info(err)
	}
	log.Info("Initialized database migrations")

	// Init base layers
	repos := persistence.NewRepositories(db)
	services := infrastructure.NewServices(
		c.AccessTokenDuration, c.RefreshTokenDuration,
		c.EmailSenderAddress, c.EmailSenderPassword)

	// Init redis task mewssaging broker
	msgBroker := redis.NewMessagingBroker(c.RedisPort, c.RedisUser, c.RedisPass, repos.UserRepo, services.Logger)

	// Init Usecases
	usecases := application.NewUsecases(repos, services, msgBroker)

	// Init handlers and API
	log.Info("Running API...")
	log.Fatal(api.RunAPI(c.PORT, services, repos, usecases))
}
