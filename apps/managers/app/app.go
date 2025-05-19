package app

import (
	"context"
	"github.com/curtis0505/bridge/apps/managers/conf"
	"github.com/curtis0505/bridge/apps/managers/handler/control"
	"github.com/curtis0505/bridge/apps/managers/handler/fxportal"
	"github.com/curtis0505/bridge/apps/managers/handler/health"
	"github.com/curtis0505/bridge/apps/managers/handler/statistics"
	"github.com/curtis0505/bridge/apps/managers/handler/validator"
	"github.com/curtis0505/bridge/apps/managers/handler/verify"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/elog"
	"github.com/curtis0505/bridge/libs/service"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"runtime/debug"
	"time"

	bridge "github.com/curtis0505/bridge/libs/types"

	"github.com/curtis0505/bridge/apps/managers/handler/event"
	"github.com/curtis0505/bridge/apps/managers/types"
)

type App struct {
	cfg    conf.Config
	client *chain.Client
	//repositories *model.Repositories
	handler []types.Handler
	router  *gin.Engine
	logger  *elog.Logger
}

func New(cfg conf.Config) *App {
	elog.InitLog(cfg.Log)

	logger := elog.NewLogger("App")
	logger.Info("event", "NewApp")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), elog.GinElogMiddleWare())

	app := App{
		cfg:     cfg,
		client:  chain.NewClientByConfig(cfg.Client),
		logger:  logger,
		handler: make([]types.Handler, 0),
		router:  r,
	}

	service.Init(cfg.Repositories, app.client)

	app.init()

	return &app
}

func (app *App) init() {
	app.logger.Info("event", "init")

	healthHandler := health.New()
	verifyHandler := verify.New(app.cfg, app.client)
	validatorHandler := validator.New(app.cfg, app.client)
	statisticsHandler := statistics.New(app.cfg, app.client)
	controlHandler := control.New(app.cfg, app.client)
	fxPortalHandler := fxportal.New(app.cfg, app.client)
	eventHandler := event.New(
		app.cfg, app.client,
		verifyHandler, validatorHandler, fxPortalHandler,
	)

	app.handler = []types.Handler{
		healthHandler,
		verifyHandler,
		validatorHandler,
		eventHandler,
		statisticsHandler,
		controlHandler,
		fxPortalHandler,
	}

	app.registerRouter()
}

func (app *App) registerRouter() *App {
	app.logger.Info("event", "registerRouter")
	for _, handler := range app.handler {
		app.logger.Info("event", "registerRouter", "handler", handler.Name())
		handler.ApiHandler(app.router)
	}

	return app
}

func (app *App) Run() {
	app.logger.Info("event", "Run")

	lo.ForEach(app.client.GetChains(), func(chain string, _ int) {
		contractsInfo, err := service.GetRegistry().ContractService().FindContract(context.Background(),
			bson.M{
				"chain":       chain,
				"contract_id": bson.M{"$in": bridgetypes.GetBridgeContractIDs()},
			})
		if err != nil {
			panic(err)
		}

		var contracts []string
		for _, contract := range contractsInfo {
			contracts = append(contracts, contract.Address)
		}

		err = app.client.Subscribe(context.Background(),
			chain,
			app.logHandler,
			contracts...,
		)
		if err != nil {
			panic(err)
		}
	})

	app.ListenAndServe()
}

//func (app *App) Close() {
//	app.client.Close()
//}

func (app *App) ListenAndServe() {
	go func() {
		api := &http.Server{
			Addr:           app.cfg.Server.Port,
			Handler:        app.router,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		app.logger.Info("event", "ListenAndServe", "addr", api.Addr)
		if err := api.ListenAndServe(); err != nil {
			app.logger.Error("event", "ListenAndServe", "err", err)
		}
	}()
}

func (app *App) logHandler(log bridge.Log) {
	defer func() {
		if v := recover(); v != nil {
			app.logger.Error("panic", v)
			debug.PrintStack()
		}
	}()

	log = bridge.NewLog(log, log.Chain())
	for _, handler := range app.handler {
		err := handler.LogHandler(log)
		if err != nil {
			app.logger.Error("handler", handler.Name(), "err", err)
		}
	}
}
