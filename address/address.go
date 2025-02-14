package address

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/lombard-finance/chain/chainid"
)

// compile time type assertion
var _ Address = &EvmAddress{}
var _ Address = &SuiAddress{}

// The Address interface serves as a unique gateway to addresses of all chains. In fact each library may treat
// addresses in different manners and with different assumptions. So we try to unify all of them.
type Address interface {
	// String is a string encoded representation of the address in the common way in which it is represented in the
	// respective ecosystem
	String() string
	// Bytes returns the bytes of the address according to the length and format of the respective ecosystem. Value is
	// a copy of the inner representation so it is safe to modify
	Bytes() []byte
	// Hex return the hex encoded version of the bytes of the address
	Hex() string
	// Length returns the lenght of the address when referencing it as an array of bytes
	Length() int
	// Ecosystem returns the ecosystem the adress belongs to according to the Lombard internal reference
	Ecosystem() chainid.Ecosystem
	// Equal verifies if the passed address is the same of the current instance
	Equal(a2 Address) bool
}

// ErrBadAddress is a generic error to return when there is an error with address validation
var ErrBadAddress = fmt.Errorf("address invalid")
var ErrUnsupportedEcosystem = fmt.Errorf("ecosystem is unsupported")

// NewAddress creates a new Address instance from a slice of bytes whose concrete implementation
// is defined by the provided ecosystem
func NewAddress(b []byte, e chainid.Ecosystem) (Address, error) {
	switch e {
	case chainid.EcosystemEVM:
		return NewEvmAddress(b)
	case chainid.EcosystemSui:
		return NewSuiAddress(b)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedEcosystem, e.String())
	}
}

// NewAddressFromHex creates a new Address instance from an hex string of bytes whose concrete implementation
// is defined by the provided ecosystem. `0x` is optional.
func NewAddressFromHex(address string, e chainid.Ecosystem) (Address, error) {
	decoded, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))
	if err != nil {
		return nil, fmt.Errorf("%w: hex decoding error: %w", ErrBadAddress, err)
	}
	return NewAddress(decoded, e)
}

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
func NewEvmAddressTruncating(b []byte) (*EvmAddress, error) {
	if len(b) < EvmAddressLength {
		return nil, fmt.Errorf("%w: invalid length, given %d, expected at least %d", ErrBadAddressEvm, len(b), EvmAddressLength)
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
	a2AsSui, ok := a2.(*SuiAddress)
	if !ok {
		return false
	}
	return bytes.Equal(a1.inner[:], a2AsSui.inner[:])
}
