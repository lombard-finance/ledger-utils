package chainid

import (
	"testing"

	"github.com/lombard-finance/ledger-utils/common"
)

func TestStarknetIdentifier(t *testing.T) {
	common.EqualStrings(t, "SN_MAIN", NewStarknetMainnetLChainId().Identifier())
	common.EqualStrings(t, "SN_SEPOLIA", NewStarknetSepoliaLChainId().Identifier())
}
