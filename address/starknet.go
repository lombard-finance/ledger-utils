package address

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/lombard-finance/ledger-utils/chainid"
)

const StarknetAddressLength = 32

// ErrBadAddressStarknet is an ErrBadAddress specialized for StarknetAddress
var ErrBadAddressStarknet = fmt.Errorf("starknet %w", ErrBadAddress)

// StarknetAddress is the address type for the Starknet L2
type StarknetAddress struct {
	inner [StarknetAddressLength]byte
}

// NewStarknetAddress creates a new StarknetAddress from a slice of bytes
func NewStarknetAddress(b []byte) (*StarknetAddress, error) {
	if len(b) != StarknetAddressLength {
		return nil, fmt.Errorf("%w: lenght error, given %d, expected %d", ErrBadAddressStarknet, len(b), StarknetAddressLength)
	}
	a := &StarknetAddress{}
	copy(a.inner[:], b)
	return a, nil
}

// NewStarknetAddressFromHex creates a new StarknetAddress from an hex string. Both string with
// and without leading 0x are supported
func NewStarknetAddressFromHex(address string) (*StarknetAddress, error) {
	b, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))
	if err != nil {
		return nil, fmt.Errorf("%w: hex decoding error %w", ErrBadAddressStarknet, err)
	}
	return NewStarknetAddress(b)
}

func (s *StarknetAddress) String() string {
	return "0x" + s.Hex()
}

func (s *StarknetAddress) Hex() string {
	return hex.EncodeToString(s.inner[:])
}

func (s *StarknetAddress) Bytes() []byte {
	buf := make([]byte, StarknetAddressLength)
	copy(buf, s.inner[:])
	return buf
}

func (s *StarknetAddress) Length() int {
	return StarknetAddressLength
}

func (s *StarknetAddress) Ecosystem() chainid.Ecosystem {
	return chainid.EcosystemStarknet
}

func (a1 *StarknetAddress) Equal(a2 Address) bool {
	if a2.Ecosystem() != a1.Ecosystem() {
		return false
	}
	return bytes.Equal(a1.inner[:], a2.Bytes())
}
