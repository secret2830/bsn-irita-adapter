package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting BSN-IRITA adapter")

	chainID := os.Getenv("BI_CHAIN_ID")
	endpointRPC := os.Getenv("BI_ENDPOINT_RPC")
	endpointGRPC := os.Getenv("BI_ENDPOINT_GRPC")

	keyPath := os.Getenv("BI_KEY_PATH")
	keyName := os.Getenv("BI_KEY_NAME")
	keyPassword := os.Getenv("BI_KEY_PASSWORD")

	endpoint := Endpoint{
		ChainID: chainID,
		RPC:     endpointRPC,
		GRPC:    endpointGRPC,
	}

	keyParams := KeyParams{
		Path:     keyPath,
		Name:     keyName,
		Password: keyPassword,
	}

	adapter, err := NewBSNIritaAdapter(endpoint, keyParams)
	if err != nil {
		fmt.Println("Failed to create the BSN-IRITA adapter:", err)
		return
	}

	RunWebserver(adapter.handle)
}
