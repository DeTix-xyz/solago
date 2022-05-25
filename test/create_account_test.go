package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DeTix-xyz/solago/src/rpc"
	"github.com/DeTix-xyz/solago/src/sdk"
	"github.com/DeTix-xyz/solago/src/sdk/account"
	"github.com/DeTix-xyz/solago/src/sdk/solana"
	"github.com/google/uuid"
)

func TestCreateAccount(t *testing.T) {
	// Sugar daddy
	payer := account.NewKeypairFromFile("/Users/trumanpurnell/.config/solana/id.json")

	// New account to be created
	newAccount := account.NewKeypairFromSeed([32]byte{})

	// Transaction to create account
	transaction := solana.Transaction{
		Signatures: solana.NewSignatures(payer.PrivateKey, newAccount.PrivateKey),
		Message: solana.Message{
			Header:           solana.NewMessageHeader(2, 0, 1),
			AccountAddresses: solana.CompactArray{3, []any{payer.PublicKey, newAccount.PublicKey, account.SystemProgram}},
			RecentBlockhash:  sdk.RecentBlockhashFromString("7xq4MaWpVTyTsRG13GGFMHYLKx2sPQkzoRFwv7SRBSTb"),
			Instructions: sdk.NewCompactArray(1, &sdk.Instruction{
				ProgramIDIndex:        sdk.SerializableUInt8(2),
				AccountAddressIndexes: sdk.NewCompactArray(2, sdk.SerializableUInt8(0), sdk.SerializableUInt8(1)),
				Data: sdk.NewCompactArray(52, &sdk.InstructionData{
					Data: struct {
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

	// [2 38 106 131 110 6 93 82 233 39 90 222 244 13 243 104 210 71 204 62 30 65 64
	//  160 233 81 146 84 42 6 140 242 54 7 208 41 165 69 226 236 4 5 78 97 90 180 157
	//  32 68 77 144 243 71 6 47 44 105 14 153 229 74 42 11 109 161 69 69
	//  0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 67 61 86 191 179 216
	//  153 102 245 212 252 150 200 156 195 15 45 89 35 48 144 246 147 69 130 91 226 165
	//  112 120 18 178 2 0 1 3 7 208 41 165 69 226 236 4 5 78 97 90 180 157 32 68 77 144
	//  243 71 6 47 44 105 14 153 229 74 42 11 109 161 67 61 86 191 179 216 153 102 245 212
	//  252 150 200 156 195 15 45 89 35 48 144 246 147 69 130 91 226 165 112 120 18 178
	//  0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 44 171 30 55 214
	//  130 1 151 253 144 209 47 109 4 158 254 100 29 63 64 152 229 107 225 250 9 156 178 117 209
	//  112 183 1 2 2 0 1 52 0 0 0 0 0 225 245 5 0 0 0 0 32
	//  0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
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
