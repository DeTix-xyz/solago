package transaction

import (
	"bytes"
	"encoding/binary"

	"github.com/mr-tron/base58"
)

type MessageHeader struct {
	NumberRequiredSignatures       uint8
	NumberReadOnlySignedAccounts   uint8
	NumberReadOnlyUnsignedAccounts uint8
}

func NewMessageHeader(requiredSignatures, readOnlySigned, readOnlyUnsigned uint8) MessageHeader {
	return MessageHeader{
		NumberRequiredSignatures:       requiredSignatures,
		NumberReadOnlySignedAccounts:   readOnlySigned,
		NumberReadOnlyUnsignedAccounts: readOnlyUnsigned,
	}
}

type Message struct {
	Header           MessageHeader
	AccountAddresses *CompactArray
	RecentBlockhash  RecentBlockhash
	Instructions     *CompactArray
}

type RecentBlockhash []byte

func RecentBlockhashFromString(hash string) RecentBlockhash {
	bytes, _ := base58.Decode(hash)

	return bytes
}

func (blockhash RecentBlockhash) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	binary.Write(buffer, binary.LittleEndian, blockhash)

	return buffer
}
