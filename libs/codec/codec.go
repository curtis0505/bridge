package codec

import (
	"encoding/json"
	"errors"
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cosmossdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	"github.com/curtis0505/bridge/libs/logger/v2"
	commontypes "github.com/curtis0505/bridge/libs/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/kaiachain/kaia/common/hexutil"
	"net/http"
	"reflect"
	"runtime"
	"sync"
)

type Codec struct {
	middleWare routeWrapper
	router     map[string]routePath

	*sync.RWMutex
	// cosmos
	txDecoder cosmossdk.TxDecoder
}

type routePath struct {
	method string
	ctx    routeWrapper
}

type routeWrapper []func(c *gin.Context)

func newRoutePath(method string, ctx ...func(c *gin.Context)) routePath {
	return routePath{
		method: method,
		ctx:    ctx,
	}
}

const DefaultContext = "default"

func NewCodec() *Codec {
	return &Codec{
		router:     make(map[string]routePath),
		middleWare: make(routeWrapper, 0),
		RWMutex:    &sync.RWMutex{},

		txDecoder: authtx.NewTxConfig(codec.NewProtoCodec(
			cosmostypes.NewInterfaceRegistry(
				cosmostypes.WithCosmosRegistry(),
				cosmostypes.WithWasmRegistry(),
				cosmostypes.WithFinschiaRegistry(),
				cosmostypes.WithIBCRegistry(),
			),
		), authtx.DefaultSignModes).TxDecoder(),
	}
}

func (cdc *Codec) RegisterAbiMethod(method string, ctx ...func(c *gin.Context)) {
	if len(method) == 0 || method == DefaultContext {
		logger.Debug(
			"RegisterAbiMethod", logger.BuildLogInput().WithMethod(DefaultContext).
				WithData("ctx", runtime.FuncForPC(reflect.ValueOf(ctx[len(ctx)-1]).Pointer()).Name()),
		)
		cdc.addRoutePath(DefaultContext, newRoutePath(DefaultContext, ctx...))
		return
	}
	sig := crypto.Keccak256([]byte(method))[:4]
	hex := hexutil.Encode(sig)

	logger.Debug(
		"RegisterAbiMethod", logger.BuildLogInput().WithMethod(method).
			WithData(
				"hex", hex,
				"ctx", runtime.FuncForPC(reflect.ValueOf(ctx[len(ctx)-1]).Pointer()).Name()),
	)

	if cdc.existRoutePath(hex) {
		logger.Error(
			"RegisterAbiMethod", logger.BuildLogInput().
				WithMethod(method).
				WithError(fmt.Errorf("duplicated hex routes")).
				WithData(
					"hex", hex,
					"ctx", runtime.FuncForPC(reflect.ValueOf(ctx[len(ctx)-1]).Pointer()).Name()),
		)
	}

	cdc.addRoutePath(hex, newRoutePath(method, ctx...))
}

func (cdc *Codec) RegisterMsgMethod(msg cosmossdk.Msg, ctx ...func(c *gin.Context)) {
	url := cosmossdk.MsgTypeURL(msg)
	logger.Debug(
		"RegisterMsgMethod", logger.BuildLogInput().WithMethod(url).
			WithData(
				"ctx", runtime.FuncForPC(reflect.ValueOf(ctx[len(ctx)-1]).Pointer()).Name()),
	)
	if cdc.existRoutePath(url) {
		logger.Error(
			"RegisterMsgMethod", logger.BuildLogInput().
				WithMethod(url).
				WithError(fmt.Errorf("duplicated hex routes")).
				WithData(
					"ctx", runtime.FuncForPC(reflect.ValueOf(ctx[len(ctx)-1]).Pointer()).Name()),
		)
	}

	cdc.addRoutePath(url, newRoutePath(url, ctx...))
}

func (cdc *Codec) RegisterCallWasmMethod(callWasm commontypes.CallWasm, ctx ...func(c *gin.Context)) {
	url := callWasm.WasmKey()
	logger.Debug(
		"registerWasmMsgMethod", logger.BuildLogInput().WithMethod(url).
			WithData(
				"ctx", runtime.FuncForPC(reflect.ValueOf(ctx[len(ctx)-1]).Pointer()).Name()),
	)

	if cdc.existRoutePath(url) {
		logger.Error(
			"registerWasmMsgMethod", logger.BuildLogInput().
				WithMethod(url).
				WithError(fmt.Errorf("duplicated CallWasm routes")).
				WithData(
					"ctx", runtime.FuncForPC(reflect.ValueOf(ctx[len(ctx)-1]).Pointer()).Name()),
		)
	}

	cdc.addRoutePath(url, newRoutePath(url, ctx...))
}

func (cdc *Codec) addRoutePath(key string, route routePath) {
	cdc.Lock()
	defer cdc.Unlock()

	cdc.router[key] = route
}

func (cdc *Codec) existRoutePath(key string) bool {
	cdc.Lock()
	defer cdc.Unlock()

	return cdc.router[key].method != ""
}

func (cdc *Codec) Use(ctx ...func(c *gin.Context)) {
	cdc.middleWare = append(cdc.middleWare, ctx...)
}

func (cdc *Codec) Route(c *gin.Context, data []byte) error {
	for _, m := range cdc.middleWare {
		if !c.IsAborted() {
			m(c)
		}
	}

	if c.IsAborted() {
		return fmt.Errorf("context aborted")
	}

	chain := c.Param("chain")

	var route string
	if len(data) == 0 {
		route = DefaultContext
	} else {
		if commontypes.GetChainType(chain) == commontypes.ChainTypeCOSMOS {
			protoTx, err := cdc.decodeCosmosTransaction(data)
			if err != nil {
				return err
			}

			if len(protoTx.GetMsgs()) == 0 {
				return errors.New("msg is nil")
			}
			message := protoTx.GetMsgs()[0]
			switch msg := message.(type) {
			case *wasmtypes.MsgExecuteContract:
				var msgExecuteContract cosmostypes.MsgCallWasm
				err = json.Unmarshal(msg.Msg, &msgExecuteContract)
				if err != nil {
					return err
				}
				route = msgExecuteContract.WasmKey()
			default:
				route = cosmossdk.MsgTypeURL(msg)
			}

		} else {
			route = hexutil.Encode(data[:4])
		}
	}

	if path, ok := cdc.router[route]; !ok {
		return fmt.Errorf("not found method: %s", route)
	} else {
		for _, ctx := range path.ctx {
			if !c.IsAborted() {
				logger.Debug(
					"Route", logger.BuildLogInput().
						WithChain(chain).
						WithMethod(path.method).
						WithData(
							"route", route,
							"ctx", runtime.FuncForPC(reflect.ValueOf(ctx).Pointer()).Name()),
				)
				ctx(c)
			} else {
				logger.Trace(
					"Route", logger.BuildLogInput().
						WithChain(chain).
						WithMethod(path.method).
						WithData(
							"route", route,
							"aborted", runtime.FuncForPC(reflect.ValueOf(ctx).Pointer()).Name()),
				)
			}
		}
	}
	return nil
}

func (cdc *Codec) RegisteredAbiMethods(c *gin.Context) {
	response := make([]map[string]string, 0)
	for hex, route := range cdc.router {
		abi := map[string]string{
			"hex":    hex,
			"method": route.method,
		}
		response = append(response, abi)
	}

	c.JSON(http.StatusOK, response)
}
