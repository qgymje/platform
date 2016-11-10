// Code generated by protoc-gen-go.
// source: room/room.proto
// DO NOT EDIT!

/*
Package room is a generated protocol buffer package.

It is generated from these files:
	room/room.proto

It has these top-level messages:
	Broadcast
	Num
	CreateRequest
	ListRequest
	RoomInfo
	Rooms
	User
	UserRoom
	Status
	BroadcastInfo
*/
package room

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

type Broadcast struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *Broadcast) Reset()                    { *m = Broadcast{} }
func (m *Broadcast) String() string            { return proto.CompactTextString(m) }
func (*Broadcast) ProtoMessage()               {}
func (*Broadcast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Num struct {
	BroadcastID string `protobuf:"bytes,1,opt,name=BroadcastID" json:"BroadcastID,omitempty"`
	Num         uint32 `protobuf:"varint,2,opt,name=Num" json:"Num,omitempty"`
}

func (m *Num) Reset()                    { *m = Num{} }
func (m *Num) String() string            { return proto.CompactTextString(m) }
func (*Num) ProtoMessage()               {}
func (*Num) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateRequest struct {
	UserID   string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Cover    string `protobuf:"bytes,4,opt,name=cover" json:"cover,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ListRequest struct {
	Num    int32  `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Size   int32  `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search" json:"search,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RoomInfo struct {
	RoomID    string         `protobuf:"bytes,1,opt,name=roomID" json:"roomID,omitempty"`
	Name      string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	UserID    string         `protobuf:"bytes,3,opt,name=userID" json:"userID,omitempty"`
	UserName  string         `protobuf:"bytes,4,opt,name=userName" json:"userName,omitempty"`
	Cover     string         `protobuf:"bytes,5,opt,name=cover" json:"cover,omitempty"`
	IsPlaying bool           `protobuf:"varint,6,opt,name=isPlaying" json:"isPlaying,omitempty"`
	IsFollow  bool           `protobuf:"varint,7,opt,name=isFollow" json:"isFollow,omitempty"`
	FollowNum int64          `protobuf:"varint,8,opt,name=followNum" json:"followNum,omitempty"`
	Broadcast *BroadcastInfo `protobuf:"bytes,9,opt,name=broadcast" json:"broadcast,omitempty"`
}

func (m *RoomInfo) Reset()                    { *m = RoomInfo{} }
func (m *RoomInfo) String() string            { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()               {}
func (*RoomInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RoomInfo) GetBroadcast() *BroadcastInfo {
	if m != nil {
		return m.Broadcast
	}
	return nil
}

type Rooms struct {
	Rooms    []*RoomInfo `protobuf:"bytes,1,rep,name=rooms" json:"rooms,omitempty"`
	TotalNum int64       `protobuf:"varint,2,opt,name=totalNum" json:"totalNum,omitempty"`
}

func (m *Rooms) Reset()                    { *m = Rooms{} }
func (m *Rooms) String() string            { return proto.CompactTextString(m) }
func (*Rooms) ProtoMessage()               {}
func (*Rooms) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Rooms) GetRooms() []*RoomInfo {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type User struct {
	UserID string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	TypeID int32  `protobuf:"varint,2,opt,name=typeID" json:"typeID,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type UserRoom struct {
	TypeID      int32  `protobuf:"varint,1,opt,name=typeID" json:"typeID,omitempty"`
	UserID      string `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	RoomID      string `protobuf:"bytes,3,opt,name=roomID" json:"roomID,omitempty"`
	BroadcastID string `protobuf:"bytes,4,opt,name=broadcastID" json:"broadcastID,omitempty"`
	Username    string `protobuf:"bytes,5,opt,name=username" json:"username,omitempty"`
	Level       int64  `protobuf:"varint,6,opt,name=level" json:"level,omitempty"`
}

func (m *UserRoom) Reset()                    { *m = UserRoom{} }
func (m *UserRoom) String() string            { return proto.CompactTextString(m) }
func (*UserRoom) ProtoMessage()               {}
func (*UserRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Status struct {
	Success     bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	RoomID      string `protobuf:"bytes,2,opt,name=roomID" json:"roomID,omitempty"`
	BroadcastID string `protobuf:"bytes,3,opt,name=broadcastID" json:"broadcastID,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type BroadcastInfo struct {
	BroadcastID     string `protobuf:"bytes,1,opt,name=broadcastID" json:"broadcastID,omitempty"`
	RoomID          string `protobuf:"bytes,2,opt,name=roomID" json:"roomID,omitempty"`
	StartTime       int64  `protobuf:"varint,3,opt,name=startTime" json:"startTime,omitempty"`
	Duration        int64  `protobuf:"varint,4,opt,name=duration" json:"duration,omitempty"`
	TotalAudience   int64  `protobuf:"varint,5,opt,name=totalAudience" json:"totalAudience,omitempty"`
	CurrentAudience int64  `protobuf:"varint,6,opt,name=currentAudience" json:"currentAudience,omitempty"`
}

func (m *BroadcastInfo) Reset()                    { *m = BroadcastInfo{} }
func (m *BroadcastInfo) String() string            { return proto.CompactTextString(m) }
func (*BroadcastInfo) ProtoMessage()               {}
func (*BroadcastInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*Broadcast)(nil), "room.Broadcast")
	proto.RegisterType((*Num)(nil), "room.Num")
	proto.RegisterType((*CreateRequest)(nil), "room.CreateRequest")
	proto.RegisterType((*ListRequest)(nil), "room.ListRequest")
	proto.RegisterType((*RoomInfo)(nil), "room.RoomInfo")
	proto.RegisterType((*Rooms)(nil), "room.Rooms")
	proto.RegisterType((*User)(nil), "room.User")
	proto.RegisterType((*UserRoom)(nil), "room.UserRoom")
	proto.RegisterType((*Status)(nil), "room.Status")
	proto.RegisterType((*BroadcastInfo)(nil), "room.BroadcastInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Room service

type RoomClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Status, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Rooms, error)
	Info(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*RoomInfo, error)
	Follow(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*Status, error)
	Unfollow(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*Status, error)
	Start(ctx context.Context, in *User, opts ...grpc.CallOption) (*BroadcastInfo, error)
	End(ctx context.Context, in *User, opts ...grpc.CallOption) (*BroadcastInfo, error)
	Enter(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*Status, error)
	Leave(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*Status, error)
}

