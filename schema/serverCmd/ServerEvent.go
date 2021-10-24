// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serverCmd

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ServerEvent struct {
	_tab flatbuffers.Table
}

func GetRootAsServerEvent(buf []byte, offset flatbuffers.UOffsetT) *ServerEvent {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ServerEvent{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsServerEvent(buf []byte, offset flatbuffers.UOffsetT) *ServerEvent {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ServerEvent{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ServerEvent) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ServerEvent) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ServerEvent) Cmd() ServerCmd {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return ServerCmd(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ServerEvent) MutateCmd(n ServerCmd) bool {
	return rcv._tab.MutateInt32Slot(4, int32(n))
}

func (rcv *ServerEvent) PayloadType() Message {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return Message(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ServerEvent) MutatePayloadType(n Message) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

func (rcv *ServerEvent) Payload(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func ServerEventStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ServerEventAddCmd(builder *flatbuffers.Builder, Cmd ServerCmd) {
	builder.PrependInt32Slot(0, int32(Cmd), 0)
}
func ServerEventAddPayloadType(builder *flatbuffers.Builder, PayloadType Message) {
	builder.PrependByteSlot(1, byte(PayloadType), 0)
}
func ServerEventAddPayload(builder *flatbuffers.Builder, Payload flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Payload), 0)
}
func ServerEventEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
