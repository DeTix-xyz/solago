package rpc

import (
	"encoding/json"

	"github.com/google/uuid"
)

func (client *Client) GetMintAccount(mintPublicKey string) string {
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

	// {
	// 	"data": {
	// 		"parsed": {
	// 			"info": {
	// 				"decimals": 0,
	// 				"freezeAuthority": null,
	// 				"isInitialized": true,
	// 				"mintAuthority": "XVxLJDE9nTh3xNegbgXS9eUSBZhEj8YxVgCtr9jHhhJ",
	// 				"supply": "0"
	// 			},
	// 			"type": "mint"
	// 		},
	// 		"program": "spl-token",
	// 		"space": 82
	// 	},
	// 	"executable": false,
	// 	"lamports": 1461600,
	// 	"owner": "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
	// 	"rentEpoch": 322
	// }

	minimumRent := &struct {
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

	json.Unmarshal(responseBytes, minimumRent)

	return minimumRent.Result.Value.Data.Parsed.Info.MintAuthority
}
