package cmd

import (
	"fmt"

	"github.com/areon-network/cognitract/compilation"
	"github.com/areon-network/cognitract/constants"
	"github.com/areon-network/cognitract/core"
	"github.com/areon-network/cognitract/prompts"
	"github.com/areon-network/cognitract/types"

	"github.com/spf13/cobra"
)

var isTestnet bool
var optimization bool = true

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Recompile & verify your smart contracts on AreonScan",
	Run: func(cmd *cobra.Command, args []string) {
		if isTestnet {
			fmt.Println("⚠️  Attention: Switched to Areon Testnet\n")
			core.SetActiveNetwork("testnet")
		}

		solcVersion := prompts.SelectPrompt("Choose the Solidity version of your contract", constants.SolidityVersions)
		contractAddress := prompts.StringPrompt("Contract Address to be verified")
		projectRootPath := prompts.StringPrompt("Path to the root of your project")
		mainContractPath := prompts.StringPrompt("Path to main contract (in project's root directory)")
		mainContractName := prompts.StringPrompt("Name of the main contract")
		isOptimized := prompts.SelectPrompt("Optimized", []bool{true, false})

		if isOptimized == "false" {
			optimization = false
		}

		fmt.Println("\n")
		compilation.DownloadCompiler(solcVersion)

		compileContractsArgs := types.CompilationArgs{
			Version:            solcVersion,
			PathToMainContract: mainContractPath,
			MainContractName:   mainContractName,
			BasePath:           projectRootPath,
			SolcBinaryPath:     constants.BinaryDirectory + "/soljson-" + solcVersion + ".bin",
			ContractAddress:    contractAddress,
			Optimization:       optimization,
			Network:            core.GetActiveNetwork(),
		}

		compilation.CompileContracts(compileContractsArgs)
	},
}

func init() {
	verifyCmd.Flags().BoolVarP(&isTestnet, "testnet", "t", false, "Runs the verification process on Testnet")
	rootCmd.AddCommand(verifyCmd)
}
