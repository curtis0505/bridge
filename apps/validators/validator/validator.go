package validator

import (
	"fmt"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/logger"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"strings"
)

type Validator struct {
	cfg      conf.Config
	client   *chain.Client
	logger   *logger.Logger
	Account  map[string]*commontypes.Account
	Contract bridgetypes.ContractAddresses

	logEvent           map[string]func(commontypes.Log) error
	PendingTransaction PendingTransaction
}

func New(client *chain.Client, cfg conf.Config, account map[string]*commontypes.Account) (*Validator, error) {
	validator := Validator{
		cfg:                cfg,
		client:             client,
		logger:             logger.NewLogger("Validator"),
		Account:            account,
		Contract:           make(bridgetypes.ContractAddresses),
		logEvent:           make(map[string]func(commontypes.Log) error),
		PendingTransaction: NewPendingTransaction(),
	}

	for chainSymbol, contract := range cfg.Contract {
		validator.Contract[commontypes.Chain(strings.ToUpper(chainSymbol))] = &bridgetypes.ContractAddress{
			NptToken:              contract.NptToken,
			Vault:                 contract.Vault,
			Minter:                contract.Minter,
			MultiSigWallet:        contract.MultiSigWallet,
			RestakeVault:          contract.RestakeVault,
			RestakeMinter:         contract.RestakeMinter,
			RestakeMultiSigWallet: contract.RestakeMultiSigWallet,
		}
	}

	validator.registerEvent()

	return &validator, nil
}

func (p *Validator) registerEvent() {
	p.logEvent[bridgetypes.EventNameDeposit] = p.Deposit
	p.logEvent[bridgetypes.EventNameDepositRestakeToken] = p.DepositRestakeToken
	p.logEvent[bridgetypes.EventNameWithdraw] = p.Withdraw
	p.logEvent[bridgetypes.EventNameBurn] = p.Burn
	p.logEvent[bridgetypes.EventNameMint] = p.Mint
}

func (p *Validator) LogHandler(log commontypes.Log) error {
	logEvent, ok := p.logEvent[bridgetypes.GetEventType(log.EventName)]
	if !ok {
		return fmt.Errorf("not found log handler : %s", log.EventName)
	} else {
		return logEvent(log)
	}
}
