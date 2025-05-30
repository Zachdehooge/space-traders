package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	/*	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		token := os.Getenv("TOKEN")

		spaceTrader := space_trader.New(token, "FLASH")
		status, err := spaceTrader.ApiStatus()
		check(err)
		fmt.Printf("API Status: %s", status)*/

	resp, err := http.Get("https://api.spacetraders.io/")
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	var data StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("JSON decode failed: %v", err)
	}

	fmt.Println("Status:", data.Status)
}
