package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	base      = "https://api.spacetraders.io/v2"
	agentURL  = "https://api.spacetraders.io/v2/my/agent"
	contracts = "https://api.spacetraders.io/v2/my/contracts"
	// TODO: Add URLs for exploration, mining, contracts, etc.
)

func agentToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("PLAYER_TOKEN")

	return token
}

func ServerStatus() string {
	type BaseURL struct {
		Status string `json:"status"`
	}

	req, err := http.NewRequest("GET", base, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Status request failed: %v", err)
	}
	defer resp.Body.Close()

	var data BaseURL
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Status decode failed: %v", err)
	}

	return data.Status
}

func ServerResetDate() string {
	type BaseURL struct {
		ResetDate string `json:"resetDate"`
	}

	req, err := http.NewRequest("GET", base, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Status request failed: %v", err)
	}
	defer resp.Body.Close()

	var data BaseURL
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Status decode failed: %v", err)
	}

	return data.ResetDate
}

func ServerNextResetDate() string {
	type BaseURL struct {
		Data struct {
			ResetDate string `json:"next"`
			Freq      string `json:"frequency"`
		} `json:"serverResets"`
	}

	req, err := http.NewRequest("GET", base, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Status request failed: %v", err)
	}
	defer resp.Body.Close()

	var data BaseURL
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Status decode failed: %v", err)
	}

	return data.Data.ResetDate + "\n" + "Reset Frequency: " + data.Data.Freq
}

func PlayerAgent() string {

	type BaseURL struct {
		Data struct {
			Player string `json:"symbol"`
		} `json:"data"`
	}

	req, err := http.NewRequest("GET", agentURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+agentToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Status request failed: %v", err)
	}
	defer resp.Body.Close()

	var data BaseURL
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Status decode failed: %v", err)
	}

	return data.Data.Player
}

func PlayerCredits() int {

	type BaseURL struct {
		Data struct {
			Credits int `json:"credits"`
		} `json:"data"`
	}

	req, err := http.NewRequest("GET", agentURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+agentToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Status request failed: %v", err)
	}
	defer resp.Body.Close()

	var data BaseURL
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Status decode failed: %v", err)
	}

	return data.Data.Credits
}

func PlayerShips() int {

	type BaseURL struct {
		Data struct {
			ShipCount int `json:"shipCount"`
		} `json:"data"`
	}

	req, err := http.NewRequest("GET", agentURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+agentToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Status request failed: %v", err)
	}
	defer resp.Body.Close()

	var data BaseURL
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Status decode failed: %v", err)
	}

	return data.Data.ShipCount
}

func PlayerHeadquarters() string {

	type BaseURL struct {
		Data struct {
			Headquarters string `json:"headquarters"`
		} `json:"data"`
	}

	req, err := http.NewRequest("GET", agentURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+agentToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Status request failed: %v", err)
	}
	defer resp.Body.Close()

	var data BaseURL
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Status decode failed: %v", err)
	}

	return data.Data.Headquarters
}

func PlayerFaction() string {

	type BaseURL struct {
		Data struct {
			Faction string `json:"startingFaction"`
		} `json:"data"`
	}

	req, err := http.NewRequest("GET", agentURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+agentToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Status request failed: %v", err)
	}
	defer resp.Body.Close()

	var data BaseURL
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Status decode failed: %v", err)
	}

	return data.Data.Faction
}

// TODO: Add Fleet Management API function

// TODO: Add Galaxy Navigation API function

// TODO: Add Market Insights API function

// TODO: Add Available Contracts API function

// TODO: Add Accepted Contracts API function
