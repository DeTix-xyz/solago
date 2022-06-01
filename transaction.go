package solana

type Transaction struct {
	Signatures CompactArray[PrivateKey]
	Message    Message
}

func getSignaturesFromAccounts(accounts []Account) CompactArray[PrivateKey] {
	signatures := []PrivateKey{}

	for _, account := range accounts {
		if account.Signer {
			signatures = append(signatures, account.Keypair.PrivateKey)
		}
	}

	return CompactArray[PrivateKey]{uint16(len(signatures)), signatures}
}

func (transaction *Transaction) Sign(accounts []Account) string {

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

	return "TODO"
}
