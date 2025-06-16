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
			"BSC Testnet",
			"0x0000000000000000000000000000000000000000000000000000000000000061",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMBinanceSmartChainTestnetLChainId() },
		},
		{
			"Sonic",
			"0x0000000000000000000000000000000000000000000000000000000000000092",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMSonicLChainId() },
		},
		{
			"Sonic Blaze Testnet",
			"0x000000000000000000000000000000000000000000000000000000000000dede",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMSonicBlazeTestnetLChainId() },
		},
		{
			"Ink",
			"0x000000000000000000000000000000000000000000000000000000000000def1",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMInkLChainId() },
		},
		{
			"Ink Sepolia",
			"0x00000000000000000000000000000000000000000000000000000000000ba5ed",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMInkSepoliaLChainId() },
		},
		{
			"Katana",
			"0x00000000000000000000000000000000000000000000000000000000000b67d2",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMKatanaLChainId() },
		},
		{
			"Katana Tatara Testnet",
			"0x000000000000000000000000000000000000000000000000000000000001f977",
			chainid.EcosystemEVM,
			func() chainid.LChainId { return chainid.NewEVMKatanaTataraTestnetLChainId() },
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
			"Solana",
			"0x02296998a6f8e2a784db5d9f95e18fc23f70441a1039446801089879b08c7ef0",
			chainid.EcosystemSolana,
			func() chainid.LChainId { return chainid.NewSolanaMainnetLChainId() },
		},
		{
			"Solana Devnet",
			"0x0259db5080fc2c6d3bcf7ca90712d3c2e5e6c28f27f0dfbb9953bdb0894c03ab",
			chainid.EcosystemSolana,
			func() chainid.LChainId { return chainid.NewSolanaDevnetLChainId() },
		},
		{
			"Cosmos - Lombard Ledger",
			"0x0387b25e8e61f2ce4838b04795b231f09ee73ffd391da018bef4bc5c4975897b",
			chainid.EcosystemCosmos,
			func() chainid.LChainId { return chainid.NewLombardLedgerLChainId() },
		},
		{
			"Cosmos - Osmosis",
			"0x038ebfb6519e8d814f1b8aee62da9a4e173f7e6898d60d962042421d18dbe4ef",
			chainid.EcosystemCosmos,
			func() chainid.LChainId { return chainid.NewOsmosisLChainId() },
		},
		{
			"Cosmos - Cosmos Hub",
			"0x03a232779a423721bfb80a99e86828034aa5726c469f770d39f29a0fb4710f9a",
			chainid.EcosystemCosmos,
			func() chainid.LChainId { return chainid.NewCosmosHubLChainId() },
		},
		{
			"Cosmos - Babylon",
			"0x0358ed74e1573257904a3b763e53361dbf356e7a01fe98b6e15e91b79c16cb80",
			chainid.EcosystemCosmos,
			func() chainid.LChainId { return chainid.NewBabylonLChainId() },
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

func TestGenericLChainId(t *testing.T) {
	unsupportedChainIdHex := "0x1100000000000000000000000000000000000000000000000000000000000001"
	mayBeGeneric, err := chainid.NewLChainIdFromHex(unsupportedChainIdHex)
	common.AssertNoError(t, err)
	genericChainId, ok := mayBeGeneric.(chainid.GenericLChainId)
	common.AssertTrue(t, ok)
	common.EqualStrings(t, genericChainId.String(), unsupportedChainIdHex)
	common.AssertTrue(t, genericChainId.Ecosystem() == 17)
}

func TestLChainIdErrorConditions(t *testing.T) {
	correctHex := "0x0000000000000000000000000000000000000000000000000000000000000001"
	longerHex := correctHex + "01"
	_, err := chainid.NewLChainIdFromHex(longerHex)
	common.AssertError(t, err, chainid.ErrLChainIdInvalid, chainid.ErrLength)
	longerBytes, _ := hex.DecodeString(longerHex)
	_, err = chainid.NewLChainId(longerBytes)
	common.AssertError(t, err, chainid.ErrLChainIdInvalid, chainid.ErrLength)
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
			"0x14a34", // Base Sepolia
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMBaseSepoliaLChainId() },
		},
		{
			"0x61", // BSC Testnet
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMBinanceSmartChainTestnetLChainId() },
		},
		{
			"0x92", // Sonic
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMSonicLChainId() },
		},
		{
			"0xdede", // Sonic Blaze Testnet
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMSonicBlazeTestnetLChainId() },
		},
		{
			"0xdef1", // Ink
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMInkLChainId() },
		},
		{
			"0xba5ed", // Ink Sepolia
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMInkSepoliaLChainId() },
		},
		{
			"0xb67d2", // Katana
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMKatanaLChainId() },
		},
		{
			"0x1f977", // Katana Tatara Testnet
			func(in string) (chainid.LChainId, error) { return chainid.NewEVMLChainId(in) },
			func() chainid.LChainId { return chainid.NewEVMKatanaTataraTestnetLChainId() },
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
		{
			"5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d",
			func(in string) (chainid.LChainId, error) { return chainid.NewSolanaLChainId(in) },
			func() chainid.LChainId { return chainid.NewSolanaMainnetLChainId() },
		},
		{
			"EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG",
			func(in string) (chainid.LChainId, error) { return chainid.NewSolanaLChainId(in) },
			func() chainid.LChainId { return chainid.NewSolanaDevnetLChainId() },
		},
		{
			"ledger-mainnet-1",
			func(in string) (chainid.LChainId, error) { return chainid.NewCosmosLChainId(in) },
			func() chainid.LChainId { return chainid.NewLombardLedgerLChainId() },
		},
		{
			"osmosis-1",
			func(in string) (chainid.LChainId, error) { return chainid.NewCosmosLChainId(in) },
			func() chainid.LChainId { return chainid.NewOsmosisLChainId() },
		},
		{
			"cosmoshub-4",
			func(in string) (chainid.LChainId, error) { return chainid.NewCosmosLChainId(in) },
			func() chainid.LChainId { return chainid.NewCosmosHubLChainId() },
		},
		{
			"bbn-1",
			func(in string) (chainid.LChainId, error) { return chainid.NewCosmosLChainId(in) },
			func() chainid.LChainId { return chainid.NewBabylonLChainId() },
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
