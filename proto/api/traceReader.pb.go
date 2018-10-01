// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/traceReader.proto

package haystack

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ExpressionTree_Operator int32

const (
	ExpressionTree_AND ExpressionTree_Operator = 0
	ExpressionTree_OR  ExpressionTree_Operator = 1
)

var ExpressionTree_Operator_name = map[int32]string{
	0: "AND",
	1: "OR",
}
var ExpressionTree_Operator_value = map[string]int32{
	"AND": 0,
	"OR":  1,
}

func (x ExpressionTree_Operator) String() string {
	return proto.EnumName(ExpressionTree_Operator_name, int32(x))
}
func (ExpressionTree_Operator) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{6, 0} }

// collection of spans belonging to a single Trace
type Trace struct {
	TraceId    string  `protobuf:"bytes,1,opt,name=traceId" json:"traceId,omitempty"`
	ChildSpans []*Span `protobuf:"bytes,2,rep,name=childSpans" json:"childSpans,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Trace) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Trace) GetChildSpans() []*Span {
	if m != nil {
		return m.ChildSpans
	}
	return nil
}

// request for fetching Trace for traceId
type TraceRequest struct {
	TraceId string `protobuf:"bytes,1,opt,name=traceId" json:"traceId,omitempty"`
}

func (m *TraceRequest) Reset()                    { *m = TraceRequest{} }
func (m *TraceRequest) String() string            { return proto.CompactTextString(m) }
func (*TraceRequest) ProtoMessage()               {}
func (*TraceRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TraceRequest) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

// request for raw traces representing list of traceIds
type RawTracesRequest struct {
	TraceId []string `protobuf:"bytes,1,rep,name=traceId" json:"traceId,omitempty"`
}

func (m *RawTracesRequest) Reset()                    { *m = RawTracesRequest{} }
func (m *RawTracesRequest) String() string            { return proto.CompactTextString(m) }
func (*RawTracesRequest) ProtoMessage()               {}
func (*RawTracesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *RawTracesRequest) GetTraceId() []string {
	if m != nil {
		return m.TraceId
	}
	return nil
}

// list of filtered traces
type RawTracesResult struct {
	Traces []*Trace `protobuf:"bytes,1,rep,name=traces" json:"traces,omitempty"`
}

func (m *RawTracesResult) Reset()                    { *m = RawTracesResult{} }
func (m *RawTracesResult) String() string            { return proto.CompactTextString(m) }
func (*RawTracesResult) ProtoMessage()               {}
func (*RawTracesResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *RawTracesResult) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// request for fetching span for give traceId and spanId
type SpanRequest struct {
	TraceId string `protobuf:"bytes,1,opt,name=traceId" json:"traceId,omitempty"`
	SpanId  string `protobuf:"bytes,2,opt,name=spanId" json:"spanId,omitempty"`
}

func (m *SpanRequest) Reset()                    { *m = SpanRequest{} }
func (m *SpanRequest) String() string            { return proto.CompactTextString(m) }
func (*SpanRequest) ProtoMessage()               {}
func (*SpanRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SpanRequest) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *SpanRequest) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

// a single operand in the expression tree
type Operand struct {
	// Types that are valid to be assigned to Operand:
	//	*Operand_Field
	//	*Operand_Expression
	Operand isOperand_Operand `protobuf_oneof:"operand"`
}

func (m *Operand) Reset()                    { *m = Operand{} }
func (m *Operand) String() string            { return proto.CompactTextString(m) }
func (*Operand) ProtoMessage()               {}
func (*Operand) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type isOperand_Operand interface {
	isOperand_Operand()
}

type Operand_Field struct {
	Field *Field `protobuf:"bytes,1,opt,name=field,oneof"`
}
type Operand_Expression struct {
	Expression *ExpressionTree `protobuf:"bytes,2,opt,name=expression,oneof"`
}

func (*Operand_Field) isOperand_Operand()      {}
func (*Operand_Expression) isOperand_Operand() {}

func (m *Operand) GetOperand() isOperand_Operand {
	if m != nil {
		return m.Operand
	}
	return nil
}

func (m *Operand) GetField() *Field {
	if x, ok := m.GetOperand().(*Operand_Field); ok {
		return x.Field
	}
	return nil
}

func (m *Operand) GetExpression() *ExpressionTree {
	if x, ok := m.GetOperand().(*Operand_Expression); ok {
		return x.Expression
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operand) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operand_OneofMarshaler, _Operand_OneofUnmarshaler, _Operand_OneofSizer, []interface{}{
		(*Operand_Field)(nil),
		(*Operand_Expression)(nil),
	}
}

func _Operand_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operand)
	// operand
	switch x := m.Operand.(type) {
	case *Operand_Field:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Field); err != nil {
			return err
		}
	case *Operand_Expression:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Expression); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operand.Operand has unexpected type %T", x)
	}
	return nil
}

func _Operand_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operand)
	switch tag {
	case 1: // operand.field
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Field)
		err := b.DecodeMessage(msg)
		m.Operand = &Operand_Field{msg}
		return true, err
	case 2: // operand.expression
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExpressionTree)
		err := b.DecodeMessage(msg)
		m.Operand = &Operand_Expression{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operand_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operand)
	// operand
	switch x := m.Operand.(type) {
	case *Operand_Field:
		s := proto.Size(x.Field)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operand_Expression:
		s := proto.Size(x.Expression)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// nested n-ary expression tree for specifying expression to filter
// represents a binary operator which will be performed on a list of operands
type ExpressionTree struct {
	Operator              ExpressionTree_Operator `protobuf:"varint,1,opt,name=operator,enum=ExpressionTree_Operator" json:"operator,omitempty"`
	Operands              []*Operand              `protobuf:"bytes,2,rep,name=operands" json:"operands,omitempty"`
	IsSpanLevelExpression bool                    `protobuf:"varint,3,opt,name=isSpanLevelExpression" json:"isSpanLevelExpression,omitempty"`
}

func (m *ExpressionTree) Reset()                    { *m = ExpressionTree{} }
func (m *ExpressionTree) String() string            { return proto.CompactTextString(m) }
func (*ExpressionTree) ProtoMessage()               {}
func (*ExpressionTree) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ExpressionTree) GetOperator() ExpressionTree_Operator {
	if m != nil {
		return m.Operator
	}
	return ExpressionTree_AND
}

func (m *ExpressionTree) GetOperands() []*Operand {
	if m != nil {
		return m.Operands
	}
	return nil
}

func (m *ExpressionTree) GetIsSpanLevelExpression() bool {
	if m != nil {
		return m.IsSpanLevelExpression
	}
	return false
}

// criteria for searching traces
type TracesSearchRequest struct {
	Fields           []*Field        `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
	StartTime        int64           `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	EndTime          int64           `protobuf:"varint,3,opt,name=endTime" json:"endTime,omitempty"`
	Limit            int32           `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	FilterExpression *ExpressionTree `protobuf:"bytes,5,opt,name=filterExpression" json:"filterExpression,omitempty"`
}

