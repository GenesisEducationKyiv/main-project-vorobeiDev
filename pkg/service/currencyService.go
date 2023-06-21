package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type CurrencyService struct {
	httpClient *http.Client
}

type UAHCoin struct {
	UAH float64 `json:"uah"`
}

type CoinGeckoResponse struct {
	Bitcoin UAHCoin `json:"bitcoin"`
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		httpClient: &http.Client{},
	}
}

func (service *CurrencyService) GetBTCPriceInUAH(ctx context.Context) (float64, error) {
	baseURL := os.Getenv("COINGECKO_BASE_URL")
	url := baseURL + "/api/v3/simple/price?ids=bitcoin&vs_currencies=uah"

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	response, err := service.httpClient.Do(request)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var data CoinGeckoResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	return data.Bitcoin.UAH, nil
}
