package chainid

import (
	"bytes"
	"encoding/hex"
	"errors"
	"strings"
	"testing"
)

func TestChainId(t *testing.T) {

	tests := []struct {
		name              string
		hexChainId        string
		ecosystem         Ecosystem
		predefinedFactory func() ChainId
	}{
		{
			"Ethereum",
			"0x0000000000000000000000000000000000000000000000000000000000000001",
			EcosystemEVM,
			func() ChainId { return NewEVMEthereumChainId() },
		},
		{
			"Ethereum Sepolia",
			"0x0000000000000000000000000000000000000000000000000000000000aa36a7",
			EcosystemEVM,
			func() ChainId { return NewEVMSepoliaChainId() },
		},
		{
			"Ethereum Holesky",
			"0x0000000000000000000000000000000000000000000000000000000000004268",
			EcosystemEVM,
			func() ChainId { return NewEVMHoleskyChainId() },
		},
		{
			"Base",
			"0x0000000000000000000000000000000000000000000000000000000000002105",
			EcosystemEVM,
			func() ChainId { return NewEVMBaseChainId() },
		},
		{
			"Base Sepolia",
			"0x0000000000000000000000000000000000000000000000000000000000014a34",
			EcosystemEVM,
			func() ChainId { return NewEVMBaseSepoliaChainId() },
		},
		{
			"BSC",
			"0x0000000000000000000000000000000000000000000000000000000000000038",
			EcosystemEVM,
			func() ChainId { return NewEVMBinanceSmartChainChainId() },
		},
		{
			"Sui",
			"0x0100000000000000000000000000000000000000000000000000000035834a8a",
			EcosystemSui,
			func() ChainId { return NewSuiMainnetChainId() },
		},
		{
			"Sui Testnet",
			"0x010000000000000000000000000000000000000000000000000000004c78adac",
			EcosystemSui,
			func() ChainId { return NewSuiTestnetChainId() },
		},
		{
			"Bitcoin",
			"0xff0000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
			EcosystemBitcoin,
			func() ChainId { return NewBitcoinChainId() },
		},
		{
			"Bitcoin Signet",
			"0xff000008819873e925422c1ff0f99f7cc9bbb232af63a077a480a3633bee1ef6",
			EcosystemBitcoin,
			func() ChainId { return NewBitcoinSignetChainId() },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chainId, err := NewChainIdFromHex(test.hexChainId)
			assertNoError(t, err)
			equalStrings(t, test.hexChainId, chainId.String())
			equalStrings(t, test.hexChainId, "0x"+chainId.Hex())
			equalEcosystem(t, test.ecosystem, chainId.Ecosystem())
			equalChainId(t, test.predefinedFactory(), chainId)
			referenceChainIdBytes, err := hex.DecodeString(strings.TrimPrefix(test.hexChainId, "0x"))
			assertNoError(t, err)
			equalBytes(t, referenceChainIdBytes, chainId.Bytes())
			fixed := chainId.FixedBytes()
			equalBytes(t, referenceChainIdBytes, fixed[:])
		})
	}
}

func TestChainIdErrorConditions(t *testing.T) {
	correctHex := "0x0000000000000000000000000000000000000000000000000000000000000001"
	longerHex := correctHex + "01"
	_, err := NewChainIdFromHex(longerHex)
	assertError(t, err, ErrChainIdInvalid, ErrLength)
	longerBytes, _ := hex.DecodeString(longerHex)
	_, err = NewChainId(longerBytes)
	assertError(t, err, ErrChainIdInvalid, ErrLength)
	unsupportedEcosystemChainId := "0x11" + correctHex[4:]
	_, err = NewChainIdFromHex(unsupportedEcosystemChainId)
	assertError(t, err, ErrChainIdInvalid, ErrUnsupportedEcosystem)
	badCharChainId := "0xY0" + correctHex[4:]
	_, err = NewChainIdFromHex(badCharChainId)
	assertError(t, err, ErrChainIdInvalid)
}

func TestChainIdFactories(t *testing.T) {
	tests := []struct {
		identifier string
		factory    func(string) (ChainId, error)
		reference  func() ChainId
	}{
		{
			"1", // no 0x
			func(in string) (ChainId, error) { return NewEVMChainId(in) },
			func() ChainId { return NewEVMEthereumChainId() },
		},
		{
			"0x4268", // with 0x
			func(in string) (ChainId, error) { return NewEVMChainId(in) },
			func() ChainId { return NewEVMHoleskyChainId() },
		},
		{
			"0x0000aa36a7", // with some zeroes
			func(in string) (ChainId, error) { return NewEVMChainId(in) },
			func() ChainId { return NewEVMSepoliaChainId() },
		},
		{
			"0x038", // with odd amount of zeroes
			func(in string) (ChainId, error) { return NewEVMChainId(in) },
			func() ChainId { return NewEVMBinanceSmartChainChainId() },
		},
		{
			"0x0000000000000000000000000000000000000000000000000000000000002105", // full 64 bytes
			func(in string) (ChainId, error) { return NewEVMChainId(in) },
			func() ChainId { return NewEVMBaseChainId() },
		},
		{
			"0x14a34",
			func(in string) (ChainId, error) { return NewEVMChainId(in) },
			func() ChainId { return NewEVMBaseSepoliaChainId() },
		},
		{
			"0x35834a8a",
			func(in string) (ChainId, error) { return NewSuiChainId(in) },
			func() ChainId { return NewSuiMainnetChainId() },
		},
		{
			"0x4c78adac",
			func(in string) (ChainId, error) { return NewSuiChainId(in) },
			func() ChainId { return NewSuiTestnetChainId() },
		},
	}
	for _, tt := range tests {
		chainIdFromFactory, err := tt.factory(tt.identifier)
		assertNoError(t, err)
		chainIdFromReference := tt.reference()
		equalChainId(t, chainIdFromReference, chainIdFromFactory)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.FailNow()
	}
}

func assertError(t *testing.T, err error, errorTypes ...error) {
	if err == nil {
		t.FailNow()
	}
	if len(errorTypes) > 0 {
		for _, target := range errorTypes {
			if !errors.Is(err, target) {
				t.Errorf("error %s is supposed to the %s as well", err.Error(), target.Error())
			}
		}
	}
}

func equalStrings(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}

func equalBytes(t *testing.T, expected []byte, actual []byte) {
	if !bytes.Equal(expected, actual) {
		t.Errorf("expected: %s actual: %s", expected, actual)
	}
}

func equalEcosystem(t *testing.T, expected Ecosystem, actual Ecosystem) {
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected.String(), actual.String())
	}
}

func equalChainId(t *testing.T, expected ChainId, actual ChainId) {
	if !expected.Equal(actual) {
		t.Errorf("expected: %s actual: %s", expected.String(), actual.String())
	}
}
