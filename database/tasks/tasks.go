package tasks

import (
	"strings"
	"time"

	"skyzar-database/constants"
	"skyzar-database/database"
	"skyzar-database/logging"
	"skyzar-database/requests"
	"skyzar-database/structs"
	"skyzar-database/utils"

	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slices"
)

var itemNames structs.HypixelReadableItemNames
var itemNameKeys []string

func StartTasks() {
	loc, e := time.LoadLocation("America/Los_Angeles")
	if e != nil {
		logging.Fatal("Error loading timezone, error: " + e.Error())
	}
	s := gocron.NewScheduler(loc)
	refreshSkyblockItemData()

	s.Cron("* * * * *").Do(refreshBazaarPriceData)
	s.Cron("*/10 * * * *").Do(refreshSkyblockItemData)
	s.Cron("0 0 * * *").Do(updateBazaarDailyData)

	s.StartAsync()
}

func refreshBazaarPriceData() {
	logging.Debug("Refreshing Bazaar price data")
	start := time.Now()

	bazaarData, err := requests.GetHypixelBazaarData()
	if err != nil {
		logging.Error("Error getting Hypixel Bazaar data, error: " + err.Error())
		return
	}

	itemNames, err = database.GetHypixelReadableNames()
	if err != nil {
		logging.Error("Error getting Hypixel readable item names, error: " + err.Error())
		return
	}

	itemNameKeys = utils.ObjectKeys(itemNames.Items)

	productModels := []mongo.WriteModel{}
	historyModels := []mongo.WriteModel{}

	for _, product := range bazaarData.Products {
		if product.ProductId == "ENCHANTED_CARROT_ON_A_STICK" {
			continue
		}
		var bazaarProduct structs.SkyblockBazaarItem = schemaBazaarItem(product, bazaarData.LastUpdated)
		
		productModels = append(productModels, 
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": bazaarProduct.Id}).
				SetUpdate(bson.M{
					"$setOnInsert": bson.M{
						"_id": bazaarProduct.Id,
						"display_name": bazaarProduct.DisplayName,
					},
					"$set": bson.M{
						"last_updated": bazaarProduct.LastUpdated,
						"sell_data": bazaarProduct.SellData,
						"buy_data": bazaarProduct.BuyData,
						"sell_price": bazaarProduct.SellPrice,
						"sell_volume": bazaarProduct.SellVolume,
						"sell_moving_week": bazaarProduct.SellMovingWeek,
						"sell_orders": bazaarProduct.SellOrders,
						"buy_price": bazaarProduct.BuyPrice,
						"buy_volume": bazaarProduct.BuyVolume,
						"buy_moving_week": bazaarProduct.BuyMovingWeek,
						"buy_orders": bazaarProduct.BuyOrders,
						"margin": bazaarProduct.Margin,
						"margin_percent": bazaarProduct.MarginPercent,
					},
				}).
				SetUpsert(true),
		)
		historyModels = append(historyModels,
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": bazaarProduct.Id}).
				SetUpdate(bson.M{
					"$setOnInsert": bson.M{
						"_id": bazaarProduct.Id,
						"history_24h": constants.Base24hHistoryData,
						"history_daily": []structs.SkyblockBazaarItemHistoryData{},
					},
					"$set": bson.M{
						"last_updated": bazaarProduct.LastUpdated,
					},
				}).SetUpsert(true),
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": bazaarProduct.Id}).
				SetUpdate(bson.M{
					"$pop": bson.M{
						"history_24h": -1,
					},
				}),
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": bazaarProduct.Id}).
				SetUpdate(bson.M{
					"$push": bson.M{
						"history_24h": bson.M{
							"time": start.Unix(),
							"sell_price": bazaarProduct.SellPrice,
							"buy_price": bazaarProduct.BuyPrice,
						},
					},
				}),
		)
	}

	err = database.BulkWriteUpdate(productModels, constants.MongoProductCollection)
	if err != nil {
		logging.Error("Error updating Bazaar items, error: " + err.Error())
		return
	}
	err = database.BulkWriteUpdate(historyModels, constants.MongoProductHistoryCollection)
	if err != nil {
		logging.Error("Error updating Bazaar items, error: " + err.Error())
		return
	}

	logging.Debug("Successfully refreshed Bazaar data in " + time.Since(start).String())
}

