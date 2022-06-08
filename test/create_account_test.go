package test

import (
	"fmt"
	"testing"

	"github.com/deezdegens/solago"
	"github.com/deezdegens/solago/system"
	"github.com/deezdegens/solago/token"
)

func TestCreateAccount(t *testing.T) {
	// Sugar daddy
	payerAccount := solago.SignerAccountFromFile("/Users/trumanpurnell/.config/solana/id.json")

	// New account to be created
	newAccount := solago.NewSignerAccount(solago.NewKeypair())

	fmt.Println(newAccount.Keypair.PublicKey.ToBase58())

	// Transaction to create account
	client := solago.NewClient("https://api.devnet.solana.com")

	confirmation := client.SendTransaction(
		system.CreateAccountInstruction{
			Payer:      payerAccount,
			NewAccount: newAccount,
			Lamports:   client.RPC.GetMinimumRent(token.SizeOfMint),
			Space:      token.SizeOfMint,
			Owner:      token.Program,
		},
	)

	fmt.Println(confirmation)
}
