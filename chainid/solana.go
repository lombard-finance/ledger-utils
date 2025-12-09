package chainid

import "github.com/lombard-finance/ledger-utils/common/base58"

const SolanaGenesisHashLength = 32

type SolanaLChainId struct {
	lChainId
}

func NewSolanaLChainId(genesisHash string) (*SolanaLChainId, error) {
	decoded, err := base58.Decode(genesisHash)
	if err != nil {
		return nil, err
	}
	if len(decoded) != SolanaGenesisHashLength {
		return nil, NewErrLength(SolanaGenesisHashLength, len(decoded))
	}
	// swap MSB with our ecosystem id
	decoded[0] = byte(EcosystemSolana)
	innerChainId, err := newLChainId(decoded)
	if err != nil {
		return nil, err
	}
	return &SolanaLChainId{lChainId: *innerChainId}, nil
}

// NewSolanaMainnetLChainId returns the ChainId for the Solana blockchain
func NewSolanaMainnetLChainId() SolanaLChainId {
    return SolanaLChainId{
        lChainId{
            inner: [32]byte{byte(EcosystemSolana), 0x29, 0x69, 0x98, 0xa6, 0xf8, 0xe2, 0xa7, 0x84, 0xdb, 0x5d, 0x9f, 0x95, 0xe1, 0x8f, 0xc2, 0x3f, 0x70, 0x44, 0x1a, 0x10, 0x39, 0x44, 0x68, 0x01, 0x08, 0x98, 0x79, 0xb0, 0x8c, 0x7e, 0xf0},
        },
    }
}

// NewSolanaDevnetLChainId returns the ChainId for the Solana Devnet blockchain
func NewSolanaDevnetLChainId() SolanaLChainId {
    return SolanaLChainId{
        lChainId{
            inner: [32]byte{byte(EcosystemSolana), 0x59, 0xdb, 0x50, 0x80, 0xfc, 0x2c, 0x6d, 0x3b, 0xcf, 0x7c, 0xa9, 0x07, 0x12, 0xd3, 0xc2, 0xe5, 0xe6, 0xc2, 0x8f, 0x27, 0xf0, 0xdf, 0xbb, 0x99, 0x53, 0xbd, 0xb0, 0x89, 0x4c, 0x03, 0xab},
        },
    }
}
