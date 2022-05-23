package metadata

import "crypto/ed25519"

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
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator
	Collection           *Collection
	Uses                 *Uses
	IsMutable            bool
}
