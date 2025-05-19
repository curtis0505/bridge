package nft

const (
	// ERC721 https://eips.ethereum.org/EIPS/eip-721
	ERC721 = 1
	KIP17  = 3
	Erc721 = "ERC-721"
	Kip17  = "KIP-17"

	// ERC1155 https://eips.ethereum.org/EIPS/eip-1155
	ERC1155 = 2
	KIP37   = 4
	//Erc1155 = "ERC-1155" // not support
	//Kip37   = "KIP-37" // not support

	// EIP5192 https://eips.ethereum.org/EIPS/eip-5192
	EIP5192 = 5
	Erc5192 = "ERC-5192"
)

const (
	URI      = "uri"
	TokenURI = "tokenURI"
)

const (
	Maintenance = "nft"
)

const (
	MethodBalanceOf   = "balanceOf"
	MethodTokenURI    = "tokenURI"
	MethodTotalSupply = "totalSupply"
	MethodOwnerOf     = "ownerOf"
)
