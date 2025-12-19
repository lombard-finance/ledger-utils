package chainid_test

import (
	"strings"
	"testing"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common"
)

func TestSuiLChainId_NewLChainIdFromHex(t *testing.T) {
	identifiers := []string{
		"35834a8a",
		"0x4c78adac",
	}

	for _, id := range identifiers {
		ch, _ := chainid.NewSuiLChainId(id)
		common.EqualStrings(t, strings.TrimPrefix(id, "0x"), ch.Identifier())
	}

	hexChainIds := []struct {
		chainId  string
		expected string
	}{
		{
			"0x0100000000000000000000000000000000000000000000000000000035834a8a",
			"35834a8a",
		},
		{
			"0x010000000000000000000000000000000000000000000000000000004c78adac",
			"4c78adac",
		},
	}

	for _, id := range hexChainIds {
		ch, _ := chainid.NewLChainIdFromHex(id.chainId)
		switch typesCh := ch.(type) {
		case chainid.SuiLChainId:
			common.EqualStrings(t, id.expected, typesCh.Identifier())
		default:
			t.FailNow()
		}
	}

}

// These tests verify that LChainId concrete types are usable as map keys.
func TestSuiLChainId_AsMapKey(t *testing.T) {
	// Use two distinct equal instances
	a := chainid.NewSuiMainnetLChainId()
	b, err := chainid.NewSuiLChainId(a.Identifier())
	common.AssertNoError(t, err)
	c, err := chainid.NewLChainIdFromHex(a.String())
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
