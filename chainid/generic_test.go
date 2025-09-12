package chainid_test

import (
	"bytes"
	"testing"

	"github.com/lombard-finance/ledger-utils/chainid"
)

func TestGenericLChainIdMarshalUnmarshal(t *testing.T) {
	genericHex := "0x1100000000000000000000000000000000000000000000000000000000000001"
	cid, err := chainid.NewLChainIdFromHex(genericHex)
	if err != nil {
		t.Fatalf("unexpected error creating generic chain id: %v", err)
	}
	gVal, ok := cid.(chainid.GenericLChainId)
	if !ok {
		t.Fatalf("expected GenericLChainId, got %T", cid)
	}
	g := &gVal

	if g.String() != genericHex {
		t.Fatalf("string mismatch: %s", g.String())
	}
	if g.Hex() != genericHex[2:] { // strip 0x
		t.Fatalf("hex mismatch: %s", g.Hex())
	}
	if g.Ecosystem() != chainid.Ecosystem(0x11) {
		t.Fatalf("ecosystem mismatch: %d", g.Ecosystem())
	}

	// Marshal
	bytes1, err := g.Marshal()
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	if len(bytes1) != chainid.ChainIdLength {
		t.Fatalf("marshal length mismatch: %d", len(bytes1))
	}

	// Size
	if g.Size() != chainid.ChainIdLength {
		t.Fatalf("size mismatch: %d", g.Size())
	}

	// Unmarshal round trip
	var g2 chainid.GenericLChainId
	if err := g2.Unmarshal(bytes1); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if g2.String() != g.String() {
		t.Fatalf("round trip mismatch: %s vs %s", g2.String(), g.String())
	}
	if !g.Equal(&g2) { // Equal is on embedded lChainId via interface methods
		t.Fatalf("expected equality after round trip")
	}
}

func TestGenericLChainIdUnmarshalErrors(t *testing.T) {
	var g chainid.GenericLChainId
	// too short
	short := make([]byte, chainid.ChainIdLength-1)
	if err := g.Unmarshal(short); err == nil {
		t.Fatalf("expected error for short length")
	}
	// too long
	long := make([]byte, chainid.ChainIdLength+1)
	if err := g.Unmarshal(long); err == nil {
		t.Fatalf("expected error for long length")
	}
}

func TestGenericLChainIdToEcosystemUnsupported(t *testing.T) {
	genericHex := "0x1100000000000000000000000000000000000000000000000000000000000001"
	cid, err := chainid.NewLChainIdFromHex(genericHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gVal, ok := cid.(chainid.GenericLChainId)
	if !ok {
		t.Fatalf("expected GenericLChainId, got %T", cid)
	}
	if _, err := (&gVal).ToEcosystem(); err == nil {
		t.Fatalf("expected unsupported ecosystem error")
	}
}

func TestGenericLChainIdMarshalTo(t *testing.T) {
	genericHex := "0x1100000000000000000000000000000000000000000000000000000000000001"
	cid, err := chainid.NewLChainIdFromHex(genericHex)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	gVal, ok := cid.(chainid.GenericLChainId)
	if !ok {
		t.Fatalf("expected GenericLChainId, got %T", cid)
	}
	g := &gVal

	buf := make([]byte, chainid.ChainIdLength)
	n, err := g.MarshalTo(buf)
	if err != nil {
		t.Fatalf("MarshalTo error: %v", err)
	}
	if n != chainid.ChainIdLength {
		t.Fatalf("written length mismatch: %d", n)
	}
	if !bytes.Equal(buf, g.Bytes()) {
		t.Fatalf("buffer content mismatch")
	}

	// small buffer error
	small := make([]byte, chainid.ChainIdLength-1)
	if _, err := g.MarshalTo(small); err == nil {
		t.Fatalf("expected error on small buffer")
	}
}
