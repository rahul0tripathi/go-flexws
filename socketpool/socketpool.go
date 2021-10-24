package socketpool

import (
	"encoding/json"
	"github.com/mailru/easygo/netpoll"
	"github.com/rahul0tripathi/fastws/client"
	"github.com/rahul0tripathi/fastws/logger"
	"github.com/rahul0tripathi/fastws/pkg/ws/ws"
	"github.com/rahul0tripathi/fastws/poller"
	"github.com/rahul0tripathi/fastws/signals"
	"github.com/rahul0tripathi/fastws/types"
	"github.com/valyala/bytebufferpool"
	"net"
	"sync"
)

type socketPoolType struct {
	pool []string
}

var (
	clientSocketPool   = make(map[string]*client.SocketClient)
	clientSocketPoolIo = sync.RWMutex{}
	roomSocketPool     = make(map[string]*socketPoolType)
	roomSocketPoolIo   = sync.RWMutex{}
	joinedUserPool = sync.Pool{
		New: func() interface{} {
			return types.JoinedUserCount{}
		},
	}
	bufferPool bytebufferpool.Pool
	headerPool = sync.Pool{
		New: func() interface{} {
			return ws.Header{}
		},
	}
)

func GetValueFromRoomSocketPool(key string) (*socketPoolType, bool) {
	roomSocketPoolIo.Lock()
	defer roomSocketPoolIo.Unlock()
	value, ok := roomSocketPool[key]
	return value, ok
}
func GetClientFromClientSocketPool(id string) (value *client.SocketClient, found bool) {
	clientSocketPoolIo.Lock()
	defer clientSocketPoolIo.Unlock()
	value, found = clientSocketPool[id]
	return

}
func AddClientToRoomPool(room string, id string) (err error) {
	roomSocketPoolIo.Lock()
	defer roomSocketPoolIo.Unlock()
	defer logger.Debug().Str("AddClientToRoomPool", id).Msgf("room: %s", room)
	if _, ok := roomSocketPool[room]; ok {
		present := false
		for _, k := range roomSocketPool[room].pool {
			if k == id {
				present = true
				break
			}
		}
		if !present {
			roomSocketPool[room].pool = append(roomSocketPool[room].pool, id)
		}
		return
	} else {
		roomSocketPool[room] = &socketPoolType{
			pool: []string{},
		}
		roomSocketPool[room].pool = append(roomSocketPool[room].pool, id)
	}
	return
}
func RemoveClientFromRoomPool(room string, id string) (err error) {
	roomSocketPoolIo.Lock()
	defer roomSocketPoolIo.Unlock()
	defer logger.Debug().Str("RemoveClientFromRoomPool", id).Msgf("%v", roomSocketPool[room].pool)
	if _, ok := roomSocketPool[room]; ok {
		for i, k := range roomSocketPool[room].pool {
			if k == id {
				roomSocketPool[room].pool = append(roomSocketPool[room].pool[:i], roomSocketPool[room].pool[i+1:]...)
				return
			}
		}
	}
	err = types.ErrNotFound
	return
}

func AddClientToSocketPool(id string, room string, conn net.Conn, desc *netpoll.Desc) (newClient *client.SocketClient, err error) {
	defer clientSocketPoolIo.Unlock()
	clientSocketPoolIo.Lock()
	err = AddClientToRoomPool(room, id)
	if err != nil {
		return nil, err
	}
	newClient = client.CreateNewClient(id, room, desc, conn)
	clientSocketPool[id] = newClient
	logger.Debug().Str("AddClientToSocketPool", id).Msgf("%v", clientSocketPool)

	return
}
func IsUserConnected(id string) (connected bool) {
	defer clientSocketPoolIo.Unlock()
	clientSocketPoolIo.Lock()
	if _, ok := clientSocketPool[id]; ok {
		connected = true
	}
	return
}
func EmitUpdatedJoinedUsers(raw []byte) {
	_joinedUsers := joinedUserPool.Get().(types.JoinedUserCount)
	err := json.Unmarshal(raw, &_joinedUsers)
	if err != nil {
		return
	}
	buf := bufferPool.Get()
	buf.B = signals.GenJoinedUsersCount(&_joinedUsers, buf.Bytes())
	header := headerPool.Get().(ws.Header)
	defer bufferPool.Put(buf)
	defer headerPool.Put(header)
	header.Length = int64(buf.Len())
	header.Masked = false
	header.OpCode = ws.OpBinary
	header.Fin = true
	if _joinedUsers.Id != "" && _joinedUsers.RoomId != "" && _joinedUsers.JoinedUsers > uint16(0) {
		if val, ok := GetValueFromRoomSocketPool(_joinedUsers.RoomId); ok {
			for _, _client := range val.pool {
				if socketClient, isClientPresent := GetClientFromClientSocketPool(_client); isClientPresent {
					_ = socketClient.WriteRaw(header, buf.Bytes())
				}
			}
		}
	}
}
func RemoveClientFromSocketPool(id string, status ws.StatusCode, reason string) (err error) {
	clientSocketPoolIo.Lock()
	defer clientSocketPoolIo.Unlock()
	defer logger.Debug().Str("RemoveClientFromSocketPool", id).Msg("Removing User")
	if _, ok := clientSocketPool[id]; ok {
		err = RemoveClientFromRoomPool(clientSocketPool[id].Room, id)
		if err != nil {
			return
		}
		_ = clientSocketPool[id].SendCloseFrame(status, reason)
		err = poller.Poller.Stop(clientSocketPool[id].Desc)
		if err != nil {
			return
		}
		err = clientSocketPool[id].Desc.Close()
		if err != nil {
			return
		}
		err = clientSocketPool[id].Conn.Close()
		if err != nil {
			return
		}
		delete(clientSocketPool, id)
		return
	}
	return
}
