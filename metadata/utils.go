package metadata

import "github.com/deezdegens/solago"

func DeriveMetadataAccountFromMint(mintAccount solago.Account) solago.Account {
	metadataPublicKey := solago.FindProgramAddress(
		[]byte("metadata"),
		Program,
		mintAccount.Keypair.PublicKey,
		Program,
	)

	return solago.NewReadWriteAccount(solago.Keypair{
		PublicKey: metadataPublicKey,
	})
}
