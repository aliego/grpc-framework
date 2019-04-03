// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/index/index.proto

package index

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message definition.
type IndexRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexRequest) Reset()         { *m = IndexRequest{} }
func (m *IndexRequest) String() string { return proto.CompactTextString(m) }
func (*IndexRequest) ProtoMessage()    {}
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5b8486092d5d3cd, []int{0}
}

func (m *IndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexRequest.Unmarshal(m, b)
}
func (m *IndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexRequest.Marshal(b, m, deterministic)
}
func (m *IndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRequest.Merge(m, src)
}
func (m *IndexRequest) XXX_Size() int {
	return xxx_messageInfo_IndexRequest.Size(m)
}
func (m *IndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRequest proto.InternalMessageInfo

func (m *IndexRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message definition.
type IndexReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexReply) Reset()         { *m = IndexReply{} }
func (m *IndexReply) String() string { return proto.CompactTextString(m) }
func (*IndexReply) ProtoMessage()    {}
func (*IndexReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5b8486092d5d3cd, []int{1}
}

func (m *IndexReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexReply.Unmarshal(m, b)
}
func (m *IndexReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexReply.Marshal(b, m, deterministic)
}
func (m *IndexReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexReply.Merge(m, src)
}
func (m *IndexReply) XXX_Size() int {
	return xxx_messageInfo_IndexReply.Size(m)
}
func (m *IndexReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexReply.DiscardUnknown(m)
}

var xxx_messageInfo_IndexReply proto.InternalMessageInfo

func (m *IndexReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*IndexRequest)(nil), "index.IndexRequest")
	proto.RegisterType((*IndexReply)(nil), "index.IndexReply")
}

func init() { proto.RegisterFile("proto/index/index.proto", fileDescriptor_b5b8486092d5d3cd) }

var fileDescriptor_b5b8486092d5d3cd = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcc, 0x4b, 0x49, 0xad, 0x80, 0x90, 0x7a, 0x60, 0x11, 0x21, 0x56, 0x30, 0x47,
	0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f,
	0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x18, 0xa2, 0x48, 0x49, 0x89, 0x8b, 0xc7, 0x13, 0xa4,
	0x2c, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x52, 0xe3, 0xe2, 0x82, 0xaa, 0x29, 0xc8,
	0xa9, 0x14, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87, 0x29, 0x82, 0x71, 0x8d,
	0x3a, 0x18, 0xb9, 0x58, 0xc1, 0x0a, 0x85, 0x1c, 0xb8, 0x58, 0x42, 0x40, 0xa6, 0x09, 0xeb, 0x41,
	0x1c, 0x84, 0x6c, 0x85, 0x94, 0x20, 0xaa, 0x60, 0x41, 0x4e, 0xa5, 0x92, 0x40, 0xd3, 0xe5, 0x27,
	0x93, 0x99, 0xb8, 0x94, 0x58, 0xf5, 0x4b, 0x52, 0x8b, 0x4b, 0xac, 0x18, 0xb5, 0x84, 0x6c, 0xb9,
	0x58, 0x1c, 0x73, 0x32, 0x53, 0x89, 0x36, 0x81, 0x17, 0x6c, 0x02, 0xbb, 0x10, 0xab, 0x7e, 0x62,
	0x4e, 0x66, 0xaa, 0x93, 0x1a, 0x97, 0x58, 0x66, 0xbe, 0x5e, 0x7a, 0x51, 0x41, 0xb2, 0x5e, 0x6a,
	0x45, 0x62, 0x6e, 0x41, 0x4e, 0x6a, 0x31, 0x44, 0x8f, 0x13, 0xc4, 0x2b, 0x01, 0x20, 0xcf, 0x07,
	0x30, 0x26, 0xb1, 0x81, 0x43, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x84, 0x56, 0xc7, 0xc0,
	0x45, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IndexClient is the client API for Index service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IndexClient interface {
	// Test
	Test(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexReply, error)
	Alie(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexReply, error)
}

type indexClient struct {
	cc *grpc.ClientConn
}

func NewIndexClient(cc *grpc.ClientConn) IndexClient {
	return &indexClient{cc}
}

func (c *indexClient) Test(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexReply, error) {
	out := new(IndexReply)
	err := c.cc.Invoke(ctx, "/index.Index/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexClient) Alie(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexReply, error) {
	out := new(IndexReply)
	err := c.cc.Invoke(ctx, "/index.Index/Alie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServer is the server API for Index service.
type IndexServer interface {
	// Test
	Test(context.Context, *IndexRequest) (*IndexReply, error)
	Alie(context.Context, *IndexRequest) (*IndexReply, error)
}

// UnimplementedIndexServer can be embedded to have forward compatible implementations.
type UnimplementedIndexServer struct {
}

func (*UnimplementedIndexServer) Test(ctx context.Context, req *IndexRequest) (*IndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (*UnimplementedIndexServer) Alie(ctx context.Context, req *IndexRequest) (*IndexReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alie not implemented")
}

func RegisterIndexServer(s *grpc.Server, srv IndexServer) {
	s.RegisterService(&_Index_serviceDesc, srv)
}

func _Index_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Test(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Index_Alie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServer).Alie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/index.Index/Alie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServer).Alie(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Index_serviceDesc = grpc.ServiceDesc{
	ServiceName: "index.Index",
	HandlerType: (*IndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Index_Test_Handler,
		},
		{
			MethodName: "Alie",
			Handler:    _Index_Alie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/index/index.proto",
}
