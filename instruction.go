package solago

import (
	"bytes"
	"encoding/binary"
)

type Instruction struct {
	ProgramIDIndex        uint8
	AccountAddressIndexes CompactArray
	Data                  CompactArray
}

func (instruction Instruction) Serialize(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.LittleEndian, instruction.ProgramIDIndex)
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
	ProgramIDIndex(AccountList) uint8
	AccountAddressIndexes(AccountList) CompactArray
	CollectAccounts() AccountList
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
		for _, account := range pseudoInstruction.CollectAccounts() {
			accounts = append(accounts, account)
		}
	}

	return accounts.Sort()
}
