package client

import (
	"github.com/mailru/easygo/netpoll"
	"github.com/rahul0tripathi/fastws/logger"
	"github.com/rahul0tripathi/fastws/pkg/ws/ws"
	"github.com/rahul0tripathi/fastws/signals"
	"github.com/valyala/bytebufferpool"
	"io"
	"net"
	"sync"
)

var (
	bufferPool bytebufferpool.Pool
	headerPool = sync.Pool{
		New: func() interface{} {
			return ws.Header{}
		},
	}
)

type SocketClient struct {
	Conn net.Conn
	Id   string
	Room string
	Desc *netpoll.Desc
	io   *sync.RWMutex
	Exit chan bool
}

func CreateNewClient(id string, room string, desc *netpoll.Desc, conn net.Conn) *SocketClient {
	return &SocketClient{
		Conn: conn,
		Id:   id,
		Room: room,
		Desc: desc,
		io:   &sync.RWMutex{},
		Exit: make(chan bool),
	}
}

func (c *SocketClient) WriteRaw(header ws.Header, payload []byte) (err error) {
	defer c.io.Unlock()
	c.io.Lock()
	if err = ws.WriteHeader(c.Conn, header); err != nil {
		return
	}
	if _, err = c.Conn.Write(payload); err != nil {
		return
	}
	return nil
}

func (c *SocketClient) Receive() (err error) {
	c.io.Lock()
	logger.Debug().Msgf("Reading User Cmd %s", c.Id)
	header := headerPool.Get().(ws.Header)
	defer headerPool.Put(header)
	header, err = ws.ReadHeader(c.Conn)
	if err != nil {
		c.io.Unlock()
		return
	}
	if header.OpCode != ws.OpBinary {
		err = nil
		c.io.Unlock()
		return
	}
	if header.Length > 0 {
		payload := bufferPool.Get()
		defer bufferPool.Put(payload)
		payload.B = make([]byte, header.Length)
		_, err = io.ReadFull(c.Conn, payload.B)
		if err != nil {
			c.io.Unlock()
			return
		}
		if header.Masked {
			ws.Cipher(payload.B, header.Mask, 0)
		}
		_ = signals.ParseClientSignal(payload.Bytes())
		// process the parsed signal

	}
	c.io.Unlock()
	return
}
func (c *SocketClient) SendCloseFrame(status ws.StatusCode, reason string) (err error) {
	frame := bufferPool.Get()
	defer bufferPool.Put(frame)
	header := headerPool.Get().(ws.Header)
	defer headerPool.Put(header)
	header.Fin = true
	header.Masked = false
	frame.B = ws.NewCloseFrameBody(status, reason)
	header.OpCode = ws.OpClose
	header.Length = int64(frame.Len())
	err = c.WriteRaw(header, frame.Bytes())
	return
}
