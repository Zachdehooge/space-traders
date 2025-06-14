/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new agent",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Error retrieving name:", err)
			return
		}
		fmt.Printf("Registered user %s with API-Key = ...\n", name)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("name", "n", "", "Name of the agent (required)")
	registerCmd.MarkFlagRequired("name")
	registerCmd.SetUsageTemplate("\nUsage: space-traders register [name]\n")

}
