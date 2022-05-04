package sdk

import (
	"bytes"
	"crypto/ed25519"
	"encoding/binary"
	"io"
)

/**
 * Maximum over-the-wire size of a Transaction
 *
 * 1280 is IPv6 minimum MTU
 * 40 bytes is the size of the IPv6 header
 * 8 bytes is the size of the fragment header
 */

const PACKET_DATA_SIZE = 1280 - 40 - 8

type Serializable interface {
	Serialize(buffer *bytes.Buffer) *bytes.Buffer
}
type SerializableUInt8 uint8
type SerializablePublicKey ed25519.PublicKey
type SerializablePrivateKey ed25519.PrivateKey

func (sui8 SerializableUInt8) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	binary.Write(buffer, binary.LittleEndian, sui8)

	return buffer
}

func (spk SerializablePrivateKey) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	binary.Write(buffer, binary.LittleEndian, spk)

	return buffer
}

func (spk SerializablePublicKey) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	binary.Write(buffer, binary.LittleEndian, spk)

	return buffer
}

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
