package chainid

import (
	"encoding/hex"
	"errors"
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
			NewEVMEthereumChainId,
		},
		{
			"Ethereum Sepolia",
			"0x0000000000000000000000000000000000000000000000000000000000aa36a7",
			EcosystemEVM,
			NewEVMSepoliaChainId,
		},
		{
			"Ethereum Holesky",
			"0x0000000000000000000000000000000000000000000000000000000000004268",
			EcosystemEVM,
			NewEVMHoleskyChainId,
		},
		{
			"Base",
			"0x0000000000000000000000000000000000000000000000000000000000002105",
			EcosystemEVM,
			NewEVMBaseChainId,
		},
		{
			"Base Sepolia",
			"0x0000000000000000000000000000000000000000000000000000000000014a34",
			EcosystemEVM,
			NewEVMBaseSepoliaChainId,
		},
		{
			"BSC",
			"0x0000000000000000000000000000000000000000000000000000000000000038",
			EcosystemEVM,
			NewEVMBinanceSmartChainChainId,
		},
		{
			"Sui",
			"0x0100000000000000000000000000000000000000000000000000000035834a8a",
			EcosystemSui,
			NewSuiMainnetChainId,
		},
		{
			"Sui Testnet",
			"0x010000000000000000000000000000000000000000000000000000004c78adac",
			EcosystemSui,
			NewSuiTestnetChainId,
		},
		{
			"Bitcoin",
			"0xff0000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
			EcosystemBitcoin,
			NewBitcoinChainId,
		},
		{
			"Bitcoin Signet",
			"0xff000008819873e925422c1ff0f99f7cc9bbb232af63a077a480a3633bee1ef6",
			EcosystemBitcoin,
			NewBitcoinSignetChainId,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chainId, err := NewChainIdFromHex(test.hexChainId)
			assertNoError(t, err)
			equalStrings(t, test.hexChainId, chainId.String())
			equalEcosystem(t, test.ecosystem, chainId.Ecosystem())
			equalChainId(t, test.predefinedFactory(), *chainId)
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

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fail()
	}
}

func assertError(t *testing.T, err error, errorTypes ...error) {
	if err == nil {
		t.Fail()
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
