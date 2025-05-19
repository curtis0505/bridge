package test_code

import (
	"context"
	"encoding/json"
	"fmt"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/libs/client/chain"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/elog"
	"github.com/curtis0505/bridge/libs/testutil"
	commontypes "github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge/abi"
	"github.com/curtis0505/bridge/libs/types/cosmos/bridge"
	"github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

func Test_BalanceAt(t *testing.T) {
	ctx := context.Background()
	address := "0x5Cb2594a92307c840aDDE612588653c4a3084413"

	config, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}

	client := new(chain.Client)
	for _, chainClient := range config.Client {
		err := client.AddClient(chainClient)
		if err != nil {
			panic(err)
		}
	}
	balance, err := client.BalanceAt(ctx, "KLAY", address, nil)
	if err != nil {

		return
	}

	t.Log("address", address, "balance", balance)
}

// Allowance
func Test_Allowance(t *testing.T) {
	ctx := context.Background()
	chainSymbol := "KLAY"
	tokenAddress := "0xA0668c2eedb4fE91d87F6a67c01C118b077fe80F"
	owner := "0xb8882ab9B22Eed4d3A464655b379036987072a58"
	spender := "0x79e6411ddd1fae66e7ea5a379f8c52cf6d0ddf1a"

	config, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}

	client := chain.NewClientByConfig(config.Client)

	var erc20ABI []map[string]interface{}
	err = json.Unmarshal([]byte(abi.ERC20Abi), &erc20ABI)
	if err != nil {
		return
	}

	msg, err := client.CallMsg(
		ctx,
		chainSymbol,
		"",
		tokenAddress,
		"allowance",
		erc20ABI,
		common.HexToAddress(chainSymbol, owner),
		common.HexToAddress(chainSymbol, spender),
	)
	if err != nil {
		return
	}

	allowance := msg[0].(*big.Int)

	t.Log("owner", owner, "spender", spender, "allowance", allowance)
}

// increaseAllowance
func Test_increaseAllowance(t *testing.T) {

	ctx := context.Background()
	chainSymbol := "BASE" //KLAY, MATIC, ETH
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(600000))

	contract := getContractAddress(testutil.DEV, chainSymbol)
	account := getUserPrivKey("lku", testutil.DEV, "")
	spender := contract.Vault // vault or minter

	config, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}

	userAccount, err := commontypes.NewAccountFromPK(account.PrivKey)
	if err != nil {
		t.Error("NewAccountFromPK", err)
		return
	}

	client := chain.NewClientByConfig(config.Client)

	data, err := commontypes.PackAbi(
		chainSymbol,
		abi.GetAbiToMap(abi.ERC20Abi),
		"increaseAllowance",
		common.HexToAddress(chainSymbol, spender),
		amount,
	)
	if err != nil {
		t.Error("PackWithChain", err)
		return
	}

	option, err := client.GetTransactionOption(ctx, chainSymbol, userAccount.Address)
	if err != nil {
		t.Error("GetTransactionOption", err)
		return
	}

	tx, err := client.GetTransactionData(chainSymbol, &commontypes.RequestTransaction{
		From:      userAccount.Address,
		To:        contract.Token,
		Nonce:     option.Nonce,
		GasPrice:  option.GasPrice,
		GasFeeCap: option.GasFeeCap,
		GasTipCap: option.GasTipCap,
		GasLimit:  commontypes.GasLimit,
		Value:     big.NewInt(0),
		Data:      data,
	})
	if err != nil {
		t.Error("GetTransactionData", err)
		return
	}

	chainId, err := client.GetChainID(ctx, chainSymbol)
	if err != nil {
		t.Error("GetChainID", err)
		return
	}

	signedTx, err := userAccount.Sign(tx, chainId)
	if err != nil {
		t.Error("Sign", err)
		return
	}

	result, err := client.RawSendTxAsyncByTx(ctx, chainSymbol, signedTx)
	if err != nil {
		t.Error("RawSendTxAsync", err)
		return
	}

	t.Log(result.TxHash.String())
}

// Vault Deposit

