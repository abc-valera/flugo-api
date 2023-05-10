package server

import (
	"time"

	"github.com/spf13/viper"
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
