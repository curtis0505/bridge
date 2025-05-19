package verify

import (
	"context"
	"fmt"
	verifytypes "github.com/curtis0505/bridge/apps/managers/handler/verify/types"
	"github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/types/bridge/abi"
)

func (p *VerifyHandler) VerifyExecution(ctx context.Context, log types.Log, event bridgetypes.EventExecution) error {
	txResp, err := p.client.CallMsg(ctx, log.Chain(), "", log.Address(), "transactions", abi.GetAbiToMap(abi.MultiSigWalletAbi), event.TransactionId)
	if err != nil {
		p.logger.Error(err.Error())
		return err
	}
	p.logger.Info("event", "VerifyExecution", "chain", log.Chain(), "txId", event.TransactionId, "txHash", log.TxHash())

	tx := verifytypes.NewMultiSigTransaction(txResp)
	if !tx.Executed {
		return fmt.Errorf("multisig: not executed tx: %s", log.TxHash())
	}

	return nil
}

func (p *VerifyHandler) VerifySubmission(ctx context.Context, log types.Log, event bridgetypes.EventSubmission) error {
	p.logger.Info("event", "VerifySubmission", "chain", "TODO", "txId", event.TransactionId)
	return nil
}
