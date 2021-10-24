package server

import (
	"flag"
	"fmt"
	"github.com/mailru/easygo/netpoll"
	"github.com/rahul0tripathi/fastws/logger"
	"github.com/rahul0tripathi/fastws/pkg/gopool"
	"github.com/rahul0tripathi/fastws/pkg/ws/ws"
	"github.com/rahul0tripathi/fastws/poller"
	"github.com/rahul0tripathi/fastws/socketpool"
	"github.com/rahul0tripathi/fastws/types"
	"net"
	"net/url"
	"time"
)

var (
	roomQuery = flag.String("roomQuery", "room", "required to connect to specific roomId socket")
	delay     = flag.Duration("server_cooldown", time.Millisecond*50, "Cooldown time before handling new connections")
)

func init() {
	flag.Parse()
}
func handler(conn net.Conn, desc *netpoll.Desc) {
	logger.Debug().Str("ipAddr::REMOTE", conn.RemoteAddr().String()).Str("ipAddr::LOCAL", conn.LocalAddr().String()).Msg("New Conn to be handled")
	var query, err, drop = func(conn net.Conn) (query url.Values, err error, drop bool) {
		drop = false
		u := ws.Upgrader{
			OnHeader: func(key, value []byte) (err error) {
				return
			},
			OnRequest: func(uri []byte) error {
				var _url, err = url.Parse(string(uri))
				if err != nil {
					return err
				}
				query = _url.Query()
				if v := query.Get(*roomQuery); v == "" {
					drop = true
					err = ws.RejectConnectionError(
						ws.RejectionReason("Invalid Request"),
						ws.RejectionStatus(400),
					)
				}
				return err
			},
		}
		if !drop {
			_, err = u.Upgrade(conn)
			return
		}
		return
	}(conn)
	if drop || err != nil {
		desc.Close()
		conn.Close()
		return
	}
	if socketpool.IsUserConnected(query.Get("id")) {
		socketpool.RemoveClientFromSocketPool(query.Get("id"), ws.StatusPolicyViolation, types.InternalServerError)
	}
	newClient, err := socketpool.AddClientToSocketPool(query.Get("id"), query.Get(*roomQuery), conn, desc)
	logger.Debug().Msg(fmt.Sprintf("+%v", newClient))
	if err != nil {
		desc.Close()
		conn.Close()
		return
	}
	err = poller.Poller.Start(newClient.Desc, func(ev netpoll.Event) {
		if ev&(netpoll.EventReadHup|netpoll.EventHup) != 0 {
			logger.Debug().Msgf("%d Event Received for user %s", ev, newClient.Id)
			socketpool.RemoveClientFromSocketPool(query.Get("id"), ws.StatusNormalClosure, types.EventHupReceived)
			return
		}
		err = gopool.ConnHandlePool.Submit(func() {
			receiverErr := newClient.Receive()
			if receiverErr != nil {
				logger.Debug().Msg(fmt.Sprintf("error occured in recieve %v , dropping connection", err))
				socketpool.RemoveClientFromSocketPool(query.Get("id"), ws.StatusInvalidFramePayloadData, types.InvalidEncoding)
				return
			}
		})
		if err != nil {
			socketpool.RemoveClientFromSocketPool(query.Get("id"), ws.StatusInternalServerError, types.InternalServerError)
			return
		}
	})
}
