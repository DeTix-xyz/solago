package system

import (
	"bytes"
	"encoding/binary"

	"github.com/DeTix-xyz/solago"
	"github.com/DeTix-xyz/solago/utils"
)

var SystemAccount = solago.NewReadOnlyAccount(
	solago.Keypair{PublicKey: SystemProgramAccount},
)

type CreateAccountInstruction struct {
	Payer      solago.Keypair
	NewAccount solago.Keypair
	Lamports   uint64
	Space      uint64
	Owner      solago.PublicKey
}

func (instruction *CreateAccountInstruction) ProgramIDIndex(accounts []solago.Account) uint8 {
	return utils.IndexOf(accounts, SystemAccount)[0]
}

func (instruction *CreateAccountInstruction) AccountAddressIndexes(accounts []solago.Account) solago.CompactArray[uint8] {
	indexes := utils.IndexOf(
		accounts,
		solago.NewSignerAccount(instruction.Payer),
		solago.NewSignerAccount(instruction.NewAccount),
	)

	return solago.CompactArray[uint8]{
		Length: uint16(len(indexes)),
		Items:  indexes,
	}
}

func (instruction *CreateAccountInstruction) CollectAccounts() []solago.Account {
	return []solago.Account{
		solago.NewSignerAccount(instruction.Payer),
		solago.NewSignerAccount(instruction.NewAccount),
		SystemAccount,
	}
}

func (instruction *CreateAccountInstruction) Data() solago.CompactArray[byte] {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.LittleEndian, CreateAccount)
	binary.Write(buffer, binary.LittleEndian, instruction.Lamports)
	binary.Write(buffer, binary.LittleEndian, instruction.Space)
	binary.Write(buffer, binary.LittleEndian, instruction.Owner)

	return solago.NewCompactArray(buffer.Bytes()...)
}
