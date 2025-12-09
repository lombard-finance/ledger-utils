package chainid_test

import (
	"testing"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common"
)

func TestStarknetLChainId_Predefined(t *testing.T) {
    common.EqualStrings(t, "SN_MAIN", chainid.NewStarknetMainnetLChainId().Identifier())
    common.EqualStrings(t, "SN_SEPOLIA", chainid.NewStarknetSepoliaLChainId().Identifier())
}

// These tests verify that LChainId concrete types are usable as map keys.
func TestStarknetLChainId_AsMapKey(t *testing.T) {
    // Use two distinct equal instances
    a := chainid.NewStarknetMainnetLChainId()
    bPtr, err := chainid.NewStarknetLChainIdFromName(a.Identifier())
    common.AssertNoError(t, err)
    b := *bPtr // both 'a' and 'b' are StarknetLChainId values

    m := map[chainid.LChainId]string{}
    m[a] = "ok"

    // same key instance
    common.EqualStrings(t, "ok", m[a])
    // distinct but equal value should still map to the same bucket if comparable by value
    // since the concrete types are value-comparable now, this should succeed
    common.EqualStrings(t, "ok", m[b])
}
