package chainid

import (
	"encoding/hex"
	"strings"
)

const SuiIdentifierLength = 4

type SuiLChainId struct {
	lChainId
}

// NewSuiMainnetLChainId returns the LChainId for the Sui mainnet blockchain
func NewSuiMainnetLChainId() SuiLChainId {
	return SuiLChainId{
		lChainId{
			inner: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x35, 0x83, 0x4a, 0x8a},
		},
	}
}

// NewSuiTestnetLChainId returns the LChainId for the Sui testnet blockchain
func NewSuiTestnetLChainId() SuiLChainId {
	return SuiLChainId{
		lChainId{
			inner: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4c, 0x78, 0xad, 0xac},
		},
	}
}

// NewSuiLChainId returns a LChainId instance given the 4 bytes identifier, hex encoded,which determines the chain
// Hex is accepted both with and without leading 0x
func NewSuiLChainId(identifier string) (*SuiLChainId, error) {
	trimmed := strings.TrimPrefix(identifier, "0x")
	if len(trimmed) != SuiIdentifierLength*2 {
		return nil, NewErrLength(SuiIdentifierLength, len(trimmed))
	}
	chainid, err := newLChainIdFromHex(EcosystemSui.ToEcosystemHexByte() + repeated64Zeros[SuiIdentifierLength*2+2:] + trimmed)
	if err != nil {
		return nil, err
	}
	return &SuiLChainId{lChainId: *chainid}, nil
}

// Identifier returns the chain identifier as it is meant in the Sui Ecosystem
// i.e. as hex encoded least significant 4 bytes of the genesis block without 0x
func (c SuiLChainId) Identifier() string {
	return hex.EncodeToString(c.inner[28:])
}
