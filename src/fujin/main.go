/*
* @Author: huang
* @Date:   2017-10-24 10:30:12
* @Last Modified by:   huang
* @Last Modified time: 2017-10-24 17:27:38
 */
package main

import (
	"comm"
	"comm/config"
	"comm/dbmgr"
	"comm/logger"
	"flag"
	"fujin/routers"
	"math/rand"
	"os"
	"time"
)

var log = logger.DefaultLogger

func main() {
	rand.Seed(time.Now().Unix())
	// parse command line
	argFile := flag.String("config", "config.json", "config file")
	argServer := flag.String("server", "fujin", "config file")
	argLog := flag.String("log", "fujin.log", "log file")

	flag.Parse()
	// load config
	config.Parse(*argFile, *argServer)
	// open log
	logger.Open(*argLog)

	// signal
	quit := make(chan int)
	comm.OnSignal(func(s os.Signal) {
		log.Warning("shutdown signal received ...")
		close(quit)
	})
	start()
	<-quit
	stop()

	// close log
	logger.Close()
}

func start() {
	// open db mgr
	dbmgr.Open()
	//routers
	routers.Routers()

	log.Info("fujin werver started")
}

func stop() {
	// close db mgr
	dbmgr.Close()

	// app stopped
	log.Info("fujin server stopped")
}
