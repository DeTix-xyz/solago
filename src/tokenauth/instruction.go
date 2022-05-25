package tokenauth

type TokenAuthorityType uint8

const (
	MintTokens TokenAuthorityType = iota
	FreezeAccount
	AccountOwner
	CloseAccount
	TransferFeeConfig
	WithheldWithdraw
	CloseMint
)
