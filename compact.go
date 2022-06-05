package solago

import (
	"bytes"
)

type CompactArray struct {
	Length uint16
	Items  *[]Serializable
}

func NewCompactArray(items ...Serializable) CompactArray {
	return CompactArray{
		Length: uint16(len(items)),
		Items:  &items,
	}
}

func (array *CompactArray) Serialize(buffer *bytes.Buffer) {
	WriteUvarint(buffer, array.Length)

	for _, item := range *array.Items {
		item.Serialize(buffer)
	}
}
