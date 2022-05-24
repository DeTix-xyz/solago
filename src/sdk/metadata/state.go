package metadata

import "github.com/DeTix-xyz/solago/src/sdk/transaction"

type Creator struct {
	Address  transaction.PublicKey
	Verified bool
	Share    uint8 // In percentages
}

type Collection struct {
	Verified bool
	Key      transaction.PublicKey
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
