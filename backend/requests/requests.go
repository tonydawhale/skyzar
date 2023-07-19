package requests

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"skyzar-backend/logging"
	"skyzar-backend/structs"
)

func GetHypixelBazaarData() (*structs.HypixelBazaarApiRes, error) {
	logging.Log("Request to Hypixel Bazaar API Initiated")
	start := time.Now()
	HypixelResp, err := http.Get("https://api.hypixel.net/skyblock/bazaar")
	if err != nil {
		logging.Error("Error getting Hypixel Bazaar data, error: " + err.Error())
		return nil, err
	}
	var HypixelBazaarApiResponse structs.HypixelBazaarApiRes
	if err := json.NewDecoder(HypixelResp.Body).Decode(&HypixelBazaarApiResponse); err != nil {
		logging.Error("Error decoding Hypixel Bazaar response, error: " + err.Error())
		return nil, err
	}
	if !HypixelBazaarApiResponse.Success {
		logging.Error("Error getting Hypixel Bazaar data, error: " + HypixelBazaarApiResponse.Cause)
		return nil, errors.New("error getting hypixel bazaar data")
	}
	logging.Log("Request to Hypixel Bazaar API Completed in " + time.Since(start).String())
	return &HypixelBazaarApiResponse, nil
}