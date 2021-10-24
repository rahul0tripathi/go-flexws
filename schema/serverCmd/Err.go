// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serverCmd

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Err struct {
	_tab flatbuffers.Table
}

func GetRootAsErr(buf []byte, offset flatbuffers.UOffsetT) *Err {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Err{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsErr(buf []byte, offset flatbuffers.UOffsetT) *Err {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Err{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Err) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Err) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Err) Message() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Err) Code() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ErrStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ErrAddMessage(builder *flatbuffers.Builder, Message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Message), 0)
}
func ErrAddCode(builder *flatbuffers.Builder, Code flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Code), 0)
}
func ErrEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
