package solana

import "bytes"

type Instruction interface {
	Size() int
	ProgramIDIndex([]Account) uint8
	AccountAddressIndexes([]Account) CompactArray
	CollectAccounts() []Account
	Serialize(buffer *bytes.Buffer) *bytes.Buffer
}
