// Code generated by protoc-gen-go. DO NOT EDIT.
// source: span.proto

/*
Package haystack is a generated protocol buffer package.

It is generated from these files:
	span.proto
	api/traceReader.proto

It has these top-level messages:
	Span
	Log
	Tag
	Batch
	Trace
	TraceRequest
	RawTracesRequest
	RawTracesResult
	SpanRequest
	Operand
	ExpressionTree
	TracesSearchRequest
	TracesSearchResult
	TraceCountsRequest
	TraceCounts
	TraceCount
	Field
	Empty
	FieldValuesRequest
	FieldNames
	FieldValues
	CallNode
	Call
	TraceCallGraph
*/
package haystack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TagType denotes the type of a Tag's value.
type Tag_TagType int32

const (
	Tag_STRING Tag_TagType = 0
	Tag_DOUBLE Tag_TagType = 1
	Tag_BOOL   Tag_TagType = 2
	Tag_LONG   Tag_TagType = 3
	Tag_BINARY Tag_TagType = 4
)

var Tag_TagType_name = map[int32]string{
	0: "STRING",
	1: "DOUBLE",
	2: "BOOL",
	3: "LONG",
	4: "BINARY",
}
var Tag_TagType_value = map[string]int32{
	"STRING": 0,
	"DOUBLE": 1,
	"BOOL":   2,
	"LONG":   3,
	"BINARY": 4,
}

func (x Tag_TagType) String() string {
	return proto.EnumName(Tag_TagType_name, int32(x))
}
func (Tag_TagType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// Span represents a unit of work performed by a service.
type Span struct {
	TraceId       string `protobuf:"bytes,1,opt,name=traceId" json:"traceId,omitempty"`
	SpanId        string `protobuf:"bytes,2,opt,name=spanId" json:"spanId,omitempty"`
	ParentSpanId  string `protobuf:"bytes,3,opt,name=parentSpanId" json:"parentSpanId,omitempty"`
	ServiceName   string `protobuf:"bytes,4,opt,name=serviceName" json:"serviceName,omitempty"`
	OperationName string `protobuf:"bytes,5,opt,name=operationName" json:"operationName,omitempty"`
	StartTime     int64  `protobuf:"varint,6,opt,name=startTime" json:"startTime,omitempty"`
	Duration      int64  `protobuf:"varint,7,opt,name=duration" json:"duration,omitempty"`
	Logs          []*Log `protobuf:"bytes,8,rep,name=logs" json:"logs,omitempty"`
	Tags          []*Tag `protobuf:"bytes,9,rep,name=tags" json:"tags,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Span) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Span) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *Span) GetParentSpanId() string {
	if m != nil {
		return m.ParentSpanId
	}
	return ""
}

func (m *Span) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Span) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *Span) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Span) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Span) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Span) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Log is a timestamped event with a set of tags.
type Log struct {
	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Fields    []*Tag `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Log) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Log) GetFields() []*Tag {
	if m != nil {
		return m.Fields
	}
	return nil
}

// Tag is a strongly typed key/value pair. We use 'oneof' protobuf attribute to represent the possible tagTypes
type Tag struct {
	Key  string      `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Type Tag_TagType `protobuf:"varint,2,opt,name=type,enum=Tag_TagType" json:"type,omitempty"`
	// Types that are valid to be assigned to Myvalue:
	//	*Tag_VStr
	//	*Tag_VLong
	//	*Tag_VDouble
	//	*Tag_VBool
	//	*Tag_VBytes
	Myvalue isTag_Myvalue `protobuf_oneof:"myvalue"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isTag_Myvalue interface {
	isTag_Myvalue()
}

type Tag_VStr struct {
	VStr string `protobuf:"bytes,3,opt,name=vStr,oneof"`
}
type Tag_VLong struct {
	VLong int64 `protobuf:"varint,4,opt,name=vLong,oneof"`
}
type Tag_VDouble struct {
	VDouble float64 `protobuf:"fixed64,5,opt,name=vDouble,oneof"`
}
type Tag_VBool struct {
	VBool bool `protobuf:"varint,6,opt,name=vBool,oneof"`
}
type Tag_VBytes struct {
	VBytes []byte `protobuf:"bytes,7,opt,name=vBytes,proto3,oneof"`
}

func (*Tag_VStr) isTag_Myvalue()    {}
func (*Tag_VLong) isTag_Myvalue()   {}
func (*Tag_VDouble) isTag_Myvalue() {}
func (*Tag_VBool) isTag_Myvalue()   {}
func (*Tag_VBytes) isTag_Myvalue()  {}

func (m *Tag) GetMyvalue() isTag_Myvalue {
	if m != nil {
		return m.Myvalue
	}
	return nil
}

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetType() Tag_TagType {
	if m != nil {
		return m.Type
	}
	return Tag_STRING
}

func (m *Tag) GetVStr() string {
	if x, ok := m.GetMyvalue().(*Tag_VStr); ok {
		return x.VStr
	}
	return ""
}

func (m *Tag) GetVLong() int64 {
	if x, ok := m.GetMyvalue().(*Tag_VLong); ok {
		return x.VLong
	}
	return 0
}

func (m *Tag) GetVDouble() float64 {
	if x, ok := m.GetMyvalue().(*Tag_VDouble); ok {
		return x.VDouble
	}
	return 0
}

func (m *Tag) GetVBool() bool {
	if x, ok := m.GetMyvalue().(*Tag_VBool); ok {
		return x.VBool
	}
	return false
}

func (m *Tag) GetVBytes() []byte {
	if x, ok := m.GetMyvalue().(*Tag_VBytes); ok {
		return x.VBytes
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Tag) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Tag_OneofMarshaler, _Tag_OneofUnmarshaler, _Tag_OneofSizer, []interface{}{
		(*Tag_VStr)(nil),
		(*Tag_VLong)(nil),
		(*Tag_VDouble)(nil),
		(*Tag_VBool)(nil),
		(*Tag_VBytes)(nil),
	}
}

func _Tag_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Tag)
	// myvalue
	switch x := m.Myvalue.(type) {
	case *Tag_VStr:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.VStr)
	case *Tag_VLong:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.VLong))
	case *Tag_VDouble:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.VDouble))
	case *Tag_VBool:
		t := uint64(0)
		if x.VBool {
			t = 1
		}
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Tag_VBytes:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.VBytes)
	case nil:
	default:
		return fmt.Errorf("Tag.Myvalue has unexpected type %T", x)
	}
	return nil
}

