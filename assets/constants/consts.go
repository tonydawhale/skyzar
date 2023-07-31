package constants

import (
	"os"
	"strconv"

	"skyzar-assets/logging"

	"github.com/joho/godotenv"
)

var (
	Port int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logging.Fatal("Error loading .env file")
	}

	Port, err = strconv.Atoi(getEnv("PORT"))
	if err != nil {
		logging.Fatal("Invalid PORT env variable")
	}
}

func LoadConsts() {
	logging.Info("Successfully loaded all environment variables")
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logging.Fatal("Environment variable " + key + " not set")
	}
	return val
}