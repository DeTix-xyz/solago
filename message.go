package solago

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

func (header MessageHeader) Serialize(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.LittleEndian, header.NumberRequiredSignatures)
	binary.Write(buffer, binary.LittleEndian, header.NumberReadOnlySignedAccounts)
	binary.Write(buffer, binary.LittleEndian, header.NumberReadOnlyUnsignedAccounts)
}

func NewMessageHeaderFromAccounts(accounts AccountList) MessageHeader {
	numRequiredSignatures := uint8(0)
	numReadOnlySigned := uint8(0)
	numReadOnlyUnsigned := uint8(0)

	for _, account := range accounts {
		readOnly := account.Read && !account.Write

		if account.Signer {
			numRequiredSignatures += 1
			if readOnly {
				numReadOnlySigned += 1
			}
		} else if readOnly {
			numReadOnlyUnsigned += 1
		}
	}

	return MessageHeader{
		NumberRequiredSignatures:       numRequiredSignatures,
		NumberReadOnlySignedAccounts:   numReadOnlySigned,
		NumberReadOnlyUnsignedAccounts: numReadOnlyUnsigned,
	}
}

type Message struct {
	Header           MessageHeader
	AccountAddresses CompactArray
	RecentBlockhash  RecentBlockhash
	Instructions     CompactArray
}

func NewMessage(blockhash RecentBlockhash, accounts AccountList, instructions InstructionList) Message {
	return Message{
		RecentBlockhash:  blockhash,
		Header:           NewMessageHeaderFromAccounts(accounts),
		AccountAddresses: NewCompactArray(accounts.ToPublicKeys()),
		Instructions:     NewCompactArray(instructions),
	}
}

func (message Message) Serialize(buffer *bytes.Buffer) {
	message.Header.Serialize(buffer)
	message.AccountAddresses.Serialize(buffer)
	message.RecentBlockhash.Serialize(buffer)
	message.Instructions.Serialize(buffer)
}

type RecentBlockhash []byte

func RecentBlockhashFromString(hash string) RecentBlockhash {
	bytes, _ := base58.Decode(hash)

	return bytes
}

func (hash RecentBlockhash) Serialize(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.LittleEndian, hash)
}
