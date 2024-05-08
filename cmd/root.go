package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "🚧👷 Cognitract helps you in your smart contract tooling!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚧👷 Areon's Cognitract helps you in your smart contract tooling!\n")
		fmt.Println("Run [-h] to see what it can do for you.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
