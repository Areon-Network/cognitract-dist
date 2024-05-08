package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/areon-network/cognitract/core"
	"github.com/areon-network/cognitract/types"
)

func UpdateAreonScan(args types.CompilationArgs, abi string) bool {
	payload := []byte(fmt.Sprintf(`{
		"contractAddress": "%s",
		"compilerVersion": "%s",
		"contractAbi": %s,
		"verified": %v,
		"optimization": %v
	}`, args.ContractAddress, args.Version, abi, true, args.Optimization))

	core.InitializeAreonScanEndpoint()
	resp, err := http.Post(core.GetAreonScanEndpoint(), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	var responseBody types.ApiResponse
	err = json.Unmarshal(body, &responseBody)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return false
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", responseBody.Result)
		return false
	}

	defer resp.Body.Close()
	return true
}
