package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type MarketDataService struct {
	apiKey     string
	apiURL     string
	pythonPath string
	scriptPath string
}

func NewMarketDataService() *MarketDataService {
	return &MarketDataService{
		apiKey:     os.Getenv("MARKET_DATA_API_KEY"),
		apiURL:     os.Getenv("MARKET_DATA_API_URL"),
		pythonPath: "python3",
		scriptPath: filepath.Join("internal", "utils", "yahoo_finance.py"),
	}
}

// Response structure for Alpha Vantage API
type MarketData struct {
	Symbol string  `json:"01. symbol"`
	Price  float64 `json:"05. price,string"`
}

// Response structure for Yahoo Finance API
type YahooFinanceResponse struct {
	CurrentData struct {
		Symbol    string  `json:"symbol"`
		Price     float64 `json:"price"`
		Volume    int64   `json:"volume"`
		Timestamp int64   `json:"timestamp"`
	} `json:"current_data"`
	HistoricalData []struct {
		Date      string  `json:"date"`
		Price     float64 `json:"price"`
		Volume    int64   `json:"volume"`
		Timestamp int64   `json:"timestamp"`
	} `json:"historical_data"`
}

// HTTP handler for web API
func (s *MarketDataService) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Processing pre check requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Obtain stock code
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		http.Error(w, "symbol is required", http.StatusBadRequest)
		return
	}

	// Get data source parameters
	source := r.URL.Query().Get("source")
	var data interface{}
	var err error

	switch source {
	case "yahoo":
		data, err = s.fetchYahooFinanceData(symbol)
	default:
		// Default use of Alpha Vantage
		data, err = s.GetMarketData(r.Context(), symbol)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Alpha Vantage API
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

// Yahoo Finance Data Acquisition Method
func (s *MarketDataService) fetchYahooFinanceData(symbol string) (*YahooFinanceResponse, error) {
	cmd := exec.Command(s.pythonPath, s.scriptPath, symbol)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("python script error: %s", string(exitErr.Stderr))
		}
		return nil, fmt.Errorf("failed to execute python script: %w", err)
	}

	var response YahooFinanceResponse
	if err := json.Unmarshal(output, &response); err != nil {
		return nil, fmt.Errorf("failed to parse python script output: %w", err)
	}

	return &response, nil
}

// GetHistoricalData
func (s *MarketDataService) GetHistoricalData(symbol string) (*YahooFinanceResponse, error) {
	return s.fetchYahooFinanceData(symbol)
}
