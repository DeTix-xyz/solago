package solana

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

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

type CreateAccountInstruction struct {
	Payer      Keypair
	NewAccount Keypair
	Lamports   uint64
	Space      uint64
	Owner      PublicKey
}

func (instruction *CreateAccountInstruction) Size() int {
	return int(reflect.TypeOf(CreateAccount).Size()) +
		int(reflect.TypeOf(instruction.Lamports).Size()) +
		int(reflect.TypeOf(instruction.Space).Size()) +
		int(reflect.TypeOf(instruction.Owner).Size())
}

func (instruction *CreateAccountInstruction) ProgramIDIndex() uint8 {
	return 0
}

func (instruction *CreateAccountInstruction) AccountAddressIndexes() CompactArray {
	return CompactArray{}
}

func (instruction *CreateAccountInstruction) CollectAccounts() []Account {
	return []Account{
		NewSignerAccount(instruction.Payer),
		NewSignerAccount(instruction.NewAccount),
		SystemProgramAccount,
	}
}

func (instruction *CreateAccountInstruction) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	binary.Write(buffer, binary.LittleEndian, CreateAccount)
	binary.Write(buffer, binary.LittleEndian, instruction.Lamports)
	binary.Write(buffer, binary.LittleEndian, instruction.Space)
	binary.Write(buffer, binary.LittleEndian, instruction.Owner)

	return buffer
}
