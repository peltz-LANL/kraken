// Code generated by protoc-gen-go. DO NOT EDIT.
// source: API.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServiceControl_Command int32

const (
	ServiceControl_STOP   ServiceControl_Command = 0
	ServiceControl_UPDATE ServiceControl_Command = 1
	ServiceControl_INIT   ServiceControl_Command = 2
)

var ServiceControl_Command_name = map[int32]string{
	0: "STOP",
	1: "UPDATE",
	2: "INIT",
}
var ServiceControl_Command_value = map[string]int32{
	"STOP":   0,
	"UPDATE": 1,
	"INIT":   2,
}

func (x ServiceControl_Command) String() string {
	return proto.EnumName(ServiceControl_Command_name, int32(x))
}
func (ServiceControl_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{3, 0}
}

type MutationControl_Type int32

const (
	MutationControl_MUTATE    MutationControl_Type = 0
	MutationControl_INTERRUPT MutationControl_Type = 1
)

var MutationControl_Type_name = map[int32]string{
	0: "MUTATE",
	1: "INTERRUPT",
}
var MutationControl_Type_value = map[string]int32{
	"MUTATE":    0,
	"INTERRUPT": 1,
}

func (x MutationControl_Type) String() string {
	return proto.EnumName(MutationControl_Type_name, int32(x))
}
func (MutationControl_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{4, 0}
}

type Query struct {
	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Query_Node
	//	*Query_Value
	Payload              isQuery_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{0}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (dst *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(dst, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

type isQuery_Payload interface {
	isQuery_Payload()
}

type Query_Node struct {
	Node *Node `protobuf:"bytes,2,opt,name=node,proto3,oneof"`
}

type Query_Value struct {
	Value *ReflectValue `protobuf:"bytes,3,opt,name=value,proto3,oneof"`
}

func (*Query_Node) isQuery_Payload() {}

func (*Query_Value) isQuery_Payload() {}

func (m *Query) GetPayload() isQuery_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Query) GetNode() *Node {
	if x, ok := m.GetPayload().(*Query_Node); ok {
		return x.Node
	}
	return nil
}

func (m *Query) GetValue() *ReflectValue {
	if x, ok := m.GetPayload().(*Query_Value); ok {
		return x.Value
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Query) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Query_OneofMarshaler, _Query_OneofUnmarshaler, _Query_OneofSizer, []interface{}{
		(*Query_Node)(nil),
		(*Query_Value)(nil),
	}
}

func _Query_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Query)
	// payload
	switch x := m.Payload.(type) {
	case *Query_Node:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Node); err != nil {
			return err
		}
	case *Query_Value:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Value); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Query.Payload has unexpected type %T", x)
	}
	return nil
}

