package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	servicesdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	"github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"
)

type Request struct {
	RequestID string      `json:"request_id"`
	Result    interface{} `json:"result"`
}

type ServiceRequestResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ServiceRequestOutput struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

type KeyParams struct {
	Path     string
	Name     string
	Password string
	Address  types.AccAddress
}

type Endpoint struct {
	ChainID string
	RPC     string
	GRPC    string
}

type BSNIritaAdapter struct {
	Client    servicesdk.ServiceClient
	KeyParams KeyParams
}

func NewBSNIritaAdapter(endpoint Endpoint, keyParams KeyParams) (*BSNIritaAdapter, error) {
	cfg := types.ClientConfig{
		ChainID:  endpoint.ChainID,
		NodeURI:  endpoint.RPC,
		GRPCAddr: endpoint.GRPC,
		KeyDAO:   store.NewFileDAO(keyParams.Path),
		Mode:     types.Commit,
	}

	client := servicesdk.NewServiceClient(cfg)

	_, address, err := client.Find(keyParams.Name, keyParams.Password)
	if err != nil {
		return nil, err
	}

	keyParams.Address = address

	return &BSNIritaAdapter{
		Client:    client,
		KeyParams: keyParams,
	}, nil
}

func (adapter BSNIritaAdapter) handle(req Request) (interface{}, error) {
	requestIDBz, err := hex.DecodeString(req.RequestID)
	if err != nil {
		return nil, err
	}

	result, output := adapter.buildServiceResponse(req)

	msg := service.MsgRespondService{
		RequestId: requestIDBz,
		Result:    result,
		Output:    output,
		Provider:  adapter.KeyParams.Address,
	}

	baseTx := types.BaseTx{
		From:     adapter.KeyParams.Name,
		Password: adapter.KeyParams.Password,
	}

	res, err := adapter.Client.BuildAndSend([]types.Msg{&msg}, baseTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %s", err.Error())
	}

	return res, nil
}

func (adapter BSNIritaAdapter) buildServiceResponse(req Request) (result, output string) {
	res := ServiceRequestResult{
		Code: 200,
	}

	if req.Result == nil {
		res.Code = 500
	}

	resBz, _ := json.Marshal(res)

	if res.Code == 200 {
		out := ServiceRequestOutput{
			Header: "{}",
			Body:   req.Result.(string),
		}

		outputBz, _ := json.Marshal(out)
		output = string(outputBz)
	}

	return string(resBz), output
}
