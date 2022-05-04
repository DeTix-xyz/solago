package test

import (
	"encoding/json"
	"testing"

	"github.com/DeTix-xyz/solago/src/rpc"
	"github.com/google/uuid"
)

func TestGetBlockhash(t *testing.T) {

	client := rpc.NewClient("https://api.devnet.solana.com", nil)

	responseBytes, err := client.Call(
		&rpc.JSONRPCRequest{
			Version: rpc.JSON_RPC_VERSION,
			ID:      uuid.NewString(),
			Method:  "getLatestBlockhash",
			Params: []map[string]string{
				{"commitment": "processed"},
			},
		},
	)

	if err != nil {
		t.Fatal("Unable to hit rpc endpoint: ", err)
	}

	blockhashResponse := &struct {
		rpc.JSONRPCResponse
		Result struct {
			Context struct {
				Slot int `json:"slot"`
			} `json:"context"`
			Value struct {
				Blockhash            string `json:"blockhash"`
				LastValidBlockHeight int    `json:"lastValidBlockHeight"`
			} `json:"value"`
		} `json:"result"`
	}{}

	err = json.Unmarshal(responseBytes, &blockhashResponse)

	if err != nil {
		t.Fatal("Unable to marshal JSON: ", err)
	}
}
