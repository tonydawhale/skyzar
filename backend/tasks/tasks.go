package tasks

import (
	"time"

	"skyzar-backend/logging"
	"skyzar-backend/requests"
	"skyzar-backend/structs"

	"github.com/go-co-op/gocron"
)

var BazaarData *structs.HypixelBazaarApiRes

func StartTasks() {
	s := gocron.NewScheduler(time.UTC)

	s.Cron("* * * * *").Do(refreshBazaarPriceData)

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