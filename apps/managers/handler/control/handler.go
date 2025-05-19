package control

import (
	"context"
	"github.com/curtis0505/bridge/apps/managers/conf"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/elog"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	_ types.Handler = &ControlHandler{}
)

type ControlHandler struct {
	cfg        conf.Config
	client     *chain.Client
	logger     *elog.Logger
	pauseCount atomic.Int32

	account      map[string]*commontypes.Account
	pending      map[string]PendingTx
	pendingLock  *sync.RWMutex
	logEvent     map[string]func(commontypes.Log) error
	logEventLock *sync.RWMutex
}

func New(cfg conf.Config, client *chain.Client) *ControlHandler {
	control := ControlHandler{
		cfg:        cfg,
		client:     client,
		pauseCount: atomic.Int32{},
		logger:     elog.NewLogger("Control"),
		account:    make(map[string]*commontypes.Account),

		pending:      make(map[string]PendingTx),
		pendingLock:  &sync.RWMutex{},
		logEvent:     make(map[string]func(commontypes.Log) error),
		logEventLock: &sync.RWMutex{},
	}

	lo.ForEach(bridgetypes.GetBridgeEVMChains(), func(chain string, _ int) {
		chainID, err := client.GetChainID(context.Background(), chain)
		if err != nil {
			panic(err)
		}

		control.account[chain], err = commontypes.NewAccount(cfg.Account, chainID)
		if err != nil {
			panic(err)
		}
	})

	for chain, account := range control.account {
		control.logger.Info("event", "ControlAccount", "chain", chain, "address", account.Address)
	}

	control.pauseCount.Store(0)
	control.registerEvent()

	go control.iterate()

	return &control
}

func (p *ControlHandler) Name() string { return "Control" }

func (p *ControlHandler) registerEvent() {
	p.logEvent[eventtypes.EventPaused] = p.Paused
	p.logEvent[eventtypes.EventUnpaused] = p.UnPaused
}

func (p *ControlHandler) ApiHandler(e *gin.Engine) {
	g := e.Group("/control")
	g.Use(p.Authorization())

	g.POST("/pause", func(ctx *gin.Context) {
		lo.ForEach(bridgetypes.GetBridgeEVMChains(), func(chain string, _ int) {
			if txHash, err := p.Pause(ctx, chain); err != nil {
				ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
				return
			} else {
				_ = txHash
			}
		})

		ctx.JSON(http.StatusOK, types.NewResponseSuccess())
	})

	g.POST("/pause/:chain", func(ctx *gin.Context) {
		if txHash, err := p.Pause(ctx, strings.ToUpper(ctx.Param("chain"))); err != nil {
			ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
			return
		} else {
			_ = txHash
		}

		ctx.JSON(http.StatusOK, types.NewResponseSuccess())
	})

	g.POST("/unpause", func(ctx *gin.Context) {
		lo.ForEach(bridgetypes.GetBridgeEVMChains(), func(chain string, _ int) {
			if txHash, err := p.Unpause(ctx, chain); err != nil {
				ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
				return
			} else {
				_ = txHash
			}
		})

		ctx.JSON(http.StatusOK, types.NewResponseSuccess())
	})

	g.POST("/unpause/:chain", func(ctx *gin.Context) {
		if txHash, err := p.Unpause(ctx, strings.ToUpper(ctx.Param("chain"))); err != nil {
			ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
			return
		} else {
			_ = txHash
		}

		ctx.JSON(http.StatusOK, types.NewResponseSuccess())
	})
}

func (p *ControlHandler) LogHandler(log commontypes.Log) error {
	p.logEventLock.Lock()
	defer p.logEventLock.Unlock()

	logEvent, ok := p.logEvent[log.EventName]

	if !ok {
		return nil
	} else {
		return logEvent(log)
	}
}

func (p *ControlHandler) iterate() {
	gasFeeTicker := time.NewTicker(time.Second * time.Duration(p.cfg.Control.BridgeGasCheckDuration))
	confirmTicker := time.NewTicker(time.Second)
	for {
		select {
		case <-confirmTicker.C:
			p.CheckConfirmTx()
		case <-gasFeeTicker.C:
			p.CheckGasPrice(commontypes.ChainETH)
		}
	}
}