func _Tag_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Tag)
	switch tag {
	case 3: // myvalue.vStr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Myvalue = &Tag_VStr{x}
		return true, err
	case 4: // myvalue.vLong
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Myvalue = &Tag_VLong{int64(x)}
		return true, err
	case 5: // myvalue.vDouble
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Myvalue = &Tag_VDouble{math.Float64frombits(x)}
		return true, err
	case 6: // myvalue.vBool
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Myvalue = &Tag_VBool{x != 0}
		return true, err
	case 7: // myvalue.vBytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Myvalue = &Tag_VBytes{x}
		return true, err
	default:
		return false, nil
	}
}

func _Tag_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Tag)
	// myvalue
	switch x := m.Myvalue.(type) {
	case *Tag_VStr:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.VStr)))
		n += len(x.VStr)
	case *Tag_VLong:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.VLong))
	case *Tag_VDouble:
		n += proto.SizeVarint(5<<3 | proto.WireFixed64)
		n += 8
	case *Tag_VBool:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += 1
	case *Tag_VBytes:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.VBytes)))
		n += len(x.VBytes)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// You can optionally use Batch to send a collection of spans. Spans may not necessarily belong to one traceId.
type Batch struct {
	Spans []*Span `protobuf:"bytes,1,rep,name=spans" json:"spans,omitempty"`
}

