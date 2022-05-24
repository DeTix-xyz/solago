package transaction

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
)

type Transaction struct {
	PrivateKeys CompactArray
	Message     Message
}

func (transaction *Transaction) Sign(buffer *bytes.Buffer) string {
	allBytes := buffer.Bytes()
	signatureCutoff := transaction.PrivateKeys.Length*ed25519.PrivateKeySize + 1

	signatures := allBytes[:signatureCutoff]
	message := allBytes[signatureCutoff:]

	for i, privateKey := range transaction.PrivateKeys.Items {
		start := i*ed25519.PrivateKeySize + 1
		end := (i+1)*ed25519.PrivateKeySize + 1
		signature := ed25519.Sign(privateKey.(ed25519.PrivateKey), message)

		copy(signatures[start:end], signature)
	}

	return base64.StdEncoding.EncodeToString(allBytes)
}
