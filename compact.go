package solago

type CompactArray[T byte | PublicKey | PrivateKey | Instruction] struct {
	Length uint16
	Items  []T
}

func NewCompactArray[T byte | PublicKey | PrivateKey | Instruction](items ...T) CompactArray[T] {
	return CompactArray[T]{
		Length: uint16(len(items)),
		Items:  items,
	}
}
