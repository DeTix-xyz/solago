package rpc

import (
	"encoding/json"

	"github.com/google/uuid"
)

func (client *Client) GetMinimumRent(size uint) uint64 {
	responseBytes, _ := client.Call(
		&Request{
			Version: JSON_RPC_VERSION,
			ID:      uuid.NewString(),
			Method:  "getMinimumBalanceForRentExemption",
			Params:  []uint{size},
		},
	)

	minimumRent := &struct {
		Response
		Result uint64 `json:"result"`
	}{}

	json.Unmarshal(responseBytes, minimumRent)

	return minimumRent.Result
}
