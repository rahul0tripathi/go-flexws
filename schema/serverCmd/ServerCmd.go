// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serverCmd

import "strconv"

type ServerCmd int32

const (
	ServerCmdEMPTY           ServerCmd = 0
	ServerCmdJOINEDUSERCOUNT ServerCmd = 1
	ServerCmdERR             ServerCmd = 2
)

var EnumNamesServerCmd = map[ServerCmd]string{
	ServerCmdEMPTY:           "EMPTY",
	ServerCmdJOINEDUSERCOUNT: "JOINEDUSERCOUNT",
	ServerCmdERR:             "ERR",
}

var EnumValuesServerCmd = map[string]ServerCmd{
	"EMPTY":           ServerCmdEMPTY,
	"JOINEDUSERCOUNT": ServerCmdJOINEDUSERCOUNT,
	"ERR":             ServerCmdERR,
}

func (v ServerCmd) String() string {
	if s, ok := EnumNamesServerCmd[v]; ok {
		return s
	}
	return "ServerCmd(" + strconv.FormatInt(int64(v), 10) + ")"
}
