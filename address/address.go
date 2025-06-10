package address

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/lombard-finance/ledger-utils/chainid"
	"github.com/lombard-finance/ledger-utils/common"
)

// compile time type assertion
var _ Address = &CosmosAddress{}
var _ Address = &EvmAddress{}
var _ Address = &SolanaAddress{}
var _ Address = &SuiAddress{}
var _ Address = &GenericAddress{}

var ErrEmptyAddress = fmt.Errorf("empty address")

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
		return NewEvmAddressTruncating(b)
	case chainid.EcosystemSui:
		return NewSuiAddress(b)
	case chainid.EcosystemSolana:
		return NewSolanaAddress(b)
	case chainid.EcosystemCosmos:
		return NewCosmosAddress(b)
	default:
		return NewGenericAddress(b, e)
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

// NewAddressFromString creates a new Address from a generic string, interpreted according to the ecosystem.
// Hex (with optional '0x') for all chains except Solana, where base58 is used.
func NewAddressFromString(address string, e chainid.Ecosystem) (Address, error) {
	switch e {
	case chainid.EcosystemSolana:
		return NewSolanaAddressFromBase58(address)
	default:
		return NewAddressFromHex(address, e)
	}
}

// GenericAddress does not enforce any ecosystem specific check, leaving to the caller
// the need to check ecosystem requirements
type GenericAddress struct {
	inner     []byte
	ecosystem chainid.Ecosystem
}

func NewGenericAddress(b []byte, e chainid.Ecosystem) (*GenericAddress, error) {
	if len(b) == 0 {
		return nil, ErrEmptyAddress
	}
	a := &GenericAddress{
		inner:     make([]byte, len(b)),
		ecosystem: e,
	}
	copy(a.inner, b)
	return a, nil
}

func NewGenericAddressFromHex(address string, e chainid.Ecosystem) (*GenericAddress, error) {
	b, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))
	if err != nil {
		return nil, fmt.Errorf("%w: hex decoding error %w", ErrBadAddress, err)
	}
	return NewGenericAddress(b, e)
}

// String returns the '0x' led hex encoding of the generic address since no assumption on the ecosystem is available
func (a *GenericAddress) String() string {
	return "0x" + a.Hex()
}

func (a *GenericAddress) Hex() string {
	return hex.EncodeToString(a.inner)
}

func (a *GenericAddress) Bytes() []byte {
	buf := make([]byte, len(a.inner))
	copy(buf, a.inner)
	return buf
}

func (a *GenericAddress) Length() int {
	return len(a.inner)
}

func (a *GenericAddress) Ecosystem() chainid.Ecosystem {
	return a.ecosystem
}

func (a1 *GenericAddress) Equal(a2 Address) bool {
	if a2.Ecosystem() != a1.Ecosystem() {
		return false
	}
	return bytes.Equal(a1.inner, a2.Bytes())
}

func NewZeroAddress(e chainid.Ecosystem) Address {
	switch e {
	case chainid.EcosystemEVM:
		addr, _ := NewEvmAddress(common.Bytes32Zeros[:EvmAddressLength])
		return addr
	default:
		addr, _ := NewAddress(common.Bytes32Zeros, e)
		return addr
	}
}
