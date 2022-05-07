package sdk

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"io/ioutil"

	"github.com/mr-tron/base58"
)

type Keypair struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

var NilPublicKey = PublicKey("11111111111111111111111111111111")

func PublicKey(key string) ed25519.PublicKey {
	publicKey, _ := base58.Decode(key)
	return publicKey
}

func PublicKeyFromPrivateKey(private ed25519.PrivateKey) ed25519.PublicKey {
	return private.Public().(ed25519.PublicKey)
}

func PrivateKey(key string) ed25519.PrivateKey {
	privateKey, _ := base58.Decode(key)

	return privateKey
}

func NewKeypair() *Keypair {
	public, private, _ := ed25519.GenerateKey(rand.Reader)

	return &Keypair{PublicKey: public, PrivateKey: private}
}

func NewKeypairFromSeed(seed [32]byte) *Keypair {
	private := ed25519.NewKeyFromSeed(seed[:])

	return &Keypair{PrivateKey: private, PublicKey: PublicKeyFromPrivateKey(private)}
}

func NewKeypairFromFile(path string) *Keypair {
	bytes, _ := ioutil.ReadFile(path)

	var keypair Keypair
	json.Unmarshal(bytes, &keypair.PrivateKey)

	keypair.PublicKey = PublicKeyFromPrivateKey(keypair.PrivateKey)

	return &keypair
}
