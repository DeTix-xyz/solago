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
