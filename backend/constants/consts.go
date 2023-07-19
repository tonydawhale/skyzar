package constants

import (
	"os"
	"strconv"

	"skyzar-backend/logging"

	"github.com/joho/godotenv"
)

var (
	Port int
	HypixelAPIKey string
	MongoURI string
	CloudflareTurnstileSecret string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logging.LogFatal("Error loading .env file")
	}

	Port, err = strconv.Atoi(getEnv("PORT"))
	if err != nil {
		logging.LogFatal("Invalid PORT env variable")
	}

	HypixelAPIKey = getEnv("HYPIXEL_API_KEY")
	MongoURI = getEnv("MONGO_URI")
	CloudflareTurnstileSecret = getEnv("CLOUDFLARE_TURNSTILE_SECRET")
}

func LoadConsts() {
	logging.Debug("Successfully loaded all environment variables")
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		logging.LogFatal("Environment variable " + key + " not set")
	}
	return val
}