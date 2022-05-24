package transaction

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
)

type Transaction struct {
	Signatures *CompactArray
	Message    Message
}

func (transaction *Transaction) Sign(buffer *bytes.Buffer) string {
	allBytes := buffer.Bytes()
	signatureCutoff := transaction.Signatures.Length*ed25519.PrivateKeySize + 1

	signatures := allBytes[:signatureCutoff]
	message := allBytes[signatureCutoff:]

	for i, privateKey := range transaction.Signatures.Items {
		start := i*ed25519.PrivateKeySize + 1
		end := (i+1)*ed25519.PrivateKeySize + 1

		copy(signatures[start:end], ed25519.Sign(privateKey.(ed25519.PrivateKey), message))
	}

	return base64.StdEncoding.EncodeToString(allBytes)
}
