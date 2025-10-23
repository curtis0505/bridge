package verify

import (
	"context"
	"github.com/curtis0505/bridge/apps/managers/conf"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/apps/managers/util"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/service"
	bridge "github.com/curtis0505/bridge/libs/types"
	"github.com/gin-gonic/gin"
	"math/big"
)

var (
	_ types.Handler = &VerifyHandler{}
)

type VerifyHandler struct {
	cfg    conf.Config
	client *chain.Client
	logger *logger.Logger

	historyService *service.HistoryService

	vaultBalance map[string]*big.Int
	minterSupply map[string]*big.Int
}

func New(cfg conf.Config, client *chain.Client) *VerifyHandler {
	verify := VerifyHandler{
		cfg:    cfg,
		client: client,
		logger: logger.NewLogger("Verify"),

		vaultBalance: make(map[string]*big.Int),
		minterSupply: make(map[string]*big.Int),

		historyService: service.GetRegistry().HistoryService(),
	}

	verify.initBalance()

	return &verify
}

func (p *VerifyHandler) Name() string { return "Verify" }

func (p *VerifyHandler) ApiHandler(e *gin.Engine) {
}

func (p *VerifyHandler) LogHandler(log bridge.Log) error {
	return nil
}

func (p *VerifyHandler) initBalance() {
	balanceList, err := p.historyService.AggregateLatestBridgeBalanceHistory(context.Background())
	if err != nil {
		p.logger.Error("event", "initBalance", "err", err)
		return
	}

	for _, history := range balanceList {
		tokenInfo, err := cache.TokenCache().GetTokenInfoByCurrencyID(history.CurrencyID)
		if err != nil {
			p.logger.Error("event", "initBalance", "err", err)
			continue
		}

		var balance string
		switch history.Name {
		case "Vault":
			p.vaultBalance[history.Chain+history.CurrencyID] = history.Balance
			balance = util.ToEtherWithDecimal(p.vaultBalance[history.Chain+history.CurrencyID], tokenInfo.Decimal).String()
		case "Minter":
			p.minterSupply[history.Chain+history.CurrencyID] = history.TotalSupply
			balance = util.ToEtherWithDecimal(p.minterSupply[history.Chain+history.CurrencyID], tokenInfo.Decimal).String()
		}

		p.logger.Info(
			"event", "initBalance",
			"chain", history.Chain,
			"name", history.Name,
			"currencyId", history.CurrencyID,
			"balance", balance,
		)
	}
}
