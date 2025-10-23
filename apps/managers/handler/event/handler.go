package event

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/apps/managers/conf"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/service"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"sync"

	bridge "github.com/curtis0505/bridge/libs/types"
)

var (
	_ types.Handler = &EventHandler{}
)

type EventHandler struct {
	cfg    conf.Config
	client *chain.Client
	logger *logger.Logger

	logEvmEvent     map[string]func(bridge.Log) error
	logEvmEventLock *sync.RWMutex

	logWasmEvent     map[string]func(bridge.Log) error
	logWasmEventLock *sync.RWMutex

	pendingConfirmation map[int64][]bridge.Log

	//itemDB          *model.ItemDB
	//cache           *model.Cache
	//historyBridgeDB *model.HistoryBridgeDB
	//configDB        *model.ConfigDB

	verifyHandler    eventtypes.VerifyHandler
	validatorHandler eventtypes.ValidatorHandler
	fxPortalHandler  eventtypes.FxPortalHandler

	historyService     *service.HistoryService
	bridgeService      *service.BridgeService
	transactionManager *service.TransactionManager

	contracts bridgetypes.ContractAddresses
}

func New(
	cfg conf.Config, client *chain.Client,
	verifyHandler eventtypes.VerifyHandler, validatorHandler eventtypes.ValidatorHandler,
	fxPortalHandler eventtypes.FxPortalHandler,
) *EventHandler {
	event := EventHandler{
		cfg:    cfg,
		client: client,
		logger: logger.NewLogger("Event"),

		logEvmEvent:     make(map[string]func(bridge.Log) error),
		logEvmEventLock: &sync.RWMutex{},

		logWasmEvent:     make(map[string]func(bridge.Log) error),
		logWasmEventLock: &sync.RWMutex{},

		pendingConfirmation: make(map[int64][]bridge.Log),

		verifyHandler:    verifyHandler,
		validatorHandler: validatorHandler,
		fxPortalHandler:  fxPortalHandler,

		historyService:     service.GetRegistry().HistoryService(),
		bridgeService:      service.GetRegistry().BridgeService(),
		transactionManager: service.GetRegistry().TransactionManager(),

		contracts: make(bridgetypes.ContractAddresses),
	}

	lo.ForEach(client.GetChains(), func(chain string, _ int) {
		contractsInfo, err := service.GetRegistry().ContractService().FindContract(context.Background(),
			bson.M{
				"chain":       chain,
				"contract_id": bson.M{"$in": bridgetypes.GetBridgeContractIDs()},
			})
		if err != nil {
			panic(err)
		}

		addresses := &bridgetypes.ContractAddress{}
		for _, contract := range contractsInfo {
			if contract.ContractID == bridgetypes.VaultContractID {
				addresses.Vault = contract.Address
			}

			if contract.ContractID == bridgetypes.MinterContractID {
				addresses.Minter = contract.Address
			}

			if contract.ContractID == "restake-vault" {
				addresses.RestakeVault = contract.Address
			}

			if contract.ContractID == "restake-minter" {
				addresses.RestakeMinter = contract.Address
			}

			event.contracts[commontypes.Chain(strings.ToUpper(chain))] = addresses
		}

	})
	event.initLogEvent()

	return &event
}

func (p *EventHandler) initLogEvent() {
	p.logger.Debug("init", "LogEvent")

	p.registerEvent()
	p.registerCosmosEvent()
}

func (p *EventHandler) registerEvent() {
	p.logEvmEvent[eventtypes.EventDeposit] = p.Deposit
	p.logEvmEvent[eventtypes.EventWithdraw] = p.Withdraw

	p.logEvmEvent[eventtypes.EventBurn] = p.Burn
	p.logEvmEvent[eventtypes.EventMint] = p.Mint

	p.logEvmEvent[eventtypes.EventSubmission] = p.Submission
	p.logEvmEvent[eventtypes.EventConfirmation] = p.Confirmation
	p.logEvmEvent[eventtypes.EventExecution] = p.Execution
	p.logEvmEvent[eventtypes.EventExecutionFailure] = p.ExecutionFailure

	p.logEvmEvent[eventtypes.EventFxDepositERC20] = p.FxDepositERC20
	p.logEvmEvent[eventtypes.EventSyncDeposit] = p.SyncDeposit
	p.logEvmEvent[eventtypes.EventFxChildWithdrawERC20] = p.FxChildWithdrawERC20
	p.logEvmEvent[eventtypes.EventFxWithdrawERC20] = p.FxWithdrawERC20

	p.logEvmEvent[eventtypes.EventNewHeaderBlock] = p.NewHeaderBlock
}

func (p *EventHandler) registerCosmosEvent() {
	p.logWasmEvent[bridgetypes.EventNameWasmDeposit] = p.WasmDeposit
	p.logWasmEvent[bridgetypes.EventNameWasmDepositCoin] = p.WasmDepositCoin
	p.logWasmEvent[bridgetypes.EventNameWasmWithdraw] = p.WasmWithdraw
	p.logWasmEvent[bridgetypes.EventNameWasmWithdrawCoin] = p.WasmWithdrawCoin

	p.logWasmEvent[bridgetypes.EventNameWasmBurn] = p.WasmBurn
	p.logWasmEvent[bridgetypes.EventNameWasmMint] = p.WasmMint
}

func (p *EventHandler) Name() string { return "Event" }

func (p *EventHandler) ApiHandler(e *gin.Engine) {}

func (p *EventHandler) LogHandler(log bridge.Log) error {
	switch commontypes.GetChainType(log.Chain()) {
	case commontypes.ChainTypeEVM:
		p.logEvmEventLock.Lock()
		defer p.logEvmEventLock.Unlock()

		logEvent, ok := p.logEvmEvent[log.EventName]
		if !ok {
			return nil
		}
		return logEvent(log)

	case commontypes.ChainTypeCOSMOS:
		logEvent, ok := p.logWasmEvent[log.EventName]
		if !ok {
			return nil
		}
		return logEvent(log)

		return nil
	default:
		return fmt.Errorf("unknown chain type: %s", log.Chain())
	}
}
