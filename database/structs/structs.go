package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HypixelSkyblockBazaarApiRes struct {
	Success 		bool `json:"success"`
	LastUpdated 	int `json:"lastUpdated"`
	Products 		map[string]HypixelBazaarProduct `json:"products"`
	Cause 			string `json:"cause,omitempty"`
}

type HypixelBazaarProduct struct {
	ProductId 		string `json:"product_id"`
	SellSummary 	[]HypixelBazaarProductBuySellSummaryItem `json:"sell_summary"`
	BuySummary 		[]HypixelBazaarProductBuySellSummaryItem `json:"buy_summary"`
	QuickStatus 	HypixelBazaarProductQuickStatus `json:"quick_status"`
}

type HypixelBazaarProductBuySellSummaryItem struct {
	Amount 			int `json:"amount"`
	PricePerUnit 	float64 `json:"pricePerUnit"`
	Orders 			int `json:"orders"`
}

type HypixelBazaarProductQuickStatus struct {
	ProductId 		string `json:"product_id"`
	SellPrice 		float64 `json:"sellPrice"`
	SellVolume 		int `json:"sellVolume"`
	SellMovingWeek 	int `json:"sellMovingWeek"`
	SellOrders 		int `json:"sellOrders"`
	BuyPrice 		float64 `json:"buyPrice"`
	BuyVolume 		int `json:"buyVolume"`
	BuyMovingWeek	int `json:"buyMovingWeek"`
	BuyOrders		int `json:"buyOrders"`
}

type HypixelSkyblockItemApiRes struct {
	Success 		bool `json:"success"`
	LastUpdated 	int `json:"lastUpdated"`
	Items 			[]HypixelSkyblockItem `json:"items"`
	Cause 			string `json:"cause,omitempty"`
}

type HypixelSkyblockItem struct {
	Id 				string `json:"id"`
	Material 		string `json:"material"`
	Name 			string `json:"name"`
	NpcSellPrice 	float64 `json:"npc_sell_price"`
	Tier 			string `json:"tier"`
	Category 		string `json:"category"`
	Skin 			string `json:"skin,omitempty"`
}

type HypixelReadableItemNames struct {
	Id 				string `json:"_id,omitempty" bson:"_id,omitempty"`
	Items			map[string]string `json:"products" bson:"products"`
}

type SkyblockItemRecipe struct {
	Id 				primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ItemID 			string `json:"itemId" bson:"itemId"`
	Costs 			map[string]int `json:"costs" bson:"costs"`
	Count			int `json:"count,omitempty" bson:"count,omitempty"`
	DeepCost 		map[string]int `json:"deepCost,omitempty" bson:"deepCost,omitempty"`
}

type SkyblockBazaarItem struct {
	Id 							string `json:"_id" bson:"_id"`
	// HypixelProductId 			string `json:"hypixel_product_id" bson:"hypixel_product_id"`
	DisplayName 				string `json:"display_name" bson:"display_name"`
	LastUpdated 				int `json:"last_updated" bson:"last_updated"`
	SellData					[]HypixelBazaarProductBuySellSummaryItem `json:"sell_data" bson:"sell_data"`
	BuyData						[]HypixelBazaarProductBuySellSummaryItem `json:"buy_data" bson:"buy_data"`
	SellPrice 					float64 `json:"sell_price" bson:"sell_price"`
	SellVolume 					int `json:"sell_volume" bson:"sell_volume"`
	SellMovingWeek 				int `json:"sell_moving_week" bson:"sell_moving_week"`
	SellOrders 					int `json:"sell_orders" bson:"sell_orders"`
	BuyPrice 					float64 `json:"buy_price" bson:"buy_price"`
	BuyVolume 					int `json:"buy_volume" bson:"buy_volume"`
	BuyMovingWeek				int `json:"buy_moving_week" bson:"buy_moving_week"`
	BuyOrders					int `json:"buy_orders" bson:"buy_orders"`
	Margin 						float64 `json:"margin" bson:"margin"`
	MarginPercent 				float64 `json:"margin_percent" bson:"margin_percent"`
}

type SkyblockBazaarItemHistory struct {
	Id 							string `json:"_id" bson:"_id"`
	// HypixelProductId 			string `json:"hypixel_product_id" bson:"hypixel_product_id"`
	LastUpdated 				int `json:"last_updated" bson:"last_updated"`
	History24h					[]SkyblockBazaarItemHistoryData `json:"history_24h" bson:"history_24h"`
	HistoryDaily				[]SkyblockBazaarItemHistoryData `json:"history_daily" bson:"history_daily"`
}

type SkyblockBazaarItemHistoryData struct {
	Time 						int64 `json:"time" bson:"time"`
	SellPrice 					float64 `json:"sell_price" bson:"sell_price"`
	BuyPrice 					float64 `json:"buy_price" bson:"buy_price"`
}

type SkyblockBazaarTopItem struct {
	Id 							string `json:"_id,omitempty" bson:"_id,omitempty"`
	// HypixelProductId 			string `json:"hypixel_product_id" bson:"hypixel_product_id"`
	DisplayName 				string `json:"display_name" bson:"display_name"`
	BuyPrice 					float64 `json:"buy_price" bson:"buy_price"`
	BuyVolume 					int `json:"buy_volume" bson:"buy_volume"`
	Margin 						float64 `json:"margin" bson:"margin"`
	MarginPercent 				float64 `json:"margin_percent" bson:"margin_percent"`
	SellPrice 					float64 `json:"sell_price" bson:"sell_price"`
	SellVolume 					int `json:"sell_volume" bson:"sell_volume"`
}

type SkyblockBazaarCraftableItem struct {
	Id 							string `json:"_id,omitempty" bson:"_id,omitempty"`
	// HypixelProductId 			string `json:"hypixel_product_id" bson:"hypixel_product_id"`
	DisplayName 				string `json:"display_name" bson:"display_name"`
	EstimatedCost 				int `json:"estimated_cost" bson:"estimated_cost"`
	EstimatedProfit 			int `json:"estimated_profit" bson:"estimated_profit"`
	EstimatedProfitPercentage 	float64 `json:"estimated_profit_percentage" bson:"estimated_profit_percentage"`
	SellPrice 					float64 `json:"sell_price" bson:"sell_price"`
	ResourcedUsed 				[]SkyblockBazaarCraftingResource `json:"resources_used" bson:"resources_used"`
}

type SkyblockBazaarCraftingResource struct {
	Id 							primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	// HypixelProductId 			string `json:"hypixel_product_id" bson:"hypixel_product_id"`
	DisplayName 				string `json:"display_name" bson:"display_name"`
	ResourceId 					string `json:"resource_id" bson:"resource_id"`
}

type SkyblockBazaarItemNameFromMongo struct {
	Id 							string `json:"_id,omitempty" bson:"_id,omitempty"`
}