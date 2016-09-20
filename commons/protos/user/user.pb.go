// Code generated by protoc-gen-go.
// source: user/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	Phone
	Code
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

type UserID struct {
	UserID string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
}

func (m *UserID) Reset()                    { *m = UserID{} }
func (m *UserID) String() string            { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()               {}
func (*UserID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Status struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Token struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type LoginInfo struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginInfo) Reset()                    { *m = LoginInfo{} }
func (m *LoginInfo) String() string            { return proto.CompactTextString(m) }
func (*LoginInfo) ProtoMessage()               {}
func (*LoginInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type RegisterInfo struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
}

func (m *RegisterInfo) Reset()                    { *m = RegisterInfo{} }
func (m *RegisterInfo) String() string            { return proto.CompactTextString(m) }
func (*RegisterInfo) ProtoMessage()               {}
func (*RegisterInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Phone)(nil), "user.Phone")
	proto.RegisterType((*Code)(nil), "user.Code")
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
	ValidCode(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*Code, error)
	Register(ctx context.Context, in *RegisterInfo, opts ...grpc.CallOption) (*UserInfo, error)
	Login(ctx context.Context, in *LoginInfo, opts ...grpc.CallOption) (*UserInfo, error)
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

func (c *userClient) ValidCode(ctx context.Context, in *Phone, opts ...grpc.CallOption) (*Code, error) {
	out := new(Code)
	err := grpc.Invoke(ctx, "/user.User/ValidCode", in, out, c.cc, opts...)
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

func (c *userClient) Login(ctx context.Context, in *LoginInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/user.User/Login", in, out, c.cc, opts...)
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
	ValidCode(context.Context, *Phone) (*Code, error)
	Register(context.Context, *RegisterInfo) (*UserInfo, error)
	Login(context.Context, *LoginInfo) (*UserInfo, error)
	Logout(context.Context, *Token) (*Status, error)
	Auth(context.Context, *Token) (*UserInfo, error)
	Info(context.Context, *UserID) (*UserInfo, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_ValidCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Phone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ValidCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/ValidCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ValidCode(ctx, req.(*Phone))
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
			MethodName: "ValidCode",
			Handler:    _User_ValidCode_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
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
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x6b, 0xe2, 0x40,
	0x14, 0x36, 0x9a, 0xc4, 0xe4, 0xad, 0xac, 0x30, 0x2c, 0x4b, 0x08, 0xbb, 0x20, 0x03, 0x2b, 0xb2,
	0x07, 0x77, 0x69, 0x8f, 0x3d, 0x95, 0x7a, 0x11, 0x3c, 0x94, 0xd4, 0xf6, 0xd0, 0x5b, 0x9a, 0x4c,
	0x63, 0xb0, 0x66, 0x64, 0x26, 0xa1, 0xf7, 0xfe, 0xac, 0xfe, 0xba, 0xf2, 0xde, 0x38, 0x51, 0xb1,
	0x5e, 0x7a, 0x91, 0xef, 0x7b, 0xdf, 0x9b, 0x99, 0xcf, 0xef, 0xbd, 0xc0, 0xb0, 0xd1, 0x42, 0xfd,
	0xc3, 0x9f, 0xe9, 0x56, 0xc9, 0x5a, 0x32, 0x17, 0x31, 0xff, 0x0d, 0xde, 0xed, 0x4a, 0x56, 0x82,
	0xfd, 0x00, 0x6f, 0x8b, 0x20, 0x72, 0x46, 0xce, 0x24, 0x4c, 0x0c, 0xe1, 0x31, 0xb8, 0x37, 0x32,
	0x17, 0x8c, 0x81, 0x9b, 0xc9, 0xdc, 0x8a, 0x84, 0xf9, 0x08, 0xfc, 0x7b, 0x2d, 0xd4, 0x7c, 0xc6,
	0x7e, 0x82, 0xdf, 0x10, 0xda, 0xe9, 0x3b, 0xc6, 0x39, 0xf8, 0x77, 0x75, 0x5a, 0x37, 0x9a, 0x45,
	0xd0, 0xd7, 0x4d, 0x96, 0x09, 0xad, 0xa9, 0x25, 0x48, 0x2c, 0x45, 0x03, 0x4b, 0xb9, 0x16, 0x15,
	0x1a, 0xa8, 0x11, 0x58, 0x03, 0x44, 0xf8, 0x15, 0x84, 0x0b, 0x59, 0x94, 0xd5, 0xbc, 0x7a, 0x96,
	0xe8, 0xa2, 0x4a, 0x37, 0xad, 0x0b, 0xc4, 0x2c, 0x86, 0x60, 0x9b, 0x6a, 0xfd, 0x2a, 0x55, 0x1e,
	0x75, 0xa9, 0xde, 0x72, 0xfe, 0x08, 0x83, 0x44, 0x14, 0xa5, 0xae, 0x85, 0xfa, 0xca, 0x79, 0xd4,
	0xaa, 0x32, 0x5b, 0xd3, 0x99, 0x9e, 0xd1, 0x2c, 0xe7, 0xef, 0x0e, 0x04, 0xf4, 0xf7, 0xf1, 0xe2,
	0x33, 0x01, 0xb4, 0x0f, 0x76, 0x8f, 0x1f, 0x3c, 0x77, 0xe9, 0x3e, 0x03, 0xf7, 0x20, 0x03, 0x0c,
	0x6f, 0x25, 0xd2, 0x7c, 0xbe, 0x29, 0x22, 0x8f, 0xea, 0x96, 0xa2, 0xa2, 0x44, 0xb1, 0x2c, 0x37,
	0x22, 0xf2, 0x47, 0xce, 0xa4, 0x97, 0x58, 0xca, 0x7e, 0x41, 0x28, 0x94, 0x92, 0x0a, 0xa7, 0x17,
	0xf5, 0xe9, 0xd4, 0xbe, 0x70, 0xf1, 0xd6, 0x05, 0x17, 0xcd, 0xb3, 0x31, 0x84, 0x0f, 0xe9, 0x4b,
	0x99, 0xd3, 0x90, 0xbf, 0x4d, 0x69, 0x3d, 0x68, 0x1f, 0x62, 0x30, 0x04, 0x05, 0xde, 0x61, 0xff,
	0x21, 0xb0, 0x49, 0x32, 0x66, 0x94, 0xc3, 0x64, 0xe3, 0xef, 0xa6, 0x66, 0x03, 0xe1, 0x1d, 0xf6,
	0x17, 0x3c, 0x1a, 0x1c, 0x1b, 0x1a, 0xa9, 0x9d, 0xe2, 0x27, 0xbd, 0x7f, 0xc0, 0x5f, 0xc8, 0x42,
	0x36, 0xb5, 0xb5, 0x40, 0x1b, 0x11, 0x0f, 0x0c, 0x31, 0x2b, 0x44, 0x6d, 0xee, 0x75, 0x53, 0xaf,
	0x8e, 0x9b, 0x4e, 0x6f, 0x1b, 0x83, 0x4b, 0x43, 0x19, 0x1c, 0x28, 0xb3, 0xd3, 0xbe, 0x27, 0x9f,
	0xbe, 0x83, 0xcb, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xb6, 0xa8, 0x5b, 0x1a, 0x03, 0x00,
	0x00,
}