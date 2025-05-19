package main

import (
	"flag"
	"github.com/curtis0505/bridge/apps/managers/app"
	"github.com/curtis0505/bridge/apps/managers/conf"
	"os"
)

var configFlag = flag.String("config", "./conf/config.toml", "toml file to use for configuration")

func main() {
	flag.Parse()
	config := conf.NewConfig(*configFlag)
	c := make(chan os.Signal)

	app := app.New(*config)
	app.Run()

	<-c
	//app.Close()
}
