package erc20

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Infura URL
const nodeURL = "https://base.dq.neopin.pmang.cloud" // base sepolia

// Contract address
const contractAddressHex = "0x2e1b23db8d75b0b220e3632ad828f5e662b0d0a8"

func Test_Allowance(t *testing.T) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		return
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddressHex)

	// 계약 인스턴스 생성
	contractInstance, err := NewErc20(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to create contract instance: %v", err)
	}

	privateKey := "d4fcd9aa16d3df5ca4a7746abf0292bb6334105141711d0b44856b8ed67259c4"
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return
	}
	address := crypto.PubkeyToAddress(pk.PublicKey)
	spender := common.HexToAddress("0x37dc564f25763b24e24fba9888cc9647773cf7dd")

	allowance, err := contractInstance.Allowance(nil, address, spender)
	if err != nil {
		return
	}
	fmt.Printf("Allowance: %s\n", allowance)
}

func Test_Approve(t *testing.T) {
	ctx := context.Background()

	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddressHex)

	// 계약 인스턴스 생성
	contractInstance, err := NewErc20(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to create contract instance: %v", err)
	}

	// EOA를 대리 계정으로 설정 (예제에는 상수값을 사용합니다. 실제 환경에서는 개인키를 사용하세요.)
	privateKey := "d4fcd9aa16d3df5ca4a7746abf0292bb6334105141711d0b44856b8ed67259c4"
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return
	}
	address := crypto.PubkeyToAddress(pk.PublicKey)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(84532)) // chainID는 사용중인 이더리움 네트워크에 따라 달라질 수 있습니다.
	if err != nil {
		t.Fatalf("Failed to create authorized transactor: %v", err)
	}
	nonce, err := client.NonceAt(ctx, address, nil)
	if err != nil {
		t.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return
	}

	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = uint64(100000)
	transactOpts.Nonce = big.NewInt(int64(nonce))

	t.Log("transactOpts", transactOpts)

	spender := common.HexToAddress("0xb8882ab9B22Eed4d3A464655b379036987072a58")
	// Deposit 함수 호출
	tx, err := contractInstance.Approve(transactOpts, spender, new(big.Int).Mul(big.NewInt(1e18), big.NewInt(60000)))
	if err != nil {
		t.Fatalf("Failed to execute deposit function: %v", err)
	}

	// 트랜잭션 해시 출력
	fmt.Printf("transaction hash: %s\n", tx.Hash().Hex())
}

func Test_Transfer(t *testing.T) {
	ctx := context.Background()

	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddressHex)

	// 계약 인스턴스 생성
	contractInstance, err := NewErc20(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to create contract instance: %v", err)
	}

	// EOA를 대리 계정으로 설정 (예제에는 상수값을 사용합니다. 실제 환경에서는 개인키를 사용하세요.)
	privateKey := "d4fcd9aa16d3df5ca4a7746abf0292bb6334105141711d0b44856b8ed67259c4"
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return
	}
	address := crypto.PubkeyToAddress(pk.PublicKey)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(84532)) // chainID는 사용중인 이더리움 네트워크에 따라 달라질 수 있습니다.
	if err != nil {
		t.Fatalf("Failed to create authorized transactor: %v", err)
	}
	nonce, err := client.NonceAt(ctx, address, nil)
	if err != nil {
		t.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return
	}

	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = uint64(100000)
	transactOpts.Nonce = big.NewInt(int64(nonce))

	t.Log("transactOpts", transactOpts)

	spender := common.HexToAddress("0xb8882ab9B22Eed4d3A464655b379036987072a58")
	// Deposit 함수 호출
	tx, err := contractInstance.Transfer(transactOpts, spender, new(big.Int).Mul(big.NewInt(1e18), big.NewInt(60000000)))
	if err != nil {
		t.Fatalf("Failed to execute deposit function: %v", err)
	}

	// 트랜잭션 해시 출력
	fmt.Printf("transaction hash: %s\n", tx.Hash().Hex())
}

func Test_BalanceOf(t *testing.T) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddressHex)

	// 계약 인스턴스 생성
	contractInstance, err := NewErc20(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to create contract instance: %v", err)
	}

	balance, err := contractInstance.BalanceOf(nil, common.HexToAddress("0x306ee01a6ba3b4a8e993fa2c1adc7ea24462000c"))
	if err != nil {
		return
	}

	fmt.Println("balance", balance)
}
