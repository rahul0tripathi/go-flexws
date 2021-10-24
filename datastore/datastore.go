package datastore

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rahul0tripathi/fastws/config"
	"github.com/rahul0tripathi/fastws/types"
	"strconv"
	"time"
)

var (
	appCache *redis.Client
)

func init() {
	appCache = redis.NewClient(&redis.Options{
		Addr:        config.AppConfig.Cache.Host,
		Password:    config.AppConfig.Cache.Password,
		MaxRetries:  4,
		DialTimeout: 10 * time.Second,
	})
}
func UpdateJoinedUsers(roomId string){
	appCache.Incr(context.Background(),types.JoinedUsersCount(roomId))
}
func GetJoinedUsers(userId string, roomId string) (joinedUsers types.JoinedUserCount) {
	_joinedUsersCountReply, _ := appCache.Get(context.Background(), types.JoinedUsersCount(roomId)).Result()
	_joinedUsersCount, _ := strconv.ParseUint(_joinedUsersCountReply, 10, 16)
	// error ignored
	joinedUsers = types.JoinedUserCount{
		Id:          userId,
		JoinedUsers: uint16(_joinedUsersCount),
		RoomId:      roomId,
	}
	return
}
