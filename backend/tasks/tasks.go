package tasks

import (
	"strings"
	"time"

	"skyzar-backend/database"
	"skyzar-backend/logging"
	"skyzar-backend/requests"
	"skyzar-backend/structs"
	"skyzar-backend/utils"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slices"
)

func StartTasks() {
	s := gocron.NewScheduler(time.UTC)

	s.Cron("* * * * *").Do(refreshBazaarPriceData)
	s.Cron("*/10 * * * *").Do(refreshSkyblockItemData)

	s.StartAsync()
}

func refreshBazaarPriceData() {
	logging.Debug("Refreshing Bazaar Data")
	start := time.Now()

	bazaarData, err := requests.GetHypixelBazaarData()
	if err != nil {
		logging.Error("Error getting Hypixel Bazaar data, error: " + err.Error())
		return
	}

	itemNames, err := database.GetHypixelReadableNames()
	if err != nil {
		logging.Error("Error getting Hypixel Readable Item Names, error: " + err.Error())
		return
	}

	itemNameKeys := utils.ObjectKeys(itemNames.Items)

	models := []mongo.WriteModel{}

	for _, product := range bazaarData.Products {
		var bazaarProduct structs.SkyblockBazaarItem = structs.SkyblockBazaarItem{
			HypixelProductId: product.ProductId,
			LastUpdated: bazaarData.LastUpdated,
			SellData: product.BuySummary,
			BuyData: product.SellSummary,
			SellVolume: product.QuickStatus.SellVolume,
			SellMovingWeek: product.QuickStatus.SellMovingWeek,
			SellOrders: product.QuickStatus.SellOrders,
			BuyVolume: product.QuickStatus.BuyVolume,
			BuyMovingWeek: product.QuickStatus.BuyMovingWeek,
			BuyOrders: product.QuickStatus.BuyOrders,
		}
		if slices.Contains(itemNameKeys, product.ProductId) {
			bazaarProduct.Id = itemNames.Items[product.ProductId]
		} else {
			bazaarProduct.Id = product.ProductId
		}
		if len(product.BuySummary) == 0 {
			product.BuySummary = []structs.HypixelBazaarProductBuySellSummaryItem{
				{
					Amount: 0,
					PricePerUnit: 0,
					Orders: 0,
				},
			}
		}
		if len(product.SellSummary) == 0 {
			product.SellSummary = []structs.HypixelBazaarProductBuySellSummaryItem{
				{
					Amount: 0,
					PricePerUnit: 0,
					Orders: 0,
				},
			}
		}
		bazaarProduct.SellPrice = product.BuySummary[0].PricePerUnit
		bazaarProduct.BuyPrice = product.SellSummary[0].PricePerUnit
		bazaarProduct.Margin = product.BuySummary[0].PricePerUnit - product.SellSummary[0].PricePerUnit
		if product.BuySummary[0].PricePerUnit == 0 {
			bazaarProduct.MarginPercent = 0
		} else {
			bazaarProduct.MarginPercent = (product.BuySummary[0].PricePerUnit - product.SellSummary[0].PricePerUnit) / product.BuySummary[0].PricePerUnit
		}
		
		models = append(models, 
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"hypixel_product_id": bazaarProduct.HypixelProductId}).
				SetUpdate(bson.M{"$set": bazaarProduct}).
				SetUpsert(true),
		)
	}

	err = database.UpdateBazaarItems(models)
	if err != nil {
		logging.Error("Error updating Bazaar Items, error: " + err.Error())
		return
	}

	logging.Debug("Successfully refreshed Bazaar data in " + time.Since(start).String())
}

func refreshSkyblockItemData() {
	logging.Debug("Refreshing Skyblock Item Data")
	start := time.Now()

	itemData, err := requests.GetHypixelSkyblockItemData()
	if err != nil {
		logging.Error("Error getting Hypixel Skyblock Item data, error: " + err.Error())
		return
	}

	var itemNames structs.HypixelReadableItemNames = structs.HypixelReadableItemNames{
		Id: "hypixelReadableItemNames",
		Items: make(map[string]string),
	}

	for _, item := range itemData.Items {
		if item.Id == "FREE_COOKIE" || item.Id == "ENCHANTED_CARROT_ON_A_STICK" {
			continue
		}
		itemNames.Items[item.Id] = strings.ReplaceAll(item.Name, " ", "_")
	}

	err = database.UpdateHypixelReadableNames(itemNames)
	if err != nil {
		logging.Error("Error updating Hypixel Readable Item Names, error: " + err.Error())
	}

	logging.Debug("Successfully refreshed Skyblock Item data in " + time.Since(start).String())
}