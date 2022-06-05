package solago

import "sort"

type Account struct {
	Read    bool
	Write   bool
	Signer  bool
	Keypair Keypair
}

type AccountList []Account

func NewSignerAccount(keypair Keypair) Account {
	return Account{
		Read:    true,
		Write:   true,
		Signer:  true,
		Keypair: keypair,
	}
}

func NewReadOnlyAccount(keypair Keypair) Account {
	return Account{
		Read:    true,
		Write:   false,
		Signer:  false,
		Keypair: keypair,
	}
}

func (accounts *AccountCollection) MapToKeys(private bool) []Serializable {
	keys := []Serializable{}

	for _, account := range *accounts {
		if private {
			keys = append(keys, account.Keypair.PrivateKey)
		} else {
			keys = append(keys, account.Keypair.PublicKey)
		}
	}

	return keys
}

func (accounts AccountList) Sort() AccountList {
	sort.SliceStable(accounts, func(a, b int) bool {
		bothSigners := accounts[a].Signer && accounts[b].Signer
		neitherSigners := !accounts[a].Signer && !accounts[b].Signer

		if bothSigners || neitherSigners {
			return accounts[a].Write || !accounts[b].Write
		} else {
			return accounts[a].Signer
		}
	})

	return accounts
}

const SizeOfMintAccount = 82
const SizeOfMultisigAccount = 355

var (
	NilPublicKey                     = NewPublicKey("11111111111111111111111111111111")
	RentProgram                      = NewPublicKey("SysvarRent111111111111111111111111111111111")
	ConfigProgram                    = NewPublicKey("Config1111111111111111111111111111111111111")
	StakeProgram                     = NewPublicKey("Stake11111111111111111111111111111111111111")
	VoteProgram                      = NewPublicKey("Vote111111111111111111111111111111111111111")
	BPFLoaderProgram                 = NewPublicKey("BPFLoader1111111111111111111111111111111111")
	Secp256k1Program                 = NewPublicKey("KeccakSecp256k11111111111111111111111111111")
	TokenProgram                     = NewPublicKey("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")
	MemoProgram                      = NewPublicKey("MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr")
	SPLAssociatedTokenAccountProgram = NewPublicKey("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")
	SPLNameServiceProgram            = NewPublicKey("namesLPneVptA9Z5rqUDD9tMTWEJwofgaYwp8cawRkX")
	MetaplexTokenMetaProgram         = NewPublicKey("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")
	ComputeBudgetProgram             = NewPublicKey("ComputeBudget111111111111111111111111111111")
)
