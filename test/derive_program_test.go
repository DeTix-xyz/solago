package test

import (
	"fmt"
	"testing"

	"github.com/deezdegens/solago"
	"github.com/deezdegens/solago/metadata"
)

func TestDeriveProgram(t *testing.T) {
	key, _ := solago.FindProgramAddress(
		[][]byte{ // seeds
			[]byte("metadata"),
			metadata.MetaplexTokenMetaProgram,
			solago.PublicKey("DwnN7Yk3i4sw4wt9VhScyYmL4EhxAYo5rqu6c4qSKxMk"),
		},
		metadata.MetaplexTokenMetaProgram, // program
	)

	fmt.Println(key)
}
