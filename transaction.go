package solago

import (
	"bytes"
	"encoding/binary"
)

type Transaction struct {
	Signatures CompactArray
	Message    Message
}

type InProcessTransaction struct {
	Buffer  *bytes.Buffer
	Message Message
	Client  Client
}

func NewTransaction(client Client, instructions ...InProcessInstruction) InProcessTransaction {
	return InProcessTransaction{
		Buffer:  new(bytes.Buffer),
		Message: NewMessage(client.GetRecentBlockhash(), instructions),
		Client:  client,
	}
}

func (transaction InProcessTransaction) Sign(accounts AccountCollection) InProcessTransaction {
	// Write private keys
	for _, privateKey := range accounts.MapToPrivateKeys() {
		binary.Write(transaction.Buffer, binary.LittleEndian, privateKey)
	}

	// Serialize message

	// allBytes := buffer.Bytes()
	// signatureCutoff := transaction.Signatures.Length*ed25519.PrivateKeySize + 1

	// signatures := allBytes[:signatureCutoff]
	// message := allBytes[signatureCutoff:]

	// for i, privateKey := range transaction.Signatures.Items {
	// 	start := i*ed25519.PrivateKeySize + 1
	// 	end := (i+1)*ed25519.PrivateKeySize + 1
	// 	signature := ed25519.Sign(privateKey.(ed25519.PrivateKey), message)

	// 	copy(signatures[start:end], signature)
	// }

	// return base64.StdEncoding.EncodeToString(allBytes)
	return transaction
}