func _Query_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Query)
	switch tag {
	case 2: // payload.node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Node)
		err := b.DecodeMessage(msg)
		m.Payload = &Query_Node{msg}
		return true, err
	case 3: // payload.value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReflectValue)
		err := b.DecodeMessage(msg)
		m.Payload = &Query_Value{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Query_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Query)
	// payload
	switch x := m.Payload.(type) {
	case *Query_Node:
		s := proto.Size(x.Node)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Query_Value:
		s := proto.Size(x.Value)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type QueryMulti struct {
	Queries              []*Query `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryMulti) Reset()         { *m = QueryMulti{} }
func (m *QueryMulti) String() string { return proto.CompactTextString(m) }
func (*QueryMulti) ProtoMessage()    {}
func (*QueryMulti) Descriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{1}
}
func (m *QueryMulti) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMulti.Unmarshal(m, b)
}
func (m *QueryMulti) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMulti.Marshal(b, m, deterministic)
}
func (dst *QueryMulti) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMulti.Merge(dst, src)
}
func (m *QueryMulti) XXX_Size() int {
	return xxx_messageInfo_QueryMulti.Size(m)
}
func (m *QueryMulti) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMulti.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMulti proto.InternalMessageInfo

func (m *QueryMulti) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type ServiceInitRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Module               string   `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceInitRequest) Reset()         { *m = ServiceInitRequest{} }
func (m *ServiceInitRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceInitRequest) ProtoMessage()    {}
func (*ServiceInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{2}
}
func (m *ServiceInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInitRequest.Unmarshal(m, b)
}
func (m *ServiceInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInitRequest.Marshal(b, m, deterministic)
}
func (dst *ServiceInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInitRequest.Merge(dst, src)
}
func (m *ServiceInitRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceInitRequest.Size(m)
}
func (m *ServiceInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInitRequest proto.InternalMessageInfo

func (m *ServiceInitRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceInitRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

type ServiceControl struct {
	Command              ServiceControl_Command `protobuf:"varint,1,opt,name=command,proto3,enum=proto.ServiceControl_Command" json:"command,omitempty"`
	Config               *any.Any               `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ServiceControl) Reset()         { *m = ServiceControl{} }
func (m *ServiceControl) String() string { return proto.CompactTextString(m) }
func (*ServiceControl) ProtoMessage()    {}
func (*ServiceControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{3}
}
func (m *ServiceControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceControl.Unmarshal(m, b)
}
func (m *ServiceControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceControl.Marshal(b, m, deterministic)
}
func (dst *ServiceControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceControl.Merge(dst, src)
}
func (m *ServiceControl) XXX_Size() int {
	return xxx_messageInfo_ServiceControl.Size(m)
}
func (m *ServiceControl) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceControl.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceControl proto.InternalMessageInfo

func (m *ServiceControl) GetCommand() ServiceControl_Command {
	if m != nil {
		return m.Command
	}
	return ServiceControl_STOP
}

func (m *ServiceControl) GetConfig() *any.Any {
	if m != nil {
		return m.Config
	}
	return nil
}

type MutationControl struct {
	Module               string               `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 MutationControl_Type `protobuf:"varint,3,opt,name=type,proto3,enum=proto.MutationControl_Type" json:"type,omitempty"`
	Cfg                  *Node                `protobuf:"bytes,4,opt,name=cfg,proto3" json:"cfg,omitempty"`
	Dsc                  *Node                `protobuf:"bytes,5,opt,name=dsc,proto3" json:"dsc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MutationControl) Reset()         { *m = MutationControl{} }
func (m *MutationControl) String() string { return proto.CompactTextString(m) }
func (*MutationControl) ProtoMessage()    {}
func (*MutationControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{4}
}
func (m *MutationControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutationControl.Unmarshal(m, b)
}
func (m *MutationControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutationControl.Marshal(b, m, deterministic)
}
func (dst *MutationControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutationControl.Merge(dst, src)
}
func (m *MutationControl) XXX_Size() int {
	return xxx_messageInfo_MutationControl.Size(m)
}
func (m *MutationControl) XXX_DiscardUnknown() {
	xxx_messageInfo_MutationControl.DiscardUnknown(m)
}

var xxx_messageInfo_MutationControl proto.InternalMessageInfo

func (m *MutationControl) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *MutationControl) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MutationControl) GetType() MutationControl_Type {
	if m != nil {
		return m.Type
	}
	return MutationControl_MUTATE
}

func (m *MutationControl) GetCfg() *Node {
	if m != nil {
		return m.Cfg
	}
	return nil
}

func (m *MutationControl) GetDsc() *Node {
	if m != nil {
		return m.Dsc
	}
	return nil
}

type DiscoveryEvent struct {
	Module               string   `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	ValueId              string   `protobuf:"bytes,3,opt,name=value_id,json=valueId,proto3" json:"value_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryEvent) Reset()         { *m = DiscoveryEvent{} }
func (m *DiscoveryEvent) String() string { return proto.CompactTextString(m) }
func (*DiscoveryEvent) ProtoMessage()    {}
func (*DiscoveryEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{5}
}
func (m *DiscoveryEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryEvent.Unmarshal(m, b)
}
func (m *DiscoveryEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryEvent.Marshal(b, m, deterministic)
}
func (dst *DiscoveryEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryEvent.Merge(dst, src)
}
func (m *DiscoveryEvent) XXX_Size() int {
	return xxx_messageInfo_DiscoveryEvent.Size(m)
}
func (m *DiscoveryEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryEvent proto.InternalMessageInfo

func (m *DiscoveryEvent) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *DiscoveryEvent) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DiscoveryEvent) GetValueId() string {
	if m != nil {
		return m.ValueId
	}
	return ""
}

type ReflectValue struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReflectValue) Reset()         { *m = ReflectValue{} }
func (m *ReflectValue) String() string { return proto.CompactTextString(m) }
func (*ReflectValue) ProtoMessage()    {}
func (*ReflectValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_API_6081312aeb8bb432, []int{6}
}
func (m *ReflectValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReflectValue.Unmarshal(m, b)
}
func (m *ReflectValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReflectValue.Marshal(b, m, deterministic)
}
func (dst *ReflectValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReflectValue.Merge(dst, src)
}
func (m *ReflectValue) XXX_Size() int {
	return xxx_messageInfo_ReflectValue.Size(m)
}
func (m *ReflectValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ReflectValue.DiscardUnknown(m)
}

var xxx_messageInfo_ReflectValue proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Query)(nil), "proto.Query")
	proto.RegisterType((*QueryMulti)(nil), "proto.QueryMulti")
	proto.RegisterType((*ServiceInitRequest)(nil), "proto.ServiceInitRequest")
	proto.RegisterType((*ServiceControl)(nil), "proto.ServiceControl")
	proto.RegisterType((*MutationControl)(nil), "proto.MutationControl")
	proto.RegisterType((*DiscoveryEvent)(nil), "proto.DiscoveryEvent")
	proto.RegisterType((*ReflectValue)(nil), "proto.ReflectValue")
	proto.RegisterEnum("proto.ServiceControl_Command", ServiceControl_Command_name, ServiceControl_Command_value)
	proto.RegisterEnum("proto.MutationControl_Type", MutationControl_Type_name, MutationControl_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	// Query language
	// TODO: create API for bulk CRUD operations
	QueryCreate(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error)
	QueryRead(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error)
	QueryReadDsc(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error)
	QueryUpdate(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error)
	QueryUpdateDsc(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error)
	QueryDelete(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error)
	QueryReadAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*QueryMulti, error)
	QueryReadAllDsc(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*QueryMulti, error)
	QueryDeleteAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*QueryMulti, error)
	// Service management
	ServiceInit(ctx context.Context, in *ServiceInitRequest, opts ...grpc.CallOption) (API_ServiceInitClient, error)
	// Mutation/Discover management
	MutationInit(ctx context.Context, in *ServiceInitRequest, opts ...grpc.CallOption) (API_MutationInitClient, error)
	// Discovery management
	DiscoveryInit(ctx context.Context, opts ...grpc.CallOption) (API_DiscoveryInitClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) QueryCreate(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error) {
	out := new(Query)
	err := c.cc.Invoke(ctx, "/proto.API/QueryCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryRead(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error) {
	out := new(Query)
	err := c.cc.Invoke(ctx, "/proto.API/QueryRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryReadDsc(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error) {
	out := new(Query)
	err := c.cc.Invoke(ctx, "/proto.API/QueryReadDsc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryUpdate(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error) {
	out := new(Query)
	err := c.cc.Invoke(ctx, "/proto.API/QueryUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryUpdateDsc(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error) {
	out := new(Query)
	err := c.cc.Invoke(ctx, "/proto.API/QueryUpdateDsc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryDelete(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Query, error) {
	out := new(Query)
	err := c.cc.Invoke(ctx, "/proto.API/QueryDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryReadAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*QueryMulti, error) {
	out := new(QueryMulti)
	err := c.cc.Invoke(ctx, "/proto.API/QueryReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryReadAllDsc(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*QueryMulti, error) {
	out := new(QueryMulti)
	err := c.cc.Invoke(ctx, "/proto.API/QueryReadAllDsc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) QueryDeleteAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*QueryMulti, error) {
	out := new(QueryMulti)
	err := c.cc.Invoke(ctx, "/proto.API/QueryDeleteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ServiceInit(ctx context.Context, in *ServiceInitRequest, opts ...grpc.CallOption) (API_ServiceInitClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/proto.API/ServiceInit", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIServiceInitClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_ServiceInitClient interface {
	Recv() (*ServiceControl, error)
	grpc.ClientStream
}

type aPIServiceInitClient struct {
	grpc.ClientStream
}

func (x *aPIServiceInitClient) Recv() (*ServiceControl, error) {
	m := new(ServiceControl)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) MutationInit(ctx context.Context, in *ServiceInitRequest, opts ...grpc.CallOption) (API_MutationInitClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[1], "/proto.API/MutationInit", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIMutationInitClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_MutationInitClient interface {
	Recv() (*MutationControl, error)
	grpc.ClientStream
}

type aPIMutationInitClient struct {
	grpc.ClientStream
}

func (x *aPIMutationInitClient) Recv() (*MutationControl, error) {
	m := new(MutationControl)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) DiscoveryInit(ctx context.Context, opts ...grpc.CallOption) (API_DiscoveryInitClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[2], "/proto.API/DiscoveryInit", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIDiscoveryInitClient{stream}
	return x, nil
}

type API_DiscoveryInitClient interface {
	Send(*DiscoveryEvent) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type aPIDiscoveryInitClient struct {
	grpc.ClientStream
}

func (x *aPIDiscoveryInitClient) Send(m *DiscoveryEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPIDiscoveryInitClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	// Query language
	// TODO: create API for bulk CRUD operations
	QueryCreate(context.Context, *Query) (*Query, error)
	QueryRead(context.Context, *Query) (*Query, error)
	QueryReadDsc(context.Context, *Query) (*Query, error)
	QueryUpdate(context.Context, *Query) (*Query, error)
	QueryUpdateDsc(context.Context, *Query) (*Query, error)
	QueryDelete(context.Context, *Query) (*Query, error)
	QueryReadAll(context.Context, *empty.Empty) (*QueryMulti, error)
	QueryReadAllDsc(context.Context, *empty.Empty) (*QueryMulti, error)
	QueryDeleteAll(context.Context, *empty.Empty) (*QueryMulti, error)
	// Service management
	ServiceInit(*ServiceInitRequest, API_ServiceInitServer) error
	// Mutation/Discover management
	MutationInit(*ServiceInitRequest, API_MutationInitServer) error
	// Discovery management
	DiscoveryInit(API_DiscoveryInitServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_QueryCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryCreate(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryRead(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryReadDsc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryReadDsc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryReadDsc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryReadDsc(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryUpdate(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryUpdateDsc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryUpdateDsc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryUpdateDsc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryUpdateDsc(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryDelete(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryReadAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryReadAllDsc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryReadAllDsc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryReadAllDsc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryReadAllDsc(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_QueryDeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).QueryDeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.API/QueryDeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).QueryDeleteAll(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ServiceInit_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServiceInitRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).ServiceInit(m, &aPIServiceInitServer{stream})
}

type API_ServiceInitServer interface {
	Send(*ServiceControl) error
	grpc.ServerStream
}

type aPIServiceInitServer struct {
	grpc.ServerStream
}

func (x *aPIServiceInitServer) Send(m *ServiceControl) error {
	return x.ServerStream.SendMsg(m)
}

func _API_MutationInit_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServiceInitRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).MutationInit(m, &aPIMutationInitServer{stream})
}

type API_MutationInitServer interface {
	Send(*MutationControl) error
	grpc.ServerStream
}

type aPIMutationInitServer struct {
	grpc.ServerStream
}

func (x *aPIMutationInitServer) Send(m *MutationControl) error {
	return x.ServerStream.SendMsg(m)
}

func _API_DiscoveryInit_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).DiscoveryInit(&aPIDiscoveryInitServer{stream})
}

type API_DiscoveryInitServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*DiscoveryEvent, error)
	grpc.ServerStream
}

type aPIDiscoveryInitServer struct {
	grpc.ServerStream
}

func (x *aPIDiscoveryInitServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPIDiscoveryInitServer) Recv() (*DiscoveryEvent, error) {
	m := new(DiscoveryEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryCreate",
			Handler:    _API_QueryCreate_Handler,
		},
		{
			MethodName: "QueryRead",
			Handler:    _API_QueryRead_Handler,
		},
		{
			MethodName: "QueryReadDsc",
			Handler:    _API_QueryReadDsc_Handler,
		},
		{
			MethodName: "QueryUpdate",
			Handler:    _API_QueryUpdate_Handler,
		},
		{
			MethodName: "QueryUpdateDsc",
			Handler:    _API_QueryUpdateDsc_Handler,
		},
		{
			MethodName: "QueryDelete",
			Handler:    _API_QueryDelete_Handler,
		},
		{
			MethodName: "QueryReadAll",
			Handler:    _API_QueryReadAll_Handler,
		},
		{
			MethodName: "QueryReadAllDsc",
			Handler:    _API_QueryReadAllDsc_Handler,
		},
		{
			MethodName: "QueryDeleteAll",
			Handler:    _API_QueryDeleteAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServiceInit",
			Handler:       _API_ServiceInit_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MutationInit",
			Handler:       _API_MutationInit_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DiscoveryInit",
			Handler:       _API_DiscoveryInit_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "API.proto",
}

func init() { proto.RegisterFile("API.proto", fileDescriptor_API_6081312aeb8bb432) }

var fileDescriptor_API_6081312aeb8bb432 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xda, 0x40,
	0x10, 0xb5, 0xb9, 0xc6, 0x03, 0x71, 0xe8, 0xb6, 0x8d, 0x08, 0x51, 0xa4, 0x64, 0x1f, 0x2a, 0xa2,
	0x44, 0x4e, 0x45, 0x2b, 0xf5, 0xa1, 0x17, 0x89, 0x00, 0x52, 0x2c, 0x35, 0x29, 0xdd, 0x98, 0xbe,
	0x56, 0x8e, 0x77, 0x41, 0x96, 0x8c, 0x97, 0x98, 0x35, 0x92, 0x3f, 0xa7, 0xfd, 0x9d, 0xfe, 0x54,
	0xe5, 0xf5, 0x3a, 0x35, 0xb9, 0x88, 0xe6, 0x09, 0xef, 0xcc, 0x99, 0x33, 0x73, 0x76, 0xcf, 0x00,
	0x46, 0x7f, 0x6c, 0x5b, 0x8b, 0x88, 0x0b, 0x8e, 0xaa, 0xf2, 0xa7, 0x03, 0x57, 0x9c, 0xb2, 0x2c,
	0xd4, 0xd9, 0x9b, 0x71, 0x3e, 0x0b, 0xd8, 0x99, 0x3c, 0xdd, 0xc4, 0xd3, 0x33, 0x37, 0x4c, 0x54,
	0x6a, 0xff, 0x7e, 0x6a, 0x34, 0x5f, 0x08, 0x95, 0xc4, 0x02, 0xaa, 0xdf, 0x63, 0x16, 0x25, 0xa8,
	0x05, 0xe5, 0x09, 0xf9, 0xda, 0xd6, 0x0f, 0xf5, 0xae, 0x41, 0xd2, 0x4f, 0x74, 0x04, 0x95, 0x90,
	0x53, 0xd6, 0x2e, 0x1d, 0xea, 0xdd, 0x46, 0xaf, 0x91, 0x15, 0x58, 0x69, 0xcf, 0x0b, 0x8d, 0xc8,
	0x14, 0x3a, 0x81, 0xea, 0xca, 0x0d, 0x62, 0xd6, 0x2e, 0x4b, 0xcc, 0x4b, 0x85, 0x21, 0x6c, 0x1a,
	0x30, 0x4f, 0xfc, 0x48, 0x53, 0x17, 0x1a, 0xc9, 0x30, 0xe7, 0x06, 0xd4, 0x17, 0x6e, 0x12, 0x70,
	0x97, 0xe2, 0xf7, 0x00, 0xb2, 0xeb, 0x65, 0x1c, 0x08, 0x1f, 0xbd, 0x81, 0xfa, 0x6d, 0xcc, 0x22,
	0x9f, 0x2d, 0xdb, 0xfa, 0x61, 0xb9, 0xdb, 0xe8, 0x35, 0x15, 0x8f, 0xc4, 0x90, 0x3c, 0x89, 0x3f,
	0x01, 0xba, 0x66, 0xd1, 0xca, 0xf7, 0x98, 0x1d, 0xfa, 0x82, 0xb0, 0xdb, 0x98, 0x2d, 0x05, 0x32,
	0xa1, 0xe4, 0x53, 0x35, 0x77, 0xc9, 0xa7, 0x68, 0x17, 0x6a, 0x73, 0x4e, 0xe3, 0x20, 0x1b, 0xdc,
	0x20, 0xea, 0x84, 0x7f, 0xeb, 0x60, 0xaa, 0xf2, 0x01, 0x0f, 0x45, 0xc4, 0x03, 0xf4, 0x01, 0xea,
	0x1e, 0x9f, 0xcf, 0xdd, 0x30, 0xab, 0x37, 0x7b, 0x07, 0xaa, 0xf1, 0x3a, 0xce, 0x1a, 0x64, 0x20,
	0x92, 0xa3, 0xd1, 0x29, 0xd4, 0x3c, 0x1e, 0x4e, 0xfd, 0x99, 0xba, 0x9c, 0x57, 0x56, 0x76, 0xc7,
	0x56, 0x7e, 0xc7, 0x56, 0x3f, 0x4c, 0x88, 0xc2, 0xe0, 0x63, 0xa8, 0x2b, 0x06, 0xb4, 0x05, 0x95,
	0x6b, 0xe7, 0xdb, 0xb8, 0xa5, 0x21, 0x80, 0xda, 0x64, 0x3c, 0xec, 0x3b, 0xa3, 0x96, 0x9e, 0x46,
	0xed, 0x2b, 0xdb, 0x69, 0x95, 0xf0, 0x1f, 0x1d, 0x76, 0x2e, 0x63, 0xe1, 0x0a, 0x9f, 0x87, 0xf9,
	0x94, 0xff, 0x04, 0xe9, 0x45, 0x41, 0x4a, 0x78, 0xe9, 0x4e, 0xf8, 0x19, 0x54, 0x44, 0xb2, 0xc8,
	0xde, 0xc2, 0xec, 0xed, 0x2b, 0x29, 0xf7, 0xd8, 0x2c, 0x27, 0x59, 0x30, 0x22, 0x81, 0xe8, 0x00,
	0xca, 0xde, 0x74, 0xd6, 0xae, 0x3c, 0x78, 0x5f, 0x92, 0xc6, 0xd3, 0x34, 0x5d, 0x7a, 0xed, 0xea,
	0x23, 0x69, 0xba, 0xf4, 0xf0, 0x11, 0x54, 0x52, 0xae, 0x54, 0xc8, 0xe5, 0xc4, 0x49, 0x85, 0x68,
	0x68, 0x1b, 0x0c, 0xfb, 0xca, 0x19, 0x11, 0x32, 0x19, 0x3b, 0x2d, 0x1d, 0x4f, 0xc0, 0x1c, 0xfa,
	0x4b, 0x8f, 0xaf, 0x58, 0x94, 0x8c, 0x56, 0x2c, 0x14, 0x4f, 0x6a, 0x69, 0x41, 0x39, 0x8e, 0x02,
	0x25, 0x26, 0xfd, 0x44, 0x7b, 0xb0, 0x25, 0x6d, 0xf3, 0xd3, 0xa7, 0x52, 0x91, 0x41, 0xea, 0xf2,
	0x6c, 0x53, 0x6c, 0x42, 0xb3, 0xe8, 0xb0, 0xde, 0xaf, 0x2a, 0x94, 0xfb, 0x63, 0x1b, 0x9d, 0x40,
	0x43, 0x3a, 0x66, 0x10, 0x31, 0x57, 0x30, 0xb4, 0xe6, 0xa2, 0xce, 0xda, 0x09, 0x6b, 0xe8, 0x18,
	0x8c, 0xcc, 0x5e, 0xcc, 0xa5, 0x1b, 0xa0, 0xa7, 0xd0, 0xbc, 0x83, 0x0e, 0x97, 0xde, 0x06, 0x74,
	0x3e, 0xc5, 0x64, 0x41, 0x37, 0x4f, 0x61, 0x81, 0x59, 0x00, 0xff, 0x3f, 0xf9, 0x90, 0x05, 0x6c,
	0x23, 0xf9, 0xc7, 0xc2, 0xdc, 0xfd, 0x20, 0x40, 0xbb, 0x0f, 0x5c, 0x2a, 0xff, 0x09, 0x3a, 0x2f,
	0x8a, 0x75, 0x72, 0x25, 0xb1, 0x86, 0xbe, 0xc0, 0x4e, 0xb1, 0x38, 0x1d, 0xed, 0x59, 0xf5, 0x9f,
	0x95, 0xb2, 0x6c, 0xd2, 0x67, 0xb7, 0x1f, 0x40, 0xa3, 0xb0, 0xeb, 0x68, 0x6f, 0x7d, 0x31, 0x0b,
	0xfb, 0xdf, 0x79, 0xfd, 0xe8, 0xce, 0x62, 0xed, 0xad, 0x8e, 0x46, 0xd0, 0xcc, 0xed, 0xbf, 0x89,
	0x65, 0xf7, 0xf1, 0x75, 0x91, 0x34, 0xe7, 0xb0, 0x7d, 0x67, 0x63, 0xc9, 0x93, 0xb7, 0x5c, 0x37,
	0x77, 0xe7, 0x09, 0x81, 0x58, 0xeb, 0xea, 0x37, 0x35, 0x19, 0x7b, 0xf7, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x50, 0x8f, 0xfa, 0x0a, 0xc6, 0x05, 0x00, 0x00,
}
