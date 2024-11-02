package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type MarketDataService struct {
    apiKey string
    apiURL string
}

func NewMarketDataService() *MarketDataService {
    return &MarketDataService{
        apiKey: os.Getenv("MARKET_DATA_API_KEY"),
        apiURL: os.Getenv("MARKET_DATA_API_URL"),
    }
}

type MarketData struct {
    Symbol string  `json:"01. symbol"`
    Price  float64 `json:"05. price,string"`
}

func (s *MarketDataService) GetMarketData(ctx context.Context, symbol string) (*MarketData, error) {
    url := fmt.Sprintf("%s?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", s.apiURL, symbol, s.apiKey)
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result map[string]map[string]string
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    dataMap, ok := result["Global Quote"]
    if !ok {
        return nil, fmt.Errorf("no data found for symbol %s", symbol)
    }

    price, err := strconv.ParseFloat(dataMap["05. price"], 64)
    if err != nil {
        return nil, err
    }

    data := &MarketData{
        Symbol: dataMap["01. symbol"],
        Price:  price,
    }

    return data, nil
}
