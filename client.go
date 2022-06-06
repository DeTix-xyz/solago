package solago

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"

	"github.com/deezdegens/solago/rpc"
)

type Client struct {
	rpc *rpc.Client
}

func NewClient(endpoint string) Client {
	return Client{
		rpc: rpc.NewClient("https://api.devnet.solana.com", nil),
	}
}

func (client Client) SendTransaction(pseudoInstructions ...PseudoInstruction) string {
	// Create the buffer
	buffer := new(bytes.Buffer)

	// Get the accounts
	accounts := PseudoInstructionList(pseudoInstructions).CollectAccounts()

	// Create real instructions from pseudos
	instructions := PseudoInstructionList(pseudoInstructions).NewInstructionList(accounts)

	// Formulate the message
	message := NewMessage(
		RecentBlockhashFromString("AM9JCV5XMnB1t2Zv8YfRxCCpVv6k898Qf8S2K4dCctbp"),
		accounts,
		instructions,
	)

	// Get signers and write their private keys for eventual signature
	signers := accounts.GetSigners()
	privateKeys := signers.ToPrivateKeys()
	privateKeys.Serialize(buffer)

	// Serialize message
	message.Serialize(buffer)

	// Sign the message
	allBytes := buffer.Bytes()
	signatureCutoff := len(signers)*ed25519.PrivateKeySize + 1

	signatureBytes := allBytes[:signatureCutoff]
	messageBytes := allBytes[signatureCutoff:]

	for i, privateKey := range privateKeys {
		start := i*ed25519.PrivateKeySize + 1
		end := (i+1)*ed25519.PrivateKeySize + 1
		signature := ed25519.Sign(ed25519.PrivateKey(privateKey), messageBytes)

		copy(signatureBytes[start:end], signature)
	}

	// Get the txn string
	transactionString := base64.StdEncoding.EncodeToString(allBytes)

	// Send the string to the cluster
	return client.rpc.SendTransaction(transactionString)
}
