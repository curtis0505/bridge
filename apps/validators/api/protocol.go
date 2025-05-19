package api

import "math/big"

type RetryTransactionRequest struct {
	TxHash string `json:"txHash" binding:"required"`
}

type SpeedUpTransactionRequest struct {
	TxHash    string   `json:"txHash" binding:"required"`
	GasPrice  *big.Int `json:"gasPrice"`
	GasFeeCap *big.Int `json:"gasFeeCap"`
	GasTip    *big.Int `json:"gasTip"`
}

type CancelTransactionRequest struct {
	Nonce    uint64   `json:"nonce" binding:"required"`
	GasPrice *big.Int `json:"gasPrice" binding:"required"`
	GasTip   *big.Int `json:"gasTip"`
}

type GenerateKeyRequest struct {
	PassName string `json:"passName" binding:"required"`
}

type MultiSigRequest struct {
	Chain         string `json:"chain"`
	TxHash        string `json:"txHash"`
	Address       string `json:"address"`
	AccountNumber uint64 `json:"accountNumber"`
	Sequence      uint64 `json:"sequence"`
}
