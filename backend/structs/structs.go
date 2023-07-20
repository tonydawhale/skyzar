package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CloudflarePost struct {
	Cf string `json:"cf-turnstile-response"`
}

type CloudflareRes struct {
	Success 		bool `json:"success"`
	Errors 			[]string `json:"error-codes"`
	Challenge_ts 	string `json:"challenge_ts"`
	Hostname 		string `json:"hostname"`
	Action 			string `json:"action"`
	Cdata 			string `json:"cdata"`
}

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
	NpcSellPrice 	int `json:"npc_sell_price"`
	Tier 			string `json:"tier"`
	Category 		string `json:"category"`
	Skin 			string `json:"skin,omitempty"`
}

type SkyblockItemRecipe struct {
	Id 				primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ItemID 			string `json:"id" bson:"itemId,omitempty"`
	Costs 			map[string]int `json:"costs" bson:"cost,omitempty"`
	DeepCost 		map[string]int `json:"deepCost,omitempty" bson:"deepCost,omitempty"`
}