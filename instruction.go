package solago

import (
	"bytes"
	"encoding/binary"
)

type ProgramIDIndex uint8

func (index ProgramIDIndex) Serialize(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.LittleEndian, index)
}

type Instruction struct {
	ProgramIDIndex        ProgramIDIndex
	AccountAddressIndexes CompactArray
	Data                  CompactArray
}

func (instruction Instruction) Serialize(buffer *bytes.Buffer) {
	instruction.ProgramIDIndex.Serialize(buffer)
	instruction.AccountAddressIndexes.Serialize(buffer)
	instruction.Data.Serialize(buffer)
}

type InstructionList []Instruction

func (instructions InstructionList) Serialize(buffer *bytes.Buffer) {
	for _, instruction := range instructions {
		instruction.Serialize(buffer)
	}
}

type PseudoInstruction interface {
	ProgramIDIndex([]Account) ProgramIDIndex
	AccountAddressIndexes([]Account) CompactArray
	CollectAccounts() []Account
	Data() CompactArray
}

type PseudoInstructionList []PseudoInstruction

func (pseudoInstructions PseudoInstructionList) NewInstructionList(accounts AccountList) InstructionList {
	instructions := InstructionList{}

	for _, pseudoInstruction := range pseudoInstructions {
		instruction := Instruction{
			ProgramIDIndex:        pseudoInstruction.ProgramIDIndex(accounts),
			AccountAddressIndexes: pseudoInstruction.AccountAddressIndexes(accounts),
			Data:                  pseudoInstruction.Data(),
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func (pseudoInstructions PseudoInstructionList) CollectAccounts() AccountList {
	accounts := AccountList{}

	for _, pseudoInstruction := range pseudoInstructions {
		accounts = append(accounts, pseudoInstruction.CollectAccounts()...)
	}

	return accounts.Sort()
}
