package chainid_test

import (
	"testing"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common"
)

func TestSolanaLChainId_NewLChainIdFromHex(t *testing.T) {
	hexChainIds := []string{
		"0x02296998a6f8e2a784db5d9f95e18fc23f70441a1039446801089879b08c7ef0",
		"0x0259db5080fc2c6d3bcf7ca90712d3c2e5e6c28f27f0dfbb9953bdb0894c03ab",
	}

	for _, id := range hexChainIds {
		ch, _ := chainid.NewLChainIdFromHex(id)
		switch ch.(type) {
		case chainid.SolanaLChainId:
		default:
			t.FailNow()
		}
	}
}

// These tests verify that LChainId concrete types are usable as map keys.
func TestSolanaLChainId_AsMapKey(t *testing.T) {
	// Use two distinct equal instances
	a := chainid.NewSolanaMainnetLChainId()
	b, err := chainid.NewLChainIdFromHex(a.String())
	common.AssertNoError(t, err)
	c, err := chainid.NewSolanaLChainId("5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d")
	common.AssertNoError(t, err)

	m := map[chainid.LChainId]string{}
	m[a] = "ok"

	// same key instance
	common.EqualStrings(t, "ok", m[a])
	// distinct but equal value should still map to the same bucket if comparable by value
	// since the concrete types are value-comparable now, this should succeed
	common.EqualStrings(t, "ok", m[b])
	common.EqualStrings(t, "ok", m[c])
}
