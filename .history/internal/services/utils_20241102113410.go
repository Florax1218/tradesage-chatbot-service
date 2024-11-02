package services

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// isMarketDataRequest checks if the message is a request for market data.
func isMarketDataRequest(message string) bool {
    keywords := []string{"price of", "quote for", "market data for", "tell me about", "what is the price of"}
    for _, keyword := range keywords {
        if strings.Contains(strings.ToLower(message), keyword) {
            return true
        }
    }
    return false
}

// extractSymbol extracts the stock symbol from the message.
func extractSymbol(message string) string {
    // Simple regex to find stock symbols (e.g., AAPL, TSLA)
    re := regexp.MustCompile(`\b[A-Z]{1,5}\b`)
    matches := re.FindAllString(message, -1)
    if len(matches) > 0 {
        return matches[len(matches)-1]
    }
    return ""
}

// formatMarketDataResponse formats the market data into a reply message.
func formatMarketDataResponse(data *MarketData) string {
    return fmt.Sprintf("The current price of %s is $%.2f.", data.Symbol, data.Price)
}

// isStockSymbol checks if a word is a valid stock symbol.
func isStockSymbol(word string) bool {
    for _, r := range word {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}
