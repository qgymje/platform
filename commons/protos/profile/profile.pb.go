// Code generated by protoc-gen-go.
// source: profile/profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile/profile.proto

It has these top-level messages:
	Friends
	Ammount
	Message
	Status
	Request
	RequestID
*/
package profile

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

type Friends struct {
	FriendIDs []string `protobuf:"bytes,1,rep,name=friendIDs" json:"friendIDs,omitempty"`
}

func (m *Friends) Reset()                    { *m = Friends{} }
func (m *Friends) String() string            { return proto.CompactTextString(m) }
func (*Friends) ProtoMessage()               {}
func (*Friends) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Ammount struct {
	UserID    string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	SnowFlake uint32 `protobuf:"varint,2,opt,name=snowFlake" json:"snowFlake,omitempty"`
	SnowBall  uint32 `protobuf:"varint,3,opt,name=snowBall" json:"snowBall,omitempty"`
	TypeID    uint32 `protobuf:"varint,4,opt,name=typeID" json:"typeID,omitempty"`
	TargetID  string `protobuf:"bytes,5,opt,name=targetID" json:"targetID,omitempty"`
}

func (m *Ammount) Reset()                    { *m = Ammount{} }
func (m *Ammount) String() string            { return proto.CompactTextString(m) }
func (*Ammount) ProtoMessage()               {}
func (*Ammount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Message struct {
	MsgID  string `protobuf:"bytes,1,opt,name=msgID" json:"msgID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Status struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	MsgID   string `protobuf:"bytes,2,opt,name=msgID" json:"msgID,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Request struct {
	FromUserID string `protobuf:"bytes,1,opt,name=fromUserID" json:"fromUserID,omitempty"`
	ToUserID   string `protobuf:"bytes,2,opt,name=toUserID" json:"toUserID,omitempty"`
	Message    string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RequestID struct {
	RequestID string `protobuf:"bytes,1,opt,name=requestID" json:"requestID,omitempty"`
}

func (m *RequestID) Reset()                    { *m = RequestID{} }
func (m *RequestID) String() string            { return proto.CompactTextString(m) }
func (*RequestID) ProtoMessage()               {}
func (*RequestID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*Friends)(nil), "profile.Friends")
	proto.RegisterType((*Ammount)(nil), "profile.Ammount")
	proto.RegisterType((*Message)(nil), "profile.Message")
	proto.RegisterType((*Status)(nil), "profile.Status")
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*RequestID)(nil), "profile.RequestID")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Profile service

