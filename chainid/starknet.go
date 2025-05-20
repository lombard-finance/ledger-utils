package chainid

import (
	"bytes"
	"strings"

	"github.com/lombard-finance/ledger-utils/common"
)

type StarknetLChainId struct {
	lChainId
}

// MAX_STARKNET_CHAIN_ID_LENGTH is the maximum length of a Starknet chain id represented in hex string.
// We accept up to 31 that fit as-is in the Lombard Chain Id, while currently available ones are only
// 0x534e5f5345504f4c4941 for sepolia and 0x534e5f4d41494e for main
const MAX_STARKNET_CHAIN_ID_LENGTH = 62

func NewStarknetLChainId(id string) (*StarknetLChainId, error) {
	trimmed := strings.TrimPrefix(id, "0x")
	if len(trimmed) > MAX_STARKNET_CHAIN_ID_LENGTH {
		return nil, NewMaxErrLength(MAX_STARKNET_CHAIN_ID_LENGTH, len(id))
	}
	innerChainId, err := newLChainIdFromHex(
		EcosystemStarknet.ToEcosystemHexByte() +
			common.Repeated64Zeros[len(trimmed)+2:] +
			trimmed,
	)
	if err != nil {
		return nil, err
	}
	return &StarknetLChainId{
		lChainId: *innerChainId,
	}, nil
}

func NewStarknetLChainIdFromName(name string) (*StarknetLChainId, error) {
	trimmed := []byte(strings.TrimSpace(name))
	if len(trimmed) > MAX_STARKNET_CHAIN_ID_LENGTH {
		return nil, NewMaxErrLength(MAX_STARKNET_CHAIN_ID_LENGTH, len(trimmed))
	}
	byteChainId := make([]byte, 32)
	byteChainId[0] = byte(EcosystemStarknet)
	copy(byteChainId[len(byteChainId)-len(trimmed):], trimmed)
	innerChainId, err := newLChainId(byteChainId)
	if err != nil {
		return nil, err
	}
	return &StarknetLChainId{
		lChainId: *innerChainId,
	}, nil
}

// NewStarknetMainnetLChainId returns the Starknet Lombard Chain Id of Starknet mainnet SN_MAIN
func NewStarknetMainnetLChainId() *StarknetLChainId {
	chId, _ := NewStarknetLChainId("0x534e5f4d41494e")
	return chId
}

// NewStarknetSepoliaLChainId returns the Starknet Lombard Chain Id of Starknet sepolia SN_SEPOLIA
func NewStarknetSepoliaLChainId() *StarknetLChainId {
	chId, _ := NewStarknetLChainId("0x534e5f5345504f4c4941")
	return chId
}

// Identifier returns the textual identifier of a Starknet network.
// They all start with SN
func (ch *StarknetLChainId) Identifier() string {
	snIndex := bytes.Index(ch.inner[1:], []byte("SN"))
	if snIndex == -1 {
		return string(ch.inner[1:])
	}
	return string(ch.inner[snIndex+1:])
}
