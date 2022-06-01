package solago

type CompactArray[T byte | PublicKey | PrivateKey | Instruction] struct {
	Length uint16
	Items  []T
}
