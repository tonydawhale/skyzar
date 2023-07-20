package requests

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"skyzar-backend/logging"
	"skyzar-backend/structs"
)

func GetHypixelBazaarData() (*structs.HypixelSkyblockBazaarApiRes, error) {
	logging.Log("Request to Hypixel Bazaar API Initiated")
	start := time.Now()
	HypixelResp, err := http.Get("https://api.hypixel.net/skyblock/bazaar")
	if err != nil {
		logging.Error("Error getting Hypixel Bazaar data, error: " + err.Error())
		return nil, err
	}
	var HypixelBazaarApiResponse structs.HypixelSkyblockBazaarApiRes
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

func GetHypixelSkyblockItemData() (*structs.HypixelSkyblockItemApiRes, error) {
	logging.Log("Request to Hypixel Skyblock Item API Initiated")
	start := time.Now()
	HypixelResp, err := http.Get("https://api.hypixel.net/skyblock/items")
	if err != nil {
		logging.Error("Error getting Hypixel Skyblock Item data, error: " + err.Error())
		return nil, err
	}
	var HypixelSkyblockItemApiResponse structs.HypixelSkyblockItemApiRes
	if err := json.NewDecoder(HypixelResp.Body).Decode(&HypixelSkyblockItemApiResponse); err != nil {
		logging.Error("Error decoding Hypixel Skyblock Item response, error: " + err.Error())
		return nil, err
	}
	if !HypixelSkyblockItemApiResponse.Success {
		logging.Error("Error getting Hypixel Skyblock Item data, error: " + HypixelSkyblockItemApiResponse.Cause)
		return nil, errors.New("error getting hypixel skyblock item data")
	}
	logging.Log("Request to Hypixel Skyblock Item API Completed in " + time.Since(start).String())
	return &HypixelSkyblockItemApiResponse, nil
}