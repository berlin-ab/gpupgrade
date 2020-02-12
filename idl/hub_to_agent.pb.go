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
	MasterBackupDir      string         `protobuf:"bytes,7,opt,name=MasterBackupDir" json:"MasterBackupDir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpgradePrimariesRequest) Reset()         { *m = UpgradePrimariesRequest{} }
func (m *UpgradePrimariesRequest) String() string { return proto.CompactTextString(m) }
func (*UpgradePrimariesRequest) ProtoMessage()    {}
func (*UpgradePrimariesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{0}
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

func (m *UpgradePrimariesRequest) GetMasterBackupDir() string {
	if m != nil {
		return m.MasterBackupDir
	}
	return ""
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
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{1}
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
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{2}
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
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{3}
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
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{4}
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

type RenamePair struct {
	Src                  string   `protobuf:"bytes,1,opt,name=Src" json:"Src,omitempty"`
	Dst                  string   `protobuf:"bytes,2,opt,name=Dst" json:"Dst,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenamePair) Reset()         { *m = RenamePair{} }
func (m *RenamePair) String() string { return proto.CompactTextString(m) }
func (*RenamePair) ProtoMessage()    {}
func (*RenamePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{5}
}
func (m *RenamePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenamePair.Unmarshal(m, b)
}
func (m *RenamePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenamePair.Marshal(b, m, deterministic)
}
func (dst *RenamePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenamePair.Merge(dst, src)
}
func (m *RenamePair) XXX_Size() int {
	return xxx_messageInfo_RenamePair.Size(m)
}
func (m *RenamePair) XXX_DiscardUnknown() {
	xxx_messageInfo_RenamePair.DiscardUnknown(m)
}

var xxx_messageInfo_RenamePair proto.InternalMessageInfo

func (m *RenamePair) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *RenamePair) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

