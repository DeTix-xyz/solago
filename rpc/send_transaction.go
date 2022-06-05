package rpc

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func (client *Client) SendTransaction(transaction string) string {
	responseBytes, _ := client.Call(
		&Request{
			Version: JSON_RPC_VERSION,
			ID:      uuid.NewString(),
			Method:  "sendTransaction",
			Params: []any{
				transaction,
				map[string]string{
					"encoding": "base64",
				},
			},
		},
	)

	sendTransactionResponse := &struct {
		Response
		Result string `json:"result"`
	}{}

	json.Unmarshal(responseBytes, sendTransactionResponse)
	fmt.Println(sendTransactionResponse.Error)

	return sendTransactionResponse.Result
}
