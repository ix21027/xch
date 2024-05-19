package service

import (
	"encoding/json"
	"net/http"
)

func GetRate() (string, error) {
	requestURL := "https://api.binance.com/api/v3/ticker/price?symbol=BTCUAH"

	response, err := http.Get(requestURL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var result struct {
		Price string `json:"price"`
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Price, nil
}
