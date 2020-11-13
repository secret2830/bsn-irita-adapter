package main

import (
	"os"

	"github.com/smartcontractkit/chainlink/core/logger"
	"go.uber.org/zap/zapcore"
)

func init() {
	logger.SetLogger(logger.CreateProductionLogger("", false, zapcore.DebugLevel, false))
}

func main() {
	logger.Info("Starting BSN-IRITA adapter")

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
		logger.Errorf("Failed to create the BSN-IRITA adapter: %s", err)
		return
	}

	RunWebServer(adapter.handle)
}
