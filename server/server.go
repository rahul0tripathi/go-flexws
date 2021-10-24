package server

import (
	"fmt"
	"github.com/mailru/easygo/netpoll"
	"github.com/rahul0tripathi/fastws/logger"
	"github.com/rahul0tripathi/fastws/pkg/gopool"
	"net"
	"time"
)

func HandleConnections(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Err(err).Msg("Failed To Accept Connection")
		}
		err = gopool.ConnHandlePool.Submit(func() {
			desc, err := netpoll.HandleRead(conn)
			if err != nil {
				logger.Err(err).Msg("failed to create descriptor")
				return
			}
			handler(conn, desc)
			return
		})
		if err != nil {
			fmt.Println("Failed To Schedule connection handler, sleeping")
			logger.Err(err).Msg("Failed To Schedule connection handler, sleeping")
			time.Sleep(*(delay))
		}
	}
}
