package chainid_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common"
)

func TestLChainId(t *testing.T) {

	tests := []struct {
		name              string
		hexChainId        string
		ecosystem         chainid.Ecosystem
		predefinedFactory func() chainid.LChainId
	}{
		{
			"Ethereum",
			"0x0000000000000000000000000000000000000000000000000000000000000001",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMEthereumLChainId() },
		},
		{
			"Ethereum Sepolia",
			"0x0000000000000000000000000000000000000000000000000000000000aa36a7",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMSepoliaLChainId() },
		},
		{
			"Ethereum Holesky",
			"0x0000000000000000000000000000000000000000000000000000000000004268",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMHoleskyLChainId() },
		},
		{
			"Base",
			"0x0000000000000000000000000000000000000000000000000000000000002105",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMBaseLChainId() },
		},
		{
			"Base Sepolia",
			"0x0000000000000000000000000000000000000000000000000000000000014a34",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMBaseSepoliaLChainId() },
		},
		{
			"BSC",
			"0x0000000000000000000000000000000000000000000000000000000000000038",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMBinanceSmartChainLChainId() },
		},
		{
			"Sui",
			"0x0100000000000000000000000000000000000000000000000000000035834a8a",
			chainid.EcosystemSui,
			func() chainid.LChainId { return chainid.NewSuiMainnetLChainId() },
		},
		{
			"Sui Testnet",
			"0x010000000000000000000000000000000000000000000000000000004c78adac",
			chainid.EcosystemSui,
			func() chainid.LChainId { return chainid.NewSuiTestnetLChainId() },
		},
		{
			"Bitcoin",
			"0xff0000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
			chainid.EcosystemBitcoin,
			func() chainid.LChainId { return chainid.NewBitcoinLChainId() },
		},
		{
			"Bitcoin Signet",
			"0xff000008819873e925422c1ff0f99f7cc9bbb232af63a077a480a3633bee1ef6",
			chainid.EcosystemBitcoin,
			func() chainid.LChainId { return chainid.NewBitcoinSignetLChainId() },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chainId, err := chainid.NewLChainIdFromHex(test.hexChainId)
			common.AssertNoError(t, err)
			common.EqualStrings(t, test.hexChainId, chainId.String())
			common.EqualStrings(t, test.hexChainId, "0x"+chainId.Hex())
			equalEcosystem(t, test.ecosystem, chainId.Ecosystem())
			equalChainId(t, test.predefinedFactory(), chainId)
			referenceChainIdBytes, err := hex.DecodeString(strings.TrimPrefix(test.hexChainId, "0x"))
			common.AssertNoError(t, err)
			common.EqualBytes(t, referenceChainIdBytes, chainId.Bytes())
			fixed := chainId.FixedBytes()
			common.EqualBytes(t, referenceChainIdBytes, fixed[:])
		})
	}
}

func TestLChainIdErrorConditions(t *testing.T) {
	correctHex := "0x0000000000000000000000000000000000000000000000000000000000000001"
	longerHex := correctHex + "01"
	_, err := chainid.NewLChainIdFromHex(longerHex)
	common.AssertError(t, err, chainid.ErrLChainIdInvalid, chainid.ErrLength)
	longerBytes, _ := hex.DecodeString(longerHex)
	_, err = chainid.NewLChainId(longerBytes)
	common.AssertError(t, err, chainid.ErrLChainIdInvalid, chainid.ErrLength)
	unsupportedEcosystemChainId := "0x11" + correctHex[4:]
	_, err = chainid.NewLChainIdFromHex(unsupportedEcosystemChainId)
	common.AssertError(t, err, chainid.ErrLChainIdInvalid, chainid.ErrUnsupportedEcosystem)
	badCharChainId := "0xY0" + correctHex[4:]
	_, err = chainid.NewLChainIdFromHex(badCharChainId)
	common.AssertError(t, err, chainid.ErrLChainIdInvalid)
}

func TestLChainIdFactories(t *testing.T) {
	tests := []struct {
		identifier string
		factory    func(string) (chainid.LChainId, error)
		reference  func() chainid.LChainId
	}{
		{
			"1", // no 0x
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMEthereumLChainId() },
		},
		{
			"0x4268", // with 0x
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMHoleskyLChainId() },
		},
		{
			"0x0000aa36a7", // with some zeroes
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMSepoliaLChainId() },
		},
		{
			"0x038", // with odd amount of zeroes
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMBinanceSmartChainLChainId() },
		},
		{
			"0x0000000000000000000000000000000000000000000000000000000000002105", // full 64 bytes
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMBaseLChainId() },
		},
		{
			"0x14a34",
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMBaseSepoliaLChainId() },
		},
		{
			"0x35834a8a",
			func(in string) (chainid.LChainId, error) { return chainid.NewSuiLChainId(in) },
			func() chainid.LChainId { return chainid.NewSuiMainnetLChainId() },
		},
		{
			"0x4c78adac",
			func(in string) (chainid.LChainId, error) { return chainid.NewSuiLChainId(in) },
			func() chainid.LChainId { return chainid.NewSuiTestnetLChainId() },
		},
	}
	for _, tt := range tests {
		chainIdFromFactory, err := tt.factory(tt.identifier)
		common.AssertNoError(t, err)
		chainIdFromReference := tt.reference()
		equalChainId(t, chainIdFromReference, chainIdFromFactory)
	}
}

func equalEcosystem(t *testing.T, expected chainid.Ecosystem, actual chainid.Ecosystem) {
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected.String(), actual.String())
	}
}

func equalChainId(t *testing.T, expected chainid.LChainId, actual chainid.LChainId) {
	if !expected.Equal(actual) {
		t.Errorf("expected: %s actual: %s", expected.String(), actual.String())
	}
}
