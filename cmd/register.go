package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new agent",
	Run: func(cmd *cobra.Command, args []string) {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		ACCOUNT_TOKEN := os.Getenv("TOKEN")

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Error retrieving name:", err)
			return
		}

		faction, err := cmd.Flags().GetString("faction")
		if err != nil {
			fmt.Println("Error retrieving faction:", err)
			return
		}

		url := "https://api.spacetraders.io/v2/register"
		payload := map[string]string{
			"symbol":  name,
			"faction": faction,
		}
		jsonData, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		req.Header.Set("Authorization", "Bearer "+ACCOUNT_TOKEN)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Request error:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		var result struct {
			Data struct {
				Token string `json:"token"`
			} `json:"data"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Error parsing JSON:", err)
			fmt.Println(string(body))
			return
		}

		fmt.Printf("Registered user [%s] with faction [%s] || API-Key = %s \n", name, faction, result.Data.Token)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("name", "n", "", "Name of the agent (required)")
	registerCmd.Flags().StringP("faction", "f", "", "Faction of the agent (required)")
	registerCmd.MarkFlagRequired("name")
	registerCmd.MarkFlagRequired("faction")
	registerCmd.SetUsageTemplate("\nUsage: space-traders register [name] [faction]\n")

}
