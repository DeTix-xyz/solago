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

	// New mint account
	mintAccount := solago.NewSignerAccount(solago.NewKeypair())
	fmt.Println(mintAccount.Keypair.PublicKey.ToBase58())

	// New metadata account (derived from mint public key)
	metadataAccount := metadata.DeriveMetadataAccountFromMint(mintAccount)

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
			PayerAccount:  payerAccount,
			MintAccount:   mintAccount,
			Decimals:      0,
			MintAuthority: mintAccount.Keypair.PublicKey,
			FreezeAccount: system.Program,
		},
		metadata.CreateMetadataAccountInstruction{
			Payer:                payerAccount,
			MintAccount:          mintAccount,
			MintAuthorityAccount: mintAccount,
			UpdateAuthority:      payerAccount,
			MetadataAccount:      metadataAccount,
			Metadata: metadata.Metadata{
				Name:                 "Duh Duh Duh",
				Symbol:               "DDD",
				URI:                  "https://gj2j5tocsgwcaymr43ta2guiei6rfctsoxexypo3qbv2r7uimhha.arweave.net/MnSezcKRrCBhkebmDRqIIj0SinJ1yXw924BrqP6IYc4",
				SellerFeeBasisPoints: 99_99,
				Creators:             nil,
				Collection:           nil,
				Uses:                 nil,
				IsMutable:            true,
			},
		},
	)

	fmt.Println(confirmation)
}
