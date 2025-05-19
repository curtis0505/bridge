package types

import (
	"time"
)

const (
	// Deprecated: Klaytn
	Klaytn = "KLAY"
	// Deprecated: Polygon
	Polygon = "MATIC"
	// Deprecated: Ethereum
	Ethereum = "ETH"
)

type TxType int

const (
	TxTypeLegacyTransaction TxType = iota
	TxTypeEthereumAccessList
	TxTypeEthereumDynamicFee
	TxTypeKlaytnTransaction
	TxTypeKlaytnFeeDelegatedTransaction
	TxTypeTronTransaction
	TxTypeCosmosTransaction
)

type AccountType string

const (
	KMS        AccountType = "KMS"
	KeyStore   AccountType = "KEYSTORE"
	PrivateKey AccountType = "PRIVATE_KEY"

	KeyStoreECS AccountType = "KEYSTORE_ECS"
)

const (
	ReceiptStatusSuccessful = uint64(1)
)

const GasLimit = 1000000

const (
	RetrySubscription = 60 * 10
	RetryDuration     = time.Second
)

const (
	Android         = "android"
	Ios             = "ios"
	NeopinExtension = "neopin_extension"
)

const (
	RedisKeyAuthAccessTokenAndroid         = "auth"
	RedisKeyAuthAccessTokenIOS             = "auth"
	RedisKeyAuthAccessTokenNeopinExtension = "auth_neopin_extension"
)

const (
	GWei = 1000000000
)

const (
	NeopinVisibleDecimal = 6
	NPTDecimal           = 18
)
