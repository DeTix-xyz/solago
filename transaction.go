package solago

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"

	"github.com/deezdegens/solago/rpc"
)

type Transaction struct {
	Buffer  *bytes.Buffer
	Message Message
	Client  rpc.Client
}

func NewTransaction(client rpc.Client, instructions ...PseudoInstruction) Transaction {
	return Transaction{
		Buffer:  new(bytes.Buffer),
		Message: NewMessage(RecentBlockhashFromString(client.GetRecentBlockhash()), instructions),
		Client:  client,
	}
}

func (transaction Transaction) SignAndSend(signers ...PrivateKey) string {
	// Write private keys for subsequent signature
	privateKeys := PrivateKeys(signers)
	privateKeys.Serialize(transaction.Buffer)

	// Serialize message
	transaction.Message.Serialize(transaction.Buffer)

	// Sign the message
	allBytes := transaction.Buffer.Bytes()
	signatureCutoff := len(signers)*ed25519.PrivateKeySize + 1

	signatures := allBytes[:signatureCutoff]
	message := allBytes[signatureCutoff:]

	for i, privateKey := range privateKeys {
		start := i*ed25519.PrivateKeySize + 1
		end := (i+1)*ed25519.PrivateKeySize + 1
		signature := ed25519.Sign(ed25519.PrivateKey(privateKey), message)

		copy(signatures[start:end], signature)
	}

	// Get the txn string
	transactionString := base64.StdEncoding.EncodeToString(allBytes)

	// Send the string to the cluster
	return transaction.Client.SendTransaction(transactionString)
}
