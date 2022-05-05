package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DeTix-xyz/solago/src/rpc"
	"github.com/DeTix-xyz/solago/src/sdk"
	"github.com/google/uuid"
)

func TestCreateAccount(t *testing.T) {
	// Sugar daddy
	payer := sdk.NewKeypairFromFile("/Users/trumanpurnell/.config/solana/id.json")

	// New account to be created
	account := sdk.NewKeypairFromSeed([32]byte{})

	// Transaction to create account
	transaction := &sdk.Transaction{
		Signatures: sdk.NewCompactArray(2,
			sdk.SerializablePrivateKey(payer.PrivateKey),
			sdk.SerializablePrivateKey(account.PrivateKey),
		),
		Message: sdk.Message{
			Header: sdk.NewMessageHeader(2, 0, 1),
			AccountAddresses: sdk.NewCompactArray(3,
				sdk.SerializablePublicKey(payer.PublicKey),
				sdk.SerializablePublicKey(account.PublicKey),
				sdk.SerializablePublicKey(sdk.SystemProgram),
			),
			RecentBlockhash: sdk.RecentBlockhashFromString("7xq4MaWpVTyTsRG13GGFMHYLKx2sPQkzoRFwv7SRBSTb"),
			Instructions: sdk.NewCompactArray(1, &sdk.Instruction{
				ProgramIDIndex:        sdk.SerializableUInt8(2),
				AccountAddressIndexes: sdk.NewCompactArray(2, sdk.SerializableUInt8(0), sdk.SerializableUInt8(1)),
				Data: sdk.NewCompactArray(52, &sdk.InstructionData{
					struct {
						Instruction sdk.SystemInstruction     // 4 +
						Lamports    uint64                    // 8 +
						Space       uint64                    // 8 +
						Owner       sdk.SerializablePublicKey // 32 == 52
					}{
						Instruction: sdk.CreateAccount,
						Lamports:    1_000_000_000 / 10,
						Space:       32,
						Owner:       sdk.SerializablePublicKey(sdk.PublicKey("7JM3jwj2hp9ULM6mqCrtX6PKeeG6C5STPPsFwXBF36CF")),
					},
				}),
			}),
		},
	}

	buffer := new(bytes.Buffer)

	transaction.Serialize(buffer)
	signedTransaction := transaction.Sign(buffer)

	fmt.Println(signedTransaction)
}

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
