package chainid

import (
	"strings"
	"testing"
)

func TestSuiIdentifier(t *testing.T) {
	identifiers := []string{
		"35834a8a",
		"0x4c78adac",
	}

	for _, id := range identifiers {
		ch, _ := NewSuiChainId(id)
		equalStrings(t, strings.TrimPrefix(id, "0x"), ch.Identifier())
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
		ch, _ := NewChainIdFromHex(id.chainId)
		switch typesCh := ch.(type) {
		case SuiChainId:
			equalStrings(t, id.expected, typesCh.Identifier())
		default:
			t.FailNow()
		}
	}

}
