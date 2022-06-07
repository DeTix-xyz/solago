package solago

import "bytes"

type Transaction struct {
	Signatures CompactArray
	Message    Message
}

func (transaction Transaction) Serialize(buffer *bytes.Buffer) {
	transaction.Signatures.Serialize(buffer)
	transaction.Message.Serialize(buffer)
}
