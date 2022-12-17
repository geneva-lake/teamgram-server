/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package idgen

import (
	"fmt"

	"github.com/teamgram/proto/mtproto"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/types"
)

//////////////////////////////////////////////////////////////////////////////////////////

var _ *types.Int32Value
var _ *mtproto.Bool
var _ fmt.GoStringer

var clazzIdRegisters2 = map[int32]func() mtproto.TLObject{
	// Constructor
	-1065859893: func() mtproto.TLObject { // 0xc07844cb
		o := MakeTLIdVal(nil)
		o.Data2.Constructor = -1065859893
		return o
	},
	473672294: func() mtproto.TLObject { // 0x1c3baa66
		o := MakeTLIdVals(nil)
		o.Data2.Constructor = 473672294
		return o
	},
	704937224: func() mtproto.TLObject { // 0x2a047d08
		o := MakeTLSeqIdVal(nil)
		o.Data2.Constructor = 704937224
		return o
	},
	-1963845268: func() mtproto.TLObject { // 0x8af2196c
		o := MakeTLInputId(nil)
		o.Data2.Constructor = -1963845268
		return o
	},
	2133352380: func() mtproto.TLObject { // 0x7f285fbc
		o := MakeTLInputIds(nil)
		o.Data2.Constructor = 2133352380
		return o
	},
	-850215987: func() mtproto.TLObject { // 0xcd52bbcd
		o := MakeTLInputSeqId(nil)
		o.Data2.Constructor = -850215987
		return o
	},
	2058448257: func() mtproto.TLObject { // 0x7ab16d81
		o := MakeTLInputNSeqId(nil)
		o.Data2.Constructor = 2058448257
		return o
	},

	// Method
	-1099886560: func() mtproto.TLObject { // 0xbe711020
		return &TLIdgenNextId{
			Constructor: -1099886560,
		}
	},
	1204121518: func() mtproto.TLObject { // 0x47c56fae
		return &TLIdgenNextIds{
			Constructor: 1204121518,
		}
	},
	-1654936704: func() mtproto.TLObject { // 0x9d5bab80
		return &TLIdgenGetCurrentSeqId{
			Constructor: -1654936704,
		}
	},
	-852747923: func() mtproto.TLObject { // 0xcd2c196d
		return &TLIdgenSetCurrentSeqId{
			Constructor: -852747923,
		}
	},
	-160339608: func() mtproto.TLObject { // 0xf6716968
		return &TLIdgenGetNextSeqId{
			Constructor: -160339608,
		}
	},
	-1479226258: func() mtproto.TLObject { // 0xa7d4cc6e
		return &TLIdgenGetNextNSeqId{
			Constructor: -1479226258,
		}
	},
	-1434062537: func() mtproto.TLObject { // 0xaa85f137
		return &TLIdgenGetNextIdValList{
			Constructor: -1434062537,
		}
	},
}

func NewTLObjectByClassID(classId int32) mtproto.TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}

//----------------------------------------------------------------------------------------------------------------

///////////////////////////////////////////////////////////////////////////////
// IdVal <--
//  + TL_IdVal
//  + TL_IdVals
//  + TL_SeqIdVal
//

func (m *IdVal) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_idVal:
		t := m.To_IdVal()
		xBuf = t.Encode(layer)
	case Predicate_idVals:
		t := m.To_IdVals()
		xBuf = t.Encode(layer)
	case Predicate_seqIdVal:
		t := m.To_SeqIdVal()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *IdVal) CalcByteSize(layer int32) int {
	return 0
}

