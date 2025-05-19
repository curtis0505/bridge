package types

import (
	"context"
	validatortypes "github.com/curtis0505/bridge/apps/managers/handler/validator/types"
	"github.com/curtis0505/bridge/apps/managers/types/fxportal"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge"
)

type VerifyHandler interface {
	VerifyDeposit(ctx context.Context, tx *types.Transaction, log types.Log, event bridge.EventDeposit) error
	VerifyBurn(ctx context.Context, tx *types.Transaction, log types.Log, event bridge.EventBurn) error

	VerifyWithdraw(ctx context.Context, tx *types.Transaction, log types.Log, event bridge.EventWithdraw) error
	VerifyMint(ctx context.Context, tx *types.Transaction, log types.Log, event bridge.EventMint) error

	VerifyExecution(ctx context.Context, log types.Log, event bridge.EventExecution) error
	VerifySubmission(ctx context.Context, log types.Log, event bridge.EventSubmission) error
}

type ValidatorHandler interface {
	ConfirmTx(tx *types.Transaction)
	AddCosmosMultiSigTx(tx validatortypes.CosmosMultiSigTx)
}

type FxPortalHandler interface {
	RegisterPendingWithdrawTx(withdrawPendingTx fxportal.WithdrawPendingTx)
}
