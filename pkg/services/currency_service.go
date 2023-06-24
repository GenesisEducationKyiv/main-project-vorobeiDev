package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CurrencyService struct{}

type CurrencyResponse map[string]map[string]float64

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{}
}

func (service *CurrencyService) GetPrice(ctx context.Context, from string, to string) (float64, error) {
	baseURL := os.Getenv("COINGECKO_BASE_URL")
	url := fmt.Sprintf("%s/api/v3/simple/price?ids=%s&vs_currencies=%s", baseURL, from, to)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var data CurrencyResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	currencyData, ok := data[from]
	if !ok {
		return 0, fmt.Errorf("currency not found: %s", from)
	}

	coinPrice, ok := currencyData[to]
	if !ok {
		return 0, fmt.Errorf("currency not found: %s", to)
	}

	return coinPrice, nil
}
