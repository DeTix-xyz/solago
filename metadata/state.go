package metadata

import "github.com/deezdegens/solago"

var Program = solago.NewPublicKey("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")
var Account = solago.NewReadOnlyAccount(solago.Keypair{PublicKey: Program})

// ---------------------------------------------------------------
// ON CHAIN METADATA
// ---------------------------------------------------------------

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
	Name                 string
	Symbol               string
	URI                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator
	Collection           *Collection
	Uses                 *Uses
	IsMutable            bool
}

// ---------------------------------------------------------------
// OFF CHAIN METADATA
// ---------------------------------------------------------------

type File struct {
	URI  string `json:"uri"`
	Type string `json:"type"`
}

type Properties struct {
	Category string `json:"category"`
	Files    []File `json:"files"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type MetadataOffChain struct {
	Image       string      `json:"image"`
	Description string      `json:"description"`
	Properties  Properties  `json:"properties"`
	Attributes  []Attribute `json:"attributes"`
}
