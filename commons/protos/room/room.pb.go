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
	ID string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *Broadcast) Reset()                    { *m = Broadcast{} }
func (m *Broadcast) String() string            { return proto.CompactTextString(m) }
func (*Broadcast) ProtoMessage()               {}
func (*Broadcast) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Num struct {
	BroadcastID string `protobuf:"bytes,1,opt,name=BroadcastID,json=broadcastID" json:"BroadcastID,omitempty"`
	Num         uint32 `protobuf:"varint,2,opt,name=Num,json=num" json:"Num,omitempty"`
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
	UserName  string         `protobuf:"bytes,3,opt,name=userName" json:"userName,omitempty"`
	Cover     string         `protobuf:"bytes,4,opt,name=cover" json:"cover,omitempty"`
	IsPlaying bool           `protobuf:"varint,5,opt,name=isPlaying" json:"isPlaying,omitempty"`
	IsFollow  bool           `protobuf:"varint,6,opt,name=isFollow" json:"isFollow,omitempty"`
	FollowNum int64          `protobuf:"varint,7,opt,name=followNum" json:"followNum,omitempty"`
	Broadcast *BroadcastInfo `protobuf:"bytes,8,opt,name=broadcast" json:"broadcast,omitempty"`
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
	UserID      string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	RoomID      string `protobuf:"bytes,2,opt,name=roomID" json:"roomID,omitempty"`
	BroadcastID string `protobuf:"bytes,3,opt,name=broadcastID" json:"broadcastID,omitempty"`
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
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x25, 0x75, 0x92, 0x25, 0x37, 0x74, 0x1b, 0x06, 0x4d, 0x51, 0xd9, 0x43, 0x14, 0x4d, 0x5a,
	0x84, 0xc4, 0x10, 0x45, 0x42, 0xe2, 0x11, 0xe8, 0x90, 0x2a, 0xa6, 0x09, 0x79, 0xec, 0x6d, 0x2f,
	0x59, 0xea, 0x42, 0xa4, 0xc6, 0x1e, 0xb6, 0x33, 0x34, 0x7e, 0x80, 0x0f, 0xe4, 0x53, 0xf8, 0x01,
	0x64, 0xc7, 0x4d, 0xda, 0xa8, 0x63, 0x93, 0xf6, 0x52, 0xf9, 0x9c, 0x9c, 0xeb, 0x7b, 0x7d, 0xcf,
	0xbd, 0x85, 0x1d, 0xc1, 0x79, 0xf5, 0x4a, 0xff, 0x1c, 0x5d, 0x09, 0xae, 0x38, 0x76, 0xf5, 0x39,
	0x7d, 0x0e, 0xe1, 0x07, 0xc1, 0xf3, 0x59, 0x91, 0x4b, 0x85, 0xb7, 0x61, 0x30, 0x9d, 0xc4, 0x4e,
	0xe2, 0x64, 0x21, 0x19, 0x94, 0x93, 0xf4, 0x1d, 0xa0, 0xd3, 0xba, 0xc2, 0x09, 0x44, 0xad, 0xa6,
	0xfd, 0x1e, 0x5d, 0x76, 0x14, 0xde, 0x35, 0xc2, 0x78, 0x90, 0x38, 0xd9, 0x90, 0x20, 0x56, 0x57,
	0x69, 0x05, 0xc3, 0x8f, 0x82, 0xe6, 0x8a, 0x12, 0xfa, 0xa3, 0xa6, 0x52, 0xe1, 0x3d, 0xf0, 0x6b,
	0x49, 0x45, 0x1b, 0x6f, 0x11, 0x1e, 0x41, 0xa0, 0x4f, 0xa7, 0x79, 0x45, 0x4d, 0x7c, 0x48, 0x5a,
	0x8c, 0x31, 0xb8, 0x4c, 0xf3, 0xc8, 0xf0, 0xe6, 0x8c, 0x9f, 0x81, 0x57, 0xf0, 0x6b, 0x2a, 0x62,
	0xd7, 0x90, 0x0d, 0x48, 0x3f, 0x43, 0x74, 0x52, 0x4a, 0xb5, 0x4c, 0xb6, 0x0b, 0xba, 0x08, 0x93,
	0xc9, 0x33, 0xf5, 0xe8, 0xab, 0x64, 0xf9, 0xab, 0x49, 0xe1, 0x11, 0x73, 0xd6, 0x25, 0x49, 0x9a,
	0x8b, 0xe2, 0xbb, 0x4d, 0x60, 0x51, 0xfa, 0xd7, 0x81, 0x80, 0x70, 0x5e, 0x4d, 0xd9, 0x9c, 0x6b,
	0x91, 0x6e, 0x54, 0x57, 0x77, 0x83, 0xda, 0xda, 0x06, 0x2b, 0xb5, 0xad, 0xbe, 0x05, 0xf5, 0xde,
	0xb2, 0xb1, 0x6e, 0xbc, 0x0f, 0x61, 0x29, 0xbf, 0x2c, 0xf2, 0x9b, 0x92, 0x7d, 0x8b, 0xbd, 0xc4,
	0xc9, 0x02, 0xd2, 0x11, 0xfa, 0xbe, 0x52, 0x7e, 0xe2, 0x8b, 0x05, 0xff, 0x19, 0xfb, 0xe6, 0x63,
	0x8b, 0x75, 0xe4, 0xdc, 0x9c, 0x74, 0xe3, 0xb7, 0x12, 0x27, 0x43, 0xa4, 0x23, 0xf0, 0x6b, 0x08,
	0x5b, 0x7f, 0xe2, 0x20, 0x71, 0xb2, 0x68, 0xfc, 0xf4, 0xc8, 0x98, 0xdf, 0x39, 0xc9, 0xe6, 0x9c,
	0x74, 0xaa, 0x74, 0x0a, 0x9e, 0x7e, 0xb4, 0xc4, 0x07, 0xe0, 0x69, 0xa5, 0x8c, 0x9d, 0x04, 0x65,
	0xd1, 0x78, 0xbb, 0x89, 0x5b, 0x36, 0x84, 0x34, 0x1f, 0x75, 0x6d, 0x8a, 0xab, 0x7c, 0xb1, 0xf4,
	0x1d, 0x91, 0x16, 0xa7, 0x6f, 0xc1, 0x3d, 0x97, 0x54, 0xdc, 0xea, 0xf9, 0x1e, 0xf8, 0xea, 0xe6,
	0x8a, 0x4e, 0x27, 0xd6, 0x0e, 0x8b, 0xd2, 0x0b, 0x08, 0x74, 0x9c, 0x4e, 0xf5, 0xbf, 0x58, 0xeb,
	0xc7, 0x60, 0xcd, 0x8f, 0x04, 0x56, 0x27, 0xd2, 0xb6, 0x7f, 0x95, 0x4a, 0x2f, 0xc0, 0x3f, 0x53,
	0xb9, 0xaa, 0x25, 0x8e, 0x61, 0x4b, 0xd6, 0x45, 0x41, 0xa5, 0x34, 0x97, 0x07, 0x64, 0x09, 0x1f,
	0x70, 0xfb, 0x1f, 0x07, 0x86, 0x6b, 0xbd, 0xed, 0xc7, 0x6c, 0x58, 0x9b, 0xdb, 0xb2, 0xed, 0x43,
	0x28, 0x55, 0x2e, 0xd4, 0xd7, 0xd2, 0x0e, 0x12, 0x22, 0x1d, 0xa1, 0x3b, 0x3f, 0xab, 0x45, 0xae,
	0x4a, 0xce, 0xcc, 0x30, 0x21, 0xd2, 0x62, 0x7c, 0x00, 0x43, 0xe3, 0xc2, 0xfb, 0x7a, 0x56, 0x52,
	0x56, 0x50, 0x33, 0x53, 0x88, 0xac, 0x93, 0x38, 0x83, 0x9d, 0xa2, 0x16, 0x82, 0x32, 0xd5, 0xea,
	0x7c, 0xa3, 0xeb, 0xd3, 0xe3, 0xdf, 0x08, 0x5c, 0x63, 0xc7, 0x4b, 0xf0, 0x9b, 0x7d, 0xc6, 0x76,
	0x8e, 0xd6, 0xb6, 0x7b, 0xf4, 0xb8, 0x21, 0x9b, 0xfe, 0xa6, 0x8f, 0x70, 0x06, 0xae, 0xde, 0x47,
	0xfc, 0xa4, 0xe1, 0x57, 0x76, 0x73, 0x14, 0x75, 0xf3, 0x64, 0x95, 0xa6, 0x5b, 0x76, 0xcc, 0x96,
	0xfe, 0x8f, 0x7a, 0x63, 0x67, 0x94, 0xbe, 0x9d, 0xfd, 0xbe, 0xb6, 0x9f, 0xfd, 0x05, 0x04, 0xe7,
	0x6c, 0x7e, 0x5f, 0xad, 0x77, 0xa6, 0x5b, 0x8b, 0xa1, 0x13, 0x8e, 0x36, 0xed, 0x8a, 0xa9, 0x00,
	0x1d, 0xb3, 0xd9, 0x7d, 0x94, 0x87, 0xe0, 0x1d, 0x33, 0x45, 0xc5, 0x9d, 0xe9, 0x0f, 0xc1, 0x3b,
	0xa1, 0xf9, 0x35, 0xbd, 0x4b, 0x78, 0xe9, 0x9b, 0x7f, 0xed, 0x37, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x76, 0xfd, 0x5a, 0x81, 0xc8, 0x05, 0x00, 0x00,
}
