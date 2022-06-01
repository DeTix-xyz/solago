package solago

import (
	"github.com/mr-tron/base58"
)

type MessageHeader struct {
	NumberRequiredSignatures       uint8
	NumberReadOnlySignedAccounts   uint8
	NumberReadOnlyUnsignedAccounts uint8
}

func NewMessageHeaderFromAccounts(accounts []Account) MessageHeader {
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
	AccountAddresses CompactArray[PublicKey]
	RecentBlockhash  RecentBlockhash
	Instructions     CompactArray[Instruction]
}

func NewMessage(blockhash RecentBlockhash, instructions ...InProcessInstruction) Message {
	// Gather all accounts
	return Message{
		RecentBlockhash: blockhash,
	}
}

type RecentBlockhash []byte

func RecentBlockhashFromString(hash string) RecentBlockhash {
	bytes, _ := base58.Decode(hash)

	return bytes
}
