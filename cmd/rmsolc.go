package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/areon-network/cognitract/constants"
	"github.com/spf13/cobra"
)

var rmsolcCmd = &cobra.Command{
	Use:   "rmsolc",
	Short: "Removes all of Solidity compiler binaries (does not affect globally installed compilers)",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := filepath.Glob(filepath.Join(constants.BinaryDirectory, "*"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(files) == 0 {
			fmt.Println("No available compilers to delete")
			return
		}

		for _, file := range files {
			// The README.md file stays in the directory for github
			if filepath.Ext(file) != ".md" {
				err := os.Remove(file)
				if err != nil {
					fmt.Println("Error removing file:", err)
				}
			}
		}

		msg := fmt.Sprintf("Removed %v compilers successfully!", len(files))
		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(rmsolcCmd)
}
