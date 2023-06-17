package service

import (
	"encoding/json"
	"io"
	"net/http"
)

const BaseUrl = "https://api.coingecko.com"

type CurrencyService struct {
	httpClient *http.Client
}

type CoinGeckoResponse struct {
	Bitcoin struct {
		UAH float64 `json:"uah"`
	} `json:"bitcoin"`
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		httpClient: &http.Client{},
	}
}

func (service *CurrencyService) GetBTCPriceInUAH() (float64, error) {
	resp, err := service.httpClient.Get(BaseUrl + "/api/v3/simple/price?ids=bitcoin&vs_currencies=uah")
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data CoinGeckoResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	return data.Bitcoin.UAH, nil
}
