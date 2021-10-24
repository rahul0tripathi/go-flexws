// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serverCmd

import "strconv"

type Message byte

const (
	MessageNONE            Message = 0
	MessageErr             Message = 1
	MessageJoinedUserCount Message = 2
)

var EnumNamesMessage = map[Message]string{
	MessageNONE:            "NONE",
	MessageErr:             "Err",
	MessageJoinedUserCount: "JoinedUserCount",
}

var EnumValuesMessage = map[string]Message{
	"NONE":            MessageNONE,
	"Err":             MessageErr,
	"JoinedUserCount": MessageJoinedUserCount,
}

func (v Message) String() string {
	if s, ok := EnumNamesMessage[v]; ok {
		return s
	}
	return "Message(" + strconv.FormatInt(int64(v), 10) + ")"
}
