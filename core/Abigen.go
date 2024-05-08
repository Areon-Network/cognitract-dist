package core

import (
	"encoding/json"
	"fmt"

	"github.com/areon-network/cognitract/gsolc"
	"github.com/areon-network/cognitract/types"
)

func Abigen(output *gsolc.Output, args types.CompilationArgs) string {
	abi := output.Contracts[args.PathToMainContract][args.MainContractName].ABI
	if len(abi) <= 0 {
		fmt.Println("This contract cannot derive an ABI")
		return "[]"
	}

	abiRaw, err := json.Marshal(abi)
	if err != nil {
		fmt.Println(err)
	}

	return string(abiRaw)
}
