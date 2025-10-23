package health

import (
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/libs/client/chain"
	bridge "github.com/curtis0505/bridge/libs/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	_ types.Handler = &Health{}
)

type Health struct {
	client *chain.Client
	logger *logger.Logger
}

func New() *Health {
	health := Health{
		logger: logger.NewLogger("Health"),
	}

	return &health
}

func (p *Health) Name() string { return "Health" }

func (p *Health) ApiHandler(e *gin.Engine) {
	e.GET("/health", p.Health)
}

func (p *Health) LogHandler(log bridge.Log) error {
	return nil
}

func (p *Health) Health(c *gin.Context) {
	c.JSON(http.StatusOK, types.NewResponseSuccess())
}
