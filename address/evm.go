package address

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/lombard-finance/ledger-utils/chainid"
)

const EvmAddressLength = 20

// ErrBadAddressEvm is an ErrBadAddress specialized for EVM chains
var ErrBadAddressEvm = fmt.Errorf("evm %w", ErrBadAddress)

// EvmAddress is the address type for EVM chains
type EvmAddress struct {
	inner [EvmAddressLength]byte
}

// NewEvmAddress creates a new EvmAddress from a byte slice. If slice is longer than 20 bytes, most significant ones
// are truncated. An error is returned if the slice is shorter than 20 bytes or if the truncated bytes are not zeroes.
func NewEvmAddress(b []byte) (*EvmAddress, error) {
	if len(b) < EvmAddressLength {
		return nil, fmt.Errorf("%w: invalid length, given %d, expected at least %d", ErrBadAddressEvm, len(b), EvmAddressLength)
	}
	if !bytes.Equal(b[:len(b)-EvmAddressLength], make([]byte, len(b)-EvmAddressLength)) {
		return nil, fmt.Errorf("%w: truncated bytes are not zeroes", ErrBadAddressEvm)
	}

	a := &EvmAddress{}
	copy(a.inner[:], b[len(b)-EvmAddressLength:])
	return a, nil
}

// NewEvmAddressFromHex creates a new EvmAddress from an hex string. Both string with
// and without leading 0x are supported
func NewEvmAddressFromHex(address string) (*EvmAddress, error) {
	decoded, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))
	if err != nil {
		return nil, fmt.Errorf("%w: hex decode error %w", ErrBadAddressEvm, err)
	}
	return NewEvmAddress(decoded)
}

func (a *EvmAddress) String() string {
	return "0x" + a.Hex()
}

func (a *EvmAddress) Hex() string {
	return hex.EncodeToString(a.inner[:])
}

func (a *EvmAddress) Bytes() []byte {
	buf := make([]byte, EvmAddressLength)
	copy(buf, a.inner[:])
	return buf
}

func (a *EvmAddress) Length() int {
	return EvmAddressLength
}

func (a *EvmAddress) Ecosystem() chainid.Ecosystem {
	return chainid.EcosystemEVM
}

func (a1 *EvmAddress) Equal(a2 Address) bool {
	if a2.Ecosystem() != a1.Ecosystem() {
		return false
	}
	return bytes.Equal(a1.inner[:], a2.Bytes())
}
