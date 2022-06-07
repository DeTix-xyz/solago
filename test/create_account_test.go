package test

import (
	"fmt"
	"testing"

	"github.com/deezdegens/solago"
	"github.com/deezdegens/solago/system"
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
			Lamports:   1_000_000_000 / 10,
			Space:      32,
			Owner:      system.Account.Keypair.PublicKey,
		},
	)

	fmt.Println(confirmation)
}
