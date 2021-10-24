package handler

import (
	"github.com/rahul0tripathi/fastws/logger"
	"github.com/rahul0tripathi/fastws/pkg/gopool"
	"github.com/rahul0tripathi/fastws/socketpool"
	"github.com/rahul0tripathi/fastws/types"
	"github.com/streadway/amqp"
)

func MessageHandler(message amqp.Delivery) {
	logger.Debug().Msgf("key:%s , body: %s", message.RoutingKey, string(message.Body))
	switch message.RoutingKey {
	case types.UpdateJoinedUserCount:
		gopool.ConnHandlePool.Submit(func() { socketpool.EmitUpdatedJoinedUsers(message.Body) })
		break
	}
}