type RenameDirectoriesRequest struct {
	Pairs                []*RenamePair `protobuf:"bytes,1,rep,name=Pairs" json:"Pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RenameDirectoriesRequest) Reset()         { *m = RenameDirectoriesRequest{} }
func (m *RenameDirectoriesRequest) String() string { return proto.CompactTextString(m) }
func (*RenameDirectoriesRequest) ProtoMessage()    {}
func (*RenameDirectoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{6}
}
func (m *RenameDirectoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameDirectoriesRequest.Unmarshal(m, b)
}
func (m *RenameDirectoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameDirectoriesRequest.Marshal(b, m, deterministic)
}
func (dst *RenameDirectoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameDirectoriesRequest.Merge(dst, src)
}
func (m *RenameDirectoriesRequest) XXX_Size() int {
	return xxx_messageInfo_RenameDirectoriesRequest.Size(m)
}
func (m *RenameDirectoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameDirectoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenameDirectoriesRequest proto.InternalMessageInfo

func (m *RenameDirectoriesRequest) GetPairs() []*RenamePair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type RenameDirectoriesReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameDirectoriesReply) Reset()         { *m = RenameDirectoriesReply{} }
func (m *RenameDirectoriesReply) String() string { return proto.CompactTextString(m) }
func (*RenameDirectoriesReply) ProtoMessage()    {}
func (*RenameDirectoriesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{7}
}
func (m *RenameDirectoriesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameDirectoriesReply.Unmarshal(m, b)
}
func (m *RenameDirectoriesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameDirectoriesReply.Marshal(b, m, deterministic)
}
func (dst *RenameDirectoriesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameDirectoriesReply.Merge(dst, src)
}
func (m *RenameDirectoriesReply) XXX_Size() int {
	return xxx_messageInfo_RenameDirectoriesReply.Size(m)
}
func (m *RenameDirectoriesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameDirectoriesReply.DiscardUnknown(m)
}

var xxx_messageInfo_RenameDirectoriesReply proto.InternalMessageInfo

type StopAgentRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAgentRequest) Reset()         { *m = StopAgentRequest{} }
func (m *StopAgentRequest) String() string { return proto.CompactTextString(m) }
func (*StopAgentRequest) ProtoMessage()    {}
func (*StopAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{8}
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
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{9}
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
	return fileDescriptor_hub_to_agent_26f98673fce61429, []int{10}
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
	proto.RegisterType((*RenamePair)(nil), "idl.RenamePair")
	proto.RegisterType((*RenameDirectoriesRequest)(nil), "idl.RenameDirectoriesRequest")
	proto.RegisterType((*RenameDirectoriesReply)(nil), "idl.RenameDirectoriesReply")
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
	RenameDirectories(ctx context.Context, in *RenameDirectoriesRequest, opts ...grpc.CallOption) (*RenameDirectoriesReply, error)
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

func (c *agentClient) RenameDirectories(ctx context.Context, in *RenameDirectoriesRequest, opts ...grpc.CallOption) (*RenameDirectoriesReply, error) {
	out := new(RenameDirectoriesReply)
	err := grpc.Invoke(ctx, "/idl.Agent/RenameDirectories", in, out, c.cc, opts...)
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
	RenameDirectories(context.Context, *RenameDirectoriesRequest) (*RenameDirectoriesReply, error)
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

func _Agent_RenameDirectories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameDirectoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).RenameDirectories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/RenameDirectories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).RenameDirectories(ctx, req.(*RenameDirectoriesRequest))
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
			MethodName: "RenameDirectories",
			Handler:    _Agent_RenameDirectories_Handler,
		},
		{
			MethodName: "StopAgent",
			Handler:    _Agent_StopAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hub_to_agent.proto",
}

func init() { proto.RegisterFile("hub_to_agent.proto", fileDescriptor_hub_to_agent_26f98673fce61429) }

var fileDescriptor_hub_to_agent_26f98673fce61429 = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0xdd, 0x34, 0x93, 0xd2, 0x9a, 0x45, 0xa5, 0xc6, 0x09, 0x95, 0xb1, 0x40, 0xca,
	0x29, 0x42, 0xa1, 0x97, 0x1e, 0x9b, 0xf8, 0x82, 0x44, 0x69, 0x70, 0x28, 0xd7, 0x6a, 0x63, 0x8f,
	0x92, 0x55, 0x12, 0xdb, 0xac, 0x37, 0x87, 0xbe, 0x0a, 0x2f, 0xc0, 0xa3, 0xf0, 0x5a, 0x68, 0x77,
	0xed, 0xc6, 0xce, 0x0f, 0xb7, 0xdd, 0x6f, 0xbe, 0x99, 0x9d, 0xf9, 0xe6, 0xb3, 0x81, 0xcc, 0xd7,
	0xd3, 0x47, 0x91, 0x3e, 0xd2, 0x19, 0x26, 0xa2, 0x9f, 0xf1, 0x54, 0xa4, 0xc4, 0x64, 0xf1, 0xd2,
	0xb5, 0xa3, 0x25, 0x93, 0x81, 0xf9, 0x7a, 0xaa, 0x61, 0xff, 0x4f, 0x03, 0x2e, 0x1f, 0xb2, 0x19,
	0xa7, 0x31, 0x8e, 0x39, 0x5b, 0x51, 0xce, 0x30, 0x0f, 0xf1, 0xd7, 0x1a, 0x73, 0x41, 0x7c, 0x38,
	0x9d, 0xa4, 0x6b, 0x1e, 0xe1, 0x90, 0x25, 0x01, 0xe3, 0x8e, 0xe1, 0x19, 0xbd, 0x56, 0x58, 0xc3,
	0x24, 0xe7, 0x07, 0xe5, 0x33, 0x14, 0x05, 0xa7, 0xa1, 0x39, 0x55, 0x8c, 0x7c, 0x80, 0x97, 0xfa,
	0xfe, 0x13, 0x79, 0xce, 0xd2, 0xc4, 0x31, 0x15, 0xa9, 0x0e, 0x92, 0x6b, 0x38, 0x0d, 0xa8, 0xa0,
	0x01, 0xe3, 0x63, 0xca, 0x78, 0xee, 0x1c, 0x79, 0x66, 0xaf, 0x3d, 0xb0, 0xfb, 0x2c, 0x5e, 0xf6,
	0x2b, 0x81, 0xb0, 0xc6, 0x22, 0x5d, 0x68, 0x8d, 0xe6, 0x18, 0x2d, 0xee, 0x93, 0xe5, 0x93, 0x63,
	0x79, 0x46, 0xef, 0x24, 0xdc, 0x00, 0xc4, 0x83, 0xf6, 0x43, 0x8e, 0x5f, 0x59, 0xb2, 0xb8, 0x4b,
	0x63, 0x74, 0x8e, 0x55, 0xbc, 0x0a, 0x91, 0x1e, 0x9c, 0xdf, 0xd1, 0x5c, 0x20, 0x1f, 0xd2, 0x68,
	0xb1, 0xce, 0xe4, 0x08, 0x4d, 0xd5, 0xdd, 0x36, 0xec, 0xff, 0x35, 0xa0, 0x5d, 0x79, 0x5a, 0x4e,
	0xa5, 0x95, 0x28, 0xc0, 0x42, 0x9e, 0x3a, 0xb8, 0x99, 0xbd, 0x64, 0x35, 0xaa, 0xb3, 0x97, 0xac,
	0x2b, 0x00, 0x9d, 0x36, 0x4e, 0xb9, 0x50, 0xf2, 0x58, 0x61, 0x05, 0x91, 0x71, 0x9d, 0xa0, 0xe2,
	0x47, 0x3a, 0xbe, 0x41, 0x88, 0x03, 0xcd, 0x51, 0x9a, 0x08, 0x4c, 0x84, 0xd2, 0xc0, 0x0a, 0xcb,
	0x2b, 0x21, 0x70, 0x14, 0x0c, 0xbf, 0x04, 0x6a, 0x74, 0x2b, 0x54, 0x67, 0xff, 0x12, 0x2e, 0x76,
	0x57, 0x9e, 0x2d, 0x9f, 0xfc, 0x1b, 0xe8, 0x8c, 0x38, 0x52, 0x81, 0x13, 0x9c, 0xad, 0x30, 0x29,
	0xdb, 0x2b, 0xfd, 0xe0, 0xc2, 0x49, 0x4c, 0x05, 0x8d, 0xe5, 0x76, 0x0c, 0xcf, 0xec, 0xb5, 0xc2,
	0xe7, 0xbb, 0xdf, 0x81, 0xb7, 0xfb, 0x53, 0x65, 0xdd, 0x4f, 0x00, 0x21, 0x26, 0x74, 0x85, 0x4a,
	0x38, 0x1b, 0xcc, 0x09, 0x8f, 0x0a, 0xb9, 0xe4, 0x51, 0x22, 0x41, 0x2e, 0x0a, 0x69, 0xe4, 0xd1,
	0xbf, 0x05, 0x47, 0x67, 0x04, 0x8c, 0x63, 0x24, 0xd2, 0xaa, 0x2d, 0x3f, 0x82, 0xa5, 0x1d, 0x62,
	0x28, 0x87, 0x9c, 0x2b, 0x87, 0x6c, 0xea, 0x87, 0x3a, 0xea, 0x3b, 0xf0, 0x66, 0x4f, 0x09, 0xd9,
	0x0e, 0x01, 0x7b, 0x22, 0xd2, 0xec, 0x56, 0x7e, 0x1d, 0x45, 0x51, 0xdf, 0x86, 0xb3, 0x0a, 0x26,
	0x59, 0x19, 0x74, 0x95, 0x91, 0xca, 0x81, 0x58, 0xbe, 0x98, 0x64, 0x34, 0xc2, 0xb2, 0x8d, 0x6b,
	0x68, 0x72, 0x7d, 0x54, 0xa3, 0xb4, 0x07, 0xae, 0x6a, 0x44, 0xe5, 0x6c, 0x93, 0xc3, 0x92, 0x5a,
	0xd3, 0xb0, 0x51, 0xd7, 0x70, 0xf0, 0xdb, 0x04, 0x4b, 0x35, 0x40, 0xee, 0xe1, 0xac, 0x5e, 0x87,
	0xbc, 0xdf, 0x14, 0x3f, 0xd0, 0x90, 0xeb, 0xec, 0x7d, 0x5f, 0x8e, 0xf2, 0x82, 0x7c, 0x03, 0x7b,
	0x7b, 0xe5, 0xa4, 0xab, 0xf8, 0x07, 0x3e, 0x7e, 0xd7, 0x3d, 0x10, 0xd5, 0xf5, 0xa6, 0xd0, 0xdd,
	0xb7, 0xee, 0x52, 0x67, 0xe2, 0xe9, 0x5e, 0x0e, 0x9b, 0xc9, 0xbd, 0xfa, 0x0f, 0x43, 0xbf, 0xf1,
	0x1d, 0x5e, 0xed, 0x2c, 0x90, 0xbc, 0xab, 0x6c, 0x7b, 0xd7, 0x1b, 0x6e, 0xe7, 0x50, 0x58, 0x97,
	0xbc, 0x81, 0xd6, 0xf3, 0x96, 0xc9, 0x85, 0xe2, 0x6e, 0x3b, 0xc1, 0x7d, 0xbd, 0x0d, 0xab, 0xd4,
	0xe9, 0xb1, 0xfa, 0x5f, 0x7e, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xea, 0x74, 0xcb, 0x5c,
	0x05, 0x00, 0x00,
}
