// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calcproto/calculator.proto

package calcproto

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

type Sum struct {
	FirstNum             int32    `protobuf:"varint,1,opt,name=first_num,json=firstNum,proto3" json:"first_num,omitempty"`
	SecondNum            int32    `protobuf:"varint,2,opt,name=second_num,json=secondNum,proto3" json:"second_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_547f3c1a3dc9b0d3, []int{0}
}
func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (dst *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(dst, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetFirstNum() int32 {
	if m != nil {
		return m.FirstNum
	}
	return 0
}

func (m *Sum) GetSecondNum() int32 {
	if m != nil {
		return m.SecondNum
	}
	return 0
}

type SumRequest struct {
	Sum                  *Sum     `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_547f3c1a3dc9b0d3, []int{1}
}
func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (dst *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(dst, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSum() *Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_547f3c1a3dc9b0d3, []int{2}
}
func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (dst *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(dst, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Sum)(nil), "greet.Sum")
	proto.RegisterType((*SumRequest)(nil), "greet.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "greet.SumResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/greet.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calcproto/calculator.proto",
}

func init() {
	proto.RegisterFile("calcproto/calculator.proto", fileDescriptor_calculator_547f3c1a3dc9b0d3)
}

var fileDescriptor_calculator_547f3c1a3dc9b0d3 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xad, 0xa5, 0xc5, 0x4e, 0x4f, 0xe6, 0x20, 0x52, 0x15, 0x24, 0x20, 0x88, 0x87, 0x08,
	0xf5, 0xea, 0x45, 0x1f, 0xa0, 0x87, 0xf6, 0xe6, 0x45, 0xba, 0xd9, 0xd9, 0xa5, 0xd0, 0x34, 0xdd,
	0x24, 0xb3, 0xcf, 0xbf, 0x74, 0x52, 0x96, 0x9e, 0x92, 0x7c, 0x7f, 0xf8, 0xfe, 0x19, 0xa8, 0x74,
	0x3f, 0xea, 0xd9, 0xd9, 0x60, 0x3f, 0x97, 0x1b, 0x8d, 0x7d, 0xb0, 0x4e, 0x31, 0x10, 0xd9, 0xd1,
	0x21, 0x06, 0xf9, 0x03, 0x69, 0x47, 0x46, 0x3c, 0x41, 0x71, 0x18, 0x9c, 0x0f, 0xff, 0x13, 0x99,
	0xc7, 0xe4, 0x35, 0x79, 0xcf, 0xda, 0x3b, 0x06, 0x0d, 0x19, 0xf1, 0x02, 0xe0, 0x51, 0xdb, 0x69,
	0xcf, 0xe9, 0x2d, 0xa7, 0x45, 0x24, 0x0d, 0x19, 0xf9, 0x01, 0xd0, 0x91, 0x69, 0xf1, 0x44, 0xe8,
	0x83, 0x78, 0x86, 0xd4, 0xaf, 0x8e, 0xb2, 0x06, 0xc5, 0x2d, 0x6a, 0xc9, 0x17, 0x2c, 0xdf, 0xa0,
	0xe4, 0xbf, 0x7e, 0xb6, 0x93, 0x47, 0xf1, 0x00, 0xb9, 0x43, 0x4f, 0x63, 0x58, 0x3b, 0xd7, 0x57,
	0xfd, 0xcd, 0xca, 0x0e, 0xdd, 0x79, 0xd0, 0x28, 0x54, 0x9c, 0xf1, 0x7e, 0x23, 0x8b, 0x65, 0x95,
	0xd8, 0xa2, 0xe8, 0x94, 0x37, 0xbf, 0xe5, 0x5f, 0x71, 0x5d, 0x7c, 0x97, 0xf3, 0xf1, 0x75, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x25, 0x63, 0xa9, 0x40, 0x0c, 0x01, 0x00, 0x00,
}