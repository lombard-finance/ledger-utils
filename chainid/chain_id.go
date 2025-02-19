package chainid

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)

const ChainIdLength = 32

// Ecosystem identifies the Ecosystem a chain Id belongs to.
type Ecosystem uint8

// Supported Ecosystems. We let the following constants match the MSB of the chain Id
// according to Lombard internal reference.
const (
	EcosystemEVM     Ecosystem = 0
	EcosystemSui     Ecosystem = 1
	EcosystemSolana  Ecosystem = 2
	EcosystemBitcoin Ecosystem = 255
)

func (t Ecosystem) String() string {
	switch t {
	case EcosystemEVM:
		return "evm"
	case EcosystemSui:
		return "sui"
	case EcosystemBitcoin:
		return "bitcoin"
	default:
		return fmt.Sprintf("ecosystem %d", t)
	}
}

// IsSupported reports whether the Ecosystem is among the supported ones, which means a
// specialized ecosystem type is available in the package
func (t Ecosystem) IsSupported() bool {
	switch t {
	case EcosystemEVM:
	case EcosystemSui:
	case EcosystemBitcoin:
	case EcosystemSolana:
	default:
		return false
	}
	return true
}

func (t Ecosystem) ToEcosystemHexByte() string {
	return hex.EncodeToString([]byte{byte(t)})
}

// LChainId is the interface of a Lombard Chain Id, modeling the chain Ids defined according to the
// Lombard internal specification
type LChainId interface {
	// String returns the hex encoding of the DestinationChainId with leading 0x
	String() string

	// Hex returns the hex encoding of the DestinationChainId without leading 0x
	Hex() string

	// Bytes returns a copy of the bytes of the ChainId (BigEndian)
	Bytes() []byte

	// FixedBytes returns a copy of the bytes of the ChainId (BigEndian) as an array rather than slice
	FixedBytes() [ChainIdLength]byte

	// Ecosystem returns the Ecosystem the ChainId belongs to.
	Ecosystem() Ecosystem

	// Equal reports whether a and be are the same
	Equal(b LChainId) bool
}

// NewLChainId creates a new ChainId instance by accepting the bytes of the chain Id encoded
// in Big Endian. Function returns an error if the chain Id is invalid or unsupported.
func NewLChainId(in []byte) (LChainId, error) {
	if err := ValidateChainIdFromBytes(in); err != nil {
		return nil, NewErrLChainIdInvalid(err)
	}
	out := make([]byte, ChainIdLength)
	copy(out, in)
	id := lChainId{
		inner: out,
	}
	switch id.Ecosystem() {
	case EcosystemEVM:
		return EVMLChainId{
			lChainId: id,
		}, nil
	case EcosystemSui:
		return SuiLChainId{
			lChainId: id,
		}, nil
	case EcosystemBitcoin:
		return BitcoinLChainId{
			lChainId: id,
		}, nil
	case EcosystemSolana:
		return SolanaLChainId{
			lChainId: id,
		}, nil
	default:
		return GenericLChainId{
			lChainId: id,
		}, nil
	}
}

// NewLChainIdFromHex creates a new ChainId instance by accepting an hex string of the chain Id.
// Hex string can be passed both with and without the leading 0x.
// Function returns an error if the chain Id is invalid or unsupported.
func NewLChainIdFromHex(s string) (LChainId, error) {
	decoded, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil {
		return nil, NewErrLChainIdInvalid(err)
	}
	return NewLChainId(decoded)
}

// lChainId is the base implementation providing basic feature all chain ids implement
type lChainId struct {
	inner []byte
}

// newLChainId creates a new ChainId instance by accepting the bytes of the chain Id encoded
// in Big Endian. Function returns an error if the chain Id is invalid or unsupported.
func newLChainId(in []byte) (*lChainId, error) {
	if err := ValidateChainIdFromBytes(in); err != nil {
		return nil, NewErrLChainIdInvalid(err)
	}
	out := make([]byte, ChainIdLength)
	copy(out, in)
	return &lChainId{
		inner: out,
	}, nil
}

// newLChainIdFromHex creates a new ChainId instance by accepting an hex string of the chain Id.
// Hex string can be passed both with and without the leading 0x.
// Function returns an error if the chain Id is invalid or unsupported.
func newLChainIdFromHex(s string) (*lChainId, error) {
	decoded, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil {
		return nil, NewErrLChainIdInvalid(err)
	}
	return newLChainId(decoded)
}

// String returns the hex encoding of the DestinationChainId with leading 0x
func (a lChainId) String() string {
	return "0x" + a.Hex()
}

// Hex returns the hex encoding of the DestinationChainId without leading 0x
func (a lChainId) Hex() string {
	return hex.EncodeToString(a.inner)
}

// Bytes returns a copy of the bytes of the ChainId (BigEndian)
func (a lChainId) Bytes() []byte {
	out := make([]byte, ChainIdLength)
	copy(out, a.inner)
	return out
}

// FixedBytes returns a copy of the bytes of the ChainId (BigEndian) as an array rather than slice
func (a lChainId) FixedBytes() [ChainIdLength]byte {
	var out [ChainIdLength]byte
	copy(out[:], a.inner)
	return out
}

// Ecosystem returns the Ecosystem the ChainId belongs to.
func (a lChainId) Ecosystem() Ecosystem {
	// Saved in big endian so MSB is in position 0
	return Ecosystem(a.inner[0])
}

// Equal reports whether a and be are the same
func (a lChainId) Equal(b LChainId) bool {
	return bytes.Equal(a.inner, b.Bytes())
}

// ValidateChainIdFromBytes validates if a slice of bytes can be used to create a valid
// ChainId instance. Ecosystem support is also verified.
func ValidateChainIdFromBytes(chainIdBytes []byte) error {
	if len(chainIdBytes) != ChainIdLength {
		return NewErrLenght(ChainIdLength, len(chainIdBytes))
	}
	return nil
}

// GenericLChainId provides base functionalities of the Lombard Chain Id without any check on the
// supported ecosystem.
type GenericLChainId struct {
	lChainId
}