type roomClient struct {
	cc *grpc.ClientConn
}

func NewRoomClient(cc *grpc.ClientConn) RoomClient {
	return &roomClient{cc}
}

func (c *roomClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/room.Room/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Rooms, error) {
	out := new(Rooms)
	err := grpc.Invoke(ctx, "/room.Room/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) Info(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*RoomInfo, error) {
	out := new(RoomInfo)
	err := grpc.Invoke(ctx, "/room.Room/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) Follow(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/room.Room/Follow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) Unfollow(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/room.Room/Unfollow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) Start(ctx context.Context, in *User, opts ...grpc.CallOption) (*BroadcastInfo, error) {
	out := new(BroadcastInfo)
	err := grpc.Invoke(ctx, "/room.Room/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) End(ctx context.Context, in *User, opts ...grpc.CallOption) (*BroadcastInfo, error) {
	out := new(BroadcastInfo)
	err := grpc.Invoke(ctx, "/room.Room/End", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) Enter(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/room.Room/Enter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) Leave(ctx context.Context, in *UserRoom, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/room.Room/Leave", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Room service

type RoomServer interface {
	Create(context.Context, *CreateRequest) (*Status, error)
	List(context.Context, *ListRequest) (*Rooms, error)
	Info(context.Context, *UserRoom) (*RoomInfo, error)
	Follow(context.Context, *UserRoom) (*Status, error)
	Unfollow(context.Context, *UserRoom) (*Status, error)
	Start(context.Context, *User) (*BroadcastInfo, error)
	End(context.Context, *User) (*BroadcastInfo, error)
	Enter(context.Context, *UserRoom) (*Status, error)
	Leave(context.Context, *UserRoom) (*Status, error)
}

func RegisterRoomServer(s *grpc.Server, srv RoomServer) {
	s.RegisterService(&_Room_serviceDesc, srv)
}

func _Room_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Info(ctx, req.(*UserRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Follow(ctx, req.(*UserRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/Unfollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Unfollow(ctx, req.(*UserRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Start(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_End_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).End(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/End",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).End(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_Enter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Enter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/Enter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Enter(ctx, req.(*UserRoom))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.Room/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Leave(ctx, req.(*UserRoom))
	}
	return interceptor(ctx, in, info, handler)
}

var _Room_serviceDesc = grpc.ServiceDesc{
	ServiceName: "room.Room",
	HandlerType: (*RoomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Room_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Room_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Room_Info_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _Room_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _Room_Unfollow_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Room_Start_Handler,
		},
		{
			MethodName: "End",
			Handler:    _Room_End_Handler,
		},
		{
			MethodName: "Enter",
			Handler:    _Room_Enter_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Room_Leave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("room/room.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0xfd, 0x12, 0x27, 0x69, 0x72, 0xf3, 0x4d, 0x5b, 0x0c, 0xaa, 0xa2, 0xa1, 0x8b, 0x28, 0xaa,
	0xd4, 0x08, 0x89, 0x22, 0x8a, 0x84, 0xc4, 0x12, 0x68, 0x91, 0x46, 0x54, 0x15, 0x72, 0xe9, 0x8e,
	0x4d, 0x9a, 0xf1, 0x40, 0xa4, 0x24, 0x2e, 0xb6, 0x53, 0x54, 0x5e, 0x80, 0x05, 0xef, 0xc1, 0x0b,
	0xf1, 0x42, 0xc8, 0x76, 0xfe, 0x26, 0xa5, 0xb4, 0x9b, 0xd1, 0xbd, 0xc7, 0xe7, 0xfa, 0xfe, 0x9c,
	0xeb, 0x0c, 0x6c, 0x71, 0xc6, 0xaa, 0x67, 0xea, 0xe7, 0xe0, 0x92, 0x33, 0xc9, 0xb0, 0xa3, 0xec,
	0xe4, 0x31, 0x04, 0x6f, 0x38, 0xcb, 0x96, 0x79, 0x26, 0x24, 0xde, 0x04, 0x7b, 0x71, 0x14, 0x59,
	0xb1, 0x95, 0x06, 0xc4, 0x5e, 0x1c, 0x25, 0xaf, 0x00, 0x9d, 0x36, 0x15, 0x8e, 0x21, 0xec, 0x39,
	0xfd, 0xf9, 0x18, 0xc2, 0xdb, 0x9a, 0x18, 0xd9, 0xb1, 0x95, 0xce, 0x88, 0x32, 0x93, 0x0a, 0x66,
	0x6f, 0x39, 0xcd, 0x24, 0x25, 0xf4, 0x6b, 0x43, 0x85, 0xc4, 0x3b, 0xe0, 0x35, 0x82, 0xf2, 0x3e,
	0xbe, 0xf5, 0xf0, 0x1c, 0x7c, 0x65, 0x9d, 0x66, 0x15, 0xd5, 0xf1, 0x01, 0xe9, 0x7d, 0x8c, 0xc1,
	0xa9, 0x15, 0x8e, 0x34, 0xae, 0x6d, 0xfc, 0x08, 0xdc, 0x9c, 0x5d, 0x51, 0x1e, 0x39, 0x1a, 0x34,
	0x4e, 0xf2, 0x1e, 0xc2, 0x93, 0x42, 0xc8, 0x2e, 0xd9, 0x36, 0xa0, 0xba, 0xa9, 0x74, 0x26, 0x97,
	0x28, 0x53, 0x5d, 0x25, 0x8a, 0xef, 0x26, 0x85, 0x4b, 0xb4, 0xad, 0x4a, 0x12, 0x34, 0xe3, 0xf9,
	0x97, 0x36, 0x41, 0xeb, 0x25, 0x3f, 0x6d, 0xf0, 0x09, 0x63, 0xd5, 0xa2, 0x5e, 0x31, 0x45, 0x52,
	0x83, 0x1a, 0xea, 0x36, 0x5e, 0x5f, 0x9b, 0x3d, 0xaa, 0x6d, 0xe8, 0x11, 0xdd, 0xda, 0xa3, 0x33,
	0xe9, 0xb1, 0xef, 0xc7, 0x1d, 0xf5, 0x83, 0x77, 0x21, 0x28, 0xc4, 0x87, 0x32, 0xbb, 0x2e, 0xea,
	0xcf, 0x91, 0x17, 0x5b, 0xa9, 0x4f, 0x06, 0x40, 0xdd, 0x57, 0x88, 0x77, 0xac, 0x2c, 0xd9, 0xb7,
	0x68, 0x43, 0x1f, 0xf6, 0xbe, 0x8a, 0x5c, 0x69, 0x4b, 0x09, 0xe2, 0xc7, 0x56, 0x8a, 0xc8, 0x00,
	0xe0, 0xe7, 0x10, 0x5c, 0x74, 0xba, 0x45, 0x41, 0x6c, 0xa5, 0xe1, 0xe1, 0xc3, 0x03, 0xbd, 0x14,
	0x83, 0x9c, 0xf5, 0x8a, 0x91, 0x81, 0x95, 0x2c, 0xc0, 0x55, 0xc3, 0x10, 0x78, 0x0f, 0x5c, 0xc5,
	0x14, 0x91, 0x15, 0xa3, 0x34, 0x3c, 0xdc, 0x34, 0x71, 0xdd, 0xa0, 0x88, 0x39, 0x54, 0xb5, 0x49,
	0x26, 0xb3, 0xb2, 0xdb, 0x07, 0x44, 0x7a, 0x3f, 0x79, 0x09, 0xce, 0xb9, 0xa0, 0xfc, 0xd6, 0x5d,
	0xd8, 0x01, 0x4f, 0x5e, 0x5f, 0xd2, 0xc5, 0x51, 0x2b, 0x53, 0xeb, 0x25, 0xbf, 0x2c, 0xf0, 0x55,
	0xa0, 0xca, 0x35, 0x22, 0x59, 0x63, 0xd2, 0xe8, 0x52, 0x7b, 0x7a, 0x69, 0x2b, 0x20, 0x5a, 0x13,
	0x30, 0x86, 0xf0, 0x62, 0xb4, 0xd5, 0x46, 0x97, 0x31, 0xd4, 0xc9, 0xa6, 0x65, 0x76, 0x07, 0xd9,
	0xba, 0x35, 0x2c, 0xe9, 0x15, 0x2d, 0xb5, 0x38, 0x88, 0x18, 0x27, 0xf9, 0x04, 0xde, 0x99, 0xcc,
	0x64, 0x23, 0x70, 0x04, 0x1b, 0xa2, 0xc9, 0x73, 0x2a, 0x84, 0x2e, 0xd3, 0x27, 0x9d, 0x3b, 0xaa,
	0xc7, 0xfe, 0x57, 0x3d, 0xe8, 0x46, 0x3d, 0xc9, 0x6f, 0x0b, 0x66, 0x6b, 0x32, 0x4d, 0x63, 0xac,
	0x9b, 0x3d, 0xdc, 0x96, 0x6d, 0x17, 0x02, 0x21, 0x33, 0x2e, 0x3f, 0x16, 0xed, 0xfb, 0x42, 0x64,
	0x00, 0x54, 0xe7, 0xcb, 0x86, 0x67, 0xb2, 0x60, 0xb5, 0x1e, 0x0c, 0x22, 0xbd, 0x8f, 0xf7, 0x60,
	0xa6, 0x05, 0x7d, 0xdd, 0x2c, 0x0b, 0x5a, 0xe7, 0x66, 0x34, 0x88, 0xac, 0x83, 0x38, 0x85, 0xad,
	0xbc, 0xe1, 0x9c, 0xd6, 0xb2, 0xe7, 0x99, 0x49, 0x4d, 0xe1, 0xc3, 0x1f, 0x08, 0x1c, 0x2d, 0xec,
	0x53, 0xf0, 0xcc, 0x27, 0x03, 0xb7, 0x2b, 0xb9, 0xf6, 0x01, 0x99, 0xff, 0x6f, 0x40, 0x33, 0xdf,
	0xe4, 0x3f, 0x9c, 0x82, 0xa3, 0x9e, 0x3c, 0x7e, 0x60, 0xf0, 0xd1, 0xf3, 0x9f, 0x87, 0xc3, 0x6a,
	0xb6, 0x4c, 0x3d, 0xad, 0x76, 0x63, 0xbb, 0x4d, 0x9a, 0x4f, 0x36, 0x58, 0x33, 0xbd, 0xf6, 0x19,
	0x4d, 0xb9, 0xd3, 0xec, 0x4f, 0xc0, 0x3f, 0xaf, 0x57, 0xf7, 0xe5, 0xba, 0x67, 0x6a, 0xb4, 0x18,
	0x06, 0xe2, 0xfc, 0x6f, 0xcf, 0x4e, 0x57, 0x80, 0x8e, 0xeb, 0xe5, 0x7d, 0x98, 0xfb, 0xe0, 0x1e,
	0xd7, 0x92, 0xf2, 0x3b, 0xd3, 0xef, 0x83, 0x7b, 0x42, 0xb3, 0x2b, 0x7a, 0x17, 0xf1, 0xc2, 0xd3,
	0x7f, 0x0c, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x59, 0xc3, 0x2b, 0x2b, 0x06, 0x00,
	0x00,
}
