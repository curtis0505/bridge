package app

import (
	"flag"
	"github.com/curtis0505/bridge/apps/managers/conf"
	"testing"
	"time"
)

var configFlag = flag.String("config", "../conf/config.toml", "toml file to use for configuration")

func TestApp(t *testing.T) {
	config := conf.NewConfig(*configFlag)

	app := New(*config)
	app.Run()

	time.Sleep(10 * time.Minute)
}
