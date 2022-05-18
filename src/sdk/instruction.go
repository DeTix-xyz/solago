package sdk

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

type Instruction struct {
	ProgramIDIndex        SerializableUInt8
	AccountAddressIndexes *CompactArray[SerializableUInt8]
	Data                  *CompactArray[*InstructionData]
}

func (instruction *Instruction) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	instruction.ProgramIDIndex.Serialize(buffer)
	instruction.AccountAddressIndexes.Serialize(buffer)
	instruction.Data.Serialize(buffer)

	return buffer
}

type InstructionData struct {
	Data interface{}
}

func (instructionData *InstructionData) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	bytes, ok := instructionData.Data.([]byte) // we may simply be passed a byte array

	if ok {
		buffer.Write(bytes)
		return buffer
	}

	// otherwise parse the struct
	structValues := reflect.ValueOf(instructionData.Data)

	for i := 0; i < structValues.NumField(); i++ {
		binary.Write(buffer, binary.LittleEndian, structValues.Field(i).Interface())
	}

	return buffer
}

type SystemInstruction uint32

const (
	CreateAccount SystemInstruction = iota
	Assign
	TransferSystem
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

type TokenInstruction uint8

const (
	InitializeMint TokenInstruction = iota
	InitializeAccount
	InitializeMultisig
	TransferToken
	Approve
	Revoke
	SetAuthority
	MintTo
	Burn
	CloseAccount
	FreezeAccount
	ThawAccount
	TransferChecked
	ApproveChecked
	MintToChecked
	BurnChecked
	InitializeAccount2
	SyncNative
	InitializeAccount3
	InitializeMultisig2
	InitializeMint2
)

type MetadataInstruction uint8

const (
	CreateMetadataAccount = iota
	UpdateMetadataAccount
	DeprecatedCreateMasterEdition
	DeprecatedMintNewEditionFromMasterEditionViaPrintingToken
	UpdatePrimarySaleHappenedViaToken
	DeprecatedSetReservationList
	DeprecatedCreateReservationList
	SignMetadata
	DeprecatedMintPrintingTokensViaToken
	DeprecatedMintPrintingTokens
	CreateMasterEdition
	MintNewEditionFromMasterEditionViaToken
	ConvertMasterEditionV1ToV2
	MintNewEditionFromMasterEditionViaVaultProxy
	PuffMetadata
	UpdateMetadataAccountV2
	CreateMetadataAccountV2
	CreateMasterEditionV3
	VerifyCollection
	Utilize
	ApproveUseAuthority
	RevokeUseAuthority
	UnverifyCollection
	ApproveCollectionAuthority
	RevokeCollectionAuthority
	SetAndVerifyCollection
	FreezeDelegatedAccount
	ThawDelegatedAccount
	RemoveCreatorVerification
)
