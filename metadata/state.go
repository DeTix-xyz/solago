package metadata

import "github.com/deezdegens/solago"

var MetaplexTokenMetaProgram = solago.NewPublicKey("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")

type Creator struct {
	Address  solago.PublicKey
	Verified bool
	Share    uint8 // In percentages
}

type Collection struct {
	Verified bool
	Key      solago.PublicKey
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
