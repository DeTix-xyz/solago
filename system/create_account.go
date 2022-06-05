package system

import (
	"bytes"
	"encoding/binary"

	"github.com/deezdegens/solago"
	"github.com/deezdegens/solago/utils"
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

func (instruction *CreateAccountInstruction) ProgramIDIndex(accounts solago.AccountList) uint8 {
	return utils.IndexOf(accounts, SystemAccount)[0]
}

func (instruction *CreateAccountInstruction) AccountAddressIndexes(accounts solago.AccountList) solago.CompactArray {
	indexes := utils.IndexOf(
		accounts,
		solago.NewSignerAccount(instruction.Payer),
		solago.NewSignerAccount(instruction.NewAccount),
	)

	return solago.NewCompactArray(indexes)
}

func (instruction *CreateAccountInstruction) CollectAccounts() solago.AccountList {
	return solago.AccountList{
		solago.NewSignerAccount(instruction.Payer),
		solago.NewSignerAccount(instruction.NewAccount),
		SystemAccount,
	}
}

func (instruction *CreateAccountInstruction) Data() solago.CompactArray {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.LittleEndian, CreateAccount)
	binary.Write(buffer, binary.LittleEndian, instruction.Lamports)
	binary.Write(buffer, binary.LittleEndian, instruction.Space)
	binary.Write(buffer, binary.LittleEndian, instruction.Owner)

	var bytes solago.ByteList = buffer.Bytes()

	return solago.NewCompactArray(bytes)
}
