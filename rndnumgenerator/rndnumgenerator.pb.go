// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rndnumgenerator.proto

package rndnumgenerator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// The request message containing the max value.
type RandomNumberRequest struct {
	MaxNumber            int32    `protobuf:"varint,1,opt,name=maxNumber,proto3" json:"maxNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomNumberRequest) Reset()         { *m = RandomNumberRequest{} }
func (m *RandomNumberRequest) String() string { return proto.CompactTextString(m) }
func (*RandomNumberRequest) ProtoMessage()    {}
func (*RandomNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2497df584b37615b, []int{0}
}

func (m *RandomNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomNumberRequest.Unmarshal(m, b)
}
func (m *RandomNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomNumberRequest.Marshal(b, m, deterministic)
}
func (m *RandomNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomNumberRequest.Merge(m, src)
}
func (m *RandomNumberRequest) XXX_Size() int {
	return xxx_messageInfo_RandomNumberRequest.Size(m)
}
func (m *RandomNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RandomNumberRequest proto.InternalMessageInfo

func (m *RandomNumberRequest) GetMaxNumber() int32 {
	if m != nil {
		return m.MaxNumber
	}
	return 0
}

// The response message containing the random number in string format
type RandomNumberReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomNumberReply) Reset()         { *m = RandomNumberReply{} }
func (m *RandomNumberReply) String() string { return proto.CompactTextString(m) }
func (*RandomNumberReply) ProtoMessage()    {}
func (*RandomNumberReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2497df584b37615b, []int{1}
}

func (m *RandomNumberReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomNumberReply.Unmarshal(m, b)
}
func (m *RandomNumberReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomNumberReply.Marshal(b, m, deterministic)
}
func (m *RandomNumberReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomNumberReply.Merge(m, src)
}
func (m *RandomNumberReply) XXX_Size() int {
	return xxx_messageInfo_RandomNumberReply.Size(m)
}
func (m *RandomNumberReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomNumberReply.DiscardUnknown(m)
}

var xxx_messageInfo_RandomNumberReply proto.InternalMessageInfo

func (m *RandomNumberReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RandomNumberRequest)(nil), "rndnumgenerator.RandomNumberRequest")
	proto.RegisterType((*RandomNumberReply)(nil), "rndnumgenerator.RandomNumberReply")
}

func init() { proto.RegisterFile("rndnumgenerator.proto", fileDescriptor_2497df584b37615b) }

var fileDescriptor_2497df584b37615b = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0xca, 0x4b, 0xc9,
	0x2b, 0xcd, 0x4d, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x56, 0x32, 0xe6, 0x12, 0x0e, 0x4a, 0xcc, 0x4b, 0xc9, 0xcf, 0xf5,
	0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe1, 0xe2,
	0xcc, 0x4d, 0xac, 0x80, 0x88, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x21, 0x04, 0x94, 0x74,
	0xb9, 0x04, 0x51, 0x35, 0x15, 0xe4, 0x54, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xa7, 0x82, 0x35, 0x70, 0x06, 0xc1, 0xb8, 0x46, 0xe5, 0x5c, 0xa2, 0xc8, 0xca, 0xdd, 0x61,
	0x96, 0x0b, 0xc5, 0x71, 0x09, 0x04, 0xa7, 0xe6, 0xa5, 0x20, 0x4b, 0x0a, 0xa9, 0xe8, 0xa1, 0xbb,
	0x1c, 0x8b, 0xfb, 0xa4, 0x94, 0x08, 0xa8, 0x2a, 0xc8, 0xa9, 0x54, 0x62, 0x70, 0x32, 0xe5, 0x92,
	0xce, 0xcc, 0xd7, 0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0xcb, 0xb6, 0x28, 0x46, 0xd7, 0xe2, 0x24, 0x12,
	0x94, 0x97, 0xe2, 0x57, 0x9a, 0x0b, 0x77, 0x4f, 0x00, 0x28, 0x88, 0x02, 0x18, 0x93, 0xd8, 0xc0,
	0x61, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xf4, 0x2f, 0x54, 0x44, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RandomNumberGeneratorClient is the client API for RandomNumberGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RandomNumberGeneratorClient interface {
	// Sends a random number
	SendRandomNumber(ctx context.Context, in *RandomNumberRequest, opts ...grpc.CallOption) (*RandomNumberReply, error)
}

type randomNumberGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewRandomNumberGeneratorClient(cc *grpc.ClientConn) RandomNumberGeneratorClient {
	return &randomNumberGeneratorClient{cc}
}

func (c *randomNumberGeneratorClient) SendRandomNumber(ctx context.Context, in *RandomNumberRequest, opts ...grpc.CallOption) (*RandomNumberReply, error) {
	out := new(RandomNumberReply)
	err := c.cc.Invoke(ctx, "/rndnumgenerator.RandomNumberGenerator/SendRandomNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RandomNumberGeneratorServer is the server API for RandomNumberGenerator service.
type RandomNumberGeneratorServer interface {
	// Sends a random number
	SendRandomNumber(context.Context, *RandomNumberRequest) (*RandomNumberReply, error)
}

func RegisterRandomNumberGeneratorServer(s *grpc.Server, srv RandomNumberGeneratorServer) {
	s.RegisterService(&_RandomNumberGenerator_serviceDesc, srv)
}

func _RandomNumberGenerator_SendRandomNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomNumberGeneratorServer).SendRandomNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rndnumgenerator.RandomNumberGenerator/SendRandomNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomNumberGeneratorServer).SendRandomNumber(ctx, req.(*RandomNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RandomNumberGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rndnumgenerator.RandomNumberGenerator",
	HandlerType: (*RandomNumberGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRandomNumber",
			Handler:    _RandomNumberGenerator_SendRandomNumber_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rndnumgenerator.proto",
}
