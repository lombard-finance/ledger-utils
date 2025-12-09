package chainid_test

import (
	"testing"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common"
)

func TestCosmosLChainId_NewCosmosLChainId(t *testing.T) {
	tests := []struct {
		name             string
		chainId          string
		expectedLChainId string
		err              error
	}{
		{
			"Lombard Ledger Mainnet",
			"ledger-mainnet-1",
			"0387b25e8e61f2ce4838b04795b231f09ee73ffd391da018bef4bc5c4975897b",
			nil,
		},
		{
			"Lombard Ledger Testnet",
			"ledger-testnet-1",
			"033bc7baf196ce32b8b9200518df11c35bad882fc6e3b6f45b4a8885f4c1281b",
			nil,
		},
		{
			"Osmosis",
			"osmosis-1",
			"038ebfb6519e8d814f1b8aee62da9a4e173f7e6898d60d962042421d18dbe4ef",
			nil,
		},
		{
			"Cosmos Hub",
			"cosmoshub-4",
			"03a232779a423721bfb80a99e86828034aa5726c469f770d39f29a0fb4710f9a",
			nil,
		},
		{
			"Cosmos Hub before snapshot",
			"cosmoshub-3",
			"03a232779a423721bfb80a99e86828034aa5726c469f770d39f29a0fb4710f9a",
			nil,
		},
		{
			"Cosmos Hub chain name",
			"cosmoshub",
			"03a232779a423721bfb80a99e86828034aa5726c469f770d39f29a0fb4710f9a",
			nil,
		},
		{
			"Chain Id with invalid counter",
			"cosmoshub-some",
			"",
			chainid.ErrInvalidCosmosChainId,
		},
		{
			"Chain Id with empty counter",
			"cosmoshub-",
			"",
			chainid.ErrInvalidCosmosChainId,
		},
		{
			"Empty chain Id",
			"",
			"",
			chainid.ErrEmptyCosmosChainId,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lchainId, err := chainid.NewCosmosLChainId(tt.chainId)
			if err != nil {
				if lchainId != nil {
					t.Error("both lchain id and err are not nil")
				}
				common.AssertError(t, err, tt.err)
			} else {
				equalEcosystem(t, chainid.EcosystemCosmos, lchainId.Ecosystem())
				common.EqualStrings(t, tt.expectedLChainId, lchainId.Hex())
			}
		})
	}
}

// These tests verify that LChainId concrete types are usable as map keys.
func TestCosmosLChainId_AsMapKey(t *testing.T) {
	// Use two distinct equal instances
	a := chainid.NewLombardLedgerLChainId()
	bPtr, err := chainid.NewLChainIdFromHex(a.String())
	common.AssertNoError(t, err)
	b := bPtr // both 'a' and 'b' are CosmosLChainId values

	m := map[chainid.LChainId]string{}
	m[a] = "ok"

	// same key instance
	common.EqualStrings(t, "ok", m[a])
	// distinct but equal value should still map to the same bucket if comparable by value
	// since the concrete types are value-comparable now, this should succeed
	common.EqualStrings(t, "ok", m[b])
}
