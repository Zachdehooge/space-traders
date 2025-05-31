package main

import (
	spacetrader "github.com/Zachdehooge/space_trader_go"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")

	spaceTrader := spacetrader.New(token, "FLASH")
	status, err := spaceTrader.ApiStatus()
	resetDate, err := spaceTrader.LastDate()
	nextReset, err := spaceTrader.NextReset()
	resetFreq, err := spaceTrader.ResetFreq()

	check(err)

	d := color.New(color.FgCyan, color.Bold)

	_, _ = d.Printf("API Status: %s\n", status)
	_, _ = d.Printf("Last Reset Date: %s\n", resetDate)
	_, _ = d.Printf("Next Reset Date: %s\n", nextReset)
	_, _ = d.Printf("Reset Frequency: %s\n", resetFreq)
}
