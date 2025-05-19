package types

type AddressFormat string

const (
	AddressFormatUnknown = AddressFormat("UNKNOWN")
	AddressFormatEVM     = AddressFormat("EVM")
	AddressFormatATOM    = AddressFormat("ATOM")
	AddressFormatFNSA    = AddressFormat("FNSA")
	AddressFormatTFNSA   = AddressFormat("TFNSA")
	AddressFormatTRX     = AddressFormat("TRX")
)

var (
	SupportedAddressFormats = []AddressFormat{
		AddressFormatEVM,
		AddressFormatATOM,
		AddressFormatFNSA,
		AddressFormatTFNSA,
		AddressFormatTRX,
	}
)

func GetPrimaryChainByAddressFormat(format AddressFormat) Chain {
	switch format {
	case AddressFormatEVM:
		return ChainETH
	case AddressFormatFNSA:
		return ChainFNSA
	case AddressFormatTFNSA:
		return ChainTFNSA
	case AddressFormatATOM:
		return ChainATOM
	case AddressFormatTRX:
		return ChainTRX
	default:
		return ""
	}
}
