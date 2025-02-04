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
		return fmt.Sprintf("unsupported ecosystem %d", t)
	}
}

// IsSupported reports whether the Ecosystem is among the supported ones.
func (t Ecosystem) IsSupported() bool {
	switch t {
	case EcosystemEVM:
	case EcosystemSui:
	case EcosystemBitcoin:
	default:
		return false
	}
	return true
}

func (t Ecosystem) ToEcosystemHexByte() string {
	return hex.EncodeToString([]byte{byte(t)})
}

type ChainId interface {
	// String returns the hex encoding of the DestinationChainId with leading 0x
	String() string

	// Hex returns the hex encoding of the DestinationChainId without leading 0x
	Hex() string

	// Bytes returns a copy of the bytes of the ChainId (BigEndian)
	Bytes() []byte

	// Ecosystem returns the Ecosystem the ChainId belongs to.
	Ecosystem() Ecosystem

	// Equal reports whether a and be are the same
	Equal(b ChainId) bool
}

// NewChainId creates a new ChainId instance by accepting the bytes of the chain Id encoded
// in Big Endian. Function returns an error if the chain Id is invalid or unsupported.
func NewChainId(in []byte) (ChainId, error) {
	if err := ValidateChainIdFromBytes(in); err != nil {
		return nil, NewErrChainIdInvalid(err)
	}
	out := make([]byte, ChainIdLength)
	copy(out, in)
	id := chainId{
		inner: out,
	}
	switch id.Ecosystem() {
	case EcosystemEVM:
		return EVMChainId{
			chainId: id,
		}, nil
	case EcosystemSui:
		return SuiChainId{
			chainId: id,
		}, nil
	case EcosystemBitcoin:
		return BitcoinChainId{
			chainId: id,
		}, nil
	default:
		return nil, fmt.Errorf("chainid package is broken, validate accepted an unsupported chain")
	}
}

// NewChainIdFromHex creates a new ChainId instance by accepting an hex string of the chain Id.
// Hex string can be passed both with and without the leading 0x.
// Function returns an error if the chain Id is invalid or unsupported.
func NewChainIdFromHex(s string) (ChainId, error) {
	decoded, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil {
		return nil, NewErrChainIdInvalid(err)
	}
	return NewChainId(decoded)
}

// chainId is the base implementation providing basic feature all chain ids implement
type chainId struct {
	inner []byte
}

// NewChainId creates a new ChainId instance by accepting the bytes of the chain Id encoded
// in Big Endian. Function returns an error if the chain Id is invalid or unsupported.
func newChainId(in []byte) (*chainId, error) {
	if err := ValidateChainIdFromBytes(in); err != nil {
		return nil, NewErrChainIdInvalid(err)
	}
	out := make([]byte, ChainIdLength)
	copy(out, in)
	return &chainId{
		inner: out,
	}, nil
}

// NewChainIdFromHex creates a new ChainId instance by accepting an hex string of the chain Id.
// Hex string can be passed both with and without the leading 0x.
// Function returns an error if the chain Id is invalid or unsupported.
func newChainIdFromHex(s string) (*chainId, error) {
	decoded, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
	if err != nil {
		return nil, NewErrChainIdInvalid(err)
	}
	return newChainId(decoded)
}

// String returns the hex encoding of the DestinationChainId with leading 0x
func (a chainId) String() string {
	return "0x" + a.Hex()
}

// Hex returns the hex encoding of the DestinationChainId without leading 0x
func (a chainId) Hex() string {
	return hex.EncodeToString(a.inner)
}

// Bytes returns a copy of the bytes of the ChainId (BigEndian)
func (a chainId) Bytes() []byte {
	out := make([]byte, ChainIdLength)
	copy(out, a.inner)
	return out
}

// Ecosystem returns the Ecosystem the ChainId belongs to.
func (a chainId) Ecosystem() Ecosystem {
	// Saved in big endian so MSB is in position 0
	return Ecosystem(a.inner[0])
}

// Equal reports whether a and be are the same
func (a chainId) Equal(b ChainId) bool {
	return bytes.Equal(a.inner, b.Bytes())
}

// ValidateChainIdFromBytes validates if a slice of bytes can be used to create a valid
// ChainId instance. Ecosystem support is also verified.
func ValidateChainIdFromBytes(chainIdBytes []byte) error {
	if len(chainIdBytes) != ChainIdLength {
		return NewErrLenght(ChainIdLength, len(chainIdBytes))
	}
	if !Ecosystem(chainIdBytes[0]).IsSupported() {
		return NewErrUnsupportedEcosystem(chainIdBytes[0])
	}
	return nil
}
