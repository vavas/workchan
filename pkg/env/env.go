package env

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func GetAppEnv() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		return "development"
	}

	return appEnv
}

func LoadEnvFileIfNeeded(environment string) error {
	var err error

	if environment == "development" {
		err = godotenv.Load(".env")
	}
	if err != nil {
		return errors.Wrap(err, "error loading the dotenv file")
	}

	return nil
}

func IsEnabled(key string) bool {
	val := os.Getenv(key)
	return strings.ToLower(val) == "true"
}
