package signals

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/rahul0tripathi/fastws/schema/serverCmd"
	"github.com/rahul0tripathi/fastws/types"
	"time"
)

var (
	Location, _ = time.LoadLocation("Asia/Kolkata")
)

func GenJoinedUsersCount(count *types.JoinedUserCount, payload []byte) []byte {
	builder := flatbuffers.NewBuilder(20)
	contestId := builder.CreateString(count.Id)
	serverCmd.JoinedUserCountStart(builder)
	serverCmd.JoinedUserCountAddJoinedUsers(builder, count.JoinedUsers)
	serverCmd.JoinedUserCountAddId(builder, contestId)
	joinedUserEvent := serverCmd.JoinedUserCountEnd(builder)
	serverCmd.ServerEventStart(builder)
	serverCmd.ServerEventAddCmd(builder, serverCmd.ServerCmdJOINEDUSERCOUNT)
	serverCmd.ServerEventAddPayloadType(builder, serverCmd.MessageJoinedUserCount)
	serverCmd.ServerEventAddPayload(builder, joinedUserEvent)
	event := serverCmd.ServerEventEnd(builder)
	builder.Finish(event)
	payload = builder.FinishedBytes()
	return payload
}
