// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dht.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FindSuccessorRequest struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindSuccessorRequest) Reset()         { *m = FindSuccessorRequest{} }
func (m *FindSuccessorRequest) String() string { return proto.CompactTextString(m) }
func (*FindSuccessorRequest) ProtoMessage()    {}
func (*FindSuccessorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{0}
}

func (m *FindSuccessorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindSuccessorRequest.Unmarshal(m, b)
}
func (m *FindSuccessorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindSuccessorRequest.Marshal(b, m, deterministic)
}
func (m *FindSuccessorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindSuccessorRequest.Merge(m, src)
}
func (m *FindSuccessorRequest) XXX_Size() int {
	return xxx_messageInfo_FindSuccessorRequest.Size(m)
}
func (m *FindSuccessorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindSuccessorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindSuccessorRequest proto.InternalMessageInfo

func (m *FindSuccessorRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type Node struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{1}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Node) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type NotifyRequest struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyRequest) Reset()         { *m = NotifyRequest{} }
func (m *NotifyRequest) String() string { return proto.CompactTextString(m) }
func (*NotifyRequest) ProtoMessage()    {}
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{2}
}

func (m *NotifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyRequest.Unmarshal(m, b)
}
func (m *NotifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyRequest.Marshal(b, m, deterministic)
}
func (m *NotifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyRequest.Merge(m, src)
}
func (m *NotifyRequest) XXX_Size() int {
	return xxx_messageInfo_NotifyRequest.Size(m)
}
func (m *NotifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyRequest proto.InternalMessageInfo

func (m *NotifyRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NotifyRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type PutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRequest) Reset()         { *m = PutRequest{} }
func (m *PutRequest) String() string { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()    {}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{4}
}

func (m *PutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRequest.Unmarshal(m, b)
}
func (m *PutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRequest.Marshal(b, m, deterministic)
}
func (m *PutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRequest.Merge(m, src)
}
func (m *PutRequest) XXX_Size() int {
	return xxx_messageInfo_PutRequest.Size(m)
}
func (m *PutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRequest proto.InternalMessageInfo

type DelRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelRequest) Reset()         { *m = DelRequest{} }
func (m *DelRequest) String() string { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()    {}
func (*DelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{5}
}

func (m *DelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelRequest.Unmarshal(m, b)
}
func (m *DelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelRequest.Marshal(b, m, deterministic)
}
func (m *DelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelRequest.Merge(m, src)
}
func (m *DelRequest) XXX_Size() int {
	return xxx_messageInfo_DelRequest.Size(m)
}
func (m *DelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelRequest proto.InternalMessageInfo

type Result struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{6}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{7}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{8}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReply) Reset()         { *m = PingReply{} }
func (m *PingReply) String() string { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()    {}
func (*PingReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{9}
}

func (m *PingReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReply.Unmarshal(m, b)
}
func (m *PingReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReply.Marshal(b, m, deterministic)
}
func (m *PingReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReply.Merge(m, src)
}
func (m *PingReply) XXX_Size() int {
	return xxx_messageInfo_PingReply.Size(m)
}
func (m *PingReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReply.DiscardUnknown(m)
}

var xxx_messageInfo_PingReply proto.InternalMessageInfo

type FindNodeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNodeRequest) Reset()         { *m = FindNodeRequest{} }
func (m *FindNodeRequest) String() string { return proto.CompactTextString(m) }
func (*FindNodeRequest) ProtoMessage()    {}
func (*FindNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{10}
}

func (m *FindNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNodeRequest.Unmarshal(m, b)
}
func (m *FindNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNodeRequest.Marshal(b, m, deterministic)
}
func (m *FindNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNodeRequest.Merge(m, src)
}
func (m *FindNodeRequest) XXX_Size() int {
	return xxx_messageInfo_FindNodeRequest.Size(m)
}
func (m *FindNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindNodeRequest proto.InternalMessageInfo

type FindNodeReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNodeReply) Reset()         { *m = FindNodeReply{} }
func (m *FindNodeReply) String() string { return proto.CompactTextString(m) }
func (*FindNodeReply) ProtoMessage()    {}
func (*FindNodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{11}
}

func (m *FindNodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNodeReply.Unmarshal(m, b)
}
func (m *FindNodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNodeReply.Marshal(b, m, deterministic)
}
func (m *FindNodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNodeReply.Merge(m, src)
}
func (m *FindNodeReply) XXX_Size() int {
	return xxx_messageInfo_FindNodeReply.Size(m)
}
func (m *FindNodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_FindNodeReply proto.InternalMessageInfo

type FindValueRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValueRequest) Reset()         { *m = FindValueRequest{} }
func (m *FindValueRequest) String() string { return proto.CompactTextString(m) }
func (*FindValueRequest) ProtoMessage()    {}
func (*FindValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{12}
}

func (m *FindValueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindValueRequest.Unmarshal(m, b)
}
func (m *FindValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindValueRequest.Marshal(b, m, deterministic)
}
func (m *FindValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindValueRequest.Merge(m, src)
}
func (m *FindValueRequest) XXX_Size() int {
	return xxx_messageInfo_FindValueRequest.Size(m)
}
func (m *FindValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindValueRequest proto.InternalMessageInfo

type FindValueReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValueReply) Reset()         { *m = FindValueReply{} }
func (m *FindValueReply) String() string { return proto.CompactTextString(m) }
func (*FindValueReply) ProtoMessage()    {}
func (*FindValueReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{13}
}

func (m *FindValueReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindValueReply.Unmarshal(m, b)
}
func (m *FindValueReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindValueReply.Marshal(b, m, deterministic)
}
func (m *FindValueReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindValueReply.Merge(m, src)
}
func (m *FindValueReply) XXX_Size() int {
	return xxx_messageInfo_FindValueReply.Size(m)
}
func (m *FindValueReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FindValueReply.DiscardUnknown(m)
}

var xxx_messageInfo_FindValueReply proto.InternalMessageInfo

type StoreRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{14}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

type StoreReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreReply) Reset()         { *m = StoreReply{} }
func (m *StoreReply) String() string { return proto.CompactTextString(m) }
func (*StoreReply) ProtoMessage()    {}
func (*StoreReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_616a434b24c97ff4, []int{15}
}

func (m *StoreReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreReply.Unmarshal(m, b)
}
func (m *StoreReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreReply.Marshal(b, m, deterministic)
}
func (m *StoreReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreReply.Merge(m, src)
}
func (m *StoreReply) XXX_Size() int {
	return xxx_messageInfo_StoreReply.Size(m)
}
func (m *StoreReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreReply.DiscardUnknown(m)
}

var xxx_messageInfo_StoreReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FindSuccessorRequest)(nil), "protos.FindSuccessorRequest")
	proto.RegisterType((*Node)(nil), "protos.Node")
	proto.RegisterType((*NotifyRequest)(nil), "protos.NotifyRequest")
	proto.RegisterType((*GetRequest)(nil), "protos.GetRequest")
	proto.RegisterType((*PutRequest)(nil), "protos.PutRequest")
	proto.RegisterType((*DelRequest)(nil), "protos.DelRequest")
	proto.RegisterType((*Result)(nil), "protos.Result")
	proto.RegisterType((*Void)(nil), "protos.Void")
	proto.RegisterType((*PingRequest)(nil), "protos.PingRequest")
	proto.RegisterType((*PingReply)(nil), "protos.PingReply")
	proto.RegisterType((*FindNodeRequest)(nil), "protos.FindNodeRequest")
	proto.RegisterType((*FindNodeReply)(nil), "protos.FindNodeReply")
	proto.RegisterType((*FindValueRequest)(nil), "protos.FindValueRequest")
	proto.RegisterType((*FindValueReply)(nil), "protos.FindValueReply")
	proto.RegisterType((*StoreRequest)(nil), "protos.StoreRequest")
	proto.RegisterType((*StoreReply)(nil), "protos.StoreReply")
}

