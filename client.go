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

	// Formulate the message
	transaction := Transaction{
		Signatures: NewCompactArray(accounts.GetSigners().ToPrivateKeys()),
		Message: NewMessage(
			RecentBlockhashFromString(client.rpc.GetRecentBlockhash()),
			accounts,
			PseudoInstructionList(pseudoInstructions).NewInstructionList(accounts),
		),
	}

	// Serialize the transaction
	transaction.Serialize(buffer)

	// Sign the message
	allBytes := buffer.Bytes()
	signatureCutoff := transaction.Signatures.Length*ed25519.PrivateKeySize + 1 // plus one byte for # of sigs

	signatureBytes := allBytes[:signatureCutoff]
	messageBytes := allBytes[signatureCutoff:]

	for i, privateKey := range transaction.Signatures.Items.(PrivateKeys) {
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