func (m *IdVal) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0xc07844cb:
		m2 := MakeTLIdVal(m)
		m2.Decode(dBuf)
	case 0x1c3baa66:
		m2 := MakeTLIdVals(m)
		m2.Decode(dBuf)
	case 0x2a047d08:
		m2 := MakeTLSeqIdVal(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *IdVal) DebugString() string {
	switch m.PredicateName {
	case Predicate_idVal:
		t := m.To_IdVal()
		return t.DebugString()
	case Predicate_idVals:
		t := m.To_IdVals()
		return t.DebugString()
	case Predicate_seqIdVal:
		t := m.To_SeqIdVal()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_IdVal
func (m *IdVal) To_IdVal() *TLIdVal {
	m.PredicateName = Predicate_idVal
	return &TLIdVal{
		Data2: m,
	}
}

// To_IdVals
func (m *IdVal) To_IdVals() *TLIdVals {
	m.PredicateName = Predicate_idVals
	return &TLIdVals{
		Data2: m,
	}
}

// To_SeqIdVal
func (m *IdVal) To_SeqIdVal() *TLSeqIdVal {
	m.PredicateName = Predicate_seqIdVal
	return &TLSeqIdVal{
		Data2: m,
	}
}

// MakeTLIdVal
func MakeTLIdVal(data2 *IdVal) *TLIdVal {
	if data2 == nil {
		return &TLIdVal{Data2: &IdVal{
			PredicateName: Predicate_idVal,
		}}
	} else {
		data2.PredicateName = Predicate_idVal
		return &TLIdVal{Data2: data2}
	}
}

func (m *TLIdVal) To_IdVal() *IdVal {
	m.Data2.PredicateName = Predicate_idVal
	return m.Data2
}

func (m *TLIdVal) SetId_INT64(v int64) { m.Data2.Id_INT64 = v }
func (m *TLIdVal) GetId_INT64() int64  { return m.Data2.Id_INT64 }

func (m *TLIdVal) GetPredicateName() string {
	return Predicate_idVal
}

func (m *TLIdVal) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xc07844cb: func() []byte {
			x.UInt(0xc07844cb)

			x.Long(m.GetId_INT64())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_idVal, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_idVal, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLIdVal) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdVal) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xc07844cb: func() error {
			m.SetId_INT64(dBuf.Long())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLIdVal) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLIdVals
func MakeTLIdVals(data2 *IdVal) *TLIdVals {
	if data2 == nil {
		return &TLIdVals{Data2: &IdVal{
			PredicateName: Predicate_idVals,
		}}
	} else {
		data2.PredicateName = Predicate_idVals
		return &TLIdVals{Data2: data2}
	}
}

func (m *TLIdVals) To_IdVal() *IdVal {
	m.Data2.PredicateName = Predicate_idVals
	return m.Data2
}

func (m *TLIdVals) SetId_VECTORINT64(v []int64) { m.Data2.Id_VECTORINT64 = v }
func (m *TLIdVals) GetId_VECTORINT64() []int64  { return m.Data2.Id_VECTORINT64 }

func (m *TLIdVals) GetPredicateName() string {
	return Predicate_idVals
}

func (m *TLIdVals) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x1c3baa66: func() []byte {
			x.UInt(0x1c3baa66)

			x.VectorLong(m.GetId_VECTORINT64())

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_idVals, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_idVals, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLIdVals) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdVals) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x1c3baa66: func() error {

			m.SetId_VECTORINT64(dBuf.VectorLong())

			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLIdVals) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLSeqIdVal
func MakeTLSeqIdVal(data2 *IdVal) *TLSeqIdVal {
	if data2 == nil {
		return &TLSeqIdVal{Data2: &IdVal{
			PredicateName: Predicate_seqIdVal,
		}}
	} else {
		data2.PredicateName = Predicate_seqIdVal
		return &TLSeqIdVal{Data2: data2}
	}
}

func (m *TLSeqIdVal) To_IdVal() *IdVal {
	m.Data2.PredicateName = Predicate_seqIdVal
	return m.Data2
}

func (m *TLSeqIdVal) SetId_INT64(v int64) { m.Data2.Id_INT64 = v }
func (m *TLSeqIdVal) GetId_INT64() int64  { return m.Data2.Id_INT64 }

func (m *TLSeqIdVal) GetPredicateName() string {
	return Predicate_seqIdVal
}

func (m *TLSeqIdVal) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x2a047d08: func() []byte {
			x.UInt(0x2a047d08)

			x.Long(m.GetId_INT64())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_seqIdVal, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_seqIdVal, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLSeqIdVal) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLSeqIdVal) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x2a047d08: func() error {
			m.SetId_INT64(dBuf.Long())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLSeqIdVal) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

///////////////////////////////////////////////////////////////////////////////
// InputId <--
//  + TL_InputId
//  + TL_InputIds
//  + TL_InputSeqId
//  + TL_InputNSeqId
//