func (m *Batch) Reset()                    { *m = Batch{} }
func (m *Batch) String() string            { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()               {}
func (*Batch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Batch) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func init() {
	proto.RegisterType((*Span)(nil), "Span")
	proto.RegisterType((*Log)(nil), "Log")
	proto.RegisterType((*Tag)(nil), "Tag")
	proto.RegisterType((*Batch)(nil), "Batch")
	proto.RegisterEnum("Tag_TagType", Tag_TagType_name, Tag_TagType_value)
}

func init() { proto.RegisterFile("span.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcd, 0x8a, 0x9c, 0x40,
	0x14, 0x85, 0xdb, 0xf6, 0xb7, 0xef, 0x74, 0x82, 0x14, 0x61, 0x28, 0x26, 0xb3, 0x10, 0x19, 0x42,
	0xaf, 0x5c, 0x4c, 0x9e, 0xa0, 0x65, 0xc2, 0x74, 0x83, 0x74, 0x87, 0x6a, 0xb3, 0x48, 0x76, 0x35,
	0x5a, 0x71, 0x64, 0xd4, 0x2a, 0xac, 0x6a, 0x89, 0xeb, 0xbc, 0x41, 0x9e, 0x38, 0x54, 0xe9, 0xfc,
	0x2d, 0x84, 0x7b, 0xbe, 0x73, 0xaf, 0xca, 0xe1, 0x00, 0x48, 0x41, 0xbb, 0x44, 0xf4, 0x5c, 0xf1,
	0xf8, 0xdf, 0x12, 0x9c, 0x93, 0xa0, 0x1d, 0xc2, 0xe0, 0xab, 0x9e, 0x16, 0x6c, 0x5f, 0x62, 0x2b,
	0xb2, 0x36, 0x2b, 0xf2, 0x2c, 0xd1, 0x25, 0x78, 0xfa, 0x60, 0x5f, 0xe2, 0xa5, 0x31, 0x66, 0x85,
	0x62, 0x58, 0x0b, 0xda, 0xb3, 0x4e, 0x9d, 0x26, 0xd7, 0x36, 0xee, 0x3b, 0x86, 0x22, 0xb8, 0x90,
	0xac, 0x1f, 0xea, 0x82, 0x1d, 0x68, 0xcb, 0xb0, 0x63, 0x56, 0xde, 0x22, 0x74, 0x03, 0x1f, 0xb8,
	0x60, 0x3d, 0x55, 0x35, 0xef, 0xcc, 0x8e, 0x6b, 0x76, 0xde, 0x43, 0x74, 0x0d, 0x2b, 0xa9, 0x68,
	0xaf, 0xf2, 0xba, 0x65, 0xd8, 0x8b, 0xac, 0x8d, 0x4d, 0x5e, 0x01, 0xba, 0x82, 0xa0, 0x3c, 0x4f,
	0xdb, 0xd8, 0x37, 0xe6, 0x8b, 0x46, 0x18, 0x9c, 0x86, 0x57, 0x12, 0x07, 0x91, 0xbd, 0xb9, 0xb8,
	0x75, 0x92, 0x8c, 0x57, 0xc4, 0x10, 0xed, 0x28, 0x5a, 0x49, 0xbc, 0x9a, 0x9d, 0x9c, 0x56, 0xc4,
	0x90, 0x78, 0x0b, 0x76, 0xc6, 0x2b, 0xfd, 0x51, 0x55, 0xb7, 0x4c, 0x2a, 0xda, 0x0a, 0x13, 0x8a,
	0x4d, 0x5e, 0x01, 0xba, 0x06, 0xef, 0x77, 0xcd, 0x9a, 0x52, 0xe2, 0xe5, 0x9b, 0x17, 0xcc, 0x2c,
	0xfe, 0xbb, 0x04, 0x3b, 0xa7, 0x15, 0x0a, 0xc1, 0x7e, 0x62, 0xe3, 0x1c, 0xa9, 0x1e, 0x51, 0x04,
	0x8e, 0x1a, 0x05, 0x33, 0x61, 0x7e, 0xbc, 0x5d, 0xeb, 0x2b, 0xfd, 0xe4, 0xa3, 0x60, 0xc4, 0x38,
	0xe8, 0x13, 0x38, 0xc3, 0x49, 0xf5, 0x53, 0xa0, 0xbb, 0x05, 0x31, 0x0a, 0x5d, 0x82, 0x3b, 0x64,
	0xbc, 0xab, 0x4c, 0x88, 0xf6, 0x6e, 0x41, 0x26, 0x89, 0xae, 0xc0, 0x1f, 0xee, 0xf8, 0xf9, 0xa1,
	0x99, 0xa2, 0xb3, 0x76, 0x0b, 0xf2, 0x0c, 0xcc, 0x4d, 0xca, 0x79, 0x63, 0x22, 0x0b, 0xcc, 0x8d,
	0x96, 0x08, 0x83, 0x37, 0xa4, 0xa3, 0x62, 0xd2, 0xc4, 0xb5, 0xde, 0x2d, 0xc8, 0xac, 0xe3, 0x2d,
	0xf8, 0xf3, 0xcf, 0x20, 0x00, 0xef, 0x94, 0x93, 0xfd, 0xe1, 0x3e, 0x5c, 0xe8, 0xf9, 0xee, 0xf8,
	0x23, 0xcd, 0xbe, 0x85, 0x16, 0x0a, 0xc0, 0x49, 0x8f, 0xc7, 0x2c, 0x5c, 0xea, 0x29, 0x3b, 0x1e,
	0xee, 0x43, 0x5b, 0xfb, 0xe9, 0xfe, 0xb0, 0x25, 0x3f, 0x43, 0x27, 0x5d, 0x81, 0xdf, 0x8e, 0x03,
	0x6d, 0xce, 0x2c, 0xbe, 0x01, 0x37, 0xa5, 0xaa, 0x78, 0x44, 0x9f, 0xc1, 0xd5, 0xad, 0x91, 0xd8,
	0x32, 0x59, 0xb9, 0x89, 0xee, 0x07, 0x99, 0x58, 0xfa, 0x05, 0x70, 0xc1, 0xdb, 0x84, 0xfd, 0x11,
	0xac, 0xac, 0x69, 0xc2, 0x05, 0xeb, 0x12, 0x5d, 0xbe, 0xba, 0xab, 0xbe, 0x5b, 0xbf, 0x82, 0x47,
	0x3a, 0x4a, 0x45, 0x8b, 0xa7, 0x07, 0xcf, 0x54, 0xf6, 0xeb, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xdd, 0x60, 0xee, 0x6b, 0xc0, 0x02, 0x00, 0x00,
}
