package types

import (
	"github.com/curtis0505/bridge/libs/types"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Name() string
	ApiHandler(e *gin.Engine)
	LogHandler(log types.Log) error
}
