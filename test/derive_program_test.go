package test

import (
	"fmt"
	"testing"

	"github.com/DeTix-xyz/solago/src/sdk"
)

func TestDeriveProgram(t *testing.T) {
	key, _ := sdk.FindProgramAddress(
		[][]byte{ // seeds
			[]byte("metadata"),
			sdk.MetaplexTokenMetaProgram,
			sdk.PublicKey("DwnN7Yk3i4sw4wt9VhScyYmL4EhxAYo5rqu6c4qSKxMk"),
		},
		sdk.MetaplexTokenMetaProgram, // program
	)

	fmt.Println(key)
}
