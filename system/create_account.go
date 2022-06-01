package system

import (
	"bytes"
	"encoding/binary"
)

var SystemAccount = NewReadOnlyAccount(Keypair{PublicKey: SystemProgramAccount})

type CreateAccountInstruction struct {
	Payer      Keypair
	NewAccount Keypair
	Lamports   uint64
	Space      uint64
	Owner      PublicKey
}

func (instruction *CreateAccountInstruction) ProgramIDIndex(accounts []Account) uint8 {
	return indexOf(accounts, SystemAccount)[0]
}

func (instruction *CreateAccountInstruction) AccountAddressIndexes(accounts []Account) CompactArray[uint8] {
	indexes := indexOf(
		accounts,
		NewSignerAccount(instruction.Payer),
		NewSignerAccount(instruction.NewAccount),
	)

	return CompactArray[uint8]{uint16(len(indexes)), indexes}
}

func (instruction *CreateAccountInstruction) CollectAccounts() []Account {
	return []Account{
		NewSignerAccount(instruction.Payer),
		NewSignerAccount(instruction.NewAccount),
		SystemAccount,
	}
}

func (instruction *CreateAccountInstruction) Data() CompactArray[byte] {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.LittleEndian, CreateAccount)
	binary.Write(buffer, binary.LittleEndian, instruction.Lamports)
	binary.Write(buffer, binary.LittleEndian, instruction.Space)
	binary.Write(buffer, binary.LittleEndian, instruction.Owner)

	bytes := buffer.Bytes()

	return CompactArray[byte]{uint16(len(bytes)), bytes}
}
