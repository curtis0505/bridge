package api

import (
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/apps/validators/validator"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/logger"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Config     *conf.Config
	Client     *chain.Client
	Controller *Controller
}

func NewApi(config *conf.Config, client *chain.Client, validator *validator.Validator) (*Api, error) {
	controller, err := NewController(config, client, validator)
	if err != nil {
		return nil, err
	}

	api := Api{
		Config:     config,
		Client:     client,
		Controller: controller,
	}

	return &api, nil
}

func (api *Api) Route(r *gin.Engine) {
	r.GET("/health", api.Controller.HealthCheck)
	r.GET("/validator", api.Controller.GetValidatorInfos)
	r.GET("/:chain/validator/address", api.Controller.ValidatorAddressByChain)

	r.POST("/:chain/recover", api.Controller.RecoverTransaction)
	r.GET("/:chain/transaction/cache", api.Controller.GetCacheTransaction)
	r.POST("/:chain/speedup", api.Controller.SpeedUpTransaction)
	r.DELETE("/:chain/transaction", api.Controller.CancelTransaction)

	r.POST("/:chain/generate_key", api.Controller.GenerateKey)

	r.GET("/:chain/pubkey", api.Controller.PubKey)
	r.POST("/sign/multisig", api.Controller.SignMultiSig)
}

func (api *Api) Run() {
	r := gin.Default()
	r.Use(CORS())
	r.Use(logger.GinElogMiddleWare())

	api.Route(r)

	go func() {
		err := r.Run(":" + api.Config.Server.Port)
		if err != nil {
			panic(err)
			return
		}
	}()

}
