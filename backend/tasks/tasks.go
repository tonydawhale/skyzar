package tasks

import (
	"time"

	"skyzar-backend/logging"
	"skyzar-backend/requests"
	"skyzar-backend/structs"

	"github.com/go-co-op/gocron"
)

var BazaarData *structs.HypixelSkyblockBazaarApiRes
var ItemData *structs.HypixelSkyblockItemApiRes

func StartTasks() {
	s := gocron.NewScheduler(time.UTC)

	s.Cron("* * * * *").Do(refreshBazaarPriceData)
	// s.Cron("*/10 * * * *").Do(refreshSkyblockItemData)

	s.StartAsync()
}

func refreshBazaarPriceData() {
	bazaarData, err := requests.GetHypixelBazaarData()
	if err != nil {
		logging.Error("Error getting Hypixel Bazaar data, error: " + err.Error())
		return
	}
	BazaarData = bazaarData

	logging.Log("Successfully refreshed Bazaar data")
}

func refreshSkyblockItemData() {
	itemData, err := requests.GetHypixelSkyblockItemData()
	if err != nil {
		logging.Error("Error getting Hypixel Skyblock Item data, error: " + err.Error())
		return
	}
	ItemData = itemData

	logging.Log("Successfully refreshed Skyblock Item data")
}