package main

import (
	"cloud/common/db"
	logger "github.com/alecthomas/log4go"
	"cloud/httpServer"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const SYSTEM_EXIST = 3

type Hu struct{
	T int64
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		logger.Info(fmt.Sprintf("receive:%v", sig))
		done <- true
	}()


	httpServer.StartUpServer()
	<-done
	logger.Info("exit")

}

func init() {
	err := db.GetEngineClient()
	if err != nil {
		logger.Error(err)
		os.Exit(SYSTEM_EXIST)
		return
	}
	err = db.GetRedisClient()
	if err != nil {
		logger.Error(err)
		os.Exit(SYSTEM_EXIST)
		return
	}
}
