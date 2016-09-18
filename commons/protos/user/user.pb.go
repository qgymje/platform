// Code generated by protoc-gen-go.
// source: user/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	UserID
	Status
	Token
	LoginInfo
	RegisterInfo
	UserInfo
*/
package user

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

type UserID struct {
	UserID string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
}

func (m *UserID) Reset()                    { *m = UserID{} }
func (m *UserID) String() string            { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()               {}
func (*UserID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Status struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Token struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type LoginInfo struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginInfo) Reset()                    { *m = LoginInfo{} }
func (m *LoginInfo) String() string            { return proto.CompactTextString(m) }
func (*LoginInfo) ProtoMessage()               {}
func (*LoginInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RegisterInfo struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
}

func (m *RegisterInfo) Reset()                    { *m = RegisterInfo{} }
func (m *RegisterInfo) String() string            { return proto.CompactTextString(m) }
func (*RegisterInfo) ProtoMessage()               {}
func (*RegisterInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type UserInfo struct {
	UserID    string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Nickname  string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	Token     string `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	HeadImg   string `protobuf:"bytes,5,opt,name=headImg" json:"headImg,omitempty"`
	RegTime   int64  `protobuf:"varint,6,opt,name=regTime" json:"regTime,omitempty"`
	ErrorCode string `protobuf:"bytes,7,opt,name=errorCode" json:"errorCode,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*UserID)(nil), "user.UserID")
	proto.RegisterType((*Status)(nil), "user.Status")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*LoginInfo)(nil), "user.LoginInfo")
	proto.RegisterType((*RegisterInfo)(nil), "user.RegisterInfo")
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for User service

type UserClient interface {
	Login(ctx context.Context, in *LoginInfo, opts ...grpc.CallOption) (*UserInfo, error)
	Register(ctx context.Context, in *RegisterInfo, opts ...grpc.CallOption) (*UserInfo, error)
	Logout(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Status, error)
	Auth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UserInfo, error)
	Info(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserInfo, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/user.User/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/user.User/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Logout(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/user.User/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Auth(ctx context.Context, in *Token, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/user.User/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Info(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/user.User/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	Login(context.Context, *LoginInfo) (*UserInfo, error)
	Register(context.Context, *RegisterInfo) (*UserInfo, error)
	Logout(context.Context, *Token) (*Status, error)
	Auth(context.Context, *Token) (*UserInfo, error)
	Info(context.Context, *UserID) (*UserInfo, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Logout(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Auth(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Info(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _User_Logout_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _User_Auth_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _User_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x2d, 0xb4, 0xa5, 0x8c, 0x44, 0x92, 0x89, 0x31, 0x0d, 0xd1, 0x84, 0x6c, 0xa2, 0x21,
	0x1e, 0xd0, 0xe8, 0xd1, 0x93, 0x91, 0x0b, 0x09, 0xa7, 0x8a, 0x17, 0x6f, 0x15, 0xc6, 0xd2, 0x10,
	0xba, 0x64, 0x77, 0x1b, 0xdf, 0xcd, 0x47, 0xf1, 0x69, 0xcc, 0xce, 0x52, 0x28, 0x41, 0x2e, 0x5e,
	0x9a, 0xf9, 0xe7, 0x9f, 0x99, 0xed, 0x7e, 0x3b, 0xd0, 0x2d, 0x35, 0xa9, 0x3b, 0xfb, 0x19, 0xae,
	0x95, 0x34, 0x12, 0x7d, 0x1b, 0x8b, 0x3e, 0x84, 0x6f, 0x9a, 0xd4, 0x78, 0x84, 0x17, 0x10, 0x96,
	0x1c, 0xc5, 0x5e, 0xdf, 0x1b, 0xb4, 0x93, 0x8d, 0x12, 0x02, 0xc2, 0x57, 0x93, 0x9a, 0x52, 0x63,
	0x0c, 0x2d, 0x5d, 0xce, 0x66, 0xa4, 0x35, 0x97, 0x44, 0x49, 0x25, 0xc5, 0x15, 0x04, 0x53, 0xb9,
	0xa4, 0x02, 0xcf, 0x21, 0x30, 0x36, 0xd8, 0xcc, 0x70, 0x42, 0x3c, 0x41, 0x7b, 0x22, 0xb3, 0xbc,
	0x18, 0x17, 0x9f, 0x12, 0x11, 0xfc, 0x22, 0x5d, 0xd1, 0xa6, 0x82, 0x63, 0xec, 0x41, 0xb4, 0x4e,
	0xb5, 0xfe, 0x92, 0x6a, 0x1e, 0x37, 0x38, 0xbf, 0xd5, 0xe2, 0x1d, 0x3a, 0x09, 0x65, 0xb9, 0x36,
	0xa4, 0xfe, 0xd3, 0x6f, 0xbd, 0x22, 0x9f, 0x2d, 0xb9, 0xa7, 0xe9, 0xbc, 0x4a, 0x8b, 0x6f, 0x0f,
	0x22, 0xbe, 0xbe, 0x1d, 0x7c, 0x04, 0xc0, 0xf6, 0xc0, 0xc6, 0xfe, 0x81, 0xc7, 0x86, 0xee, 0x18,
	0xf8, 0x35, 0x06, 0x16, 0xde, 0x82, 0xd2, 0xf9, 0x78, 0x95, 0xc5, 0x01, 0xe7, 0x2b, 0x69, 0x1d,
	0x45, 0xd9, 0x34, 0x5f, 0x51, 0x1c, 0xf6, 0xbd, 0x41, 0x33, 0xa9, 0x24, 0x5e, 0x42, 0x9b, 0x94,
	0x92, 0xea, 0x45, 0xce, 0x29, 0x6e, 0x71, 0xd7, 0x2e, 0xf1, 0xf0, 0xe3, 0x81, 0x6f, 0x7f, 0x1e,
	0x6f, 0x21, 0x60, 0xbc, 0xd8, 0x1d, 0xf2, 0xfb, 0x6e, 0x59, 0xf7, 0xce, 0x5c, 0xa2, 0xba, 0xa2,
	0x38, 0xc1, 0x7b, 0x88, 0x2a, 0x9a, 0x88, 0xce, 0xad, 0xd3, 0xfd, 0xa3, 0xe3, 0x1a, 0xc2, 0x89,
	0xcc, 0x64, 0x69, 0xf0, 0xd4, 0x79, 0xfc, 0xd2, 0xbd, 0x8e, 0x13, 0x6e, 0x35, 0xb8, 0xcc, 0x7f,
	0x2e, 0xcd, 0x62, 0xbf, 0xe8, 0x70, 0xda, 0x0d, 0xf8, 0x0c, 0xbb, 0x53, 0x73, 0x46, 0x87, 0x75,
	0x1f, 0x21, 0x2f, 0xe9, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x70, 0x4b, 0x38, 0xb7,
	0x02, 0x00, 0x00,
}
