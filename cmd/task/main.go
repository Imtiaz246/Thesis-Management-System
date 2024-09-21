package main

import (
	"context"
	"flag"
	"github.com/Imtiaz246/Thesis-Management-System/cmd/task/wire"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/config"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/log"
)

func main() {
	var envConf = flag.String("conf", "config/config.yml", "config path, eg: -conf ./config/config.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)
	logger.Info("start task")
	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
