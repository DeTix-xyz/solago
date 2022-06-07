package system

import "github.com/deezdegens/solago"

var Program = solago.NewPublicKey("11111111111111111111111111111111")
var Account = solago.NewReadOnlyAccount(solago.Keypair{PublicKey: Program})