func (m *InputId) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	var (
		xBuf []byte
	)

	switch predicateName {
	case Predicate_inputId:
		t := m.To_InputId()
		xBuf = t.Encode(layer)
	case Predicate_inputIds:
		t := m.To_InputIds()
		xBuf = t.Encode(layer)
	case Predicate_inputSeqId:
		t := m.To_InputSeqId()
		xBuf = t.Encode(layer)
	case Predicate_inputNSeqId:
		t := m.To_InputNSeqId()
		xBuf = t.Encode(layer)

	default:
		// logx.Errorf("invalid predicate error: %s",  m.PredicateName)
		return []byte{}
	}

	return xBuf
}

func (m *InputId) CalcByteSize(layer int32) int {
	return 0
}

func (m *InputId) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Constructor = TLConstructor(dBuf.Int())
	switch uint32(m.Constructor) {
	case 0x8af2196c:
		m2 := MakeTLInputId(m)
		m2.Decode(dBuf)
	case 0x7f285fbc:
		m2 := MakeTLInputIds(m)
		m2.Decode(dBuf)
	case 0xcd52bbcd:
		m2 := MakeTLInputSeqId(m)
		m2.Decode(dBuf)
	case 0x7ab16d81:
		m2 := MakeTLInputNSeqId(m)
		m2.Decode(dBuf)

	default:
		return fmt.Errorf("invalid constructorId: 0x%x", uint32(m.Constructor))
	}
	return dBuf.GetError()
}

func (m *InputId) DebugString() string {
	switch m.PredicateName {
	case Predicate_inputId:
		t := m.To_InputId()
		return t.DebugString()
	case Predicate_inputIds:
		t := m.To_InputIds()
		return t.DebugString()
	case Predicate_inputSeqId:
		t := m.To_InputSeqId()
		return t.DebugString()
	case Predicate_inputNSeqId:
		t := m.To_InputNSeqId()
		return t.DebugString()

	default:
		return "{}"
	}
}

// To_InputId
func (m *InputId) To_InputId() *TLInputId {
	m.PredicateName = Predicate_inputId
	return &TLInputId{
		Data2: m,
	}
}

// To_InputIds
func (m *InputId) To_InputIds() *TLInputIds {
	m.PredicateName = Predicate_inputIds
	return &TLInputIds{
		Data2: m,
	}
}

// To_InputSeqId
func (m *InputId) To_InputSeqId() *TLInputSeqId {
	m.PredicateName = Predicate_inputSeqId
	return &TLInputSeqId{
		Data2: m,
	}
}

// To_InputNSeqId
func (m *InputId) To_InputNSeqId() *TLInputNSeqId {
	m.PredicateName = Predicate_inputNSeqId
	return &TLInputNSeqId{
		Data2: m,
	}
}

// MakeTLInputId
func MakeTLInputId(data2 *InputId) *TLInputId {
	if data2 == nil {
		return &TLInputId{Data2: &InputId{
			PredicateName: Predicate_inputId,
		}}
	} else {
		data2.PredicateName = Predicate_inputId
		return &TLInputId{Data2: data2}
	}
}

func (m *TLInputId) To_InputId() *InputId {
	m.Data2.PredicateName = Predicate_inputId
	return m.Data2
}

func (m *TLInputId) GetPredicateName() string {
	return Predicate_inputId
}

