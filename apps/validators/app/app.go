package app

import (
	"context"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/apps/validators/validator"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/elog"
	"github.com/curtis0505/bridge/libs/types"
	"runtime/debug"
)

type App struct {
	cfg       conf.Config
	client    *chain.Client
	logger    *elog.Logger
	Validator *validator.Validator
}

func New(client *chain.Client, config conf.Config, validator *validator.Validator) (*App, error) {
	app := &App{
		cfg:       config,
		client:    client,
		Validator: validator,
		logger:    elog.NewLogger("app"),
	}
	return app, nil
}

func (app *App) Run() {
	for _, chainSymbol := range app.client.GetChains() {
		contractAddress, err := app.Validator.Contract.GetChain(types.Chain(chainSymbol))
		if err != nil {
			elog.Error("Run", "method", "GetChain", "err", err)
			return
		}

		err = app.client.Subscribe(context.Background(),
			chainSymbol,
			app.logHandler,
			contractAddress.Vault,
			contractAddress.Minter,
			contractAddress.MultiSigWallet,
			contractAddress.RestakeVault,
			contractAddress.RestakeMinter,
			contractAddress.RestakeMultiSigWallet,
		)
		if err != nil {
			elog.Error("Run", "method", "Subscribe", "err", err)
			return
		}
	}

	go app.Validator.CheckTransactionHistory()
}

func (app *App) logHandler(log types.Log) {

	defer func() {
		if v := recover(); v != nil {
			app.logger.Error("txHash", log.TxHash(), "event", log.EventName, "panic", v, "stack", string(debug.Stack()))
		}
	}()

	err := app.Validator.LogHandler(log)
	if err != nil {
		app.logger.Warn("txHash", log.TxHash(), "event", log.EventName, "LogHandler", err)
	}
}
