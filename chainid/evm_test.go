package chainid_test

import (
	"testing"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common"
)

func TestEVMLChainId_NewLChainIdFromHex(t *testing.T) {
	hexChainIds := []string{
		"0x0000000000000000000000000000000000000000000000000000000000000001", // Ethereum Mainnet
		"0x0000000000000000000000000000000000000000000000000000000000aa36a7", // Ethereum Sepolia
		"0x0000000000000000000000000000000000000000000000000000000000004268", // Ethereum Holesky
		"0x0000000000000000000000000000000000000000000000000000000000002105", // Base
		"0x0000000000000000000000000000000000000000000000000000000000014a34", // Base Sepolia
		"0x0000000000000000000000000000000000000000000000000000000000000038", // BSC
		"0x0000000000000000000000000000000000000000000000000000000000000061", // BSC Testnet
		"0x0000000000000000000000000000000000000000000000000000000000000092", // Sonic
		"0x000000000000000000000000000000000000000000000000000000000000dede", // Sonic Blaze Testnet
		"0x000000000000000000000000000000000000000000000000000000000000def1", // Ink
		"0x00000000000000000000000000000000000000000000000000000000000ba5ed", // Ink Sepolia
		"0x00000000000000000000000000000000000000000000000000000000000b67d2", // Katana
		"0x000000000000000000000000000000000000000000000000000000000001f977", // Katana Tatara Testnet
	}

	for _, id := range hexChainIds {
		ch, _ := chainid.NewLChainIdFromHex(id)
		switch ch.(type) {
		case chainid.EVMLChainId:
		default:
			t.FailNow()
		}
	}
}

func TestEVMLChainId_AsMapKey(t *testing.T) {
	a := chainid.NewEVMEthereumLChainId()
	b, err := chainid.NewLChainIdFromHex(a.String())
	common.AssertNoError(t, err)

	// Ensure b is an EVM concrete type too
	if _, ok := b.(chainid.EVMLChainId); !ok {
		t.Fatalf("expected EVM LChainId, got %T", b)
	}

	m := map[chainid.LChainId]int{}
	m[a] = 42
	if m[a] != 42 {
		t.Fatalf("expected 42, got %d", m[a])
	}
	if m[b] != 42 {
		t.Fatalf("expected retrieval by equal key to return 42, got %d", m[b])
	}
}
