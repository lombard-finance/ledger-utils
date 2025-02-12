package chainid_test

import (
	"strings"
	"testing"

	"github.com/lombard-finance/chain/chainid"
	"github.com/lombard-finance/chain/common"
)

func TestSuiIdentifier(t *testing.T) {
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