func updateBazaarDailyData() {
	logging.Debug("Refreshing Bazaar daily data")
	start := time.Now()

	models := []mongo.WriteModel{}

	history, err := database.GetEntireBazaarHistory()
	if err != nil {
		logging.Error("Error getting Bazaar history, error: " + err.Error())
		return
	}

	for _, item := range history {
		var todaysHistory structs.SkyblockBazaarItemHistoryData = structs.SkyblockBazaarItemHistoryData{
			Time: start.Unix(),
		}
		var sell = 0.0
		var buy = 0.0
		for _, historyData := range item.History24h {
			sell += historyData.SellPrice
			buy += historyData.BuyPrice
		}
		todaysHistory.SellPrice = sell / float64(len(item.History24h))
		todaysHistory.BuyPrice = buy / float64(len(item.History24h))

		models = append(models,
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"_id": item.Id}).
				SetUpdate(bson.M{
					"$push": bson.M{
						"history_daily": todaysHistory,
					},
				}),
		)
	}

	err = database.BulkWriteUpdate(models, constants.MongoProductHistoryCollection)
	if err != nil {
		logging.Error("Error updating Bazaar items, error: " + err.Error())
		return
	}

	logging.Debug("Successfully updated Bazaar daily data in " + time.Since(start).String())
}

func refreshSkyblockItemData() {
	logging.Debug("Refreshing Skyblock item data")
	start := time.Now()

	itemData, err := requests.GetHypixelSkyblockItemData()
	if err != nil {
		logging.Error("Error getting Hypixel Skyblock item data, error: " + err.Error())
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
		logging.Error("Error updating Hypixel readable item names, error: " + err.Error())
	}

	logging.Debug("Successfully refreshed Skyblock item data in " + time.Since(start).String())
}

func schemaBazaarItem(data structs.HypixelBazaarProduct, lastUpdated int) structs.SkyblockBazaarItem {
	var product structs.SkyblockBazaarItem = structs.SkyblockBazaarItem{
		Id: data.ProductId,
		DisplayName: createDisplayName(data.ProductId),
		LastUpdated: lastUpdated,
		SellData: data.BuySummary,
		BuyData: data.SellSummary,
		SellVolume: data.QuickStatus.SellVolume,
		SellMovingWeek: data.QuickStatus.SellMovingWeek,
		SellOrders: data.QuickStatus.SellOrders,
		BuyVolume: data.QuickStatus.BuyVolume,
		BuyMovingWeek: data.QuickStatus.BuyMovingWeek,
		BuyOrders: data.QuickStatus.BuyOrders,
	}
	if len(data.BuySummary) == 0 {
		data.BuySummary = []structs.HypixelBazaarProductBuySellSummaryItem{
			{
				Amount: 0,
				PricePerUnit: 0,
				Orders: 0,
			},
		}
	}
	if len(data.SellSummary) == 0 {
		data.SellSummary = []structs.HypixelBazaarProductBuySellSummaryItem{
			{
				Amount: 0,
				PricePerUnit: 0,
				Orders: 0,
			},
		}
	}
	product.SellPrice = data.BuySummary[0].PricePerUnit
	product.BuyPrice = data.SellSummary[0].PricePerUnit
	product.Margin = data.BuySummary[0].PricePerUnit - data.SellSummary[0].PricePerUnit
	if data.BuySummary[0].PricePerUnit == 0 {
		product.MarginPercent = 0
	} else {
		product.MarginPercent = (data.BuySummary[0].PricePerUnit - data.SellSummary[0].PricePerUnit) / data.BuySummary[0].PricePerUnit
	}
	return product
}

func createDisplayName(hypixel_product_id string) string {
	var displayName string
	if slices.Contains(itemNameKeys, hypixel_product_id) {
		hypixel_product_id = itemNames.Items[hypixel_product_id]
	}
	if strings.HasPrefix(hypixel_product_id, "ENCHANTMENT_") {
		if strings.HasPrefix(hypixel_product_id, "ENCHANTMENT_ULTIMATE_") {
			displayName = strings.Replace(hypixel_product_id, "ENCHANTMENT_ULTIMATE_", "", 1)
		} else {
			displayName = strings.Replace(hypixel_product_id, "ENCHANTMENT_", "", 1)
		}
	} else if strings.HasPrefix(hypixel_product_id, "ESSENCE_") {
		displayName = strings.Replace(hypixel_product_id, "ESSENCE_", "", 1) + " Essence"
	} else {
		displayName = hypixel_product_id
	}
	return utils.ToProperCase(strings.ReplaceAll(displayName, "_", " "))
}