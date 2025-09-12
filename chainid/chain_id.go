package chainid

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/lombard-finance/ledger-utils/common"
)

const ChainIdLength = 32

// ChainIdAvailableLength is the amount of bytes available to distinguish a chain within an ecosystem
const ChainIdAvailableLength = 31

// Ecosystem identifies the Ecosystem a chain Id belongs to.
type Ecosystem uint8

// Supported Ecosystems. We let the following constants match the MSB of the chain Id
// according to Lombard internal reference.
const (
	EcosystemEVM      Ecosystem = 0
	EcosystemSui      Ecosystem = 1
	EcosystemSolana   Ecosystem = 2
	EcosystemCosmos   Ecosystem = 3
	EcosystemStarknet Ecosystem = 4
	EcosystemUnknown  Ecosystem = 254
	EcosystemBitcoin  Ecosystem = 255
)

func (t Ecosystem) String() string {
	switch t {
	case EcosystemEVM:
		return "evm"
	case EcosystemSui:
		return "sui"
	case EcosystemSolana:
		return "solana"
	case EcosystemCosmos:
		return "cosmos"
	case EcosystemStarknet:
		return "starknet"
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
	case EcosystemSolana:
	case EcosystemBitcoin:
	case EcosystemCosmos:
	case EcosystemStarknet:
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
	case EcosystemCosmos:
		return CosmosLChainId{
			lChainId: id,
		}, nil
	case EcosystemStarknet:
		return StarknetLChainId{
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
// ChainId instance.
func ValidateChainIdFromBytes(chainIdBytes []byte) error {
	if len(chainIdBytes) != ChainIdLength {
		return NewErrLength(ChainIdLength, len(chainIdBytes))
	}
	return nil
}

var _ LChainId = &GenericLChainId{}
var _ common.GogoprotoCustomType = &GenericLChainId{}

// GenericLChainId provides base functionalities of the Lombard Chain Id without any check on the
// supported ecosystem.
type GenericLChainId struct {
	lChainId
}

func NewGenericLChainId(in []byte) (*GenericLChainId, error) {
	inner, err := newLChainId(in)
	if err != nil {
		return nil, err
	}
	return &GenericLChainId{
		lChainId: *inner,
	}, nil
}

func NewGenericLChainIdFromLChainId(in LChainId) (*GenericLChainId, error) {
	return &GenericLChainId{
		lChainId: lChainId{
			inner: in.Bytes(),
		},
	}, nil
}

// ToEcosystem returns a specialized instance of the LChainId if the ecosystem is supported,
// otherwise it returns ErrUnsupportedEcosystem
func (g *GenericLChainId) ToEcosystem() (LChainId, error) {
	switch g.Ecosystem() {
	case EcosystemEVM:
		return EVMLChainId{
			lChainId: g.lChainId,
		}, nil
	case EcosystemSui:
		return SuiLChainId{
			lChainId: g.lChainId,
		}, nil
	case EcosystemSolana:
		return SolanaLChainId{
			lChainId: g.lChainId,
		}, nil
	case EcosystemCosmos:
		return CosmosLChainId{
			lChainId: g.lChainId,
		}, nil
	case EcosystemStarknet:
		return StarknetLChainId{
			lChainId: g.lChainId,
		}, nil
	case EcosystemBitcoin:
		return BitcoinLChainId{
			lChainId: g.lChainId,
		}, nil
	default:
		return nil, NewErrUnsupportedEcosystem(g.Ecosystem())
	}
}

// Marshal implements common.GogoprotoCustomType.
func (g *GenericLChainId) Marshal() ([]byte, error) {
	return g.Bytes(), nil
}

// MarshalTo implements common.GogoprotoCustomType.
func (g *GenericLChainId) MarshalTo(data []byte) (int, error) {
	if len(data) < ChainIdLength {
		return 0, NewErrLength(ChainIdLength, len(data))
	}
	copy(data, g.inner)
	return ChainIdLength, nil
}

// Size implements common.GogoprotoCustomType.
func (g *GenericLChainId) Size() int {
	return ChainIdLength
}

// Unmarshal implements common.GogoprotoCustomType.
func (g *GenericLChainId) Unmarshal(data []byte) error {
	if len(data) != ChainIdLength {
		return NewErrLength(ChainIdLength, len(data))
	}
	g.inner = make([]byte, ChainIdLength)
	copy(g.inner[:], data)
	return nil
}
