package solago

type Instruction struct {
	ProgramIDIndex        uint8
	AccountAddressIndexes CompactArray[uint8]
	Data                  CompactArray[byte]
}

type InProcessInstruction interface {
	ProgramIDIndex([]Account) uint8
	AccountAddressIndexes([]Account) CompactArray[uint8]
	CollectAccounts() []Account
	Data() CompactArray[byte]
}

type InProcessInstructionCollection []InProcessInstruction

func (collection InProcessInstructionCollection) CollectAccounts() AccountCollection {
	accounts := []Account{}

	for _, instruction := range collection {
		accounts = append(accounts, instruction.CollectAccounts()...)
	}

	return accounts
}

func (collection InProcessInstructionCollection) MapToRaw() []Instruction {
	instructions := []Instruction{}
	sortedAccounts := collection.CollectAccounts().Sort()

	for _, instruction := range collection {
		instructions = append(instructions, Instruction{
			ProgramIDIndex:        instruction.ProgramIDIndex(sortedAccounts),
			AccountAddressIndexes: instruction.AccountAddressIndexes(sortedAccounts),
			Data:                  instruction.Data(),
		})
	}

	return instructions
}
