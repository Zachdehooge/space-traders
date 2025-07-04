/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Zachdehooge/space-traders/api"
	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Gets basic information about the agent",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("credits") {
			fmt.Printf("Agent Credits = %d \n", api.PlayerCredits())
		}
	},
}

func init() {
	rootCmd.AddCommand(agentCmd)

	agentCmd.Flags().StringP("credits", "c", "", "Credits")
	//agentCmd.SetUsageTemplate("\nUsage: space-traders agent [flag]\n")
}
