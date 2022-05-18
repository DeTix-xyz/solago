package sdk

import (
	"bytes"
	"encoding/binary"

	"github.com/mr-tron/base58"
)

type MessageHeader struct {
	NumberRequiredSignatures       SerializableUInt8
	NumberReadOnlySignedAccounts   SerializableUInt8
	NumberReadOnlyUnsignedAccounts SerializableUInt8
}

func NewMessageHeader(requiredSignatures, readOnlySigned, readOnlyUnsigned uint8) MessageHeader {
	return MessageHeader{
		NumberRequiredSignatures:       SerializableUInt8(requiredSignatures),
		NumberReadOnlySignedAccounts:   SerializableUInt8(readOnlySigned),
		NumberReadOnlyUnsignedAccounts: SerializableUInt8(readOnlyUnsigned),
	}
}

func (header *MessageHeader) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	header.NumberRequiredSignatures.Serialize(buffer)
	header.NumberReadOnlySignedAccounts.Serialize(buffer)
	header.NumberReadOnlyUnsignedAccounts.Serialize(buffer)

	return buffer
}

type Message struct {
	Header           MessageHeader
	AccountAddresses *CompactArray[SerializablePublicKey]
	RecentBlockhash  RecentBlockhash
	Instructions     *CompactArray[*Instruction]
}

func (message *Message) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	message.Header.Serialize(buffer)
	message.AccountAddresses.Serialize(buffer)
	message.RecentBlockhash.Serialize(buffer)
	message.Instructions.Serialize(buffer)

	return buffer
}

type RecentBlockhash SerializablePublicKey

func RecentBlockhashFromString(hash string) RecentBlockhash {
	bytes, _ := base58.Decode(hash)

	return bytes
}

func (blockhash RecentBlockhash) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	binary.Write(buffer, binary.LittleEndian, blockhash)

	return buffer
}
