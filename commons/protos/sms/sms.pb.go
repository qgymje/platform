// Code generated by protoc-gen-go.
// source: sms/sms.proto
// DO NOT EDIT!

/*
Package sms is a generated protocol buffer package.

It is generated from these files:
	sms/sms.proto

It has these top-level messages:
	Phone
	Code
	PhoneCode
	Status
*/
package sms

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

type Phone struct {
	Phone string `protobuf:"bytes,1,opt,name=phone" json:"phone,omitempty"`
}

func (m *Phone) Reset()                    { *m = Phone{} }
func (m *Phone) String() string            { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()               {}
func (*Phone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Code struct {
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (m *Code) Reset()                    { *m = Code{} }
func (m *Code) String() string            { return proto.CompactTextString(m) }
func (*Code) ProtoMessage()               {}
func (*Code) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PhoneCode struct {
	Phone string `protobuf:"bytes,1,opt,name=phone" json:"phone,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
}

func (m *PhoneCode) Reset()                    { *m = PhoneCode{} }
func (m *PhoneCode) String() string            { return proto.CompactTextString(m) }
func (*PhoneCode) ProtoMessage()               {}
func (*PhoneCode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Status struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Phone)(nil), "sms.Phone")
	proto.RegisterType((*Code)(nil), "sms.Code")
	proto.RegisterType((*PhoneCode)(nil), "sms.PhoneCode")
	proto.RegisterType((*Status)(nil), "sms.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for SMS service

type SMSClient interface {
	Verify(ctx context.Context, in *PhoneCode, opts ...grpc.CallOption) (*Status, error)
}

type sMSClient struct {
	cc *grpc.ClientConn
}

func NewSMSClient(cc *grpc.ClientConn) SMSClient {
	return &sMSClient{cc}
}

func (c *sMSClient) Verify(ctx context.Context, in *PhoneCode, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/sms.SMS/Verify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SMS service

type SMSServer interface {
	Verify(context.Context, *PhoneCode) (*Status, error)
}

func RegisterSMSServer(s *grpc.Server, srv SMSServer) {
	s.RegisterService(&_SMS_serviceDesc, srv)
}

func _SMS_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sms.SMS/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServer).Verify(ctx, req.(*PhoneCode))
	}
	return interceptor(ctx, in, info, handler)
}

var _SMS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sms.SMS",
	HandlerType: (*SMSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _SMS_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("sms/sms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xce, 0x2d, 0xd6,
	0x2f, 0xce, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0xce, 0x2d, 0x56, 0x92,
	0xe5, 0x62, 0x0d, 0xc8, 0xc8, 0xcf, 0x4b, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x00, 0x31, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x29, 0x2e, 0x16, 0xe7, 0xfc, 0x94, 0x54, 0x21,
	0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x14, 0x98, 0x24, 0x98, 0xad, 0x64, 0xca, 0xc5, 0x09, 0xd6, 0x0a,
	0x56, 0x80, 0x55, 0x3b, 0x5c, 0x1b, 0x13, 0x92, 0x36, 0x25, 0x2e, 0xb6, 0xe0, 0x92, 0xc4, 0x92,
	0xd2, 0x62, 0x21, 0x09, 0x2e, 0xf6, 0xe2, 0xd2, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0xb0, 0x2e, 0x8e,
	0x20, 0x18, 0xd7, 0x48, 0x8f, 0x8b, 0x39, 0xd8, 0x37, 0x58, 0x48, 0x9d, 0x8b, 0x2d, 0x2c, 0xb5,
	0x28, 0x33, 0xad, 0x52, 0x88, 0x4f, 0x0f, 0xe4, 0x6e, 0xb8, 0x75, 0x52, 0xdc, 0x60, 0x3e, 0xc4,
	0x1c, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x8f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xce,
	0x45, 0x8c, 0xe2, 0x00, 0x00, 0x00,
}
