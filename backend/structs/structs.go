package structs

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

type HypixelBazaarApiRes struct {
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