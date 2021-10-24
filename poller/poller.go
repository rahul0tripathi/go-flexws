package poller

import (
	"github.com/mailru/easygo/netpoll"
	"log"
)

var (
	Poller netpoll.Poller
)

func init() {
	var err error
	if Poller, err = netpoll.New(nil); err != nil {
		log.Fatal(err)
	}
}
