package sdk

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
)

type Transaction struct {
	Signatures *CompactArray[SerializablePrivateKey]
	Message    Message
}

func (transaction *Transaction) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	transaction.Signatures.Serialize(buffer)
	transaction.Message.Serialize(buffer)

	return buffer
}

func (transaction *Transaction) Sign(buffer *bytes.Buffer) string {
	allBytes := buffer.Bytes()
	signatureCutoff := transaction.Signatures.Length*ed25519.PrivateKeySize + 1

	signatures := allBytes[:signatureCutoff]
	message := allBytes[signatureCutoff:]

	for i, privateKey := range transaction.Signatures.Items {
		start := i * ed25519.PrivateKeySize
		end := (i + 1) * ed25519.PrivateKeySize

		copy(signatures[start+1:end+1], ed25519.Sign(ed25519.PrivateKey(privateKey), message))
	}

	return base64.StdEncoding.EncodeToString(allBytes)
}
