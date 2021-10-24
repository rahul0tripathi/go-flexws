package main

import (
	"fmt"
	"github.com/rahul0tripathi/fastws/amqp"
	"github.com/rahul0tripathi/fastws/config"
	"github.com/rahul0tripathi/fastws/logger"
	"github.com/rahul0tripathi/fastws/server"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"syscall"
)

func init() {
	go func() {
		http.ListenAndServe(":2233", nil)
	}()
	instance := os.Getenv("INSTANCEID")
	if instance == "" {
		os.Setenv("INSTANCEID", "01")
	}
}
func main() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		logger.Err(err).Msg("Failed to Get RLimit")
		return
	}
	logger.Info().Str("runtime", runtime.GOOS).Msg("current runtime")
	if runtime.GOOS == "darwin" {
		rLimit.Cur = rLimit.Max
	} else {
		rLimit.Cur = rLimit.Max
	}
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		logger.Err(err).Msg("Failed to Set RLimit")
		return
	}
	ln, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		logger.Err(err).Msg("Failed to Listen server")
		return
	}
	go amqp.RoomQueueHandler(config.AppConfig.Mq.Username, config.AppConfig.Mq.Password, config.AppConfig.Mq.Host, config.AppConfig.Mq.Port, config.AppConfig.RoomQueue.Name+os.Getenv("INSTANCEID"), config.AppConfig.RoomQueue.Bind, config.AppConfig.RoomQueue.Exchange, config.AppConfig.RoomQueue.ExchangeType, config.AppConfig.RoomQueue.Consumer+os.Getenv("INSTANCEID"))
	fmt.Printf("🚀 Server ready at ::%s", os.Getenv("PORT"))
	server.HandleConnections(ln)
}
