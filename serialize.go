package solago

import (
	"bytes"
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
