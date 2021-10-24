// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serverCmd

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type JoinedUserCount struct {
	_tab flatbuffers.Table
}

func GetRootAsJoinedUserCount(buf []byte, offset flatbuffers.UOffsetT) *JoinedUserCount {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &JoinedUserCount{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsJoinedUserCount(buf []byte, offset flatbuffers.UOffsetT) *JoinedUserCount {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &JoinedUserCount{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *JoinedUserCount) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *JoinedUserCount) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *JoinedUserCount) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *JoinedUserCount) JoinedUsers() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *JoinedUserCount) MutateJoinedUsers(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func JoinedUserCountStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func JoinedUserCountAddId(builder *flatbuffers.Builder, Id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Id), 0)
}
func JoinedUserCountAddJoinedUsers(builder *flatbuffers.Builder, JoinedUsers uint16) {
	builder.PrependUint16Slot(1, JoinedUsers, 0)
}
func JoinedUserCountEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
