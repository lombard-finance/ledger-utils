package chainid

import (
	"bytes"
	"encoding/hex"
	"errors"
	"strings"
	"testing"
)

func TestLChainId(t *testing.T) {

	tests := []struct {
		name              string
		hexChainId        string
		ecosystem         Ecosystem
		predefinedFactory func() LChainId
	}{
		{
			"Ethereum",
			"0x0000000000000000000000000000000000000000000000000000000000000001",
			EcosystemEVM,
			func() LChainId { return NewEVMEthereumLChainId() },
		},
		{
			"Ethereum Sepolia",
			"0x0000000000000000000000000000000000000000000000000000000000aa36a7",
			EcosystemEVM,
			func() LChainId { return NewEVMSepoliaLChainId() },
		},
		{
			"Ethereum Holesky",
			"0x0000000000000000000000000000000000000000000000000000000000004268",
			EcosystemEVM,
			func() LChainId { return NewEVMHoleskyLChainId() },
		},
		{
			"Base",
			"0x0000000000000000000000000000000000000000000000000000000000002105",
			EcosystemEVM,
			func() LChainId { return NewEVMBaseLChainId() },
		},
		{
			"Base Sepolia",
			"0x0000000000000000000000000000000000000000000000000000000000014a34",
			EcosystemEVM,
			func() LChainId { return NewEVMBaseSepoliaLChainId() },
		},
		{
			"BSC",
			"0x0000000000000000000000000000000000000000000000000000000000000038",
			EcosystemEVM,
			func() LChainId { return NewEVMBinanceSmartChainLChainId() },
		},
		{
			"Sui",
			"0x0100000000000000000000000000000000000000000000000000000035834a8a",
			EcosystemSui,
			func() LChainId { return NewSuiMainnetLChainId() },
		},
		{
			"Sui Testnet",
			"0x010000000000000000000000000000000000000000000000000000004c78adac",
			EcosystemSui,
			func() LChainId { return NewSuiTestnetLChainId() },
		},
		{
			"Bitcoin",
			"0xff0000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
			EcosystemBitcoin,
			func() LChainId { return NewBitcoinLChainId() },
		},
		{
			"Bitcoin Signet",
			"0xff000008819873e925422c1ff0f99f7cc9bbb232af63a077a480a3633bee1ef6",
			EcosystemBitcoin,
			func() LChainId { return NewBitcoinSignetLChainId() },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chainId, err := NewLChainIdFromHex(test.hexChainId)
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

func TestLChainIdErrorConditions(t *testing.T) {
	correctHex := "0x0000000000000000000000000000000000000000000000000000000000000001"
	longerHex := correctHex + "01"
	_, err := NewLChainIdFromHex(longerHex)
	assertError(t, err, ErrLChainIdInvalid, ErrLength)
	longerBytes, _ := hex.DecodeString(longerHex)
	_, err = NewLChainId(longerBytes)
	assertError(t, err, ErrLChainIdInvalid, ErrLength)
	unsupportedEcosystemChainId := "0x11" + correctHex[4:]
	_, err = NewLChainIdFromHex(unsupportedEcosystemChainId)
	assertError(t, err, ErrLChainIdInvalid, ErrUnsupportedEcosystem)
	badCharChainId := "0xY0" + correctHex[4:]
	_, err = NewLChainIdFromHex(badCharChainId)
	assertError(t, err, ErrLChainIdInvalid)
}

func TestLChainIdFactories(t *testing.T) {
	tests := []struct {
		identifier string
		factory    func(string) (LChainId, error)
		reference  func() LChainId
	}{
		{
			"1", // no 0x
			func(in string) (LChainId, error) { return NewEVMLChainId(in) },
			func() LChainId { return NewEVMEthereumLChainId() },
		},
		{
			"0x4268", // with 0x
			func(in string) (LChainId, error) { return NewEVMLChainId(in) },
			func() LChainId { return NewEVMHoleskyLChainId() },
		},
		{
			"0x0000aa36a7", // with some zeroes
			func(in string) (LChainId, error) { return NewEVMLChainId(in) },
			func() LChainId { return NewEVMSepoliaLChainId() },
		},
		{
			"0x038", // with odd amount of zeroes
			func(in string) (LChainId, error) { return NewEVMLChainId(in) },
			func() LChainId { return NewEVMBinanceSmartChainLChainId() },
		},
		{
			"0x0000000000000000000000000000000000000000000000000000000000002105", // full 64 bytes
			func(in string) (LChainId, error) { return NewEVMLChainId(in) },
			func() LChainId { return NewEVMBaseLChainId() },
		},
		{
			"0x14a34",
			func(in string) (LChainId, error) { return NewEVMLChainId(in) },
			func() LChainId { return NewEVMBaseSepoliaLChainId() },
		},
		{
			"0x35834a8a",
			func(in string) (LChainId, error) { return NewSuiLChainId(in) },
			func() LChainId { return NewSuiMainnetLChainId() },
		},
		{
			"0x4c78adac",
			func(in string) (LChainId, error) { return NewSuiLChainId(in) },
			func() LChainId { return NewSuiTestnetLChainId() },
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

func equalChainId(t *testing.T, expected LChainId, actual LChainId) {
	if !expected.Equal(actual) {
		t.Errorf("expected: %s actual: %s", expected.String(), actual.String())
	}
}
