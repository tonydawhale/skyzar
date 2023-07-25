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

	productModels := []mongo.WriteModel{}
	historyModels := []mongo.WriteModel{}

	for _, product := range bazaarData.Products {
		if product.ProductId == "ENCHANTED_CARROT_ON_A_STICK" {
			continue
		}
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
		
		productModels = append(productModels, 
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"hypixel_product_id": bazaarProduct.HypixelProductId}).
				SetUpdate(bson.M{
					"$setOnInsert": bson.M{
						"_id": bazaarProduct.Id,
						"hypixel_product_id": bazaarProduct.HypixelProductId,
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
				SetFilter(bson.M{"hypixel_product_id": bazaarProduct.HypixelProductId}).
				SetUpdate(bson.M{
					"$setOnInsert": bson.M{
						"_id": bazaarProduct.Id,
						"hypixel_product_id": bazaarProduct.HypixelProductId,
						"history_24h": constants.Base24hHistoryData,
						"history_daily": []structs.SkyblockBazaarItemHistoryData{},
					},
					"$set": bson.M{
						"last_updated": bazaarProduct.LastUpdated,
					},
				}).SetUpsert(true),
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"hypixel_product_id": bazaarProduct.HypixelProductId}).
				SetUpdate(bson.M{
					"$pop": bson.M{
						"history_24h": -1,
					},
				}),
			mongo.NewUpdateOneModel().
				SetFilter(bson.M{"hypixel_product_id": bazaarProduct.HypixelProductId}).
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
		logging.Error("Error updating Bazaar Items, error: " + err.Error())
		return
	}
	err = database.BulkWriteUpdate(historyModels, constants.MongoProductHistoryCollection)
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