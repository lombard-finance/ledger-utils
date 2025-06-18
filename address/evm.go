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

// NewEvmAddress creates a new EvmAddress from a 20-bytes array
func NewEvmAddress(address []byte) (*EvmAddress, error) {
	if len(address) != EvmAddressLength {
		return nil, fmt.Errorf("%w: invalid length, given %d, expected %d", ErrBadAddressEvm, len(address), EvmAddressLength)
	}

	a := &EvmAddress{}
	copy(a.inner[:], address)
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

// NewEvmAddressTruncating creates a new EvmAddress from a byte slice which is longer than
// EvmAddressLength by truncating most significant bytes. This is useful when creating an
// address from a event topic which is always 32 bytes long.
// It returns an error if the byte slice is shorter than EvmAddressLength or truncated bytes
// are not zeroes.
func NewEvmAddressTruncating(b []byte) (*EvmAddress, error) {
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
	a2AsEvm, ok := a2.(*EvmAddress)
	if !ok {
		return false
	}
	return bytes.Equal(a1.inner[:], a2AsEvm.inner[:])
}
