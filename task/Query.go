// automatically generated, do not modify

package task

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Query struct {
	_tab flatbuffers.Table
}

func (rcv *Query) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Query) Attr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Query) Uids(j int) uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint64(a + flatbuffers.UOffsetT(j * 8))
	}
	return 0
}

func (rcv *Query) UidsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func QueryStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func QueryAddAttr(builder *flatbuffers.Builder, attr flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(attr), 0) }
func QueryAddUids(builder *flatbuffers.Builder, uids flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(uids), 0) }
func QueryStartUidsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(8, numElems, 8)
}
func QueryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
