package api

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	base     = "https://api.spacetraders.io/v2"
	agentURL = "https://api.spacetraders.io/v2/my/agent"
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
