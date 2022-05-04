package sdk

import (
	"bytes"
)

type CompactArray[T Serializable] struct {
	Length uint16
	Items  []T
}

func (compactArray *CompactArray[T]) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	WriteUvarint(buffer, compactArray.Length)

	for _, serializable := range compactArray.Items {
		serializable.Serialize(buffer)
	}

	return buffer
}

func NewCompactArray[T Serializable](Length uint16, Items ...T) *CompactArray[T] {
	return &CompactArray[T]{Length, Items}
}
