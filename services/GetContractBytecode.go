package services

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

func GetContractBytecode(network string, contractAddress string) string {
	var bytecode string
	rpcEndpoint := fmt.Sprintf("https://%v-rpc.areon.network", network)
	client, err := rpc.DialHTTP(rpcEndpoint)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	err = client.CallContext(context.Background(), &bytecode, "eth_getCode", contractAddress, "latest")
	if err != nil {
		log.Fatal(err)
	}

	return bytecode
}
