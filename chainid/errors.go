package chainid

import "fmt"

var ErrLChainIdInvalid = fmt.Errorf("invalid chain id")

func NewErrLChainIdInvalid(reason error) error {
	return fmt.Errorf("%w: %w", ErrLChainIdInvalid, reason)
}

var ErrLength = fmt.Errorf("wrong length")

func NewErrLength(expected int, actual int) error {
	return NewErrLChainIdInvalid(fmt.Errorf(
		"%w: expected %d bytes, got %d",
		ErrLength, expected, actual,
	))
}

var ErrUnsupportedEcosystem = fmt.Errorf("unsupported ecosystem")

func NewErrUnsupportedEcosystem(e byte) error {
	return NewErrLChainIdInvalid(
		fmt.Errorf("%w: %d", ErrUnsupportedEcosystem, e),
	)
}
