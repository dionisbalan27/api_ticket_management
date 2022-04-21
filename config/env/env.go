package env

import (
	"api_ticket/models"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Config models.ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		logrus.Fatal(err, " config/env: load config")
	}
}

func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal(err, " config/env: load gotdotenv")
	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.PostgresConfig)
	if err != nil {
		return err
	}
	return err
}
