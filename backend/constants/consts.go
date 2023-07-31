package constants

import (
	"os"
	"strconv"

	"skyzar-backend/logging"

	"github.com/joho/godotenv"
)

var (
	Port int

	ApiAuthToken string

	HypixelAPIKey string

	MongoURI string
	MongoDatabase string
	MongoProductCollection string
	MongoProductHistoryCollection string
	MongoReadableNamesCollection string 

	CloudflareTurnstileSecret string

	BazaarTopMarginQuota float64
	BazaarTopMarginPercentQuota float64
	BazaarTopDemandQuota float64
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

	ApiAuthToken = getEnv("HYPIXEL_API_KEY")

	HypixelAPIKey = getEnv("HYPIXEL_API_KEY")

	MongoURI = getEnv("MONGO_URI")
	MongoDatabase = getEnv("MONGO_DATABASE")
	MongoProductCollection = getEnv("MONGO_PRODUCT_COLLECTION")
	MongoProductHistoryCollection = getEnv("MONGO_PRODUCT_HISTORY_COLLECTION")
	MongoReadableNamesCollection = getEnv("MONGO_READABLE_NAMES_COLLECTION")

	CloudflareTurnstileSecret = getEnv("CLOUDFLARE_TURNSTILE_SECRET")

	BazaarTopMarginQuota = 10000
	BazaarTopMarginPercentQuota = 0.33
	BazaarTopDemandQuota = 50000
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