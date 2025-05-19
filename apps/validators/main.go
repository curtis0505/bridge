package main

import (
	"flag"
	"github.com/curtis0505/bridge/apps/validators/api"
	"github.com/curtis0505/bridge/apps/validators/app"
	"github.com/curtis0505/bridge/apps/validators/conf"
	"github.com/curtis0505/bridge/apps/validators/validator"
	"github.com/curtis0505/bridge/libs/client/chain"
	"github.com/curtis0505/bridge/libs/elog"
	"time"
)

func main() {
	configFilePath := flag.String("config", "./conf/config.toml", "config file path, default - ./conf/config.toml")
	config, err := conf.NewConfig(*configFilePath)
	if err != nil {
		panic(err)
	}

	elog.InitLog(config.Log)
	elog.SetAppName("validator")

	clientInstance := chain.NewClientByConfig(config.Client)

	validatorInstance := &validator.Validator{}
	if config.Server.SubscribeStatus {
		account, err := validator.NewAccount(clientInstance, config)
		if err != nil {
			panic(err)
		}

		validatorInstance, err = validator.New(clientInstance, *config, account)
		if err != nil {
			panic(err)
		}

		appInstance, err := app.New(clientInstance, *config, validatorInstance)
		if err != nil {
			panic(err)
		}

		appInstance.Run()
	}

	// API 서버 실행
	apiInstance, err := api.NewApi(config, clientInstance, validatorInstance)
	if err != nil {
		panic(err)
	}

	apiInstance.Run()

	for {
		time.Sleep(time.Second * 30)
	}
}
