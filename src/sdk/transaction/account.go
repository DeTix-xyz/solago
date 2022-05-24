package transaction

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"

	"filippo.io/edwards25519"
	"github.com/mr-tron/base58"
)

type PublicKey ed25519.PublicKey
type PrivateKey ed25519.PrivateKey

type Keypair struct {
	PublicKey  PublicKey
	PrivateKey PrivateKey
}

const SizeOfMintAccount = 82
const SizeOfMultisigAccount = 355

var PDAMarker = []byte("ProgramDerivedAddress")
var NilPublicKey = NewPublicKey("11111111111111111111111111111111")

func NewPublicKey(key string) PublicKey {
	publicKey, _ := base58.Decode(key)
	return publicKey
}

func PublicKeyFromPrivateKey(private PrivateKey) PublicKey {
	return ed25519.PrivateKey(private).Public().(PublicKey)
}

func NewPrivateKey(key string) PrivateKey {
	privateKey, _ := base58.Decode(key)

	return privateKey
}

func NewKeypair() *Keypair {
	public, private, _ := ed25519.GenerateKey(rand.Reader)

	return &Keypair{PublicKey: PublicKey(public), PrivateKey: PrivateKey(private)}
}

func NewKeypairFromSeed(seed [32]byte) *Keypair {
	private := ed25519.NewKeyFromSeed(seed[:])

	return &Keypair{
		PrivateKey: PrivateKey(private),
		PublicKey:  PublicKeyFromPrivateKey(PrivateKey(private)),
	}
}

func NewKeypairFromFile(path string) *Keypair {
	bytes, _ := ioutil.ReadFile(path)

	var keypair Keypair
	json.Unmarshal(bytes, &keypair.PrivateKey)

	keypair.PublicKey = PublicKeyFromPrivateKey(keypair.PrivateKey)

	return &keypair
}

func CreateProgramAddress(seeds [][]byte, program PublicKey) (PublicKey, error) {
	buffer := bytes.Buffer{}

	for _, seed := range seeds {
		buffer.Write(seed)
	}

	buffer.Write(program)
	buffer.Write(PDAMarker)

	hash := sha256.Sum256(buffer.Bytes())

	if isOnCurve(hash[:]) {
		return nil, errors.New("invalid seeds; address must fall off the curve")
	}

	return PublicKey(hash[:]), nil
}

func isOnCurve(b []byte) bool {
	_, err := new(edwards25519.Point).SetBytes(b)
	isOnCurve := err == nil
	return isOnCurve
}

func FindProgramAddress(seed [][]byte, programID PublicKey) (PublicKey, error) {
	bumpSeed := uint8(math.MaxUint8)

	for bumpSeed != 0 {
		address, err := CreateProgramAddress(append(seed, []byte{byte(bumpSeed)}), programID)
		if err == nil {
			return address, nil
		}
		bumpSeed--
	}

	return nil, errors.New("could not find valid key for given seeds")
}
