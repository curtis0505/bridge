package contractstakingV2

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Infura URL
// const nodeURL = "ws://node2.dq.neopin.pmang.cloud:16200"
const nodeURL = "https://base.dq.neopin.pmang.cloud" // base sepolia

// Contract address
const contractAddress = "0x37dc564f25763b24e24fba9888cc9647773cf7dd"

func Test_Deposit(t *testing.T) {
	ctx := context.Background()

	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddress)

	// 계약 인스턴스 생성
	contractInstance, err := NewContractstakingV2(contractAddress, client)
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
	transactOpts.GasLimit = uint64(1000000)
	transactOpts.Nonce = big.NewInt(int64(nonce))

	t.Log("transactOpts", transactOpts)

	// Deposit 함수 호출
	tx, err := contractInstance.Deposit(transactOpts, big.NewInt(0), big.NewInt(1e18), []byte{})
	if err != nil {
		t.Fatalf("Failed to execute deposit function: %v", err)
	}

	// 트랜잭션 해시 출력
	fmt.Printf("Deposit transaction hash: %s\n", tx.Hash().Hex())

	length, err := contractInstance.PoolLength(nil)
	if err != nil {
		return
	}
	fmt.Printf("Pool length: %d\n", length)

	info, err := contractInstance.PoolInfo(nil, big.NewInt(0))
	if err != nil {
		return
	}
	fmt.Printf("Pool info: %v\n", info)

}

func Test_EnterStaking(t *testing.T) {
	ctx := context.Background()

	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddress)

	// 계약 인스턴스 생성
	contractInstance, err := NewContractstakingV2(contractAddress, client)
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
	transactOpts.GasLimit = uint64(1000000)
	transactOpts.Nonce = big.NewInt(int64(nonce))

	t.Log("transactOpts", transactOpts)

	tx, err := contractInstance.EnterStaking(transactOpts, big.NewInt(1e18), []byte{})
	if err != nil {
		return
	}

	// 트랜잭션 해시 출력
	fmt.Printf("Deposit transaction hash: %s\n", tx.Hash().Hex())

	length, err := contractInstance.PoolLength(nil)
	if err != nil {
		return
	}
	fmt.Printf("Pool length: %d\n", length)

	info, err := contractInstance.PoolInfo(nil, big.NewInt(0))
	if err != nil {
		return
	}
	fmt.Printf("Pool info: %v\n", info)

}

func Test_LeaveStaking(t *testing.T) {
	ctx := context.Background()

	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddress)

	// 계약 인스턴스 생성
	contractInstance, err := NewContractstakingV2(contractAddress, client)
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
	transactOpts.GasLimit = uint64(1000000)
	transactOpts.Nonce = big.NewInt(int64(nonce))

	t.Log("transactOpts", transactOpts)

	tx, err := contractInstance.LeaveStaking(transactOpts, big.NewInt(1e18), []byte{})
	if err != nil {
		return
	}

	// 트랜잭션 해시 출력
	fmt.Printf("Deposit transaction hash: %s\n", tx.Hash().Hex())

	length, err := contractInstance.PoolLength(nil)
	if err != nil {
		return
	}
	fmt.Printf("Pool length: %d\n", length)

	info, err := contractInstance.PoolInfo(nil, big.NewInt(0))
	if err != nil {
		return
	}
	fmt.Printf("Pool info: %v\n", info)

}

func Test_PoolInfo(t *testing.T) {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 스마트 계약 주소
	contractAddress := common.HexToAddress(contractAddress)

	// 계약 인스턴스 생성
	contractInstance, err := NewContractstakingV2(contractAddress, client)
	if err != nil {
		t.Fatalf("Failed to create contract instance: %v", err)
	}

	length, err := contractInstance.PoolLength(nil)
	if err != nil {
		return
	}
	fmt.Printf("Pool length: %d\n", length)

	info, err := contractInstance.PoolInfo(nil, big.NewInt(0))
	if err != nil {
		return
	}
	//fmt.Printf("Pool info: %v\n", info)
	printStructFields(info)
}

// printStructFields 함수는 구조체의 필드 이름과 값을 출력합니다.
func printStructFields(info interface{}) {
	v := reflect.ValueOf(info)
	t := reflect.TypeOf(info)

	fmt.Println("Pool info:")
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("  %s: %v\n", field.Name, value)
	}
}
