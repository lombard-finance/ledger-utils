package chainid

type BitcoinLChainId struct {
	lChainId
}

// NewBitcoinLChainId returns the LChainId for the Bitcoin blockchain
func NewBitcoinLChainId() BitcoinLChainId {
	return BitcoinLChainId{
		lChainId{
			inner: []byte{0xff, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, 0x9c, 0x08, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1, 0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f},
		},
	}
}

// NewBitcoinSignetLChainId returns the LChainId for the Bitcoin Signet blockchain
func NewBitcoinSignetLChainId() BitcoinLChainId {
	return BitcoinLChainId{
		lChainId{
			inner: []byte{0xff, 0x00, 0x00, 0x08, 0x81, 0x98, 0x73, 0xe9, 0x25, 0x42, 0x2c, 0x1f, 0xf0, 0xf9, 0x9f, 0x7c, 0xc9, 0xbb, 0xb2, 0x32, 0xaf, 0x63, 0xa0, 0x77, 0xa4, 0x80, 0xa3, 0x63, 0x3b, 0xee, 0x1e, 0xf6},
		},
	}
}