func init() { proto.RegisterFile("dht.proto", fileDescriptor_616a434b24c97ff4) }

var fileDescriptor_616a434b24c97ff4 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x86, 0x29, 0xed, 0x36, 0xf6, 0x6c, 0xe9, 0xb2, 0xc7, 0xdd, 0x95, 0x34, 0xc6, 0x90, 0xb9,
	0x30, 0x44, 0x12, 0x54, 0xb8, 0xd5, 0x78, 0x21, 0x11, 0x13, 0x13, 0xd2, 0x14, 0xc3, 0x3d, 0x32,
	0xa3, 0x34, 0x36, 0x0c, 0xb6, 0xd3, 0x8b, 0xbe, 0x83, 0x4f, 0xe2, 0xab, 0xf9, 0x12, 0x66, 0x3a,
	0x9d, 0x76, 0x84, 0xee, 0x15, 0x73, 0xce, 0xff, 0xff, 0xa7, 0x33, 0xdf, 0x0c, 0xe0, 0xd1, 0x83,
	0x98, 0x9d, 0x32, 0x2e, 0x38, 0xba, 0xd5, 0x4f, 0x4e, 0x5e, 0xc2, 0xdd, 0xa7, 0xe4, 0x48, 0x37,
	0xc5, 0x7e, 0xcf, 0xf2, 0x9c, 0x67, 0x31, 0xfb, 0x55, 0xb0, 0x5c, 0x60, 0x00, 0xfd, 0x84, 0x8e,
	0xac, 0xb1, 0x35, 0xf1, 0xe3, 0x7e, 0x42, 0xc9, 0x2b, 0x70, 0xd6, 0x9c, 0xb2, 0xf3, 0x3e, 0x22,
	0x38, 0x3b, 0x4a, 0xb3, 0x51, 0x7f, 0x6c, 0x4d, 0xbc, 0xb8, 0x5a, 0x93, 0x05, 0x0c, 0xd6, 0x5c,
	0x24, 0xdf, 0xcb, 0x47, 0x86, 0x75, 0x86, 0x5e, 0x00, 0xac, 0x98, 0xd0, 0x89, 0x21, 0xd8, 0x3f,
	0x59, 0x59, 0x45, 0xbc, 0x58, 0x2e, 0x89, 0x0f, 0x10, 0x15, 0x5a, 0x97, 0xd5, 0x92, 0xa5, 0xba,
	0x1a, 0x83, 0x1b, 0xb3, 0xbc, 0x48, 0x05, 0x3e, 0x80, 0x9b, 0x55, 0xab, 0x3a, 0x5a, 0x57, 0xc4,
	0x05, 0x67, 0xcb, 0x13, 0x4a, 0x06, 0x70, 0x1d, 0x25, 0xc7, 0x1f, 0x3a, 0x78, 0x0d, 0x9e, 0x2a,
	0x4f, 0x69, 0x49, 0x6e, 0xe1, 0x46, 0xa2, 0x90, 0xc7, 0xd4, 0xfa, 0x0d, 0x0c, 0xda, 0x96, 0xf4,
	0x20, 0x0c, 0x65, 0x63, 0xbb, 0x4b, 0x8b, 0xc6, 0x34, 0x84, 0xc0, 0xe8, 0x49, 0x57, 0x00, 0xfe,
	0x46, 0xf0, 0x8c, 0x19, 0xbb, 0xad, 0xeb, 0x53, 0x5a, 0xce, 0xff, 0x58, 0x70, 0xf5, 0xf1, 0xc0,
	0x33, 0x8a, 0xef, 0xd5, 0xf8, 0x06, 0x3e, 0x3e, 0x57, 0xb7, 0x93, 0xcf, 0xba, 0xee, 0x24, 0xf4,
	0xb5, 0x2a, 0xf7, 0x43, 0x7a, 0xf8, 0x16, 0x5c, 0xc5, 0x19, 0xef, 0x5b, 0xc5, 0xe0, 0x1e, 0x06,
	0xba, 0xad, 0xe8, 0x90, 0x1e, 0xbe, 0x56, 0x67, 0x8c, 0x32, 0x46, 0x59, 0xfd, 0xcd, 0x66, 0xaa,
	0x04, 0x74, 0xfe, 0x8d, 0xf9, 0x5f, 0x0b, 0xec, 0x2f, 0x3b, 0x8a, 0x6f, 0xc0, 0x91, 0xa4, 0xf0,
	0xa9, 0xd6, 0x0d, 0x8c, 0xe1, 0xed, 0xff, 0x4d, 0x89, 0xa0, 0x87, 0xef, 0xe0, 0x89, 0x66, 0x87,
	0xcf, 0xcc, 0x73, 0x19, 0x80, 0xc3, 0xfb, 0x4b, 0x41, 0xa5, 0x3f, 0x80, 0xd7, 0x40, 0xc5, 0x91,
	0xe9, 0x32, 0xd9, 0x87, 0x0f, 0x1d, 0x8a, 0x1a, 0xb0, 0x80, 0xab, 0x8a, 0x39, 0xde, 0x69, 0x8b,
	0x79, 0x25, 0x21, 0x9e, 0x75, 0xab, 0xd0, 0xfc, 0xb7, 0x05, 0xf6, 0xf2, 0xf3, 0x57, 0x9c, 0x82,
	0xbd, 0x62, 0x02, 0x1b, 0x53, 0xfb, 0x32, 0x3b, 0x98, 0x4e, 0xc1, 0x8e, 0x0a, 0xc3, 0xdc, 0x3e,
	0xd3, 0x6e, 0xf3, 0x92, 0xa5, 0xad, 0xb9, 0x7d, 0xc5, 0x97, 0xe6, 0x6f, 0xea, 0x4f, 0xba, 0xf8,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x13, 0x8a, 0x76, 0xb8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChordClient is the client API for Chord service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChordClient interface {
	FindSuccessor(ctx context.Context, in *FindSuccessorRequest, opts ...grpc.CallOption) (*Node, error)
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*Result, error)
	FindPredecessor(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Node, error)
}

