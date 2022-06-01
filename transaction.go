package solago

import "bytes"

type Transaction struct {
	Signatures CompactArray[PrivateKey]
	Message    Message
}

type InProcessTransaction struct {
	Buffer  *bytes.Buffer
	Message Message
	Client  Client
}

func NewTransaction(client Client, instructions ...InProcessInstruction) InProcessInstruction {
	return InProcessTransaction{
		Buffer:  new(bytes.Buffer),
		Message: NewMessage(client.GetRecentBlockhash(), instructions),
		Client:  client,
	}
}

func (transaction InProcessTransaction) Sign(accounts ...Account) InProcessTransaction {

	//
	// NEEDS to compute message header
	//

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

	return nil
}
