package rpc

import (
	"encoding/json"

	"github.com/google/uuid"
)

func (client *Client) GetRecentBlockhash() string {
	responseBytes, _ := client.Call(
		&Request{
			Version: JSON_RPC_VERSION,
			ID:      uuid.NewString(),
			Method:  "getLatestBlockhash",
			Params: []map[string]string{
				{"commitment": "finalized"},
			},
		},
	)

	blockhashResponse := &struct {
		Response
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

	json.Unmarshal(responseBytes, blockhashResponse)

	return blockhashResponse.Result.Value.Blockhash
}
