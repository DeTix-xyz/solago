package system

import "github.com/DeTix-xyz/solago"

type SystemInstruction uint32

const (
	CreateAccount SystemInstruction = iota
	Assign
	Transfer
	CreateAccountWithSeed
	AdvanceNonceAccount
	WithdrawNonceAccount
	InitializeNonceAccount
	AuthorizeNonceAccount
	Allocate
	AllocateWithSeed
	AssignWithSeed
	TransferWithSeed
)

var SystemProgramAccount = solago.NewPublicKey("11111111111111111111111111111111")
