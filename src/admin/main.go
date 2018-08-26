package main

import (
	_ "admin/routers"
	"comm/config"
	"comm/dbmgr"
	"comm/logger"
	"flag"
	"github.com/astaxie/beego"
)

var log = logger.DefaultLogger

func main() {

	argFile := flag.String("config", "config.json", "config file")
	argServer := flag.String("server", "admin", "config file")
	argLog := flag.String("log", "admin.log", "log file")

	flag.Parse()
	// load config
	config.Parse(*argFile, *argServer)
	logger.Open(*argLog)

	dbmgr.Open()

	beego.Run()

	log.Info("admin start")
}
