// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hub_to_agent.proto

package idl

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpgradePrimariesRequest struct {
	SourceBinDir         string         `protobuf:"bytes,1,opt,name=SourceBinDir" json:"SourceBinDir,omitempty"`
	TargetBinDir         string         `protobuf:"bytes,2,opt,name=TargetBinDir" json:"TargetBinDir,omitempty"`
	TargetVersion        string         `protobuf:"bytes,3,opt,name=TargetVersion" json:"TargetVersion,omitempty"`
	DataDirPairs         []*DataDirPair `protobuf:"bytes,4,rep,name=DataDirPairs" json:"DataDirPairs,omitempty"`
	CheckOnly            bool           `protobuf:"varint,5,opt,name=CheckOnly" json:"CheckOnly,omitempty"`
	UseLinkMode          bool           `protobuf:"varint,6,opt,name=UseLinkMode" json:"UseLinkMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpgradePrimariesRequest) Reset()         { *m = UpgradePrimariesRequest{} }
func (m *UpgradePrimariesRequest) String() string { return proto.CompactTextString(m) }
func (*UpgradePrimariesRequest) ProtoMessage()    {}
func (*UpgradePrimariesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_aad269811b68e3c9, []int{0}
}
func (m *UpgradePrimariesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradePrimariesRequest.Unmarshal(m, b)
}
func (m *UpgradePrimariesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradePrimariesRequest.Marshal(b, m, deterministic)
}
func (dst *UpgradePrimariesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradePrimariesRequest.Merge(dst, src)
}
func (m *UpgradePrimariesRequest) XXX_Size() int {
	return xxx_messageInfo_UpgradePrimariesRequest.Size(m)
}
func (m *UpgradePrimariesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradePrimariesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradePrimariesRequest proto.InternalMessageInfo

func (m *UpgradePrimariesRequest) GetSourceBinDir() string {
	if m != nil {
		return m.SourceBinDir
	}
	return ""
}

func (m *UpgradePrimariesRequest) GetTargetBinDir() string {
	if m != nil {
		return m.TargetBinDir
	}
	return ""
}

func (m *UpgradePrimariesRequest) GetTargetVersion() string {
	if m != nil {
		return m.TargetVersion
	}
	return ""
}

func (m *UpgradePrimariesRequest) GetDataDirPairs() []*DataDirPair {
	if m != nil {
		return m.DataDirPairs
	}
	return nil
}

func (m *UpgradePrimariesRequest) GetCheckOnly() bool {
	if m != nil {
		return m.CheckOnly
	}
	return false
}

func (m *UpgradePrimariesRequest) GetUseLinkMode() bool {
	if m != nil {
		return m.UseLinkMode
	}
	return false
}

type DataDirPair struct {
	SourceDataDir        string   `protobuf:"bytes,1,opt,name=SourceDataDir" json:"SourceDataDir,omitempty"`
	TargetDataDir        string   `protobuf:"bytes,2,opt,name=TargetDataDir" json:"TargetDataDir,omitempty"`
	SourcePort           int32    `protobuf:"varint,3,opt,name=SourcePort" json:"SourcePort,omitempty"`
	TargetPort           int32    `protobuf:"varint,4,opt,name=TargetPort" json:"TargetPort,omitempty"`
	Content              int32    `protobuf:"varint,5,opt,name=Content" json:"Content,omitempty"`
	DBID                 int32    `protobuf:"varint,6,opt,name=DBID" json:"DBID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataDirPair) Reset()         { *m = DataDirPair{} }
func (m *DataDirPair) String() string { return proto.CompactTextString(m) }
func (*DataDirPair) ProtoMessage()    {}
func (*DataDirPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_aad269811b68e3c9, []int{1}
}
func (m *DataDirPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataDirPair.Unmarshal(m, b)
}
func (m *DataDirPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataDirPair.Marshal(b, m, deterministic)
}
func (dst *DataDirPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataDirPair.Merge(dst, src)
}
func (m *DataDirPair) XXX_Size() int {
	return xxx_messageInfo_DataDirPair.Size(m)
}
func (m *DataDirPair) XXX_DiscardUnknown() {
	xxx_messageInfo_DataDirPair.DiscardUnknown(m)
}

var xxx_messageInfo_DataDirPair proto.InternalMessageInfo

func (m *DataDirPair) GetSourceDataDir() string {
	if m != nil {
		return m.SourceDataDir
	}
	return ""
}

func (m *DataDirPair) GetTargetDataDir() string {
	if m != nil {
		return m.TargetDataDir
	}
	return ""
}

func (m *DataDirPair) GetSourcePort() int32 {
	if m != nil {
		return m.SourcePort
	}
	return 0
}

func (m *DataDirPair) GetTargetPort() int32 {
	if m != nil {
		return m.TargetPort
	}
	return 0
}

func (m *DataDirPair) GetContent() int32 {
	if m != nil {
		return m.Content
	}
	return 0
}

func (m *DataDirPair) GetDBID() int32 {
	if m != nil {
		return m.DBID
	}
	return 0
}

type UpgradePrimariesReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpgradePrimariesReply) Reset()         { *m = UpgradePrimariesReply{} }
func (m *UpgradePrimariesReply) String() string { return proto.CompactTextString(m) }
func (*UpgradePrimariesReply) ProtoMessage()    {}
func (*UpgradePrimariesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_aad269811b68e3c9, []int{2}
}
func (m *UpgradePrimariesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradePrimariesReply.Unmarshal(m, b)
}
func (m *UpgradePrimariesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradePrimariesReply.Marshal(b, m, deterministic)
}
func (dst *UpgradePrimariesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradePrimariesReply.Merge(dst, src)
}
func (m *UpgradePrimariesReply) XXX_Size() int {
	return xxx_messageInfo_UpgradePrimariesReply.Size(m)
}
func (m *UpgradePrimariesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradePrimariesReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradePrimariesReply proto.InternalMessageInfo

type CreateSegmentDataDirRequest struct {
	Datadirs             []string `protobuf:"bytes,1,rep,name=datadirs" json:"datadirs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSegmentDataDirRequest) Reset()         { *m = CreateSegmentDataDirRequest{} }
func (m *CreateSegmentDataDirRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSegmentDataDirRequest) ProtoMessage()    {}
func (*CreateSegmentDataDirRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_aad269811b68e3c9, []int{3}
}
func (m *CreateSegmentDataDirRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSegmentDataDirRequest.Unmarshal(m, b)
}
func (m *CreateSegmentDataDirRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSegmentDataDirRequest.Marshal(b, m, deterministic)
}
func (dst *CreateSegmentDataDirRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSegmentDataDirRequest.Merge(dst, src)
}
func (m *CreateSegmentDataDirRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSegmentDataDirRequest.Size(m)
}
func (m *CreateSegmentDataDirRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSegmentDataDirRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSegmentDataDirRequest proto.InternalMessageInfo

func (m *CreateSegmentDataDirRequest) GetDatadirs() []string {
	if m != nil {
		return m.Datadirs
	}
	return nil
}

type CreateSegmentDataDirReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSegmentDataDirReply) Reset()         { *m = CreateSegmentDataDirReply{} }
func (m *CreateSegmentDataDirReply) String() string { return proto.CompactTextString(m) }
func (*CreateSegmentDataDirReply) ProtoMessage()    {}
func (*CreateSegmentDataDirReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_aad269811b68e3c9, []int{4}
}
func (m *CreateSegmentDataDirReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSegmentDataDirReply.Unmarshal(m, b)
}
func (m *CreateSegmentDataDirReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSegmentDataDirReply.Marshal(b, m, deterministic)
}
func (dst *CreateSegmentDataDirReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSegmentDataDirReply.Merge(dst, src)
}
func (m *CreateSegmentDataDirReply) XXX_Size() int {
	return xxx_messageInfo_CreateSegmentDataDirReply.Size(m)
}
func (m *CreateSegmentDataDirReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSegmentDataDirReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSegmentDataDirReply proto.InternalMessageInfo

type StopAgentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAgentRequest) Reset()         { *m = StopAgentRequest{} }
func (m *StopAgentRequest) String() string { return proto.CompactTextString(m) }
func (*StopAgentRequest) ProtoMessage()    {}
func (*StopAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_aad269811b68e3c9, []int{5}
}
func (m *StopAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAgentRequest.Unmarshal(m, b)
}
func (m *StopAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAgentRequest.Marshal(b, m, deterministic)
}
func (dst *StopAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAgentRequest.Merge(dst, src)
}
func (m *StopAgentRequest) XXX_Size() int {
	return xxx_messageInfo_StopAgentRequest.Size(m)
}
func (m *StopAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopAgentRequest proto.InternalMessageInfo

type StopAgentReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAgentReply) Reset()         { *m = StopAgentReply{} }
func (m *StopAgentReply) String() string { return proto.CompactTextString(m) }
func (*StopAgentReply) ProtoMessage()    {}
func (*StopAgentReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_aad269811b68e3c9, []int{6}
}
func (m *StopAgentReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAgentReply.Unmarshal(m, b)
}
func (m *StopAgentReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAgentReply.Marshal(b, m, deterministic)
}
func (dst *StopAgentReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAgentReply.Merge(dst, src)
}
func (m *StopAgentReply) XXX_Size() int {
	return xxx_messageInfo_StopAgentReply.Size(m)
}
func (m *StopAgentReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAgentReply.DiscardUnknown(m)
}

var xxx_messageInfo_StopAgentReply proto.InternalMessageInfo

type CheckSegmentDiskSpaceRequest struct {
	Request              *CheckDiskSpaceRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Datadirs             []string               `protobuf:"bytes,2,rep,name=datadirs" json:"datadirs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CheckSegmentDiskSpaceRequest) Reset()         { *m = CheckSegmentDiskSpaceRequest{} }
func (m *CheckSegmentDiskSpaceRequest) String() string { return proto.CompactTextString(m) }
func (*CheckSegmentDiskSpaceRequest) ProtoMessage()    {}
func (*CheckSegmentDiskSpaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_aad269811b68e3c9, []int{7}
}
func (m *CheckSegmentDiskSpaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSegmentDiskSpaceRequest.Unmarshal(m, b)
}
func (m *CheckSegmentDiskSpaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSegmentDiskSpaceRequest.Marshal(b, m, deterministic)
}
func (dst *CheckSegmentDiskSpaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSegmentDiskSpaceRequest.Merge(dst, src)
}
func (m *CheckSegmentDiskSpaceRequest) XXX_Size() int {
	return xxx_messageInfo_CheckSegmentDiskSpaceRequest.Size(m)
}
func (m *CheckSegmentDiskSpaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSegmentDiskSpaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSegmentDiskSpaceRequest proto.InternalMessageInfo

func (m *CheckSegmentDiskSpaceRequest) GetRequest() *CheckDiskSpaceRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *CheckSegmentDiskSpaceRequest) GetDatadirs() []string {
	if m != nil {
		return m.Datadirs
	}
	return nil
}

func init() {
	proto.RegisterType((*UpgradePrimariesRequest)(nil), "idl.UpgradePrimariesRequest")
	proto.RegisterType((*DataDirPair)(nil), "idl.DataDirPair")
	proto.RegisterType((*UpgradePrimariesReply)(nil), "idl.UpgradePrimariesReply")
	proto.RegisterType((*CreateSegmentDataDirRequest)(nil), "idl.CreateSegmentDataDirRequest")
	proto.RegisterType((*CreateSegmentDataDirReply)(nil), "idl.CreateSegmentDataDirReply")
	proto.RegisterType((*StopAgentRequest)(nil), "idl.StopAgentRequest")
	proto.RegisterType((*StopAgentReply)(nil), "idl.StopAgentReply")
	proto.RegisterType((*CheckSegmentDiskSpaceRequest)(nil), "idl.CheckSegmentDiskSpaceRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agent service

type AgentClient interface {
	CheckDiskSpace(ctx context.Context, in *CheckSegmentDiskSpaceRequest, opts ...grpc.CallOption) (*CheckDiskSpaceReply, error)
	UpgradePrimaries(ctx context.Context, in *UpgradePrimariesRequest, opts ...grpc.CallOption) (*UpgradePrimariesReply, error)
	CreateSegmentDataDirectories(ctx context.Context, in *CreateSegmentDataDirRequest, opts ...grpc.CallOption) (*CreateSegmentDataDirReply, error)
	StopAgent(ctx context.Context, in *StopAgentRequest, opts ...grpc.CallOption) (*StopAgentReply, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CheckDiskSpace(ctx context.Context, in *CheckSegmentDiskSpaceRequest, opts ...grpc.CallOption) (*CheckDiskSpaceReply, error) {
	out := new(CheckDiskSpaceReply)
	err := grpc.Invoke(ctx, "/idl.Agent/CheckDiskSpace", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpgradePrimaries(ctx context.Context, in *UpgradePrimariesRequest, opts ...grpc.CallOption) (*UpgradePrimariesReply, error) {
	out := new(UpgradePrimariesReply)
	err := grpc.Invoke(ctx, "/idl.Agent/UpgradePrimaries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CreateSegmentDataDirectories(ctx context.Context, in *CreateSegmentDataDirRequest, opts ...grpc.CallOption) (*CreateSegmentDataDirReply, error) {
	out := new(CreateSegmentDataDirReply)
	err := grpc.Invoke(ctx, "/idl.Agent/CreateSegmentDataDirectories", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) StopAgent(ctx context.Context, in *StopAgentRequest, opts ...grpc.CallOption) (*StopAgentReply, error) {
	out := new(StopAgentReply)
	err := grpc.Invoke(ctx, "/idl.Agent/StopAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agent service

type AgentServer interface {
	CheckDiskSpace(context.Context, *CheckSegmentDiskSpaceRequest) (*CheckDiskSpaceReply, error)
	UpgradePrimaries(context.Context, *UpgradePrimariesRequest) (*UpgradePrimariesReply, error)
	CreateSegmentDataDirectories(context.Context, *CreateSegmentDataDirRequest) (*CreateSegmentDataDirReply, error)
	StopAgent(context.Context, *StopAgentRequest) (*StopAgentReply, error)
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_CheckDiskSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSegmentDiskSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CheckDiskSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/CheckDiskSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CheckDiskSpace(ctx, req.(*CheckSegmentDiskSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpgradePrimaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradePrimariesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpgradePrimaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/UpgradePrimaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpgradePrimaries(ctx, req.(*UpgradePrimariesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CreateSegmentDataDirectories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSegmentDataDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateSegmentDataDirectories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/CreateSegmentDataDirectories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateSegmentDataDirectories(ctx, req.(*CreateSegmentDataDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_StopAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).StopAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/StopAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).StopAgent(ctx, req.(*StopAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckDiskSpace",
			Handler:    _Agent_CheckDiskSpace_Handler,
		},
		{
			MethodName: "UpgradePrimaries",
			Handler:    _Agent_UpgradePrimaries_Handler,
		},
		{
			MethodName: "CreateSegmentDataDirectories",
			Handler:    _Agent_CreateSegmentDataDirectories_Handler,
		},
		{
			MethodName: "StopAgent",
			Handler:    _Agent_StopAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hub_to_agent.proto",
}

func init() { proto.RegisterFile("hub_to_agent.proto", fileDescriptor_hub_to_agent_aad269811b68e3c9) }

var fileDescriptor_hub_to_agent_aad269811b68e3c9 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x41, 0x6f, 0xda, 0x30,
	0x14, 0x5e, 0x80, 0xb4, 0xe5, 0xd1, 0x55, 0x91, 0xa7, 0xaa, 0x59, 0x8a, 0xaa, 0x2c, 0xda, 0x81,
	0x13, 0x07, 0xd6, 0x4b, 0x8f, 0x2b, 0xb9, 0x4c, 0xda, 0x56, 0x64, 0xd6, 0x5d, 0x2b, 0x93, 0x3c,
	0x81, 0x45, 0x1a, 0x67, 0x8e, 0x39, 0xf4, 0x17, 0xf6, 0x57, 0x4d, 0x9a, 0x6c, 0x13, 0x48, 0x18,
	0xe9, 0xcd, 0xfe, 0xde, 0xf7, 0x3e, 0xbf, 0xef, 0x7d, 0x10, 0x20, 0xab, 0xcd, 0xe2, 0x49, 0x89,
	0x27, 0xb6, 0xc4, 0x5c, 0x8d, 0x0b, 0x29, 0x94, 0x20, 0x5d, 0x9e, 0x66, 0x81, 0x97, 0x64, 0x5c,
	0x17, 0x56, 0x9b, 0x85, 0x85, 0xa3, 0xbf, 0x0e, 0x5c, 0x3d, 0x16, 0x4b, 0xc9, 0x52, 0x9c, 0x49,
	0xfe, 0xcc, 0x24, 0xc7, 0x92, 0xe2, 0x9f, 0x0d, 0x96, 0x8a, 0x44, 0x70, 0x3e, 0x17, 0x1b, 0x99,
	0xe0, 0x3d, 0xcf, 0x63, 0x2e, 0x7d, 0x27, 0x74, 0x46, 0x7d, 0xda, 0xc0, 0x34, 0xe7, 0x17, 0x93,
	0x4b, 0x54, 0x5b, 0x4e, 0xc7, 0x72, 0xea, 0x18, 0xf9, 0x0c, 0xef, 0xed, 0xfd, 0x37, 0xca, 0x92,
	0x8b, 0xdc, 0xef, 0x1a, 0x52, 0x13, 0x24, 0xb7, 0x70, 0x1e, 0x33, 0xc5, 0x62, 0x2e, 0x67, 0x8c,
	0xcb, 0xd2, 0xef, 0x85, 0xdd, 0xd1, 0x60, 0xe2, 0x8d, 0x79, 0x9a, 0x8d, 0x6b, 0x05, 0xda, 0x60,
	0x91, 0x21, 0xf4, 0xa7, 0x2b, 0x4c, 0xd6, 0x0f, 0x79, 0xf6, 0xe2, 0xbb, 0xa1, 0x33, 0x3a, 0xa3,
	0x7b, 0x80, 0x84, 0x30, 0x78, 0x2c, 0xf1, 0x3b, 0xcf, 0xd7, 0x3f, 0x44, 0x8a, 0xfe, 0x89, 0xa9,
	0xd7, 0xa1, 0xe8, 0xd5, 0x81, 0x41, 0x4d, 0x50, 0xcf, 0x6a, 0xfd, 0x6d, 0xc1, 0xad, 0xe9, 0x26,
	0xb8, 0x77, 0x54, 0xb1, 0x3a, 0x75, 0x47, 0x15, 0xeb, 0x06, 0xc0, 0xb6, 0xcd, 0x84, 0x54, 0xc6,
	0xb4, 0x4b, 0x6b, 0x88, 0xae, 0xdb, 0x06, 0x53, 0xef, 0xd9, 0xfa, 0x1e, 0x21, 0x3e, 0x9c, 0x4e,
	0x45, 0xae, 0x30, 0x57, 0xc6, 0x99, 0x4b, 0xab, 0x2b, 0x21, 0xd0, 0x8b, 0xef, 0xbf, 0xc5, 0xc6,
	0x90, 0x4b, 0xcd, 0x39, 0xba, 0x82, 0xcb, 0xff, 0x83, 0x2c, 0xb2, 0x97, 0xe8, 0x0e, 0xae, 0xa7,
	0x12, 0x99, 0xc2, 0x39, 0x2e, 0x9f, 0x31, 0xaf, 0xc6, 0xab, 0x52, 0x0e, 0xe0, 0x2c, 0x65, 0x8a,
	0xa5, 0x7a, 0xe7, 0x4e, 0xd8, 0x1d, 0xf5, 0xe9, 0xee, 0x1e, 0x5d, 0xc3, 0xc7, 0xe3, 0xad, 0x5a,
	0x97, 0x80, 0x37, 0x57, 0xa2, 0xf8, 0xaa, 0x7f, 0x64, 0x5b, 0xb1, 0xc8, 0x83, 0x8b, 0x1a, 0xa6,
	0x59, 0x05, 0x0c, 0x4d, 0x1e, 0x95, 0x02, 0x2f, 0xd7, 0xf3, 0x82, 0x25, 0x58, 0x3d, 0x7f, 0x0b,
	0xa7, 0xd2, 0x1e, 0xcd, 0xaa, 0x07, 0x93, 0xc0, 0x24, 0x6e, 0x7a, 0x0e, 0xc9, 0xb4, 0xa2, 0x36,
	0x86, 0xee, 0x34, 0x87, 0x9e, 0xbc, 0x76, 0xc0, 0x35, 0x03, 0x90, 0x07, 0xb8, 0x68, 0xea, 0x90,
	0x4f, 0x7b, 0xf1, 0x96, 0x81, 0x02, 0xff, 0xe8, 0xfb, 0xda, 0xca, 0x3b, 0xf2, 0x13, 0xbc, 0xc3,
	0x1d, 0x93, 0xa1, 0xe1, 0xb7, 0xfc, 0x87, 0x82, 0xa0, 0xa5, 0x6a, 0xf5, 0x16, 0x30, 0x3c, 0xb6,
	0x5f, 0x4c, 0x94, 0x30, 0xda, 0xa1, 0x9d, 0xa5, 0x3d, 0xbd, 0xe0, 0xe6, 0x0d, 0x86, 0x7d, 0xe3,
	0x0e, 0xfa, 0xbb, 0x48, 0xc8, 0xa5, 0xa1, 0x1f, 0xc6, 0x16, 0x7c, 0x38, 0x84, 0x4d, 0xeb, 0xe2,
	0xc4, 0x7c, 0x23, 0xbe, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xf0, 0xee, 0xdd, 0x50, 0x04,
	0x00, 0x00,
}
