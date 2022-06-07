package system

import (
	"bytes"
	"encoding/binary"

	"github.com/deezdegens/solago"
	"github.com/deezdegens/solago/utils"
)

type CreateAccountInstruction struct {
	Payer      solago.Account
	NewAccount solago.Account
	Lamports   uint64
	Space      uint64
	Owner      solago.PublicKey
}

func (instruction CreateAccountInstruction) ProgramIDIndex(accounts solago.AccountList) uint8 {
	return utils.IndexOf(accounts, Account)[0]
}

func (instruction CreateAccountInstruction) AccountAddressIndexes(accounts solago.AccountList) solago.CompactArray {
	indexes := utils.IndexOf(
		accounts,
		instruction.Payer,
		instruction.NewAccount,
	)

	return solago.NewCompactArray(indexes)
}

func (instruction CreateAccountInstruction) CollectAccounts() solago.AccountList {
	return solago.AccountList{
		instruction.Payer,
		instruction.NewAccount,
		Account,
	}
}

func (instruction CreateAccountInstruction) Data() solago.CompactArray {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.LittleEndian, CreateAccount)
	binary.Write(buffer, binary.LittleEndian, instruction.Lamports)
	binary.Write(buffer, binary.LittleEndian, instruction.Space)
	binary.Write(buffer, binary.LittleEndian, instruction.Owner)

	var bytes solago.ByteList = buffer.Bytes()

	return solago.NewCompactArray(bytes)
}
