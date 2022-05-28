package solana

type CompactArray struct {
	Length uint16
	Items  []any
}

func NewCompactArray(items ...any) CompactArray {
	numKeys := uint16(len(items))

	return CompactArray{numKeys, items}
}
