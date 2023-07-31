package requests

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"skyzar-database/logging"
	"skyzar-database/structs"
)

func hypixelRequest(endpoint string) (*http.Response, error) {
	logging.Log("Request to Hypixel API Initiated - Endpoint: " + endpoint)
	start := time.Now()
	HypixelResp, err := http.Get("https://api.hypixel.net/" + endpoint)
	if err != nil {
		logging.Error("Error getting Hypixel data, error: " + err.Error())
		return nil, err
	}
	logging.Log("Request to Hypixel API Completed in " + time.Since(start).String())
	return HypixelResp, nil
}

func GetHypixelBazaarData() (*structs.HypixelSkyblockBazaarApiRes, error) {
	HypixelResp, err := hypixelRequest("skyblock/bazaar")
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
	return &HypixelBazaarApiResponse, nil
}

func GetHypixelSkyblockItemData() (*structs.HypixelSkyblockItemApiRes, error) {
	HypixelResp, err := hypixelRequest("resources/skyblock/items")
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
	return &HypixelSkyblockItemApiResponse, nil
}