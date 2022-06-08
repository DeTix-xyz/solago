package test

import (
	"fmt"
	"testing"

	"github.com/deezdegens/solago"
	"github.com/deezdegens/solago/metadata"
	"github.com/deezdegens/solago/system"
	"github.com/deezdegens/solago/token"
)

func TestCreateMetadata(t *testing.T) {
	// Sugar daddy
	payerAccount := solago.SignerAccountFromFile("/Users/trumanpurnell/.config/solana/id.json")

	// New mint to be created
	mintAccount := solago.NewSignerAccount(solago.NewKeypair())

	// Metadata account must be derived from mint public key
	metadataAccountPublicKey, _ := solago.FindProgramAddress(
		[]byte("metadata"),
		metadata.Program,
		instruction.MintAccount.Keypair.PublicKey,
		metadata.Program,
	)

	fmt.Println(mintAccount.Keypair.PublicKey.ToBase58())

	// Transaction to create the mint
	client := solago.NewClient("https://api.devnet.solana.com")

	confirmation := client.SendTransaction(
		system.CreateAccountInstruction{
			Payer:      payerAccount,
			NewAccount: mintAccount,
			Lamports:   client.RPC.GetMinimumRent(token.SizeOfMint),
			Space:      token.SizeOfMint,
			Owner:      token.Program,
		},
		token.InitializeMint2Instruction{
			MintAccount:   mintAccount,
			Decimals:      0,
			MintAuthority: payerAccount.Keypair.PublicKey,
			FreezeAccount: system.Program,
		},
	)

	fmt.Println(confirmation)
}
