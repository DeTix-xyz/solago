package token

import (
	"bytes"
	"encoding/binary"

	"github.com/deezdegens/solago"
	"github.com/deezdegens/solago/utils"
)

type InitializeMint2Instruction struct {
	MintAccount   solago.Account
	Decimals      uint8
	MintAuthority solago.PublicKey
	FreezeAccount solago.PublicKey
}

func (instruction InitializeMint2Instruction) ProgramIDIndex(accounts solago.AccountList) uint8 {
	return utils.IndexOf(accounts, Account)[0]
}

func (instruction InitializeMint2Instruction) AccountAddressIndexes(accounts solago.AccountList) solago.CompactArray {
	indexes := utils.IndexOf(
		accounts,
		instruction.MintAccount,
	)

	return solago.NewCompactArray(indexes)
}

func (instruction InitializeMint2Instruction) CollectAccounts() solago.AccountList {
	return solago.AccountList{
		instruction.MintAccount,
		Account,
	}
}

func (instruction InitializeMint2Instruction) Data() solago.CompactArray {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.LittleEndian, InitializeMint2)
	binary.Write(buffer, binary.LittleEndian, instruction.Decimals)
	binary.Write(buffer, binary.LittleEndian, instruction.MintAuthority)
	binary.Write(buffer, binary.LittleEndian, instruction.FreezeAccount)

	var bytes solago.ByteList = buffer.Bytes()

	return solago.NewCompactArray(bytes)
}
