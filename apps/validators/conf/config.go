package conf

import (
	"github.com/curtis0505/bridge/libs/client/chain/conf"
	"github.com/curtis0505/bridge/libs/elog"
	commontypes "github.com/curtis0505/bridge/libs/types"
	"github.com/naoina/toml"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

type Config struct {
	Server   Server                    `toml:"server"`
	Client   conf.Config               `toml:"client"`
	Account  commontypes.AccountConfig `toml:"account"`
	Contract map[string]ContractConfig `toml:"contract"`
	Log      elog.Config
}

type Server struct {
	Port            string `toml:"port"`
	ProofURL        string `toml:"proofURL"`
	Monitor         string `toml:"monitor"`
	SubscribeStatus bool   `toml:"subscribeStatus"`
}

type ContractConfig struct {
	NptToken              string `toml:"nptToken"`
	Vault                 string `toml:"vault"`
	Minter                string `toml:"minter"`
	MultiSigWallet        string `toml:"multiSigWallet"`
	RestakeVault          string `toml:"restakeVault"`
	RestakeMinter         string `toml:"restakeMinter"`
	RestakeMultiSigWallet string `toml:"restakeMultiSigWallet"`
}

// NewConfig
// filePath 없으면 config.go 파일과 같은 디렉토리에 config.toml 파일 참조(기본값 "./config.toml")
func NewConfig(filePath ...string) (*Config, error) {
	var configFilePath string
	if len(filePath) > 0 {
		configFilePath = filePath[0]
	} else {
		_, filename, _, _ := runtime.Caller(0)
		configFilePath = filepath.Join(filepath.Dir(filename), "config.toml")
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		elog.Error("Error reading config file", "error", err)
		return nil, err
	}
	var config Config
	customToml := toml.DefaultConfig
	customToml.MissingField = func(rt reflect.Type, field string) error {
		elog.Warn("MissingField", "type", rt, "field", field)
		return nil
	}
	err = customToml.Unmarshal(configFile, &config)
	if err != nil {
		elog.Error("Unmarshal", "err", err)
		return nil, err
	}

	return &config, nil
}
