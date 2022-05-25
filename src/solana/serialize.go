package solana

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"github.com/DeTix-xyz/solago/src/metadata"
	"github.com/near/borsh-go"
)

/**
 * https://github.com/golang/go/issues/29010
 */

func WriteUvarint(buffer *bytes.Buffer, value uint16) error {
	for value >= 0x80 {
		buffer.WriteByte(byte(value) | 0x80)
		value >>= 7
	}

	return buffer.WriteByte(byte(value))
}

func Serialize(value any) *bytes.Buffer {
	buffer := new(bytes.Buffer)

	serialize(reflect.ValueOf(value), buffer)

	return buffer
}

func serialize(value reflect.Value, buffer *bytes.Buffer) {
	switch value.Kind() {
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			serialize(value.Index(i), buffer)
		}
	case reflect.Struct:
		switch value.Interface().(type) {
		case CompactArray:
			WriteUvarint(buffer, uint16(value.Field(0).Uint()))
			serialize(value.Field(1), buffer)
		case metadata.Metadata:
			bytes, _ := borsh.Serialize(value.Interface())
			buffer.Write(bytes)
		default:
			for i := 0; i < value.NumField(); i++ {
				serialize(value.Field(i), buffer)
			}
		}
	default:
		binary.Write(buffer, binary.LittleEndian, value.Interface())
	}
}
