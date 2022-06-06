package solago

import (
	"bytes"
	"reflect"
)

type CompactArray struct {
	Length uint16
	Items  Serializable
}

func NewCompactArray(items Serializable) CompactArray {
	return CompactArray{
		Length: uint16(reflect.ValueOf(items).Len()),
		Items:  items,
	}
}

func (array CompactArray) Serialize(buffer *bytes.Buffer) {
	WriteUvarint(buffer, array.Length)
	array.Items.Serialize(buffer)
}
