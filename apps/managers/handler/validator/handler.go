package validator

import (
	"context"
	"github.com/curtis0505/bridge/apps/managers/conf"
	validatortypes "github.com/curtis0505/bridge/apps/managers/handler/validator/types"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/model"
	"github.com/curtis0505/bridge/libs/service"
	commontypes "github.com/curtis0505/bridge/libs/types"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

var (
	_ types.Handler = &ValidatorHandler{}
)

type ValidatorHandler struct {
	cfg    conf.Config
	client *chain.Client
	logger *logger.Logger

	cache   *model.Cache
	redSync *redsync.Redsync

	configService  *service.ConfigService
	bridgeService  *service.BridgeService
	redisService   *service.RedisService
	historyService *service.HistoryService

	validatorList    []*validatortypes.ValidatorInfo
	pendingTxMapLock *sync.RWMutex
	pendingTxMap     map[string]*validatortypes.PendingTx

	multiSigTxChan chan validatortypes.CosmosMultiSigTx
	multiSigTxLock *redsync.Mutex
}

func New(cfg conf.Config, client *chain.Client) *ValidatorHandler {
	validator := ValidatorHandler{
		cfg:    cfg,
		client: client,
		logger: logger.NewLogger("Validator"),

		pendingTxMapLock: &sync.RWMutex{},
		pendingTxMap:     make(map[string]*validatortypes.PendingTx),

		multiSigTxChan: make(chan validatortypes.CosmosMultiSigTx, 200),

		bridgeService:  service.GetRegistry().BridgeService(),
		configService:  service.GetRegistry().ConfigService(),
		redisService:   service.GetRegistry().RedisService(),
		historyService: service.GetRegistry().HistoryService(),
	}

	pool := goredis.NewPool(validator.redisService.Conn())
	validator.redSync = redsync.New(pool)
	validator.multiSigTxLock = validator.redSync.NewMutex(redisKeyLockMultiSigLock, redsync.WithExpiry(time.Minute))

	validator.initValidatorList()

	go validator.iterate()

	return &validator
}

func (p *ValidatorHandler) Name() string { return "Validator" }

func (p *ValidatorHandler) ApiHandler(e *gin.Engine) {
	e.GET("/tx/:txHash", p.Tx)

	g := e.Group("/validator")

	g.GET("/info", p.ValidatorInfo)
	g.GET("/statistics", p.ValidatorStatistics)
	g.GET("/pending", p.ValidatorPendingTxs)
	g.GET("/multisig/:chain/address", p.MultiSigAddress)

	g.POST("/pending/add", p.PendingTxRequest)
	g.POST("/recover/:chain/:txHash", p.Recover)
	g.POST("/recover/multisig/:chain/:txHash", p.RecoverMultiSig)
}

func (p *ValidatorHandler) LogHandler(log commontypes.Log) error {
	return nil
}

func (p *ValidatorHandler) iterate() {
	healthTicker := time.Tick(time.Minute)
	balanceTicker := time.Tick(time.Minute * 10)
	speedUpTicker := time.Tick(time.Second)
	recoverTicker := time.Tick(time.Hour)
	multiSigTicker := time.Tick(time.Second)

	go func() {
		for {
			select {
			case <-healthTicker:
				p.CheckHealth()
			case <-balanceTicker:
				p.CheckBalance()
			case <-speedUpTicker:
				p.CheckSpeedUpTx()
			case <-recoverTicker:
				p.CheckRecoverTx()

			case <-multiSigTicker:
				go p.CosmosMultiSigTx()
			}
		}
	}()
}

func (p *ValidatorHandler) initValidatorList() {
	validatorList, err := p.configService.FindValidatorInfo(context.Background())
	if err != nil {
		p.logger.Error("GetValidatorList", err)
		return
	}

	p.logger.Debug("init", "ValidatorList", "count", len(validatorList))

	for _, v := range validatorList {
		validator := validatortypes.NewValidatorInfo(v)
		p.logger.Debug(
			"name", validator.Name, "desc", validator.Description, "url", validator.Url,
			commontypes.ChainKLAY, validator.AddressInfo[commontypes.ChainKLAY].Address,
			commontypes.ChainMATIC, validator.AddressInfo[commontypes.ChainMATIC].Address,
			commontypes.ChainETH, validator.AddressInfo[commontypes.ChainETH].Address,
		)

		p.validatorList = append(p.validatorList, validator)
	}

	p.CheckHealth()
	p.CheckBalance()
}

func (p *ValidatorHandler) GetValidatorInfo(chain, address string) *validatortypes.ValidatorInfo {
	for _, v := range p.validatorList {
		if v.AddressInfo[chain].Address == address {
			return v
		}
	}
	return nil
}
