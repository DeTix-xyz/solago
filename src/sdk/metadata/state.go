package metadata

import (
	"bytes"
	"crypto/ed25519"

	"github.com/near/borsh-go"
)

type Creator struct {
	Address  ed25519.PublicKey
	Verified bool
	Share    uint8 // In percentages
}

type Collection struct {
	Verified bool
	Key      ed25519.PublicKey
}

type UseMethod uint8

const (
	Burn UseMethod = iota
	Multiple
	Single
)

type Uses struct {
	UseMethod UseMethod
	Remaining uint64
	Total     uint64
}

type Metadata struct {
	Instruction          MetadataInstruction
	Name                 string
	Symbol               string
	URI                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator
	Collection           *Collection
	Uses                 *Uses
	IsMutable            bool
}

func (metadata *Metadata) Serialize(buffer *bytes.Buffer) *bytes.Buffer {
	bytes, _ := borsh.Serialize(metadata)

	buffer.Write(bytes)

	return buffer
}
