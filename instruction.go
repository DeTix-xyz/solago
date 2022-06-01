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

	for _, entry := range collection {
		accounts = append(accounts, entry.CollectAccounts()...)
	}

	return accounts
}
