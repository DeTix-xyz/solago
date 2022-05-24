package transaction

import "io"

/**
 * https://github.com/golang/go/issues/29010
 */

func WriteUvarint(w io.ByteWriter, x uint16) error {
	for x >= 0x80 {
		w.WriteByte(byte(x) | 0x80)
		x >>= 7
	}

	return w.WriteByte(byte(x))
}
