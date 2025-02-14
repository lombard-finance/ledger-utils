package address_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/lombard-finance/chain/address"
	"github.com/lombard-finance/chain/chainid"
	"github.com/lombard-finance/chain/common"
)

func TestNewAddress(t *testing.T) {
	evmAddressString := "0x8236a87084f8B84306f72007F36F2618A5634494"
	suiAddressString := "0xbfde966bacd4260852155f7b523ef157f0b75a0e1e8a0784e463c3ef0bb69deb"

	evmAddress, err := address.NewAddressFromHex(evmAddressString, chainid.EcosystemEVM)
	common.AssertNoError(t, err)
	_, ok := evmAddress.(*address.EvmAddress)
	common.AssertTrue(t, ok)

	suiAddress, err := address.NewAddressFromHex(suiAddressString, chainid.EcosystemSui)
	common.AssertNoError(t, err)
	_, ok = suiAddress.(*address.SuiAddress)
	common.AssertTrue(t, ok)
}

func TestEVMAddress(t *testing.T) {
	validAddressString := "0x8236a87084f8B84306f72007F36F2618A5634494"
	anotherValidAddressString := "0xA1Bc65eCf8BC7B2FAA22c53bcC49b0376Da3845A"

	t.Run("should create address from valid hex addresses", func(t *testing.T) {
		addr, err := address.NewEvmAddressFromHex(validAddressString)
		common.AssertNoError(t, err)
		common.EqualStrings(t, strings.ToLower(validAddressString), addr.String())
		equalEcosystem(t, chainid.EcosystemEVM, addr.Ecosystem())
		common.AssertTrue(t, address.EvmAddressLength == addr.Length())
		// Same address without leading 0x
		addrNo0x, err := address.NewEvmAddressFromHex(validAddressString[2:])
		common.AssertNoError(t, err)
		common.AssertTrue(t, addr.Equal(addrNo0x))
		// Check with different addresses
		differentAddr, err := address.NewEvmAddressFromHex(anotherValidAddressString)
		common.AssertNoError(t, err)
		common.AssertFalse(t, differentAddr.Equal(addr))
		// Check address checksum or case does not matter
		noChecksumAddr, err := address.NewEvmAddressFromHex(strings.ToLower(validAddressString))
		common.AssertNoError(t, err)
		common.AssertTrue(t, addr.Equal(noChecksumAddr))
	})

	t.Run("should reject invalid addresses", func(t *testing.T) {
		// shorter address
		_, err := address.NewEvmAddressFromHex(validAddressString[5:])
		common.AssertError(t, err, address.ErrBadAddress, address.ErrBadAddressEvm)
		// longer address
		_, err = address.NewEvmAddressFromHex(validAddressString + anotherValidAddressString)
		common.AssertError(t, err, address.ErrBadAddress, address.ErrBadAddressEvm)
		// invalid hex char
		_, err = address.NewEvmAddressFromHex(validAddressString[:5] + "K" + validAddressString[6:])
		common.AssertError(t, err, address.ErrBadAddress, address.ErrBadAddressEvm)
	})

	t.Run("should create valid addresses from bytes", func(t *testing.T) {
		validAddrBytes, _ := hex.DecodeString(validAddressString[2:])
		addrByte, err := address.NewEvmAddress(validAddrBytes)
		common.AssertNoError(t, err)
		addr, _ := address.NewEvmAddressFromHex(validAddressString)
		common.AssertTrue(t, addr.Equal(addrByte))
	})
}

func TestSuiAddress(t *testing.T) {
	validAddressString := "0xbfde966bacd4260852155f7b523ef157f0b75a0e1e8a0784e463c3ef0bb69deb"
	anotherValidAddressString := "0x3e8e9423d80e1774a7ca128fccd8bf5f1f7753be658c5e645929037f7c819040"

	t.Run("should create address from valid hex addresses", func(t *testing.T) {
		addr, err := address.NewSuiAddressFromHex(validAddressString)
		common.AssertNoError(t, err)
		common.EqualStrings(t, validAddressString, addr.String())
		equalEcosystem(t, chainid.EcosystemSui, addr.Ecosystem())
		common.AssertTrue(t, address.SuiAddressLength == addr.Length())
		// Same address without leading 0x
		addrNo0x, err := address.NewSuiAddressFromHex(validAddressString[2:])
		common.AssertNoError(t, err)
		common.AssertTrue(t, addr.Equal(addrNo0x))
		// Check with different addresses
		differentAddr, err := address.NewSuiAddressFromHex(anotherValidAddressString)
		common.AssertNoError(t, err)
		common.AssertFalse(t, differentAddr.Equal(addr))
		// Check address case does not matter
		noChecksumAddr, err := address.NewSuiAddressFromHex(strings.ToUpper(validAddressString[2:]))
		common.AssertNoError(t, err)
		common.AssertTrue(t, addr.Equal(noChecksumAddr))
	})

	t.Run("should reject invalid addresses", func(t *testing.T) {
		// shorter address
		_, err := address.NewSuiAddressFromHex(validAddressString[5:])
		common.AssertError(t, err, address.ErrBadAddress, address.ErrBadAddressSui)
		// longer address
		_, err = address.NewSuiAddressFromHex(validAddressString + anotherValidAddressString)
		common.AssertError(t, err, address.ErrBadAddress, address.ErrBadAddressSui)
		// arbitrary char
		_, err = address.NewSuiAddressFromHex(validAddressString[:5] + "K" + validAddressString[6:])
		common.AssertError(t, err, address.ErrBadAddress, address.ErrBadAddressSui)
	})
}

func equalEcosystem(t *testing.T, expected chainid.Ecosystem, actual chainid.Ecosystem) {
	if expected != actual {
		t.Errorf("expected: %s actual: %s", expected.String(), actual.String())
	}
}
