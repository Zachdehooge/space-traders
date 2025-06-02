package api

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

	token := os.Getenv("TOKEN")

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

	return "Reset Date: " + data.Data.ResetDate + "\n" + "\nReset Frequency: " + data.Data.Freq
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

// Contract structures
type Contract struct {
	ID               string `json:"id"`
	FactionSymbol    string `json:"factionSymbol"`
	Type             string `json:"type"`
	Terms            Terms  `json:"terms"`
	Accepted         bool   `json:"accepted"`
	Fulfilled        bool   `json:"fulfilled"`
	Expiration       string `json:"expiration"`
	DeadlineToAccept string `json:"deadlineToAccept"`
}

type Terms struct {
	Deadline string    `json:"deadline"`
	Payment  Payment   `json:"payment"`
	Deliver  []Deliver `json:"deliver"`
}

type Payment struct {
	OnAccepted  int `json:"onAccepted"`
	OnFulfilled int `json:"onFulfilled"`
}

type Deliver struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int    `json:"unitsRequired"`
	UnitsFulfilled    int    `json:"unitsFulfilled"`
}

type Meta struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ContractsResponse struct {
	Data []Contract `json:"data"`
	Meta Meta       `json:"meta"`
}

func FetchAndPrintContracts() string {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", contracts, nil)

	req.Header.Set("Authorization", "Bearer "+agentToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("HTTP error! status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v", err)
	}

	var contractData ContractsResponse
	err = json.Unmarshal(body, &contractData)
	if err != nil {
		return fmt.Sprintf("Error parsing JSON: %v", err)
	}

	var result strings.Builder

	// Add total contracts from meta
	result.WriteString(fmt.Sprintf("Total Contracts: %d\n", contractData.Meta.Total))
	result.WriteString(strings.Repeat("=", 50) + "\n")

	// Loop through each contract
	for i, contract := range contractData.Data {
		result.WriteString(fmt.Sprintf("Contract %d:\n", i+1))
		result.WriteString(fmt.Sprintf("  Type: %s\n", contract.Type))
		result.WriteString(fmt.Sprintf("  Deadline: %s\n", contract.Terms.Deadline))
		result.WriteString(fmt.Sprintf("  Payment On Accepted: %d\n", contract.Terms.Payment.OnAccepted))
		result.WriteString(fmt.Sprintf("  Payment On Fulfilled: %d\n", contract.Terms.Payment.OnFulfilled))
		result.WriteString(fmt.Sprintf("  Accepted: %t\n", contract.Accepted))
		result.WriteString(fmt.Sprintf("  Fulfilled: %t\n", contract.Fulfilled))

		// Add delivery information for each item
		for j, delivery := range contract.Terms.Deliver {
			result.WriteString(fmt.Sprintf("  Delivery %d:\n", j+1))
			result.WriteString(fmt.Sprintf("    Trade Symbol: %s\n", delivery.TradeSymbol))
			result.WriteString(fmt.Sprintf("    Destination Symbol: %s\n", delivery.DestinationSymbol))
			result.WriteString(fmt.Sprintf("    Units Required: %d\n", delivery.UnitsRequired))
			result.WriteString(fmt.Sprintf("    Units Fulfilled: %d\n", delivery.UnitsFulfilled))
		}

		result.WriteString(strings.Repeat("-", 30) + "\n")
	}

	return result.String()
}

// TODO: Add Accepted Contracts API function
