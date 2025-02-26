package chainid_test

import (
	"testing"

	"github.com/lombard-finance/ledger-utils/chainid"
)

func TestSolanaIdentifier(t *testing.T) {
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
