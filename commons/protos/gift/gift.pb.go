// Code generated by protoc-gen-go.
// source: gift/gift.proto
// DO NOT EDIT!

/*
Package gift is a generated protocol buffer package.

It is generated from these files:
	gift/gift.proto

It has these top-level messages:
	Page
	SendGift
	Status
	GiftInfo
	Gifts
*/
package gift

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

type Page struct {
	Num  int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SendGift struct {
	GiftID      string `protobuf:"bytes,1,opt,name=giftID" json:"giftID,omitempty"`
	UserID      string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	ToUserID    string `protobuf:"bytes,3,opt,name=toUserID" json:"toUserID,omitempty"`
	BroadcastID string `protobuf:"bytes,4,opt,name=broadcastID" json:"broadcastID,omitempty"`
	Number      int32  `protobuf:"varint,5,opt,name=number" json:"number,omitempty"`
}

func (m *SendGift) Reset()                    { *m = SendGift{} }
func (m *SendGift) String() string            { return proto.CompactTextString(m) }
func (*SendGift) ProtoMessage()               {}
func (*SendGift) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Status struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GiftInfo struct {
	GiftID    string `protobuf:"bytes,1,opt,name=giftID" json:"giftID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Image     string `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	SnowBall  string `protobuf:"bytes,4,opt,name=snowBall" json:"snowBall,omitempty"`
	SnowFlake string `protobuf:"bytes,5,opt,name=snowFlake" json:"snowFlake,omitempty"`
}

func (m *GiftInfo) Reset()                    { *m = GiftInfo{} }
func (m *GiftInfo) String() string            { return proto.CompactTextString(m) }
func (*GiftInfo) ProtoMessage()               {}
func (*GiftInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Gifts struct {
	List []*GiftInfo `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *Gifts) Reset()                    { *m = Gifts{} }
func (m *Gifts) String() string            { return proto.CompactTextString(m) }
func (*Gifts) ProtoMessage()               {}
func (*Gifts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Gifts) GetList() []*GiftInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Page)(nil), "gift.Page")
	proto.RegisterType((*SendGift)(nil), "gift.SendGift")
	proto.RegisterType((*Status)(nil), "gift.Status")
	proto.RegisterType((*GiftInfo)(nil), "gift.GiftInfo")
	proto.RegisterType((*Gifts)(nil), "gift.Gifts")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Gift service

type GiftClient interface {
	List(ctx context.Context, in *Page, opts ...grpc.CallOption) (*Gifts, error)
	Send(ctx context.Context, in *SendGift, opts ...grpc.CallOption) (*Status, error)
}

type giftClient struct {
	cc *grpc.ClientConn
}

func NewGiftClient(cc *grpc.ClientConn) GiftClient {
	return &giftClient{cc}
}

func (c *giftClient) List(ctx context.Context, in *Page, opts ...grpc.CallOption) (*Gifts, error) {
	out := new(Gifts)
	err := grpc.Invoke(ctx, "/gift.Gift/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftClient) Send(ctx context.Context, in *SendGift, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/gift.Gift/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gift service

type GiftServer interface {
	List(context.Context, *Page) (*Gifts, error)
	Send(context.Context, *SendGift) (*Status, error)
}

func RegisterGiftServer(s *grpc.Server, srv GiftServer) {
	s.RegisterService(&_Gift_serviceDesc, srv)
}

func _Gift_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gift.Gift/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).List(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gift_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGift)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gift.Gift/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftServer).Send(ctx, req.(*SendGift))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gift_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gift.Gift",
	HandlerType: (*GiftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Gift_List_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Gift_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("gift/gift.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x9b, 0x7f, 0x37, 0xfd, 0x27, 0x53, 0x51, 0x19, 0x44, 0x42, 0xf1, 0x50, 0xf7, 0x20,
	0x05, 0xa5, 0x42, 0x7d, 0x03, 0x29, 0x4a, 0xc1, 0x83, 0xa6, 0xf8, 0x00, 0xdb, 0x76, 0x5b, 0x82,
	0x49, 0x56, 0x32, 0x1b, 0x04, 0x1f, 0xc0, 0xb3, 0x8f, 0x2c, 0x33, 0xdd, 0xb6, 0x5e, 0xbc, 0x84,
	0xef, 0x9b, 0xc9, 0x0c, 0xbf, 0x9d, 0x0f, 0x4e, 0x36, 0xc5, 0xda, 0xdf, 0xf2, 0x67, 0xfc, 0xde,
	0x38, 0xef, 0x50, 0xb1, 0xd6, 0x37, 0xa0, 0x9e, 0xcd, 0xc6, 0xe2, 0x29, 0x74, 0xeb, 0xb6, 0xca,
	0xa2, 0x61, 0x34, 0x8a, 0x73, 0x96, 0x88, 0xa0, 0xa8, 0xf8, 0xb4, 0xd9, 0x3f, 0x29, 0x89, 0xd6,
	0xdf, 0x11, 0x24, 0x73, 0x5b, 0xaf, 0x1e, 0x8b, 0xb5, 0xc7, 0x73, 0xe8, 0xf1, 0x8a, 0xd9, 0x54,
	0xa6, 0xd2, 0x3c, 0x38, 0xae, 0xb7, 0x64, 0x9b, 0xd9, 0x54, 0x46, 0xd3, 0x3c, 0x38, 0x1c, 0x40,
	0xe2, 0xdd, 0xeb, 0xb6, 0xd3, 0x95, 0xce, 0xde, 0xe3, 0x10, 0xfa, 0x8b, 0xc6, 0x99, 0xd5, 0xd2,
	0x10, 0x2f, 0x54, 0xd2, 0xfe, 0x5d, 0xe2, 0xad, 0x75, 0x5b, 0x2d, 0x6c, 0x93, 0xc5, 0x02, 0x14,
	0x9c, 0xd6, 0xd0, 0x9b, 0x7b, 0xe3, 0x5b, 0xc2, 0x0c, 0xfe, 0x53, 0xbb, 0x5c, 0x5a, 0x22, 0x01,
	0x4a, 0xf2, 0x9d, 0xd5, 0x5f, 0x11, 0x24, 0x8c, 0x3c, 0xab, 0xd7, 0xee, 0x4f, 0x6c, 0x04, 0x55,
	0x9b, 0xca, 0x06, 0x68, 0xd1, 0x78, 0x06, 0x71, 0x51, 0x99, 0x8d, 0x0d, 0xbc, 0x5b, 0xc3, 0x0f,
	0xa1, 0xda, 0x7d, 0xdc, 0x9b, 0xb2, 0x0c, 0xa4, 0x7b, 0x8f, 0x17, 0x90, 0xb2, 0x7e, 0x28, 0xcd,
	0x9b, 0x15, 0xd2, 0x34, 0x3f, 0x14, 0xf4, 0x35, 0xc4, 0xcc, 0x41, 0xa8, 0x41, 0x95, 0x05, 0xf9,
	0x2c, 0x1a, 0x76, 0x47, 0xfd, 0xc9, 0xf1, 0x58, 0x72, 0xd9, 0x21, 0xe6, 0xd2, 0x9b, 0xbc, 0x80,
	0x92, 0x3b, 0x5f, 0x82, 0x7a, 0x2a, 0xc8, 0x23, 0x6c, 0xff, 0xe2, 0xb8, 0x06, 0xfd, 0xc3, 0x04,
	0xe9, 0x0e, 0x5e, 0x81, 0xe2, 0x58, 0x30, 0x2c, 0xda, 0x45, 0x34, 0x38, 0x0a, 0x5e, 0x0e, 0xa4,
	0x3b, 0x8b, 0x9e, 0x44, 0x7f, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x15, 0x9f, 0x42, 0x0d,
	0x02, 0x00, 0x00,
}
