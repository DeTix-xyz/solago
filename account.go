package solago

import "sort"

type Account struct {
	Read    bool
	Write   bool
	Signer  bool
	Keypair Keypair
}

type AccountList []Account

const SizeOfMintAccount = 82
const SizeOfMultisigAccount = 355

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

func NewReadWriteAccount(keypair Keypair) Account {
	return Account{
		Read:    true,
		Write:   true,
		Signer:  false,
		Keypair: keypair,
	}
}

func (accounts AccountList) ToPublicKeys() PublicKeys {
	publicKeys := PublicKeys{}

	for _, account := range accounts {
		publicKeys = append(publicKeys, account.Keypair.PublicKey)
	}

	return publicKeys
}

func (accounts AccountList) ToPrivateKeys() PrivateKeys {
	publicKeys := PrivateKeys{}

	for _, account := range accounts {
		publicKeys = append(publicKeys, account.Keypair.PrivateKey)
	}

	return publicKeys
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
