package codec

import (
	"encoding/hex"
	"errors"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
)

func (cdc *Codec) DecodeCosmosTransaction(data string) (*txtypes.Tx, string, error) {
	var txHash string
	bz, err := hex.DecodeString(data)
	if err != nil {
		return nil, txHash, err
	}
	txHash = cosmoscommon.NewTxHash(bz).String()
	tx, err := cdc.decodeCosmosTransaction(bz)
	if err != nil {
		return nil, txHash, err
	}

	return tx, txHash, nil
}

func (cdc *Codec) decodeCosmosTransaction(data []byte) (*txtypes.Tx, error) {
	decodedTx, err := cdc.txDecoder(data)
	if err != nil {
		return nil, err
	}

	protoTxProvider, ok := decodedTx.(cosmostypes.ProtoTxProvider)
	if !ok {
		return nil, errors.New("invalid proto tx")
	}

	protoTx := protoTxProvider.GetProtoTx()
	if len(protoTx.GetMsgs()) == 0 {
		return nil, errors.New("msg is nil")
	}

	return protoTx, nil
}