func (m *TLInputId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x8af2196c: func() []byte {
			x.UInt(0x8af2196c)

			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_inputId, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_inputId, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLInputId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInputId) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x8af2196c: func() error {
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLInputId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLInputIds
func MakeTLInputIds(data2 *InputId) *TLInputIds {
	if data2 == nil {
		return &TLInputIds{Data2: &InputId{
			PredicateName: Predicate_inputIds,
		}}
	} else {
		data2.PredicateName = Predicate_inputIds
		return &TLInputIds{Data2: data2}
	}
}

func (m *TLInputIds) To_InputId() *InputId {
	m.Data2.PredicateName = Predicate_inputIds
	return m.Data2
}

func (m *TLInputIds) SetNum(v int32) { m.Data2.Num = v }
func (m *TLInputIds) GetNum() int32  { return m.Data2.Num }

func (m *TLInputIds) GetPredicateName() string {
	return Predicate_inputIds
}

func (m *TLInputIds) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x7f285fbc: func() []byte {
			x.UInt(0x7f285fbc)

			x.Int(m.GetNum())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_inputIds, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_inputIds, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLInputIds) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInputIds) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x7f285fbc: func() error {
			m.SetNum(dBuf.Int())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLInputIds) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLInputSeqId
func MakeTLInputSeqId(data2 *InputId) *TLInputSeqId {
	if data2 == nil {
		return &TLInputSeqId{Data2: &InputId{
			PredicateName: Predicate_inputSeqId,
		}}
	} else {
		data2.PredicateName = Predicate_inputSeqId
		return &TLInputSeqId{Data2: data2}
	}
}

func (m *TLInputSeqId) To_InputId() *InputId {
	m.Data2.PredicateName = Predicate_inputSeqId
	return m.Data2
}

func (m *TLInputSeqId) SetKey(v string) { m.Data2.Key = v }
func (m *TLInputSeqId) GetKey() string  { return m.Data2.Key }

func (m *TLInputSeqId) GetPredicateName() string {
	return Predicate_inputSeqId
}

func (m *TLInputSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0xcd52bbcd: func() []byte {
			x.UInt(0xcd52bbcd)

			x.String(m.GetKey())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_inputSeqId, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_inputSeqId, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLInputSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInputSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0xcd52bbcd: func() error {
			m.SetKey(dBuf.String())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLInputSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// MakeTLInputNSeqId
func MakeTLInputNSeqId(data2 *InputId) *TLInputNSeqId {
	if data2 == nil {
		return &TLInputNSeqId{Data2: &InputId{
			PredicateName: Predicate_inputNSeqId,
		}}
	} else {
		data2.PredicateName = Predicate_inputNSeqId
		return &TLInputNSeqId{Data2: data2}
	}
}

func (m *TLInputNSeqId) To_InputId() *InputId {
	m.Data2.PredicateName = Predicate_inputNSeqId
	return m.Data2
}

func (m *TLInputNSeqId) SetKey(v string) { m.Data2.Key = v }
func (m *TLInputNSeqId) GetKey() string  { return m.Data2.Key }

func (m *TLInputNSeqId) SetN(v int32) { m.Data2.N = v }
func (m *TLInputNSeqId) GetN() int32  { return m.Data2.N }

func (m *TLInputNSeqId) GetPredicateName() string {
	return Predicate_inputNSeqId
}

func (m *TLInputNSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x7ab16d81: func() []byte {
			x.UInt(0x7ab16d81)

			x.String(m.GetKey())
			x.Int(m.GetN())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_inputNSeqId, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_inputNSeqId, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}

func (m *TLInputNSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLInputNSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	var decodeF = map[uint32]func() error{
		0x7ab16d81: func() error {
			m.SetKey(dBuf.String())
			m.SetN(dBuf.Int())
			return dBuf.GetError()
		},
	}

	if f, ok := decodeF[uint32(m.Data2.Constructor)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.Constructor))
	}
}

func (m *TLInputNSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// TLIdgenNextId
///////////////////////////////////////////////////////////////////////////////

func (m *TLIdgenNextId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_nextId))

	switch uint32(m.Constructor) {
	case 0xbe711020:
		x.UInt(0xbe711020)

		// no flags

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenNextId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenNextId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xbe711020:

		// not has flags

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenNextId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenNextIds
///////////////////////////////////////////////////////////////////////////////

func (m *TLIdgenNextIds) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_nextIds))

	switch uint32(m.Constructor) {
	case 0x47c56fae:
		x.UInt(0x47c56fae)

		// no flags

		x.Int(m.GetNum())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenNextIds) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenNextIds) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x47c56fae:

		// not has flags

		m.Num = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenNextIds) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenGetCurrentSeqId
///////////////////////////////////////////////////////////////////////////////

func (m *TLIdgenGetCurrentSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_getCurrentSeqId))

	switch uint32(m.Constructor) {
	case 0x9d5bab80:
		x.UInt(0x9d5bab80)

		// no flags

		x.String(m.GetKey())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenGetCurrentSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenGetCurrentSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0x9d5bab80:

		// not has flags

		m.Key = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenGetCurrentSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenSetCurrentSeqId
///////////////////////////////////////////////////////////////////////////////

func (m *TLIdgenSetCurrentSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_setCurrentSeqId))

	switch uint32(m.Constructor) {
	case 0xcd2c196d:
		x.UInt(0xcd2c196d)

		// no flags

		x.String(m.GetKey())
		x.Long(m.GetId())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenSetCurrentSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenSetCurrentSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xcd2c196d:

		// not has flags

		m.Key = dBuf.String()
		m.Id = dBuf.Long()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenSetCurrentSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenGetNextSeqId
///////////////////////////////////////////////////////////////////////////////

func (m *TLIdgenGetNextSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_getNextSeqId))

	switch uint32(m.Constructor) {
	case 0xf6716968:
		x.UInt(0xf6716968)

		// no flags

		x.String(m.GetKey())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenGetNextSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenGetNextSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xf6716968:

		// not has flags

		m.Key = dBuf.String()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenGetNextSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenGetNextNSeqId
///////////////////////////////////////////////////////////////////////////////

func (m *TLIdgenGetNextNSeqId) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_getNextNSeqId))

	switch uint32(m.Constructor) {
	case 0xa7d4cc6e:
		x.UInt(0xa7d4cc6e)

		// no flags

		x.String(m.GetKey())
		x.Int(m.GetN())

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenGetNextNSeqId) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenGetNextNSeqId) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xa7d4cc6e:

		// not has flags

		m.Key = dBuf.String()
		m.N = dBuf.Int()
		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenGetNextNSeqId) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// TLIdgenGetNextIdValList
///////////////////////////////////////////////////////////////////////////////

func (m *TLIdgenGetNextIdValList) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	// x.Int(int32(CRC32_idgen_getNextIdValList))

	switch uint32(m.Constructor) {
	case 0xaa85f137:
		x.UInt(0xaa85f137)

		// no flags

		x.Int(int32(mtproto.CRC32_vector))
		x.Int(int32(len(m.GetId())))
		for _, v := range m.GetId() {
			x.Bytes((*v).Encode(layer))
		}

	default:
		// log.Errorf("")
	}

	return x.GetBuf()
}

func (m *TLIdgenGetNextIdValList) CalcByteSize(layer int32) int {
	return 0
}

func (m *TLIdgenGetNextIdValList) Decode(dBuf *mtproto.DecodeBuf) error {
	switch uint32(m.Constructor) {
	case 0xaa85f137:

		// not has flags

		c1 := dBuf.Int()
		if c1 != int32(mtproto.CRC32_vector) {
			// dBuf.err = fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
			return fmt.Errorf("invalid mtproto.CRC32_vector, c%d: %d", 1, c1)
		}
		l1 := dBuf.Int()
		v1 := make([]*InputId, l1)
		for i := int32(0); i < l1; i++ {
			v1[i] = &InputId{}
			v1[i].Decode(dBuf)
		}
		m.Id = v1

		return dBuf.GetError()

	default:
		// log.Errorf("")
	}
	return dBuf.GetError()
}

func (m *TLIdgenGetNextIdValList) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

//----------------------------------------------------------------------------------------------------------------
// Vector_Long
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_Long) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.VectorLong(m.Datas)

	return x.GetBuf()
}

func (m *Vector_Long) Decode(dBuf *mtproto.DecodeBuf) error {
	m.Datas = dBuf.VectorLong()

	return dBuf.GetError()
}

func (m *Vector_Long) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_Long) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}

// Vector_IdVal
///////////////////////////////////////////////////////////////////////////////
func (m *Vector_IdVal) Encode(layer int32) []byte {
	x := mtproto.NewEncodeBuf(512)
	x.Int(int32(mtproto.CRC32_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.Bytes((*v).Encode(layer))
	}

	return x.GetBuf()
}

func (m *Vector_IdVal) Decode(dBuf *mtproto.DecodeBuf) error {
	dBuf.Int() // TODO(@benqi): Check crc32 invalid
	l1 := dBuf.Int()
	m.Datas = make([]*IdVal, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = new(IdVal)
		(*m.Datas[i]).Decode(dBuf)
	}

	return dBuf.GetError()
}

func (m *Vector_IdVal) CalcByteSize(layer int32) int {
	return 0
}

func (m *Vector_IdVal) DebugString() string {
	jsonm := &jsonpb.Marshaler{OrigName: true}
	dbgString, _ := jsonm.MarshalToString(m)
	return dbgString
}
