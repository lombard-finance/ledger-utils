package chainid

import (
	"encoding/hex"
	"strings"
)

const SuiIdentifierLength = 4

type SuiChainId struct {
	chainId
}

// NewSuiMainnetChainId returns the ChainId for the Sui mainnet blockchain
func NewSuiMainnetChainId() SuiChainId {
	return SuiChainId{
		chainId{
			inner: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x35, 0x83, 0x4a, 0x8a},
		},
	}
}

// NewSuiTestnetChainId returns the ChainId for the Sui testnet blockchain
func NewSuiTestnetChainId() SuiChainId {
	return SuiChainId{
		chainId{
			inner: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4c, 0x78, 0xad, 0xac},
		},
	}
}

// NewSuiChainId returns a ChainId instance given the 4 bytes identifier, hex encoded,which determines the chain
// Hex is accepted both with and without leading 0x
func NewSuiChainId(identifier string) (*SuiChainId, error) {
	trimmed := strings.TrimPrefix(identifier, "0x")
	if len(trimmed) != SuiIdentifierLength*2 {
		return nil, NewErrLenght(SuiIdentifierLength, len(trimmed))
	}
	chainid, err := newChainIdFromHex(EcosystemSui.ToEcosystemHexByte() + repeated64Zeros[SuiIdentifierLength*2+2:] + trimmed)
	if err != nil {
		return nil, err
	}
	return &SuiChainId{chainId: *chainid}, nil
}

// Identifier returns the chain identifier as it is meant in the Sui Ecosystem
// i.e. as hex encoded least significant 4 bytes of the genesis block without 0x
func (c SuiChainId) Identifier() string {
	return hex.EncodeToString(c.inner[28:])
}