// FNSA Deposit
func Test_Deposit_FNSA(t *testing.T) {
	ctx := context.Background()
	toChain := "KLAY"
	fromAddress := "link1tnfhl7gh5g2n982zwcuypy4sxfaluaq6q7px3z"
	toAddress := "0xb8882ab9B22Eed4d3A464655b379036987072a58"
	toTokenAddress := "0x93de1ec8a5dab7cdf31e013837391cc2dff356b3"                    // DQ - "0xb0c0a345b3bf609e8f5f33d5e0d2a908e03157f0" // nFNSA
	vaultAddress := "link1sqctf7u55yhu4s00d4f33nj7k2a09tjx4x0ev74u83zvpqkhvfqsvqqnpz" //  "link1fvw0rt94gl5eyeq36qdhj5x7lunv3xpuqcjxa0llhdssvqtcmrnq8g5x2c"
	//amount := big.NewInt(1e6)

	userPrivateKey := "983ca545f79478ce7ecdfd5d9eb0abb007b67d753abb4eff702e17a3f310369f"
	userAccount, err := commontypes.NewAccountFromPK(userPrivateKey)
	if err != nil {
		t.Error("NewAccountFromPK", err)
		return
	}

	config, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}

	client := chain.NewClientByConfig(config.Client)

	for _, cosmosClient := range client.CosmosClients() {
		cosmosClient.SetAccountPrefix()

		//fmt.Println("userAccount.Address", userAccount.Address)
		at, err := cosmosClient.BalanceAt(ctx, fromAddress, nil)
		if err != nil {
			t.Error("BalanceAt", err)
			continue
		}
		fmt.Println("at", at)
		minGasPrice, err := cosmosClient.GetMinimumGasPrice(ctx)
		if err != nil {
			t.Error("GetMinimumGasPrice", err)
			continue
		}
		gasLimit := int64(2_000_000)
		gasPrice := util.ToDecimal(minGasPrice, 0).Mul(decimal.NewFromInt(gasLimit)).Round(0).IntPart()
		//txBuilder.SetFeeAmount(cosmossdk.NewCoins(cosmossdk.NewInt64Coin(token.DenomByChain(cosmosClient.Chain()), gasPrice)))
		transaction, err := cosmosClient.SendTransaction(
			ctx,
			userAccount,
			cosmostypes.WithMsgs(
				cosmostypes.NewMsgExecuteContract(
					fromAddress,
					vaultAddress,
					bridge.MsgDepositCoin{
						ChildChainName: toChain,
						ChildTokenAddr: toTokenAddress,
						ToAddr:         toAddress,
					},
					cosmossdk.NewCoins(cosmossdk.NewInt64Coin(token.DenomByChain(cosmosClient.Chain()), 4000000)),
				),
			),
			cosmostypes.WithFeeAmount(cosmossdk.NewInt64Coin(token.DenomByChain(cosmosClient.Chain()), gasPrice)),
			cosmostypes.WithBroadcastMode(txtypes.BroadcastMode_BROADCAST_MODE_BLOCK),
		)
		if err != nil {
			t.Error("SendTransaction", err)
			continue
		}
		t.Log(transaction.TxHash.String())
		t.Log(transaction.Result)
		t.Log(transaction.Hash)
		t.Log(transaction.Error)

		break
	}
}

type Contract struct {
	Token  string
	Vault  string
	Minter string
}