type ProfileClient interface {
	Withdraw(ctx context.Context, in *Ammount, opts ...grpc.CallOption) (*Status, error)
	WithdrawRollback(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
	WithdrawCommit(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error)
	FriendList(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Friends, error)
	FriendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RequestID, error)
	FriendAgree(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*Status, error)
	FriendRefuse(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*Status, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) Withdraw(ctx context.Context, in *Ammount, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/profile.Profile/Withdraw", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) WithdrawRollback(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/profile.Profile/WithdrawRollback", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) WithdrawCommit(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/profile.Profile/WithdrawCommit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) FriendList(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Friends, error) {
	out := new(Friends)
	err := grpc.Invoke(ctx, "/profile.Profile/FriendList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) FriendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RequestID, error) {
	out := new(RequestID)
	err := grpc.Invoke(ctx, "/profile.Profile/FriendRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) FriendAgree(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/profile.Profile/FriendAgree", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) FriendRefuse(ctx context.Context, in *RequestID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/profile.Profile/FriendRefuse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileServer interface {
	Withdraw(context.Context, *Ammount) (*Status, error)
	WithdrawRollback(context.Context, *Message) (*Status, error)
	WithdrawCommit(context.Context, *Message) (*Status, error)
	FriendList(context.Context, *Message) (*Friends, error)
	FriendRequest(context.Context, *Request) (*RequestID, error)
	FriendAgree(context.Context, *RequestID) (*Status, error)
	FriendRefuse(context.Context, *RequestID) (*Status, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ammount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Withdraw(ctx, req.(*Ammount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_WithdrawRollback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).WithdrawRollback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/WithdrawRollback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).WithdrawRollback(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_WithdrawCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).WithdrawCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/WithdrawCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).WithdrawCommit(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_FriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).FriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/FriendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).FriendList(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_FriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).FriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/FriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).FriendRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_FriendAgree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).FriendAgree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/FriendAgree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).FriendAgree(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_FriendRefuse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).FriendRefuse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/FriendRefuse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).FriendRefuse(ctx, req.(*RequestID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Withdraw",
			Handler:    _Profile_Withdraw_Handler,
		},
		{
			MethodName: "WithdrawRollback",
			Handler:    _Profile_WithdrawRollback_Handler,
		},
		{
			MethodName: "WithdrawCommit",
			Handler:    _Profile_WithdrawCommit_Handler,
		},
		{
			MethodName: "FriendList",
			Handler:    _Profile_FriendList_Handler,
		},
		{
			MethodName: "FriendRequest",
			Handler:    _Profile_FriendRequest_Handler,
		},
		{
			MethodName: "FriendAgree",
			Handler:    _Profile_FriendAgree_Handler,
		},
		{
			MethodName: "FriendRefuse",
			Handler:    _Profile_FriendRefuse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("profile/profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6b, 0xe2, 0x40,
	0x14, 0xc7, 0x8d, 0xae, 0xc6, 0xbc, 0x5d, 0x77, 0x65, 0xd8, 0x5d, 0x82, 0x2c, 0x8b, 0xcc, 0x65,
	0xdd, 0x8b, 0x82, 0xbb, 0xc5, 0x5e, 0x6d, 0x83, 0x10, 0x68, 0xa1, 0xa4, 0x48, 0x8f, 0x25, 0xea,
	0x24, 0x0d, 0x26, 0x8e, 0x9d, 0x99, 0x20, 0xbd, 0xf7, 0xde, 0x7f, 0xb9, 0x64, 0x7e, 0x24, 0xa1,
	0xe4, 0xe0, 0x49, 0x3f, 0xdf, 0x99, 0xef, 0x7b, 0xdf, 0x79, 0x33, 0x81, 0x1f, 0x47, 0x46, 0xa3,
	0x24, 0x25, 0x33, 0xfd, 0x3b, 0x3d, 0x32, 0x2a, 0x28, 0xb2, 0x35, 0xe2, 0x3f, 0x60, 0xaf, 0x58,
	0x42, 0x0e, 0x3b, 0x8e, 0x7e, 0x81, 0x13, 0xc9, 0xbf, 0xbe, 0xc7, 0x5d, 0x6b, 0xdc, 0x99, 0x38,
	0x41, 0x25, 0xe0, 0x37, 0x0b, 0xec, 0x65, 0x96, 0xd1, 0xfc, 0x20, 0xd0, 0x4f, 0xe8, 0xe5, 0x9c,
	0x30, 0xdf, 0x73, 0xad, 0xb1, 0x35, 0x71, 0x02, 0x4d, 0x45, 0x05, 0x7e, 0xa0, 0xa7, 0x55, 0x1a,
	0xee, 0x89, 0xdb, 0x1e, 0x5b, 0x93, 0x41, 0x50, 0x09, 0x68, 0x04, 0xfd, 0x02, 0xae, 0xc2, 0x34,
	0x75, 0x3b, 0x72, 0xb1, 0xe4, 0xa2, 0xa2, 0x78, 0x39, 0x12, 0xdf, 0x73, 0x3f, 0xc9, 0x15, 0x4d,
	0x85, 0x47, 0x84, 0x2c, 0x26, 0xc2, 0xf7, 0xdc, 0xae, 0xec, 0x55, 0x32, 0x5e, 0x80, 0x7d, 0x4b,
	0x38, 0x0f, 0x63, 0x82, 0xbe, 0x43, 0x37, 0xe3, 0x71, 0x99, 0x47, 0x41, 0x2d, 0x66, 0xbb, 0x1e,
	0x13, 0x5f, 0x42, 0xef, 0x5e, 0x84, 0x22, 0xe7, 0xc8, 0x05, 0x9b, 0xe7, 0xdb, 0x2d, 0xe1, 0x5c,
	0x3a, 0xfb, 0x81, 0xc1, 0xaa, 0x62, 0xbb, 0x56, 0x11, 0x3f, 0x82, 0x1d, 0x90, 0xe7, 0x9c, 0x70,
	0x81, 0x7e, 0x03, 0x44, 0x8c, 0x66, 0xeb, 0xfa, 0x1c, 0x6a, 0x8a, 0x4c, 0x4e, 0xd7, 0xf5, 0xf6,
	0x25, 0x17, 0x6d, 0x33, 0x95, 0x5c, 0x0e, 0xc2, 0x09, 0x0c, 0xe2, 0xbf, 0xe0, 0xe8, 0x06, 0x6a,
	0x9c, 0xcc, 0x80, 0xee, 0x50, 0x09, 0xf3, 0xd7, 0x0e, 0xd8, 0x77, 0xea, 0x16, 0xd1, 0x0c, 0xfa,
	0x0f, 0x89, 0x78, 0xda, 0xb1, 0xf0, 0x84, 0x86, 0x53, 0x73, 0xd5, 0xfa, 0xba, 0x46, 0xdf, 0x4a,
	0x45, 0x1d, 0x1b, 0xb7, 0xd0, 0x02, 0x86, 0xc6, 0x10, 0xd0, 0x34, 0xdd, 0x84, 0xdb, 0x7d, 0xcd,
	0xa8, 0xc7, 0xda, 0x64, 0xbc, 0x80, 0xaf, 0xc6, 0x78, 0x4d, 0xb3, 0x2c, 0x11, 0xe7, 0xd9, 0xe6,
	0x00, 0xea, 0x99, 0xdd, 0x24, 0xbc, 0xc9, 0x52, 0x29, 0xfa, 0x35, 0xca, 0x8c, 0x03, 0x05, 0x66,
	0xe4, 0xd5, 0x26, 0xad, 0x8c, 0xd0, 0x47, 0xc5, 0xf7, 0x70, 0x0b, 0xfd, 0x87, 0xcf, 0xca, 0xb8,
	0x8c, 0x19, 0x21, 0xa8, 0x61, 0x53, 0xf3, 0xc9, 0xbe, 0x98, 0x76, 0x51, 0xce, 0xcf, 0xb5, 0x6d,
	0x7a, 0xf2, 0x83, 0xfa, 0xf7, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x13, 0x51, 0x30, 0x69, 0x03,
	0x00, 0x00,
}
