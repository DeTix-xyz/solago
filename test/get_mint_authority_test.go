package test

import (
	"fmt"
	"testing"

	"github.com/deezdegens/solago"
)

func TestGetMintAccount(t *testing.T) {
	client := solago.NewClient("https://api.devnet.solana.com")
	result := client.RPC.GetMintAccount("GkeGmiFDwnP74X73LcdfJMXd5RoRfy5NKWkwwohDRkwX")

	fmt.Println(result)
}
