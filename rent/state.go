package rent

import "github.com/deezdegens/solago"

var Program = solago.NewPublicKey("SysvarRent111111111111111111111111111111111")
var Account = solago.NewReadOnlyAccount(solago.Keypair{PublicKey: Program})