func (m *TracesSearchRequest) Reset()                    { *m = TracesSearchRequest{} }
func (m *TracesSearchRequest) String() string            { return proto.CompactTextString(m) }
func (*TracesSearchRequest) ProtoMessage()               {}
func (*TracesSearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *TracesSearchRequest) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *TracesSearchRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TracesSearchRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *TracesSearchRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TracesSearchRequest) GetFilterExpression() *ExpressionTree {
	if m != nil {
		return m.FilterExpression
	}
	return nil
}

// list of filtered traces
type TracesSearchResult struct {
	Traces []*Trace `protobuf:"bytes,1,rep,name=traces" json:"traces,omitempty"`
}

func (m *TracesSearchResult) Reset()                    { *m = TracesSearchResult{} }
func (m *TracesSearchResult) String() string            { return proto.CompactTextString(m) }
func (*TracesSearchResult) ProtoMessage()               {}
func (*TracesSearchResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *TracesSearchResult) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

// request for fetching trace count of search result per interval
type TraceCountsRequest struct {
	Fields           []*Field        `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty"`
	StartTime        int64           `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	EndTime          int64           `protobuf:"varint,3,opt,name=endTime" json:"endTime,omitempty"`
	Interval         int64           `protobuf:"varint,4,opt,name=interval" json:"interval,omitempty"`
	FilterExpression *ExpressionTree `protobuf:"bytes,5,opt,name=filterExpression" json:"filterExpression,omitempty"`
}

func (m *TraceCountsRequest) Reset()                    { *m = TraceCountsRequest{} }
func (m *TraceCountsRequest) String() string            { return proto.CompactTextString(m) }
func (*TraceCountsRequest) ProtoMessage()               {}
func (*TraceCountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *TraceCountsRequest) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *TraceCountsRequest) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TraceCountsRequest) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *TraceCountsRequest) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *TraceCountsRequest) GetFilterExpression() *ExpressionTree {
	if m != nil {
		return m.FilterExpression
	}
	return nil
}

// trace count list
type TraceCounts struct {
	TraceCount []*TraceCount `protobuf:"bytes,1,rep,name=traceCount" json:"traceCount,omitempty"`
}

func (m *TraceCounts) Reset()                    { *m = TraceCounts{} }
func (m *TraceCounts) String() string            { return proto.CompactTextString(m) }
func (*TraceCounts) ProtoMessage()               {}
func (*TraceCounts) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *TraceCounts) GetTraceCount() []*TraceCount {
	if m != nil {
		return m.TraceCount
	}
	return nil
}

// count of traces for an interval
type TraceCount struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Count     int64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *TraceCount) Reset()                    { *m = TraceCount{} }
func (m *TraceCount) String() string            { return proto.CompactTextString(m) }
func (*TraceCount) ProtoMessage()               {}
func (*TraceCount) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *TraceCount) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TraceCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// Field is a general abstraction on data associated with a span
// It can represent any indexed span attribute such as tag, log, spanName, or operationName
type Field struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// An empty message type for rq/rs
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

// query for fetching values for given field
type FieldValuesRequest struct {
	FieldName string   `protobuf:"bytes,1,opt,name=fieldName" json:"fieldName,omitempty"`
	Filters   []*Field `protobuf:"bytes,2,rep,name=filters" json:"filters,omitempty"`
}

func (m *FieldValuesRequest) Reset()                    { *m = FieldValuesRequest{} }
func (m *FieldValuesRequest) String() string            { return proto.CompactTextString(m) }
func (*FieldValuesRequest) ProtoMessage()               {}
func (*FieldValuesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *FieldValuesRequest) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *FieldValuesRequest) GetFilters() []*Field {
	if m != nil {
		return m.Filters
	}
	return nil
}

type FieldNames struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *FieldNames) Reset()                    { *m = FieldNames{} }
func (m *FieldNames) String() string            { return proto.CompactTextString(m) }
func (*FieldNames) ProtoMessage()               {}
func (*FieldNames) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *FieldNames) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type FieldValues struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *FieldValues) Reset()                    { *m = FieldValues{} }
func (m *FieldValues) String() string            { return proto.CompactTextString(m) }
func (*FieldValues) ProtoMessage()               {}
func (*FieldValues) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *FieldValues) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type CallNode struct {
	ServiceName            string `protobuf:"bytes,1,opt,name=serviceName" json:"serviceName,omitempty"`
	OperationName          string `protobuf:"bytes,2,opt,name=operationName" json:"operationName,omitempty"`
	InfrastructureProvider string `protobuf:"bytes,3,opt,name=infrastructureProvider" json:"infrastructureProvider,omitempty"`
	InfrastructureLocation string `protobuf:"bytes,4,opt,name=infrastructureLocation" json:"infrastructureLocation,omitempty"`
	Duration               string `protobuf:"bytes,5,opt,name=duration" json:"duration,omitempty"`
}

func (m *CallNode) Reset()                    { *m = CallNode{} }
func (m *CallNode) String() string            { return proto.CompactTextString(m) }
func (*CallNode) ProtoMessage()               {}
func (*CallNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *CallNode) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *CallNode) GetOperationName() string {
	if m != nil {
		return m.OperationName
	}
	return ""
}

func (m *CallNode) GetInfrastructureProvider() string {
	if m != nil {
		return m.InfrastructureProvider
	}
	return ""
}

func (m *CallNode) GetInfrastructureLocation() string {
	if m != nil {
		return m.InfrastructureLocation
	}
	return ""
}

func (m *CallNode) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

type Call struct {
	From         *CallNode `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To           *CallNode `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	NetworkDelta int64     `protobuf:"varint,3,opt,name=networkDelta" json:"networkDelta,omitempty"`
}

func (m *Call) Reset()                    { *m = Call{} }
func (m *Call) String() string            { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()               {}
func (*Call) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

func (m *Call) GetFrom() *CallNode {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Call) GetTo() *CallNode {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Call) GetNetworkDelta() int64 {
	if m != nil {
		return m.NetworkDelta
	}
	return 0
}

type TraceCallGraph struct {
	Calls []*Call `protobuf:"bytes,1,rep,name=calls" json:"calls,omitempty"`
}

func (m *TraceCallGraph) Reset()                    { *m = TraceCallGraph{} }
func (m *TraceCallGraph) String() string            { return proto.CompactTextString(m) }
func (*TraceCallGraph) ProtoMessage()               {}
func (*TraceCallGraph) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

func (m *TraceCallGraph) GetCalls() []*Call {
	if m != nil {
		return m.Calls
	}
	return nil
}

func init() {
	proto.RegisterType((*Trace)(nil), "Trace")
	proto.RegisterType((*TraceRequest)(nil), "TraceRequest")
	proto.RegisterType((*RawTracesRequest)(nil), "RawTracesRequest")
	proto.RegisterType((*RawTracesResult)(nil), "RawTracesResult")
	proto.RegisterType((*SpanRequest)(nil), "SpanRequest")
	proto.RegisterType((*Operand)(nil), "Operand")
	proto.RegisterType((*ExpressionTree)(nil), "ExpressionTree")
	proto.RegisterType((*TracesSearchRequest)(nil), "TracesSearchRequest")
	proto.RegisterType((*TracesSearchResult)(nil), "TracesSearchResult")
	proto.RegisterType((*TraceCountsRequest)(nil), "TraceCountsRequest")
	proto.RegisterType((*TraceCounts)(nil), "TraceCounts")
	proto.RegisterType((*TraceCount)(nil), "TraceCount")
	proto.RegisterType((*Field)(nil), "Field")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*FieldValuesRequest)(nil), "FieldValuesRequest")
	proto.RegisterType((*FieldNames)(nil), "FieldNames")
	proto.RegisterType((*FieldValues)(nil), "FieldValues")
	proto.RegisterType((*CallNode)(nil), "CallNode")
	proto.RegisterType((*Call)(nil), "Call")
	proto.RegisterType((*TraceCallGraph)(nil), "TraceCallGraph")
	proto.RegisterEnum("ExpressionTree_Operator", ExpressionTree_Operator_name, ExpressionTree_Operator_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TraceReader service

type TraceReaderClient interface {
	SearchTraces(ctx context.Context, in *TracesSearchRequest, opts ...grpc.CallOption) (*TracesSearchResult, error)
	GetTraceCounts(ctx context.Context, in *TraceCountsRequest, opts ...grpc.CallOption) (*TraceCounts, error)
	GetTrace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*Trace, error)
	GetRawTrace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*Trace, error)
	GetRawSpan(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*Span, error)
	GetFieldNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FieldNames, error)
	GetFieldValues(ctx context.Context, in *FieldValuesRequest, opts ...grpc.CallOption) (*FieldValues, error)
	GetTraceCallGraph(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceCallGraph, error)
	GetRawTraces(ctx context.Context, in *RawTracesRequest, opts ...grpc.CallOption) (*RawTracesResult, error)
}

type traceReaderClient struct {
	cc *grpc.ClientConn
}

func NewTraceReaderClient(cc *grpc.ClientConn) TraceReaderClient {
	return &traceReaderClient{cc}
}

func (c *traceReaderClient) SearchTraces(ctx context.Context, in *TracesSearchRequest, opts ...grpc.CallOption) (*TracesSearchResult, error) {
	out := new(TracesSearchResult)
	err := grpc.Invoke(ctx, "/TraceReader/searchTraces", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetTraceCounts(ctx context.Context, in *TraceCountsRequest, opts ...grpc.CallOption) (*TraceCounts, error) {
	out := new(TraceCounts)
	err := grpc.Invoke(ctx, "/TraceReader/getTraceCounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetTrace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*Trace, error) {
	out := new(Trace)
	err := grpc.Invoke(ctx, "/TraceReader/getTrace", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetRawTrace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*Trace, error) {
	out := new(Trace)
	err := grpc.Invoke(ctx, "/TraceReader/getRawTrace", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetRawSpan(ctx context.Context, in *SpanRequest, opts ...grpc.CallOption) (*Span, error) {
	out := new(Span)
	err := grpc.Invoke(ctx, "/TraceReader/getRawSpan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetFieldNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*FieldNames, error) {
	out := new(FieldNames)
	err := grpc.Invoke(ctx, "/TraceReader/getFieldNames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetFieldValues(ctx context.Context, in *FieldValuesRequest, opts ...grpc.CallOption) (*FieldValues, error) {
	out := new(FieldValues)
	err := grpc.Invoke(ctx, "/TraceReader/getFieldValues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetTraceCallGraph(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (*TraceCallGraph, error) {
	out := new(TraceCallGraph)
	err := grpc.Invoke(ctx, "/TraceReader/getTraceCallGraph", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceReaderClient) GetRawTraces(ctx context.Context, in *RawTracesRequest, opts ...grpc.CallOption) (*RawTracesResult, error) {
	out := new(RawTracesResult)
	err := grpc.Invoke(ctx, "/TraceReader/getRawTraces", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TraceReader service

type TraceReaderServer interface {
	SearchTraces(context.Context, *TracesSearchRequest) (*TracesSearchResult, error)
	GetTraceCounts(context.Context, *TraceCountsRequest) (*TraceCounts, error)
	GetTrace(context.Context, *TraceRequest) (*Trace, error)
	GetRawTrace(context.Context, *TraceRequest) (*Trace, error)
	GetRawSpan(context.Context, *SpanRequest) (*Span, error)
	GetFieldNames(context.Context, *Empty) (*FieldNames, error)
	GetFieldValues(context.Context, *FieldValuesRequest) (*FieldValues, error)
	GetTraceCallGraph(context.Context, *TraceRequest) (*TraceCallGraph, error)
	GetRawTraces(context.Context, *RawTracesRequest) (*RawTracesResult, error)
}

func RegisterTraceReaderServer(s *grpc.Server, srv TraceReaderServer) {
	s.RegisterService(&_TraceReader_serviceDesc, srv)
}

func _TraceReader_SearchTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TracesSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).SearchTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/SearchTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).SearchTraces(ctx, req.(*TracesSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetTraceCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetTraceCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/GetTraceCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetTraceCounts(ctx, req.(*TraceCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/GetTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetTrace(ctx, req.(*TraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetRawTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetRawTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/GetRawTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetRawTrace(ctx, req.(*TraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetRawSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetRawSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/GetRawSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetRawSpan(ctx, req.(*SpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetFieldNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetFieldNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/GetFieldNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetFieldNames(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetFieldValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FieldValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetFieldValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/GetFieldValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetFieldValues(ctx, req.(*FieldValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetTraceCallGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetTraceCallGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/GetTraceCallGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetTraceCallGraph(ctx, req.(*TraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceReader_GetRawTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceReaderServer).GetRawTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TraceReader/GetRawTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceReaderServer).GetRawTraces(ctx, req.(*RawTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceReader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TraceReader",
	HandlerType: (*TraceReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "searchTraces",
			Handler:    _TraceReader_SearchTraces_Handler,
		},
		{
			MethodName: "getTraceCounts",
			Handler:    _TraceReader_GetTraceCounts_Handler,
		},
		{
			MethodName: "getTrace",
			Handler:    _TraceReader_GetTrace_Handler,
		},
		{
			MethodName: "getRawTrace",
			Handler:    _TraceReader_GetRawTrace_Handler,
		},
		{
			MethodName: "getRawSpan",
			Handler:    _TraceReader_GetRawSpan_Handler,
		},
		{
			MethodName: "getFieldNames",
			Handler:    _TraceReader_GetFieldNames_Handler,
		},
		{
			MethodName: "getFieldValues",
			Handler:    _TraceReader_GetFieldValues_Handler,
		},
		{
			MethodName: "getTraceCallGraph",
			Handler:    _TraceReader_GetTraceCallGraph_Handler,
		},
		{
			MethodName: "getRawTraces",
			Handler:    _TraceReader_GetRawTraces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/traceReader.proto",
}

func init() { proto.RegisterFile("api/traceReader.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdb, 0x6e, 0x1b, 0x37,
	0x10, 0xf5, 0x4a, 0x5e, 0x5d, 0x66, 0x15, 0x5b, 0xa1, 0x93, 0x40, 0x55, 0xdc, 0x40, 0x21, 0x1c,
	0x40, 0xe8, 0x85, 0x85, 0x65, 0xa3, 0x40, 0xdb, 0x87, 0xb6, 0xce, 0xa5, 0x0e, 0x10, 0xd8, 0x01,
	0x23, 0xf4, 0xa1, 0x6f, 0xec, 0x2e, 0x2d, 0x2f, 0xb2, 0xda, 0xdd, 0x92, 0x94, 0x92, 0xfc, 0x54,
	0x7f, 0xa1, 0x2f, 0x7d, 0xe9, 0x4f, 0xf4, 0x5b, 0x0a, 0x5e, 0x76, 0x45, 0x4b, 0x32, 0x5a, 0xa0,
	0x40, 0xde, 0x76, 0xce, 0x39, 0x1c, 0x72, 0x86, 0x33, 0xc3, 0x85, 0xfb, 0xac, 0x4c, 0xbf, 0x52,
	0x82, 0xc5, 0x9c, 0x72, 0x96, 0x70, 0x41, 0x4a, 0x51, 0xa8, 0x62, 0x08, 0xb2, 0x64, 0xb9, 0xfd,
	0xc6, 0xe7, 0x10, 0x4e, 0xb5, 0x00, 0x0d, 0xa0, 0x6d, 0x94, 0x2f, 0x93, 0x41, 0x30, 0x0a, 0xc6,
	0x5d, 0x5a, 0x99, 0xe8, 0x09, 0x40, 0x7c, 0x9d, 0x66, 0xc9, 0x9b, 0x92, 0xe5, 0x72, 0xd0, 0x18,
	0x35, 0xc7, 0xd1, 0x24, 0x24, 0xda, 0xa2, 0x1e, 0x81, 0xc7, 0xd0, 0x9b, 0xda, 0xad, 0x7e, 0x5b,
	0x70, 0xa9, 0x6e, 0x77, 0x88, 0xbf, 0x80, 0x3e, 0x65, 0xef, 0x8c, 0x58, 0x6e, 0x55, 0x37, 0x7d,
	0xf5, 0x31, 0xec, 0x7b, 0x6a, 0xb9, 0xc8, 0x14, 0x7a, 0x04, 0x2d, 0xc3, 0x4a, 0xa3, 0x8d, 0x26,
	0x2d, 0x62, 0x77, 0x76, 0x28, 0xfe, 0x1e, 0x22, 0x73, 0xbc, 0x7f, 0x3b, 0x09, 0x7a, 0x00, 0x2d,
	0x9d, 0x8b, 0x97, 0xc9, 0xa0, 0x61, 0x08, 0x67, 0xe1, 0x19, 0xb4, 0x2f, 0x4b, 0x2e, 0x58, 0x9e,
	0xa0, 0x47, 0x10, 0x5e, 0xa5, 0x3c, 0xb3, 0x4b, 0xf5, 0x56, 0x2f, 0xb4, 0x75, 0xbe, 0x43, 0x2d,
	0x8c, 0x8e, 0x01, 0xf8, 0xfb, 0x52, 0x70, 0x29, 0xd3, 0x22, 0x37, 0x6e, 0xa2, 0xc9, 0x3e, 0x79,
	0x5e, 0x43, 0x53, 0xc1, 0xf9, 0xf9, 0x0e, 0xf5, 0x44, 0x67, 0x5d, 0x68, 0x17, 0xd6, 0x3b, 0xfe,
	0x23, 0x80, 0xbd, 0x9b, 0x5a, 0x74, 0x0a, 0x1d, 0xc3, 0xaa, 0x42, 0x98, 0x3d, 0xf7, 0x26, 0x83,
	0x35, 0x77, 0xe4, 0xd2, 0xf1, 0xb4, 0x56, 0xa2, 0x23, 0xb7, 0x2a, 0x4f, 0xaa, 0x2b, 0xea, 0x10,
	0x17, 0x02, 0xad, 0x19, 0x74, 0x0a, 0xf7, 0x53, 0xa9, 0x53, 0xf3, 0x8a, 0x2f, 0x79, 0xb6, 0xf2,
	0x3a, 0x68, 0x8e, 0x82, 0x71, 0x87, 0x6e, 0x27, 0xf1, 0x43, 0xe8, 0x54, 0x3b, 0xa2, 0x36, 0x34,
	0x7f, 0xbc, 0x78, 0xd6, 0xdf, 0x41, 0x2d, 0x68, 0x5c, 0xd2, 0x7e, 0x80, 0xff, 0x0c, 0xe0, 0xc0,
	0x5e, 0xce, 0x1b, 0xce, 0x44, 0x7c, 0x5d, 0x25, 0x1d, 0x43, 0xcb, 0x24, 0x68, 0x75, 0x47, 0x26,
	0x71, 0x67, 0x8d, 0x41, 0x40, 0x1d, 0x83, 0x0e, 0xa1, 0x2b, 0x15, 0x13, 0x6a, 0x9a, 0xce, 0xb9,
	0x49, 0x5d, 0x93, 0xae, 0x00, 0x7d, 0x6d, 0x3c, 0x4f, 0x0c, 0xd7, 0x34, 0x5c, 0x65, 0xa2, 0x7b,
	0x10, 0x66, 0xe9, 0x3c, 0x55, 0x83, 0xdd, 0x51, 0x30, 0x0e, 0xa9, 0x35, 0xd0, 0x77, 0xd0, 0xbf,
	0x4a, 0x33, 0xc5, 0x85, 0x17, 0x57, 0xb8, 0xf5, 0x3e, 0xe8, 0x86, 0x10, 0x9f, 0x02, 0xba, 0x19,
	0xc5, 0x7f, 0x2a, 0xb4, 0xbf, 0x02, 0xb7, 0xec, 0x69, 0xb1, 0xc8, 0x95, 0xfc, 0x18, 0xb1, 0x0f,
	0xa1, 0x93, 0xe6, 0x8a, 0x8b, 0x25, 0xcb, 0x4c, 0xf8, 0x4d, 0x5a, 0xdb, 0xff, 0x2f, 0x03, 0xdf,
	0x42, 0xe4, 0x85, 0x82, 0x3e, 0x07, 0x50, 0xb5, 0xe9, 0xe2, 0x88, 0xc8, 0x4a, 0x41, 0x3d, 0x1a,
	0xff, 0x00, 0xb0, 0x62, 0x74, 0x68, 0x2a, 0x9d, 0x73, 0xa9, 0xd8, 0xbc, 0x34, 0x25, 0xdc, 0xa4,
	0x2b, 0x40, 0x5f, 0x5e, 0x6c, 0x7c, 0xda, 0xa0, 0xad, 0x81, 0x8f, 0x21, 0x34, 0x39, 0x42, 0x08,
	0x76, 0x73, 0x36, 0xe7, 0xae, 0x53, 0xcd, 0xb7, 0x5e, 0xb2, 0x64, 0xd9, 0x82, 0xbb, 0x2e, 0xb5,
	0x06, 0x6e, 0x43, 0xf8, 0x7c, 0x5e, 0xaa, 0x0f, 0x78, 0x0a, 0xc8, 0xac, 0xfd, 0x59, 0xc3, 0xf5,
	0x25, 0x1c, 0x42, 0xd7, 0xa4, 0xfa, 0x62, 0xe5, 0x6d, 0x05, 0xa0, 0x11, 0xb4, 0x6d, 0x06, 0xaa,
	0x76, 0x71, 0x77, 0x44, 0x2b, 0x18, 0x63, 0x80, 0x17, 0x95, 0x5c, 0xea, 0x23, 0xe8, 0xa3, 0x48,
	0x37, 0x9d, 0xac, 0x81, 0x9f, 0x40, 0xe4, 0xed, 0xac, 0xc7, 0x89, 0x39, 0x5a, 0xa5, 0x72, 0x16,
	0xfe, 0x3b, 0x80, 0xce, 0x53, 0x96, 0x65, 0x17, 0x45, 0xa2, 0x77, 0x8e, 0x24, 0x17, 0xcb, 0x34,
	0xe6, 0xde, 0xc9, 0x7c, 0x08, 0x1d, 0xc1, 0x1d, 0xdb, 0xd7, 0x69, 0x91, 0x1b, 0x8d, 0x0d, 0xfb,
	0x26, 0x88, 0xbe, 0x86, 0x07, 0x69, 0x7e, 0x25, 0x98, 0x54, 0x62, 0x11, 0xab, 0x85, 0xe0, 0xaf,
	0x45, 0xb1, 0x4c, 0x13, 0x2e, 0x4c, 0xc5, 0x74, 0xe9, 0x2d, 0xec, 0xe6, 0xba, 0x57, 0x45, 0x6c,
	0xbc, 0x9a, 0x72, 0xda, 0x58, 0x57, 0xb1, 0xba, 0xf0, 0x92, 0x85, 0xdd, 0xdf, 0x14, 0x55, 0x97,
	0xd6, 0x36, 0x4e, 0x60, 0x57, 0xc7, 0x87, 0x3e, 0x85, 0xdd, 0x2b, 0x51, 0xcc, 0xdd, 0xac, 0xec,
	0x92, 0x2a, 0x68, 0x6a, 0x60, 0xf4, 0x09, 0x34, 0x54, 0xe1, 0x66, 0xa4, 0x47, 0x36, 0x54, 0x81,
	0x30, 0xf4, 0x72, 0xae, 0xde, 0x15, 0xe2, 0xed, 0x33, 0x9e, 0x29, 0xe6, 0xaa, 0xfe, 0x06, 0x86,
	0xbf, 0x84, 0x3d, 0x5b, 0x65, 0x2c, 0xcb, 0x7e, 0x12, 0xac, 0xbc, 0x46, 0x0f, 0x21, 0x8c, 0x59,
	0x96, 0x55, 0x7d, 0x16, 0x1a, 0x9f, 0xd4, 0x62, 0x93, 0xdf, 0x9b, 0xae, 0xa2, 0xed, 0xe3, 0x87,
	0xbe, 0x81, 0x9e, 0x34, 0xcd, 0x6d, 0x1b, 0x1d, 0xdd, 0x23, 0x5b, 0xe6, 0xd6, 0xf0, 0x80, 0x6c,
	0x99, 0x03, 0x27, 0xb0, 0x37, 0xe3, 0xca, 0x6f, 0x8f, 0x03, 0xb2, 0xd9, 0xf7, 0xc3, 0x9e, 0x0f,
	0xa2, 0xc7, 0xd0, 0xa9, 0x16, 0xa1, 0x3b, 0xc4, 0x7f, 0x1b, 0x87, 0x6e, 0x8e, 0xa0, 0x23, 0x88,
	0x66, 0x5c, 0x55, 0xcf, 0xdb, 0x6d, 0xaa, 0xc7, 0x00, 0x56, 0xa5, 0x87, 0x33, 0xea, 0x11, 0xef,
	0x6d, 0x1b, 0xda, 0x87, 0x58, 0x97, 0xcc, 0x8c, 0x2b, 0xaf, 0x5e, 0x5b, 0xc4, 0xf4, 0xc6, 0x30,
	0x22, 0x1e, 0x68, 0xc3, 0xf0, 0x2b, 0xf6, 0x80, 0x6c, 0x76, 0xce, 0xb0, 0xe7, 0x83, 0xe8, 0x04,
	0xee, 0xd6, 0xb1, 0xd7, 0x89, 0x5f, 0x3b, 0xe9, 0x3e, 0x59, 0xe3, 0x4f, 0xa0, 0xe7, 0x05, 0x26,
	0xd1, 0x5d, 0xb2, 0xfe, 0xe2, 0x0f, 0xfb, 0x64, 0xed, 0x59, 0x3f, 0xfb, 0x0c, 0x0e, 0xe3, 0x62,
	0x4e, 0xf8, 0xfb, 0x92, 0x27, 0x29, 0x23, 0x45, 0xc9, 0x73, 0xa2, 0x87, 0x4c, 0x9a, 0xcf, 0x08,
	0x2b, 0xd3, 0xd7, 0xc1, 0x2f, 0x9d, 0x6b, 0xf6, 0x41, 0x2a, 0x16, 0xbf, 0xfd, 0xb5, 0x65, 0x7e,
	0x5f, 0x4e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x69, 0xa2, 0xd9, 0xff, 0xe3, 0x08, 0x00, 0x00,
}
