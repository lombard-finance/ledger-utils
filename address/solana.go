package address

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common/base58"
)

const SolanaAddressLength = 32

// ErrBadAddressSolana is an ErrBadAddress specialized for Solana
var ErrBadAddressSolana = fmt.Errorf("solana %w", ErrBadAddress)

// SolanaAddress is the address type for the Solana blockchain
type SolanaAddress struct {
	inner [SolanaAddressLength]byte
}

// NewSolanaAddress creates a new SolanaAddress from a slice of bytes
func NewSolanaAddress(b []byte) (*SolanaAddress, error) {
	if len(b) != SolanaAddressLength {
		return nil, fmt.Errorf("%w: length error, given %d, expected %d", ErrBadAddressSolana, len(b), SolanaAddressLength)
	}
	a := &SolanaAddress{}
	copy(a.inner[:], b)
	return a, nil
}

// NewSolanaAddressFromHex creates a new SolanaAddress from an hex string. Both string with
// and without leading 0x are supported
func NewSolanaAddressFromHex(address string) (*SolanaAddress, error) {
	b, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))
	if err != nil {
		return nil, fmt.Errorf("%w: hex decoding error %w", ErrBadAddressSolana, err)
	}
	return NewSolanaAddress(b)
}

// NewSolanaAddressFromBase58 creates a new SolanaAddress from a base58 string.
func NewSolanaAddressFromBase58(address string) (*SolanaAddress, error) {
	b, err := base58.Decode(address)
	if err != nil {
		return nil, fmt.Errorf("%w: base58 decoding error %w", ErrBadAddressSolana, err)
	}
	return NewSolanaAddress(b)
}

// String returns the base58 encoding of the address as common in the Solana ecosystem
func (s *SolanaAddress) String() string {
	return base58.Encode(s.inner[:])
}

func (s *SolanaAddress) Hex() string {
	return hex.EncodeToString(s.inner[:])
}

func (s *SolanaAddress) Bytes() []byte {
	buf := make([]byte, SolanaAddressLength)
	copy(buf, s.inner[:])
	return buf
}

func (s *SolanaAddress) Length() int {
	return SolanaAddressLength
}

func (s *SolanaAddress) Ecosystem() chainid.Ecosystem {
	return chainid.EcosystemSolana
}

func (a1 *SolanaAddress) Equal(a2 Address) bool {
	if a2.Ecosystem() != a1.Ecosystem() {
		return false
	}
	a2AsSolana, ok := a2.(*SolanaAddress)
	if !ok {
		return false
	}
	return bytes.Equal(a1.inner[:], a2AsSolana.inner[:])
}
