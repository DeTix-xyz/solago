package metadata

import (
	"bytes"
	"encoding/binary"

	"github.com/deezdegens/solago"
	"github.com/deezdegens/solago/rent"
	"github.com/deezdegens/solago/system"
	"github.com/deezdegens/solago/utils"
	"github.com/near/borsh-go"
)

type CreateMetadataAccountInstruction struct {
	Payer                solago.Account
	MintAccount          solago.Account
	MintAuthorityAccount solago.Account // leave Init to fetch
	UpdateAuthority      solago.Account
	MetadataAccount      solago.Account // leave Init to derive
	Metadata             Metadata
}

func (instruction CreateMetadataAccountInstruction) ProgramIDIndex(accounts solago.AccountList) uint8 {
	return utils.IndexOf(accounts, Account)[0]
}

func (instruction CreateMetadataAccountInstruction) AccountAddressIndexes(accounts solago.AccountList) solago.CompactArray {
	indexes := utils.IndexOf(
		accounts,
		instruction.MetadataAccount,
		instruction.MintAccount,
		instruction.MintAuthorityAccount,
		instruction.Payer,
		instruction.UpdateAuthority,
		system.Account,
		rent.Account,
	)

	return solago.NewCompactArray(indexes)
}

func (instruction CreateMetadataAccountInstruction) CollectAccounts() solago.AccountList {
	return solago.AccountList{
		instruction.MetadataAccount,
		instruction.MintAccount,
		instruction.MintAuthorityAccount,
		instruction.Payer,
		instruction.UpdateAuthority,
		system.Account,
		rent.Account,
		Account,
	}
}

func (instruction CreateMetadataAccountInstruction) Data() solago.CompactArray {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.LittleEndian, CreateMetadataAccountV2)

	data, _ := borsh.Serialize(instruction.Metadata)
	buffer.Write(data)

	return solago.NewCompactArray(solago.ByteList(buffer.Bytes()))
}
