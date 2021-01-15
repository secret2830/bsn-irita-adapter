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

	chainID := os.Getenv("BA_CHAIN_ID")
	endpointRPC := os.Getenv("BA_ENDPOINT_RPC")
	endpointGRPC := os.Getenv("BA_ENDPOINT_GRPC")
	mnemonic := os.Getenv("BA_KEY_MNEMONIC")
	listenAddr := os.Getenv("BA_LISTEN_ADDR")

	endpoint := Endpoint{
		ChainID: chainID,
		RPC:     endpointRPC,
		GRPC:    endpointGRPC,
	}

	keyParams := KeyParams{
		Mnemonic: mnemonic,
		Name:     DefaultKeyName,
		Password: DefaultKeyPass,
	}

	adapter, err := NewBSNIritaAdapter(endpoint, keyParams)
	if err != nil {
		logger.Errorf("Failed to create the BSN-IRITA adapter: %s", err)
		return
	}

	RunWebServer(listenAddr, adapter.handle)
}
