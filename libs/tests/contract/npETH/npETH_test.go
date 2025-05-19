package npETH

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Infura URL
const nodeURL = "https://eth-api.neopin.io" // base sepolia

// Contract address
const contractAddressHex = "0x841d3B6660663Ed4B0D9b9EDAEe6642e05A4E182"

func Test_UserWithdrawAmount(t *testing.T) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddressHex)

	// 계약 인스턴스 생성
	contractInstance, err := NewNpETH(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to create contract instance: %v", err)
	}

	length, err := contractInstance.WithdrawReqLogLength(nil)
	if err != nil {
		return
	}

	t.Log("length", length)

	withdrawReq, err := contractInstance.WithdrawReqLog(nil, big.NewInt(162))
	if err != nil {
		return
	}
	t.Logf("%+v\n", withdrawReq)

	users, err := contractInstance.Users(nil, 1, common.HexToAddress("0x7955E5862C1f9AE310de7e658Cbd8Fdb4bd33B99"))
	if err != nil {
		return
	}
	t.Logf("%+v\n", users)
}
