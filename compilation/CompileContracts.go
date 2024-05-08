package compilation

import (
	"fmt"
	"os"

	"github.com/areon-network/cognitract/core"
	"github.com/areon-network/cognitract/gsolc"
	"github.com/areon-network/cognitract/services"
	"github.com/areon-network/cognitract/types"
)

func printResult(result bool, out *gsolc.Output, args types.CompilationArgs) {
	if result {
		core.CreateArtifacts(out)
		mainContractAbi := core.Abigen(out, args)
		if mainContractAbi != "[]" {
			core.CreateAbiFolder([]byte(mainContractAbi))
		}

		serviceRes := services.UpdateAreonScan(args, mainContractAbi)
		if serviceRes {
			fmt.Println("...\n")
			fmt.Println("- Verification complete ✅")
			return
		}
	}

	fmt.Println("...\n")
	fmt.Println("- Contract verification failed ❌")
}

func concludeVerification(args types.CompilationArgs, out *gsolc.Output) {
	bytecodeFromNetwork := services.GetContractBytecode(args.Network, args.ContractAddress)
	bytecodeFromCompiler := "0x" + out.Contracts[args.PathToMainContract][args.MainContractName].EVM.DeployedBytecode.Object

	if bytecodeFromNetwork != "0x" && bytecodeFromCompiler != "0x" {
		res := bytecodeFromNetwork == bytecodeFromCompiler
		printResult(res, out, args)
		return
	}

	fmt.Println("This contract is not available on chain. Please check your contract address")
}

func CompileContracts(args types.CompilationArgs) {
	compiler, err := gsolc.NewFromFile(args.SolcBinaryPath, args.Version)
	if err != nil {
		panic(err)
	}

	input := &gsolc.Input{
		Language: "Solidity",
		Sources:  make(map[string]gsolc.SourceIn),
		Settings: gsolc.Settings{
			Optimizer: gsolc.Optimizer{
				Enabled: args.Optimization,
			},
			OutputSelection: map[string]map[string][]string{
				"*": {"*": []string{"*"}},
			},
		},
	}

	os.Chdir(args.BasePath) // Go to project's path

	// This control here checks if it was deployed using Remix or not
	if _, err := os.Stat("node_modules"); os.IsNotExist(err) {
		fmt.Println("🚀 Deployment stage detected: Remix IDE")
		for f, d := range core.GetAllSources("./") {
			input.Sources[f] = gsolc.SourceIn{Content: string(d)}
		}

		out, err := compiler.Compile(input)
		if err != nil {
			panic(err)
		}

		concludeVerification(args, out)
		return
	}

	fmt.Println("🚀 Deployment stage detected: local environment")

	var libPaths []string
	for f, d := range GetSourceCode("./") {
		fmt.Println("🔎 Analyzing imports in your code...\n")
		input.Sources[f] = gsolc.SourceIn{Content: string(d)}
		libPaths = append(libPaths, core.ResolveSolidityImports(string(d))...)
	}

	fmt.Printf("Found %v import(s) in total\n", len(libPaths))
	for f, d := range GetNodeModuleCompiled(libPaths) {
		input.Sources[f] = gsolc.SourceIn{Content: string(d)}
	}

	out, err := compiler.Compile(input)
	if err != nil {
		panic(err)
	}

	concludeVerification(args, out)
}