type chordClient struct {
	cc *grpc.ClientConn
}

func NewChordClient(cc *grpc.ClientConn) ChordClient {
	return &chordClient{cc}
}

func (c *chordClient) FindSuccessor(ctx context.Context, in *FindSuccessorRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/protos.Chord/FindSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/protos.Chord/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chordClient) FindPredecessor(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/protos.Chord/FindPredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChordServer is the server API for Chord service.
type ChordServer interface {
	FindSuccessor(context.Context, *FindSuccessorRequest) (*Node, error)
	Notify(context.Context, *NotifyRequest) (*Result, error)
	FindPredecessor(context.Context, *Void) (*Node, error)
}

// UnimplementedChordServer can be embedded to have forward compatible implementations.
type UnimplementedChordServer struct {
}

func (*UnimplementedChordServer) FindSuccessor(ctx context.Context, req *FindSuccessorRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (*UnimplementedChordServer) Notify(ctx context.Context, req *NotifyRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedChordServer) FindPredecessor(ctx context.Context, req *Void) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPredecessor not implemented")
}

func RegisterChordServer(s *grpc.Server, srv ChordServer) {
	s.RegisterService(&_Chord_serviceDesc, srv)
}

func _Chord_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSuccessorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindSuccessor(ctx, req.(*FindSuccessorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chord_FindPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChordServer).FindPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Chord/FindPredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChordServer).FindPredecessor(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chord_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Chord",
	HandlerType: (*ChordServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSuccessor",
			Handler:    _Chord_FindSuccessor_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Chord_Notify_Handler,
		},
		{
			MethodName: "FindPredecessor",
			Handler:    _Chord_FindPredecessor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dht.proto",
}

// KadClient is the client API for Kad service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KadClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	FindNode(ctx context.Context, in *FindNodeRequest, opts ...grpc.CallOption) (*FindNodeReply, error)
	FindValue(ctx context.Context, in *FindValueRequest, opts ...grpc.CallOption) (*FindValueReply, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreReply, error)
}

type kadClient struct {
	cc *grpc.ClientConn
}

func NewKadClient(cc *grpc.ClientConn) KadClient {
	return &kadClient{cc}
}

func (c *kadClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, "/protos.Kad/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kadClient) FindNode(ctx context.Context, in *FindNodeRequest, opts ...grpc.CallOption) (*FindNodeReply, error) {
	out := new(FindNodeReply)
	err := c.cc.Invoke(ctx, "/protos.Kad/FindNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kadClient) FindValue(ctx context.Context, in *FindValueRequest, opts ...grpc.CallOption) (*FindValueReply, error) {
	out := new(FindValueReply)
	err := c.cc.Invoke(ctx, "/protos.Kad/FindValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kadClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreReply, error) {
	out := new(StoreReply)
	err := c.cc.Invoke(ctx, "/protos.Kad/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KadServer is the server API for Kad service.
type KadServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	FindNode(context.Context, *FindNodeRequest) (*FindNodeReply, error)
	FindValue(context.Context, *FindValueRequest) (*FindValueReply, error)
	Store(context.Context, *StoreRequest) (*StoreReply, error)
}

// UnimplementedKadServer can be embedded to have forward compatible implementations.
type UnimplementedKadServer struct {
}

func (*UnimplementedKadServer) Ping(ctx context.Context, req *PingRequest) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedKadServer) FindNode(ctx context.Context, req *FindNodeRequest) (*FindNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNode not implemented")
}
func (*UnimplementedKadServer) FindValue(ctx context.Context, req *FindValueRequest) (*FindValueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindValue not implemented")
}
func (*UnimplementedKadServer) Store(ctx context.Context, req *StoreRequest) (*StoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}

func RegisterKadServer(s *grpc.Server, srv KadServer) {
	s.RegisterService(&_Kad_serviceDesc, srv)
}

func _Kad_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KadServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Kad/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KadServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kad_FindNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KadServer).FindNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Kad/FindNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KadServer).FindNode(ctx, req.(*FindNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kad_FindValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KadServer).FindValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Kad/FindValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KadServer).FindValue(ctx, req.(*FindValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kad_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KadServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Kad/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KadServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kad_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Kad",
	HandlerType: (*KadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Kad_Ping_Handler,
		},
		{
			MethodName: "FindNode",
			Handler:    _Kad_FindNode_Handler,
		},
		{
			MethodName: "FindValue",
			Handler:    _Kad_FindValue_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Kad_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dht.proto",
}

// DHTClient is the client API for DHT service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DHTClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Result, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Result, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*Result, error)
}

type dHTClient struct {
	cc *grpc.ClientConn
}

func NewDHTClient(cc *grpc.ClientConn) DHTClient {
	return &dHTClient{cc}
}

func (c *dHTClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/protos.DHT/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHTClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/protos.DHT/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHTClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/protos.DHT/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DHTServer is the server API for DHT service.
type DHTServer interface {
	Get(context.Context, *GetRequest) (*Result, error)
	Put(context.Context, *PutRequest) (*Result, error)
	Del(context.Context, *DelRequest) (*Result, error)
}

// UnimplementedDHTServer can be embedded to have forward compatible implementations.
type UnimplementedDHTServer struct {
}

func (*UnimplementedDHTServer) Get(ctx context.Context, req *GetRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDHTServer) Put(ctx context.Context, req *PutRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedDHTServer) Del(ctx context.Context, req *DelRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}

func RegisterDHTServer(s *grpc.Server, srv DHTServer) {
	s.RegisterService(&_DHT_serviceDesc, srv)
}

func _DHT_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.DHT/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHT_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.DHT/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHT_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.DHT/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DHT_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.DHT",
	HandlerType: (*DHTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DHT_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _DHT_Put_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _DHT_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dht.proto",
}