func getContractAddress(env testutil.TestENV, chain string) Contract {
	switch env {
	case testutil.DEV:
		switch chain {
		case commontypes.ChainKLAY: //kairos
			return Contract{
				Token:  "0xa0668c2eedb4fe91d87f6a67c01c118b077fe80f",
				Vault:  "0x79e6411DdD1FaE66e7eA5a379F8c52cf6D0ddF1a",
				Minter: "0x0c3Ea2587eC105CA447e7474Db6AEA2a996C1546",
			}
		case commontypes.ChainMATIC: // amoy
			return Contract{
				Token:  "0x9721ad50abc816531afbb3fed3530e9c124d176e",
				Vault:  "0x867545B6487543a5b195d8B2C32D115403e5bcF0",
				Minter: "0x3E957AFdd8Cc382528161eE1a2a8B1c2c99DFDe1",
			}
		case commontypes.ChainETH: //sepolia
			return Contract{
				Token:  "0x9721ad50abc816531afbb3fed3530e9c124d176e",
				Vault:  "0x16f9599359325b5A4630E1D16F02E472A245f010",
				Minter: "0xa3753fc45c4373C2f0583E5593252A6b3375283c",
			}
		case commontypes.ChainBASE: //sepolia
			return Contract{
				Token:  "0x2e1b23db8d75b0b220e3632ad828f5e662b0d0a8",
				Vault:  "0x28209a5cE46a691D5c1b2ED920b3ef0F5965C0c5",
				Minter: "0xea7b95a8ca34c98E87dcFe95ff147b35F2Efb162",
			}
		}
	case testutil.DEV2:
		switch chain {
		case commontypes.ChainKLAY: //kairos
			return Contract{
				Token:  "0xa0668c2eedb4fe91d87f6a67c01c118b077fe80f",
				Vault:  "0x4f9748ca344ba98f44318f10a2184f3dd6d44601",
				Minter: "0x1b371ec0cd411dff250f86c435cb7eabc8a1b167",
			}
		case commontypes.ChainMATIC: //amoy
			return Contract{
				Token:  "0x9721ad50abc816531afbb3fed3530e9c124d176e",
				Vault:  "0x5b8e7021762d6a283269c7b8f731595851d68d49",
				Minter: "0xf2e76a26da76bc17c96188a6cabb5eb14644cd36",
			}
		case commontypes.ChainETH: //holsky
			return Contract{
				Token:  "0x7a5d7eab9c3ea0d9c3cdd981200693c23e9e539a",
				Vault:  "0x2453eed7c0e2d5dc17bdfc833398441ee78d8a4f",
				Minter: "0xff3a8e1c1683e0bac0ddff2fc4aa9c0b101402ce",
			}
		}
	default:
	}

	return Contract{}
}

type User struct {
	PrivKey   string
	ToAddress []byte
}

func getUserPrivKey(name string, env testutil.TestENV, toChain string) *User {
	switch name {
	case "youngho":
		switch env {
		case testutil.DEV:
			user := &User{
				PrivKey:   "fd524812db2d92ec5a27c35e402c48044bfba644494ccbeff3efee2960163068",
				ToAddress: common.HexToAddress(toChain, "0xbD1649F9a1599301639d6525Bb0C87e5081e28d3").Bytes(),
			}

			if toChain == commontypes.ChainTFNSA || toChain == commontypes.ChainFNSA {
				user.ToAddress = common.HexToAddress(toChain, "link1lgrnfcmkqy023wfgdlykykk4497h2rlxvxjtwk").Bytes()
			}

			return user
		case testutil.DQ:
			user := &User{
				PrivKey:   "58d7ad6ece2c88ce01648aaf0799e27eb760e37f11d30b937f36529bdee926db",
				ToAddress: common.HexToAddress(toChain, "0xaD1fd213d73932c049974676cBE0D0a2efB383C6").Bytes(),
			}

			if toChain == commontypes.ChainTFNSA || toChain == commontypes.ChainFNSA {
				user.ToAddress = common.HexToAddress(toChain, "tlink15vfaymegl4kc8rau32gz6ezz7u4uhkcxveu2ud").Bytes()
			}

			return user
		}
	case "lku":
		switch env {
		case testutil.DEV:
			user := &User{
				PrivKey:   "d4fcd9aa16d3df5ca4a7746abf0292bb6334105141711d0b44856b8ed67259c4",
				ToAddress: common.HexToAddress(toChain, "0xb8882ab9B22Eed4d3A464655b379036987072a58").Bytes(),
			}

			if toChain == commontypes.ChainTFNSA || toChain == commontypes.ChainFNSA {
				user.ToAddress = common.HexToAddress(toChain, "link1tnfhl7gh5g2n982zwcuypy4sxfaluaq6q7px3z").Bytes()
			}

			return user
		case testutil.DQ:
			user := &User{
				PrivKey:   "",
				ToAddress: common.HexToAddress(toChain, "").Bytes(),
			}

			if toChain == commontypes.ChainTFNSA || toChain == commontypes.ChainFNSA {
				user.ToAddress = common.HexToAddress(toChain, "").Bytes()
			}
		}
	}
	return nil
}

