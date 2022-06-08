package rpc

import (
	"encoding/json"

	"github.com/google/uuid"
)

func (client *Client) GetMintAccountAuthority(mintPublicKey string) string {
	responseBytes, _ := client.Call(
		&Request{
			Version: JSON_RPC_VERSION,
			ID:      uuid.NewString(),
			Method:  "getAccountInfo",
			Params: []any{
				mintPublicKey,
				map[string]string{
					"encoding": "jsonParsed",
				},
			},
		},
	)

	mintAccount := &struct {
		Response
		Result struct {
			Value struct {
				Data struct {
					Parsed struct {
						Info struct {
							Decimals        int    `json:"decimals"`
							FreezeAuthority string `json:"freezeAuthority"`
							IsInitialized   bool   `json:"isInitialized"`
							MintAuthority   string `json:"mintAuthority"`
							Supply          int    `json:"supply"`
						} `json:"info"`
						Type string `json:"type"`
					} `json:"parsed"`
					Program string `json:"program"`
					Space   int    `json:"space"`
				} `json:"data"`
			} `json:"value"`
		} `json:"result"`
	}{}

	json.Unmarshal(responseBytes, mintAccount)

	return mintAccount.Result.Value.Data.Parsed.Info.MintAuthority
}
