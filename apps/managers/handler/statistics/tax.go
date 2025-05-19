package statistics

import (
	"github.com/curtis0505/bridge/apps/managers/types"
	"github.com/curtis0505/bridge/libs/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *StatisticsHandler) TaxInfo(ctx *gin.Context) {
	taxInfoList, err := service.GetRegistry().HistoryService().AggregateBridgeTax(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.NewResponseHeader(types.Failed, err))
		return
	}

	ctx.JSON(http.StatusOK, taxInfoList)
}
