package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

type AuthServiceResponse struct {
    UserID string `json:"user_id"`
    Valid  bool   `json:"valid"`
}

// ValidateAuthToken validates the authentication token with the Authentication Service.
func ValidateAuthToken(token string) (string, error) {
    authServiceURL := os.Getenv("AUTH_SERVICE_URL")
    if authServiceURL == "" {
        return "", errors.New("AUTH_SERVICE_URL not set")
    }

    // Remove "Bearer " prefix if present
    token = strings.TrimPrefix(token, "Bearer ")

    req, err := http.NewRequest("GET", authServiceURL, nil)
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+token)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", errors.New("invalid token")
    }

    var authResp AuthServiceResponse
    if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
        return "", err
    }

    if !authResp.Valid {
        return "", errors.New("invalid token")
    }

    return authResp.UserID, nil
}
