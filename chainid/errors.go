package chainid

import "fmt"

var ErrChainIdInvalid = fmt.Errorf("invalid chain id")

func NewErrChainIdInvalid(reason error) error {
	return fmt.Errorf("%w: %w", ErrChainIdInvalid, reason)
}

var ErrLength = fmt.Errorf("wrong length")

func NewErrLenght(expected int, actual int) error {
	return NewErrChainIdInvalid(fmt.Errorf(
		"%w: expected %d bytes, got %d",
		ErrLength, expected, actual,
	))
}

var ErrUnsupportedEcosystem = fmt.Errorf("unsupported ecosystem")

func NewErrUnsupportedEcosystem(e byte) error {
	return NewErrChainIdInvalid(
		fmt.Errorf("%w: %d", ErrUnsupportedEcosystem, e),
	)
}
