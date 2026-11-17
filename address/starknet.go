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

// starknetFieldPrime is the Starknet field prime 2^251 + 17*2^192 + 1, big-endian encoded
var starknetFieldPrime = [StarknetAddressLength]byte{0: 0x08, 7: 0x11, 31: 0x01}

// StarknetAddress is the address type for the Starknet L2
type StarknetAddress struct {
	inner [StarknetAddressLength]byte
}

// NewStarknetAddress creates a new StarknetAddress from a slice of bytes
func NewStarknetAddress(b []byte) (*StarknetAddress, error) {
	if len(b) != StarknetAddressLength {
		return nil, fmt.Errorf("%w: lenght error, given %d, expected %d", ErrBadAddressStarknet, len(b), StarknetAddressLength)
	}
	// An address is a field element. Values not lower than the prime are reduced modulo the prime
	// by Starknet tooling, so they would be a second byte representation of an existing address.
	if bytes.Compare(b, starknetFieldPrime[:]) >= 0 {
		return nil, fmt.Errorf("%w: not lower than the field prime: %x", ErrBadAddressStarknet, b)
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
	a2AsStark, ok := a2.(*StarknetAddress)
	if !ok {
		return false
	}
	return bytes.Equal(a1.inner[:], a2AsStark.inner[:])
}
