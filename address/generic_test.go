package address_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/lombard-finance/ledger-utils/address"
	"github.com/lombard-finance/ledger-utils/chainid"
)

func TestNewGenericAddress(t *testing.T) {
	// empty slice error
	_, err := address.NewGenericAddress([]byte{}, chainid.EcosystemEVM)
	if err == nil || err != address.ErrEmptyAddress {
		t.Fatalf("expected ErrEmptyAddress, got %v", err)
	}

	payload := []byte{0x01, 0x02, 0x03}
	g, err := address.NewGenericAddress(payload, chainid.EcosystemCosmos)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// ensure copy semantics
	payload[0] = 0xFF
	if g.Hex() != "010203" {
		t.Fatalf("expected internal bytes unchanged, got %s", g.Hex())
	}

	if g.Length() != 3 {
		t.Fatalf("length mismatch: %d", g.Length())
	}
	if g.Ecosystem() != chainid.EcosystemCosmos {
		t.Fatalf("ecosystem mismatch")
	}
}

func TestGenericAddressStringAndHex(t *testing.T) {
	g, _ := address.NewGenericAddress([]byte{0xAB, 0xCD, 0xEF}, chainid.EcosystemEVM)
	if g.Hex() != "abcdef" {
		t.Fatalf("hex mismatch: %s", g.Hex())
	}
	if g.String() != "0xabcdef" {
		t.Fatalf("string mismatch: %s", g.String())
	}
}

func TestGenericAddressBytesIsolation(t *testing.T) {
	g, _ := address.NewGenericAddress([]byte{0x10, 0x20}, chainid.EcosystemEVM)
	b := g.Bytes()
	if !bytes.Equal(b, []byte{0x10, 0x20}) {
		t.Fatalf("bytes mismatch")
	}
	b[0] = 0xFF
	b2 := g.Bytes()
	if b2[0] != 0x10 {
		t.Fatalf("mutation of returned slice affected internal state")
	}
}

func TestGenericAddressEqual(t *testing.T) {
	g1, _ := address.NewGenericAddress([]byte{0x01, 0x02}, chainid.EcosystemCosmos)
	g2, _ := address.NewGenericAddress([]byte{0x01, 0x02}, chainid.EcosystemCosmos)
	g3, _ := address.NewGenericAddress([]byte{0x01, 0x03}, chainid.EcosystemCosmos)
	g4, _ := address.NewGenericAddress([]byte{0x01, 0x02}, chainid.EcosystemEVM)

	if !g1.Equal(g2) {
		t.Fatalf("expected equality")
	}
	if g1.Equal(g3) {
		t.Fatalf("different bytes should not be equal")
	}
	if g1.Equal(g4) {
		t.Fatalf("different ecosystems should not be equal")
	}
}

func TestGenericAddressMarshalUnmarshal(t *testing.T) {
	payload := []byte{0xDE, 0xAD, 0xBE, 0xEF}
	g, _ := address.NewGenericAddress(payload, chainid.EcosystemSui)
	ser, err := g.Marshal()
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}
	if !bytes.Equal(ser, payload) {
		t.Fatalf("marshal bytes mismatch")
	}

	var g2 address.GenericAddress
	if err := g2.Unmarshal(ser); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if g2.Ecosystem() != chainid.EcosystemUnknown {
		t.Fatalf("expected unknown ecosystem after unmarshal")
	}
	if g2.Hex() != hex.EncodeToString(payload) {
		t.Fatalf("payload mismatch after unmarshal")
	}

	var g3 address.GenericAddress
	if err := g3.Unmarshal([]byte{}); err == nil {
		t.Fatalf("expected error on empty data")
	}
}

func TestGenericAddressFromHex(t *testing.T) {
	g1, err := address.NewGenericAddressFromHex("0x0102", chainid.EcosystemEVM)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g1.Hex() != "0102" || g1.String() != "0x0102" {
		t.Fatalf("hex/string mismatch")
	}

	g2, err := address.NewGenericAddressFromHex("0102", chainid.EcosystemEVM)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !g1.Equal(g2) {
		t.Fatalf("expected equality for with/without 0x")
	}

	if _, err := address.NewGenericAddressFromHex("0xZZ", chainid.EcosystemEVM); err == nil {
		t.Fatalf("expected hex decode error")
	}
}

func TestGenericAddressToEcosystem(t *testing.T) {
	bytes20 := make([]byte, 20)
	for i := range bytes20 {
		bytes20[i] = byte(i)
	}
	g, _ := address.NewGenericAddress(bytes20, chainid.EcosystemCosmos)
	converted, err := g.ToEcosystem(chainid.EcosystemEVM)
	if err != nil {
		t.Fatalf("ToEcosystem error: %v", err)
	}
	if converted.Ecosystem() != chainid.EcosystemEVM {
		t.Fatalf("ecosystem mismatch after conversion")
	}
	if !bytes.Equal(converted.Bytes(), bytes20) {
		t.Fatalf("bytes mismatch after conversion")
	}
}

func TestGenericAddressMarshalTo(t *testing.T) {
	payload := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	g, err := address.NewGenericAddress(payload, chainid.EcosystemEVM)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	buf := make([]byte, len(payload))
	n, err := g.MarshalTo(buf)
	if err != nil {
		t.Fatalf("MarshalTo error: %v", err)
	}
	if n != len(payload) {
		t.Fatalf("written length mismatch: %d", n)
	}
	if !bytes.Equal(buf, payload) {
		t.Fatalf("buffer content mismatch")
	}

	// buffer too small
	small := make([]byte, len(payload)-1)
	if _, err := g.MarshalTo(small); err == nil {
		t.Fatalf("expected error on small buffer")
	}
}

func TestNewGenericAddressFromAddress(t *testing.T) {
	baseHex := "0x8236a87084f8B84306f72007F36F2618A5634494"
	base, err := address.NewEvmAddressFromHex(baseHex)
	if err != nil {
		t.Fatalf("unexpected error creating base address: %v", err)
	}

	wrapped := address.NewGenericAddressFromAddress(base)
	if wrapped.Ecosystem() != base.Ecosystem() {
		t.Fatalf("ecosystem mismatch: %d vs %d", wrapped.Ecosystem(), base.Ecosystem())
	}
	if !wrapped.Equal(base) || !base.Equal(wrapped) {
		t.Fatalf("expected equality between wrapped generic and base address")
	}

	// copy semantics: mutate a copy of base bytes and ensure wrapped unchanged
	b := base.Bytes()
	b[0] ^= 0xFF
	if !wrapped.Equal(base) {
		t.Fatalf("mutation of external slice should not affect wrapped")
	}
}
