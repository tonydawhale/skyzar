package constants

import (
	"os"

	"skyzar-database/logging"
	"skyzar-database/structs"

	"github.com/joho/godotenv"
)

var (
	HypixelAPIKey string

	MongoURI string
	MongoDatabase string
	MongoProductCollection string
	MongoProductHistoryCollection string
	MongoReadableNamesCollection string 

	Base24hHistoryData []structs.SkyblockBazaarItemHistoryData

	BazaarTopMarginQuota float64
	BazaarTopMarginPercentQuota float64
	BazaarTopDemandQuota float64
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logging.Fatal("Error loading .env file")
	}

	HypixelAPIKey = getEnv("HYPIXEL_API_KEY")

	MongoURI = getEnv("MONGO_URI")
	MongoDatabase = getEnv("MONGO_DATABASE")
	MongoProductCollection = getEnv("MONGO_PRODUCT_COLLECTION")
	MongoProductHistoryCollection = getEnv("MONGO_PRODUCT_HISTORY_COLLECTION")
	MongoReadableNamesCollection = getEnv("MONGO_READABLE_NAMES_COLLECTION")

	Base24hHistoryData = make([]structs.SkyblockBazaarItemHistoryData, 1440)

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