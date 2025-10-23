package conf

import (
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	modelconf "github.com/curtis0505/bridge/libs/model/conf/v2"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port      string
		ServiceId string
	}

	Account types.AccountConfig

	Key struct {
		Control string
	}

	Validator struct {
		SpeedUpTimeOut int
	}

	Control struct {
		BridgeGasCheckDuration int

		BridgePauseGasPrice int
		BridgePauseCount    int
		BridgePauseCountMax int
		BridgePauseTimeOut  int
	}

	Repositories modelconf.Config

	Log logger.Config

	Client   conf.Config
	FxPortal struct {
		MaticApiUrl       string
		RootTokenAddress  string
		ChildTokenAddress string
	}
}

func NewConfig(file string) *Config {
	c := new(Config)

	if file, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
