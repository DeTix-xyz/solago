package solana

type Instruction struct {
	ProgramIDIndex        uint8
	AccountAddressIndexes CompactArray[uint8]
	Data                  CompactArray[byte]
}

type instruction interface {
	ProgramIDIndex([]Account) uint8
	AccountAddressIndexes([]Account) CompactArray[uint8]
	CollectAccounts() []Account
	Data() CompactArray[byte]
}
