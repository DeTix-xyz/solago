package test

import (
	"fmt"
	"testing"

	"github.com/DeTix-xyz/solago/src/sdk/transaction"
)

func TestSerializeBool(t *testing.T) {
	buffer := transaction.Serialize(false)
	fmt.Println(buffer.Bytes())
}