// Vault Deposit Token
func Test_DepositToken(t *testing.T) {
	ctx := context.Background()
	env := testutil.DEV

	fromChain := "KLAY"
	toChain := "BASE"
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(3))

	contract := getContractAddress(env, fromChain)
	account := getUserPrivKey("lku", env, toChain)

	config, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}

	elog.InitLog(config.Log)

	userAccount, err := commontypes.NewAccountFromPK(account.PrivKey)
	if err != nil {
		t.Error("NewAccountFromPK", err)
		return
	}

	client := chain.NewClientByConfig(config.Client)

	data, err := commontypes.PackAbi(
		fromChain,
		abi.GetAbiToMap(abi.VaultAbi),
		"depositToken",
		common.HexToAddress(fromChain, contract.Token),
		toChain,
		account.ToAddress,
		amount,
		[]byte{},
	)
	if err != nil {
		t.Error("PackWithChain", err)
		return
	}

	option, err := client.GetTransactionOption(ctx, fromChain, userAccount.Address)
	if err != nil {
		t.Error("GetTransactionOption", err)
		return
	}

	var vaultABI []map[string]interface{}
	err = json.Unmarshal([]byte(abi.VaultAbi), &vaultABI)
	if err != nil {
		return
	}

	msg, err := client.CallMsg(ctx,
		fromChain,
		"",
		contract.Vault,
		"getChainId",
		vaultABI,
		toChain,
	)
	if err != nil {
		return
	}

	toChainID := msg[0].([32]byte)

	msg, err = client.CallMsg(ctx,
		fromChain,
		"",
		contract.Vault,
		"getChainFee",
		vaultABI,
		toChainID,
	)
	if err != nil {
		return
	}
	chainFee := msg[0].(*big.Int)

	tx, err := client.GetTransactionData(fromChain, &commontypes.RequestTransaction{
		From:      userAccount.Address,
		To:        contract.Vault,
		Nonce:     option.Nonce,
		GasPrice:  option.GasPrice,
		GasFeeCap: option.GasFeeCap,
		GasTipCap: option.GasTipCap,
		GasLimit:  commontypes.GasLimit,
		Value:     chainFee,
		Data:      data,
	})
	if err != nil {
		t.Error("GetTransactionData", err)
		return
	}

	chainId, err := client.GetChainID(ctx, fromChain)
	if err != nil {
		t.Error("GetChainID", err)
		return
	}

	signedTx, err := userAccount.Sign(tx, chainId)
	if err != nil {
		t.Error("Sign", err)
		return
	}

	at, err := client.BalanceAt(ctx, fromChain, userAccount.Address, nil)
	if err != nil {
		return
	}
	t.Log("balance", at, "chainId", chainId, "fromChain", fromChain)

	result, err := client.RawSendTxAsyncByTx(ctx, fromChain, signedTx)
	if err != nil {
		t.Error("RawSendTxAsync", err)
		return
	}

	t.Log(result.TxHash.String())
}

// Vault DepositRestakeToken
func Test_DepositRestakeToken(t *testing.T) {
	ctx := context.Background()
	fromChain := "ETH"
	toChain := "KLAY"
	tokenAddress := "0xAd98FB3967B9A2962c2051dA00bF5C8DF6454b4c"
	vaultAddress := "0xf665b3cb701ce76346f79d41fb86960621d4ed27"
	userPrivateKey := "d4fcd9aa16d3df5ca4a7746abf0292bb6334105141711d0b44856b8ed67259c4" // 0xb8882ab9B22Eed4d3A464655b379036987072a58
	amount := new(big.Int).Mul(big.NewInt(1e16), big.NewInt(3))

	config, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}

	userAccount, err := commontypes.NewAccountFromPK(userPrivateKey)
	if err != nil {
		t.Error("NewAccountFromPK", err)
		return
	}

	client := chain.NewClientByConfig(config.Client)

	data, err := commontypes.PackAbi(
		fromChain,
		abi.GetAbiToMap(abi.RestakeVaultAbi),
		"depositRestakeToken",
		common.HexToAddress(fromChain, tokenAddress),
		toChain,
		common.HexToAddress(toChain, userAccount.Address).Bytes(),
		[]byte{},
	)
	if err != nil {
		t.Error("PackWithChain", err)
		return
	}

	option, err := client.GetTransactionOption(ctx, fromChain, userAccount.Address)
	if err != nil {
		t.Error("GetTransactionOption", err)
		return
	}

	var vaultABI []map[string]interface{}
	err = json.Unmarshal([]byte(abi.RestakeVaultAbi), &vaultABI)
	if err != nil {
		return
	}

	msg, err := client.CallMsg(ctx,
		fromChain,
		"",
		vaultAddress,
		"getChainId",
		vaultABI,
		toChain,
	)
	if err != nil {
		return
	}

	toChainID := msg[0].([32]byte)

	msg, err = client.CallMsg(ctx,
		fromChain,
		"",
		vaultAddress,
		"getChainFee",
		vaultABI,
		toChainID,
	)
	if err != nil {
		return
	}
	chainFee := msg[0].(*big.Int)

	tx, err := client.GetTransactionData(fromChain, &commontypes.RequestTransaction{
		From:      userAccount.Address,
		To:        vaultAddress,
		Nonce:     option.Nonce,
		GasPrice:  option.GasPrice,
		GasFeeCap: option.GasFeeCap,
		GasTipCap: option.GasTipCap,
		GasLimit:  commontypes.GasLimit,
		Value:     new(big.Int).Add(amount, chainFee),
		Data:      data,
	})
	if err != nil {
		t.Error("GetTransactionData", err)
		return
	}

	chainId, err := client.GetChainID(ctx, fromChain)
	if err != nil {
		t.Error("GetChainID", err)
		return
	}

	signedTx, err := userAccount.Sign(tx, chainId)
	if err != nil {
		t.Error("Sign", err)
		return
	}

	at, err := client.BalanceAt(ctx, fromChain, userAccount.Address, nil)
	if err != nil {
		return
	}
	t.Log("balance", at, "chainId", chainId, "fromChain", fromChain)

	result, err := client.RawSendTxAsyncByTx(ctx, fromChain, signedTx)
	if err != nil {
		t.Error("RawSendTxAsync", err)
		return
	}

	t.Log(result.TxHash.String())
}

