package test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/DeTix-xyz/solago"
	"github.com/DeTix-xyz/solago/system"
)

func TestCreateAccount2(t *testing.T) {
	// Sugar daddy
	payer := solago.NewKeypairFromFile("/Users/trumanpurnell/.config/solana/id.json")

	// New account to be created
	newAccount := solago.NewKeypairFromSeed([32]byte{})

	// Transaction to create account
	client := solago.NewClient("https://api.devnet.solana.com", nil)
	
	signedTransaction := client.
		NewTransaction(system.CreateAccountInstruction{
			Payer:      *payer,
			NewAccount: *newAccount,
			Lamports:   1_000_000_000 / 10,
			Space:      32,
			Owner:      payer.PublicKey,
		})
		.Sign(payer.PrivateKey, newnewAccount.PrivateKey)


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

