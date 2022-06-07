package token

import "github.com/deezdegens/solago"

const SizeOfMint = 82

var Program = solago.NewPublicKey("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
var Account = solago.NewReadOnlyAccount(solago.Keypair{PublicKey: Program})
