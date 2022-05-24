package transaction

type Instruction struct {
	ProgramIDIndex        uint8
	AccountAddressIndexes *CompactArray
	Data                  *CompactArray
}
