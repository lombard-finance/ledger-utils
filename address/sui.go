package address

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/lombard-finance/ledger-utils/chainid"
)

const SuiAddressLength = 32

// ErrBadAddressSui is an ErrBadAddress specialized for Sui
var ErrBadAddressSui = fmt.Errorf("sui %w", ErrBadAddress)

// SuiAddress is the address type for the Sui blockchain
type SuiAddress struct {
	inner [SuiAddressLength]byte
}

// NewSuiAddress creates a new SuiAddress from a slice of bytes
func NewSuiAddress(b []byte) (*SuiAddress, error) {
	if len(b) != SuiAddressLength {
		return nil, fmt.Errorf("%w: lenght error, given %d, expected %d", ErrBadAddressSui, len(b), SuiAddressLength)
	}
	a := &SuiAddress{}
	copy(a.inner[:], b)
	return a, nil
}

// NewSuiAddressFromHex creates a new SuiAddress from an hex string. Both string with
// and without leading 0x are supported
func NewSuiAddressFromHex(address string) (*SuiAddress, error) {
	b, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))
	if err != nil {
		return nil, fmt.Errorf("%w: hex decoding error %w", ErrBadAddressSui, err)
	}
	return NewSuiAddress(b)
}

func (s *SuiAddress) String() string {
	return "0x" + s.Hex()
}

func (s *SuiAddress) Hex() string {
	return hex.EncodeToString(s.inner[:])
}

func (s *SuiAddress) Bytes() []byte {
	buf := make([]byte, SuiAddressLength)
	copy(buf, s.inner[:])
	return buf
}

func (s *SuiAddress) Length() int {
	return SuiAddressLength
}

func (s *SuiAddress) Ecosystem() chainid.Ecosystem {
	return chainid.EcosystemSui
}

func (a1 *SuiAddress) Equal(a2 Address) bool {
	if a2.Ecosystem() != a1.Ecosystem() {
		return false
	}
	return bytes.Equal(a1.inner[:], a2.Bytes())
}
