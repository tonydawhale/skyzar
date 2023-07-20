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
	MongoDatabase string
	CloudflareTurnstileSecret string
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

	HypixelAPIKey = getEnv("HYPIXEL_API_KEY")
	MongoURI = getEnv("MONGO_URI")
	MongoDatabase = getEnv("MONGO_DATABASE")
	CloudflareTurnstileSecret = getEnv("CLOUDFLARE_TURNSTILE_SECRET")
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