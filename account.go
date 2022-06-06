package solago

import (
	"bytes"
	"encoding/binary"
	"sort"
)

type Account struct {
	Read    bool
	Write   bool
	Signer  bool
	Keypair Keypair
}

type AccountList []Account
type AccountIndexes []uint8

func (indexes AccountIndexes) Serialize(buffer *bytes.Buffer) {
	for _, index := range indexes {
		binary.Write(buffer, binary.LittleEndian, index)
	}
}

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

func NewSignerAccountFromSeed(seed [32]byte) Account {
	keypair := NewKeypairFromSeed(seed)

	return NewSignerAccount(keypair)
}

func SignerAccountFromFile(path string) Account {
	keypair := KeypairFromFile(path)

	return NewSignerAccount(keypair)
}

func (accounts AccountList) ToPublicKeys() PublicKeys {
	publicKeys := PublicKeys{}

	for _, account := range accounts {
		publicKeys = append(publicKeys, account.Keypair.PublicKey)
	}

	return publicKeys
}

func (accounts AccountList) ToPrivateKeys() PrivateKeys {
	privateKeys := PrivateKeys{}

	for _, account := range accounts {
		privateKeys = append(privateKeys, account.Keypair.PrivateKey)
	}

	return privateKeys
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

func (accounts AccountList) GetSigners() AccountList {
	signers := AccountList{}

	for _, account := range accounts {
		if account.Signer {
			signers = append(signers, account)
		}
	}

	return signers
}
