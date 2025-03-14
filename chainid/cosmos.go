package chainid

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var ErrEmptyCosmosChainId = fmt.Errorf("cannot create Lombard chain Id from empty cosmos chain id")
var ErrInvalidCosmosChainId = fmt.Errorf("chain id is not valid")

type CosmosLChainId struct {
	lChainId
}

// NewCosmosLChainId generates a new Lombard Chain Id for a Cosmos chain given its chain id. Note that since chain ids
// in the Cosmos ecosystem may change based on the amount of time a chain has been restarted (e.g. cosmoshub-4) we only
// consider what is referred as chain name, i.e., the chain id without the trailing dash and incrementing counter.
// The resulting Lombard Chain Id is the hash of such chain name with its MSB replaced by the ecosystem byte
func NewCosmosLChainId(chainId string) (*CosmosLChainId, error) {
	if chainId == "" {
		return nil, fmt.Errorf("%w: %w", ErrInvalidCosmosChainId, ErrEmptyCosmosChainId)
	}
	lastDashPosition := strings.LastIndex(chainId, "-")
	if lastDashPosition == -1 {
		return nil, fmt.Errorf("%w: cannot find counter in provided chain id", ErrInvalidCosmosChainId)
	}
	chainName := chainId[0:strings.LastIndex(chainId, "-")]
	hashedChainName := sha256.Sum256([]byte(chainName))
	// Replace MSB with cosmos ecosystem byte
	hashedChainName[0] = byte(EcosystemCosmos)
	innerChainId, err := newLChainId(hashedChainName[:])
	if err != nil {
		return nil, err
	}
	return &CosmosLChainId{
		lChainId: *innerChainId,
	}, nil
}

func NewLombardLedgerLChainId() *CosmosLChainId {
	return &CosmosLChainId{
		lChainId: lChainId{
			inner: []byte{0x03, 0x87, 0xb2, 0x5e, 0x8e, 0x61, 0xf2, 0xce, 0x48, 0x38, 0xb0, 0x47, 0x95, 0xb2, 0x31, 0xf0, 0x9e, 0xe7, 0x3f, 0xfd, 0x39, 0x1d, 0xa0, 0x18, 0xbe, 0xf4, 0xbc, 0x5c, 0x49, 0x75, 0x89, 0x7b},
		},
	}
}

func NewOsmosisLChainId() *CosmosLChainId {
	return &CosmosLChainId{
		lChainId: lChainId{
			inner: []byte{0x03, 0x8e, 0xbf, 0xb6, 0x51, 0x9e, 0x8d, 0x81, 0x4f, 0x1b, 0x8a, 0xee, 0x62, 0xda, 0x9a, 0x4e, 0x17, 0x3f, 0x7e, 0x68, 0x98, 0xd6, 0x0d, 0x96, 0x20, 0x42, 0x42, 0x1d, 0x18, 0xdb, 0xe4, 0xef},
		},
	}
}

func NewCosmosHubLChainId() *CosmosLChainId {
	return &CosmosLChainId{
		lChainId: lChainId{
			inner: []byte{0x03, 0xa2, 0x32, 0x77, 0x9a, 0x42, 0x37, 0x21, 0xbf, 0xb8, 0x0a, 0x99, 0xe8, 0x68, 0x28, 0x03, 0x4a, 0xa5, 0x72, 0x6c, 0x46, 0x9f, 0x77, 0x0d, 0x39, 0xf2, 0x9a, 0x0f, 0xb4, 0x71, 0x0f, 0x9a},
		},
	}
}
