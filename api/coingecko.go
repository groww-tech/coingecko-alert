package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// CoinGeckoAPI struct
type CoinGeckoAPI struct {
	client  *http.Client
	baseUrl string
}

// NewCoinGecko new instance
func NewCoinGeckoAPI(client *http.Client, baseUrl string) *CoinGeckoAPI {
	return &CoinGeckoAPI{client, baseUrl}
}

// GetSimplePrice get simple price in USD
func (api *CoinGeckoAPI) GetSimplePrice(id string) (float64, error) {
	resp, err := api.client.Get(api.baseUrl + fmt.Sprintf("simple/price?ids=%s&vs_currencies=usd", id))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	m := make(map[string]UsdResponse)
	err = decoder.Decode(&m)
	if err != nil {
		return 0, err
	}

	result, ok := m[id]
	if !ok {
		return 0, errors.New("Cannot find USD price")
	}

	return result.USD, nil
}
