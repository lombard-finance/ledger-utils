package chainid

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var ErrEmptyChainId = fmt.Errorf("cannot create Lombard chain Id from empty chain id")

type CosmosLChainId struct {
	lChainId
}

// NewCosmosLChainId generates a new Lombard Chain Id for a Cosmos chain given its chain id. Note that since chain ids
// in the Cosmos ecosystem may change based on the amount of time a chain has been restarted (e.g. cosmoshub-4) we only
// consider what is referred as chain name, i.e., the chain id without the trailing dash and incrementing counter.
// The resulting Lombard Chain Id is the hash of such chain name with its MSB replaced by the ecosystem byte
func NewCosmosLChainId(chainId string) (*CosmosLChainId, error) {
	if chainId == "" {
		return nil, ErrEmptyChainId
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
