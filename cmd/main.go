package main

import (
	"log"
	"time"

	"github.com/abc-valera/flugo-api/internal/application"
	"github.com/abc-valera/flugo-api/internal/infrastructure/pkg"
	"github.com/abc-valera/flugo-api/internal/infrastructure/port/rest/api"
	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Contains all configuration variables
type Config struct {
	PORT                 string        `mapstructure:"PORT"`
	DatabaseDriver       string        `mapstructure:"DATABASE_DRIVER"`
	DatabaseUrl          string        `mapstructure:"DATABASE_URL"`
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

	// init migrations
	m, err := migrate.New("file://migration", c.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
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
	pkgs := pkg.NewPackages(
		c.AccessTokenDuration, c.RefreshTokenDuration,
		c.EmailSenderAddress, c.EmailSenderPassword)
	services := application.NewServices(repos, pkgs)

	// Init app
	log.Fatal(api.RunAPI(c.PORT, pkgs, repos, services))
}