// Minter Burn
func Test_Burn(t *testing.T) {
	ctx := context.Background()
	env := testutil.DEV

	fromChain := "MATIC"
	toChain := "KLAY"
	amount := big.NewInt(1e17)

	contract := getContractAddress(testutil.DEV, fromChain)
	account := getUserPrivKey("youngho", env, toChain)

	config, err := conf.NewConfig("../conf/config.toml")
	if err != nil {
		panic(err)
	}

	userAccount, err := commontypes.NewAccountFromPK(account.PrivKey)
	if err != nil {
		t.Error("NewAccountFromPK", err)
		return
	}

	client := chain.NewClientByConfig(config.Client)

	data, err := commontypes.PackAbi(
		fromChain,
		abi.GetAbiToMap(abi.MinterAbi),
		"burn",
		common.HexToAddress(fromChain, contract.Token),
		toChain,
		account.ToAddress,
		amount,
		[]byte{},
	)
	if err != nil {
		t.Error("PackWithChain", err)
		return
	}

	option, err := client.GetTransactionOption(ctx, fromChain, userAccount.Address)
	if err != nil {
		t.Error("GetTransactionOption", err)
		return
	}

	var minterABI []map[string]interface{}
	err = json.Unmarshal([]byte(abi.MinterAbi), &minterABI)
	if err != nil {
		return
	}

	msg, err := client.CallMsg(ctx,
		fromChain,
		"",
		contract.Minter,
		"getChainId",
		minterABI,
		toChain,
	)
	if err != nil {
		return
	}

	toChainID := msg[0].([32]byte)

	msg, err = client.CallMsg(ctx,
		fromChain,
		"",
		contract.Minter,
		"getChainFee",
		minterABI,
		toChainID,
	)
	if err != nil {
		return
	}
	chainFee := msg[0].(*big.Int)
	t.Log("chainFee", chainFee)

	tx, err := client.GetTransactionData(fromChain, &commontypes.RequestTransaction{
		From:      userAccount.Address,
		To:        contract.Minter,
		Nonce:     option.Nonce,
		GasPrice:  option.GasPrice,
		GasFeeCap: option.GasFeeCap,
		GasTipCap: option.GasTipCap,
		GasLimit:  commontypes.GasLimit,
		Value:     chainFee,
		Data:      data,
	})
	if err != nil {
		t.Error("GetTransactionData", err)
		return
	}

	chainId, err := client.GetChainID(ctx, fromChain)
	if err != nil {
		t.Error("GetChainID", err)
		return
	}

	signedTx, err := userAccount.Sign(tx, chainId)
	if err != nil {
		t.Error("Sign", err)
		return
	}

	result, err := client.RawSendTxAsyncByTx(ctx, fromChain, signedTx)
	if err != nil {
		t.Error("RawSendTxAsync", err)
		return
	}

	t.Log(result.TxHash.String())
}
