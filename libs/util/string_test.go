package util

import (
	"fmt"
	"testing"
)

func Test_TxHash_Abbreviation(t *testing.T) {
	transactionHash := "A7A7A9E7B4B4B4A2A2A2B8B8B8C3C3C3D1D1D1E17FF8"
	abbreviation := AbbreviateTxHash(transactionHash, 8, 6)
	fmt.Printf("%s\n", abbreviation)

	transactionHash2 := ""
	abbreviation2 := AbbreviateTxHash(transactionHash2, 8, 6)
	fmt.Printf("%s\n", abbreviation2)

	transactionHash3 := "A7A7A9E7B4B4B4A2A2A2B8B8B8C3C3C3D1D1D1E17FF8"
	abbreviation3 := AbbreviateTxHash(transactionHash3, 100, 6)
	fmt.Printf("%s\n", abbreviation3)

	transactionHash4 := "A7A7A9E7B4B4B4A2A2A2B8B8B8C3C3C3D1D1D1E17FF8"
	abbreviation4 := AbbreviateTxHash(transactionHash4, 1, 3)
	fmt.Printf("%s\n", abbreviation4)

	transactionHash5 := "A7A7A9E7B4B4B4A2A2A2B8B8B8C3C3C3D1D1D1E17FF8"
	abbreviation5 := AbbreviateTxHash(transactionHash5, 0, 5)
	fmt.Printf("%s\n", abbreviation5)
}
