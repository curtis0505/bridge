package statistics

import (
	"github.com/curtis0505/bridge/apps/managers/conf"
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/libs/client/chain"
	bridge "github.com/curtis0505/bridge/libs/types"
	"github.com/gin-gonic/gin"
)

var (
	_ types.Handler = &StatisticsHandler{}
)

type StatisticsHandler struct {
	cfg    conf.Config
	client *chain.Client
	logger *logger.Logger
}

func New(cfg conf.Config, client *chain.Client) *StatisticsHandler {
	statistics := StatisticsHandler{
		cfg:    cfg,
		client: client,
		logger: logger.NewLogger("StatisticsHandler"),
	}

	return &statistics
}

func (p *StatisticsHandler) Name() string { return "StatisticsHandler" }

func (p *StatisticsHandler) ApiHandler(e *gin.Engine) {
	g := e.Group("/statistics")

	g.GET("/tax", p.TaxInfo)
}

func (p *StatisticsHandler) LogHandler(log bridge.Log) error {
	return nil
}
