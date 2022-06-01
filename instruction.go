package solago

type Instruction struct {
	ProgramIDIndex        uint8
	AccountAddressIndexes CompactArray[uint8]
	Data                  CompactArray[byte]
}

type PreInstruction interface {
	ProgramIDIndex([]Account) uint8
	AccountAddressIndexes([]Account) CompactArray[uint8]
	CollectAccounts() []Account
	Data() CompactArray[byte]
}
