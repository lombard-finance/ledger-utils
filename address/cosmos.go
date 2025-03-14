package address

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common"
)

// CosmWasmAddressLength is the length of a contract address in CosmWasm without prefix and checksum
const CosmWasmAddressLength = 32

// CosmosSdkAddressLength is the length of an address in a Cosmos SDK chain without prefix and checksum
const CosmosSdkAddressLength = 20

// DifferenceWasmSdkLength is the length difference between the two accepted formats
const DifferenceWasmSdkLength = CosmWasmAddressLength - CosmosSdkAddressLength

// ErrBadAddressCosmos is an ErrBadAddress specialized for Cosmos chains
var ErrBadAddressCosmos = fmt.Errorf("evm %w", ErrBadAddress)

// CosmosAddress is the address type generic for Cosmos chains. It is NOT tied to a particular chain.
// It includes very basic functionalities. Any additional property like verifying bech32 prefix
// and checksum should be conducted with the Cosmos SDK.
type CosmosAddress struct {
	inner []byte
}

func NewCosmosAddress(addressBytes []byte) (*CosmosAddress, error) {
	if len(addressBytes) != CosmWasmAddressLength && len(addressBytes) != CosmosSdkAddressLength {
		return nil, fmt.Errorf(
			"%w: invalid address length, expected %d or %d (CosmWasm) but given %d",
			ErrBadAddressCosmos,
			CosmosSdkAddressLength,
			CosmWasmAddressLength,
			len(addressBytes),
		)
	}
	// if leading bytes are all zeros then trim them and reduce to SDK address
	if bytes.Equal(addressBytes[:DifferenceWasmSdkLength], common.Bytes32Zeros[:DifferenceWasmSdkLength]) {
		addressBytes = addressBytes[DifferenceWasmSdkLength:]
	}
	a := CosmosAddress{
		inner: make([]byte, len(addressBytes)),
	}
	copy(a.inner[:], addressBytes)
	return &a, nil
}

// Bytes implements Address.
func (c *CosmosAddress) Bytes() []byte {
	buf := make([]byte, len(c.inner))
	copy(buf, c.inner)
	return buf
}

// Ecosystem implements Address.
func (c *CosmosAddress) Ecosystem() chainid.Ecosystem {
	return chainid.EcosystemCosmos
}

// Equal implements Address.
func (c *CosmosAddress) Equal(a2 Address) bool {
	if a2 == nil {
		return false
	}
	if c.Ecosystem() != a2.Ecosystem() {
		return false
	}
	a2AsCosmos, ok := a2.(*CosmosAddress)
	if !ok {
		return false
	}
	return bytes.Equal(a2AsCosmos.inner, c.inner)
}

// Hex implements Address.
func (c *CosmosAddress) Hex() string {
	return hex.EncodeToString(c.inner)
}

// Length implements Address.
func (c *CosmosAddress) Length() int {
	return len(c.inner)
}

// String implements Address.
func (c *CosmosAddress) String() string {
	return "0x" + c.Hex()
}
