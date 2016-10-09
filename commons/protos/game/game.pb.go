// Code generated by protoc-gen-go.
// source: game/game.proto
// DO NOT EDIT!

/*
Package game is a generated protocol buffer package.

It is generated from these files:
	game/game.proto

It has these top-level messages:
	Status
	PreferenceConfig
	UserGame
	GameVM
	Games
	GameInfo
	Page
*/
package game

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

type Status struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PreferenceConfig struct {
	Json string `protobuf:"bytes,1,opt,name=json" json:"json,omitempty"`
}

func (m *PreferenceConfig) Reset()                    { *m = PreferenceConfig{} }
func (m *PreferenceConfig) String() string            { return proto.CompactTextString(m) }
func (*PreferenceConfig) ProtoMessage()               {}
func (*PreferenceConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type UserGame struct {
	UserID string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	GameID string `protobuf:"bytes,2,opt,name=gameID" json:"gameID,omitempty"`
}

func (m *UserGame) Reset()                    { *m = UserGame{} }
func (m *UserGame) String() string            { return proto.CompactTextString(m) }
func (*UserGame) ProtoMessage()               {}
func (*UserGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GameVM struct {
	IP   string `protobuf:"bytes,1,opt,name=IP,json=iP" json:"IP,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
}

func (m *GameVM) Reset()                    { *m = GameVM{} }
func (m *GameVM) String() string            { return proto.CompactTextString(m) }
func (*GameVM) ProtoMessage()               {}
func (*GameVM) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Games struct {
	Games    []*GameInfo `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
	TotalNum int64       `protobuf:"varint,2,opt,name=totalNum" json:"totalNum,omitempty"`
}

func (m *Games) Reset()                    { *m = Games{} }
func (m *Games) String() string            { return proto.CompactTextString(m) }
func (*Games) ProtoMessage()               {}
func (*Games) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Games) GetGames() []*GameInfo {
	if m != nil {
		return m.Games
	}
	return nil
}

type GameInfo struct {
	GameID       string   `protobuf:"bytes,1,opt,name=gameID" json:"gameID,omitempty"`
	CompanyID    string   `protobuf:"bytes,2,opt,name=companyID" json:"companyID,omitempty"`
	Name         string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	GameTypeID   int32    `protobuf:"varint,4,opt,name=gameTypeID" json:"gameTypeID,omitempty"`
	GameTypeName string   `protobuf:"bytes,5,opt,name=gameTypeName" json:"gameTypeName,omitempty"`
	Description  string   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Cover        string   `protobuf:"bytes,7,opt,name=cover" json:"cover,omitempty"`
	Screenshots  []string `protobuf:"bytes,8,rep,name=screenshots" json:"screenshots,omitempty"`
	PlayTimes    int64    `protobuf:"varint,9,opt,name=playTimes" json:"playTimes,omitempty"`
	PlayerNum    int64    `protobuf:"varint,10,opt,name=playerNum" json:"playerNum,omitempty"`
	IsFree       bool     `protobuf:"varint,11,opt,name=isFree" json:"isFree,omitempty"`
	Charge       float64  `protobuf:"fixed64,12,opt,name=charge" json:"charge,omitempty"`
	PayStatus    bool     `protobuf:"varint,13,opt,name=payStatus" json:"payStatus,omitempty"`
}

func (m *GameInfo) Reset()                    { *m = GameInfo{} }
func (m *GameInfo) String() string            { return proto.CompactTextString(m) }
func (*GameInfo) ProtoMessage()               {}
func (*GameInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Page struct {
	Num        int32  `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Size       int32  `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	GameTypeID int32  `protobuf:"varint,3,opt,name=gameTypeID" json:"gameTypeID,omitempty"`
	Search     string `protobuf:"bytes,4,opt,name=search" json:"search,omitempty"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Status)(nil), "game.Status")
	proto.RegisterType((*PreferenceConfig)(nil), "game.PreferenceConfig")
	proto.RegisterType((*UserGame)(nil), "game.UserGame")
	proto.RegisterType((*GameVM)(nil), "game.GameVM")
	proto.RegisterType((*Games)(nil), "game.Games")
	proto.RegisterType((*GameInfo)(nil), "game.GameInfo")
	proto.RegisterType((*Page)(nil), "game.Page")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Game service

type GameClient interface {
	Create(ctx context.Context, in *GameInfo, opts ...grpc.CallOption) (*Status, error)
	Start(ctx context.Context, in *UserGame, opts ...grpc.CallOption) (*GameVM, error)
	End(ctx context.Context, in *UserGame, opts ...grpc.CallOption) (*Status, error)
	List(ctx context.Context, in *Page, opts ...grpc.CallOption) (*Games, error)
	Preference(ctx context.Context, in *UserGame, opts ...grpc.CallOption) (*PreferenceConfig, error)
	UpdatePreference(ctx context.Context, in *UserGame, opts ...grpc.CallOption) (*Status, error)
}

type gameClient struct {
	cc *grpc.ClientConn
}

func NewGameClient(cc *grpc.ClientConn) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Create(ctx context.Context, in *GameInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/game.Game/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Start(ctx context.Context, in *UserGame, opts ...grpc.CallOption) (*GameVM, error) {
	out := new(GameVM)
	err := grpc.Invoke(ctx, "/game.Game/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) End(ctx context.Context, in *UserGame, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/game.Game/End", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) List(ctx context.Context, in *Page, opts ...grpc.CallOption) (*Games, error) {
	out := new(Games)
	err := grpc.Invoke(ctx, "/game.Game/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Preference(ctx context.Context, in *UserGame, opts ...grpc.CallOption) (*PreferenceConfig, error) {
	out := new(PreferenceConfig)
	err := grpc.Invoke(ctx, "/game.Game/Preference", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) UpdatePreference(ctx context.Context, in *UserGame, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/game.Game/UpdatePreference", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Game service

type GameServer interface {
	Create(context.Context, *GameInfo) (*Status, error)
	Start(context.Context, *UserGame) (*GameVM, error)
	End(context.Context, *UserGame) (*Status, error)
	List(context.Context, *Page) (*Games, error)
	Preference(context.Context, *UserGame) (*PreferenceConfig, error)
	UpdatePreference(context.Context, *UserGame) (*Status, error)
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.Game/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Create(ctx, req.(*GameInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.Game/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Start(ctx, req.(*UserGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_End_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).End(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.Game/End",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).End(ctx, req.(*UserGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Page)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.Game/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).List(ctx, req.(*Page))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Preference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Preference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.Game/Preference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Preference(ctx, req.(*UserGame))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_UpdatePreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).UpdatePreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.Game/UpdatePreference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).UpdatePreference(ctx, req.(*UserGame))
	}
	return interceptor(ctx, in, info, handler)
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Game_Create_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Game_Start_Handler,
		},
		{
			MethodName: "End",
			Handler:    _Game_End_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Game_List_Handler,
		},
		{
			MethodName: "Preference",
			Handler:    _Game_Preference_Handler,
		},
		{
			MethodName: "UpdatePreference",
			Handler:    _Game_UpdatePreference_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("game/game.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0x6d, 0x9a, 0x3f, 0xdb, 0xde, 0xf6, 0xb7, 0xbf, 0x32, 0x48, 0x19, 0x8a, 0x48, 0x1d, 0xfc,
	0xd3, 0x07, 0x59, 0xa1, 0x82, 0x0f, 0xbe, 0xee, 0xaa, 0x14, 0x74, 0x29, 0xd9, 0x5d, 0xdf, 0xc7,
	0xf4, 0xb6, 0x8d, 0x6c, 0x66, 0xc2, 0xcc, 0x54, 0xa8, 0x5f, 0xc2, 0x2f, 0xec, 0x83, 0xdc, 0x49,
	0xd2, 0xd4, 0x2e, 0xfa, 0x52, 0xee, 0x39, 0xf7, 0x9c, 0x9b, 0x9b, 0xc3, 0x4d, 0xe1, 0xff, 0x8d,
	0x2c, 0xf0, 0x35, 0xfd, 0x5c, 0x94, 0x46, 0x3b, 0xcd, 0x22, 0xaa, 0x85, 0x80, 0xe4, 0xc6, 0x49,
	0xb7, 0xb3, 0x8c, 0xc3, 0x99, 0xdd, 0x65, 0x19, 0x5a, 0xcb, 0x83, 0x69, 0x30, 0xeb, 0xa5, 0x0d,
	0x14, 0x2f, 0x60, 0xb4, 0x34, 0xb8, 0x46, 0x83, 0x2a, 0xc3, 0x4b, 0xad, 0xd6, 0xf9, 0x86, 0x31,
	0x88, 0xbe, 0x59, 0xad, 0xbc, 0xb4, 0x9f, 0xfa, 0x5a, 0xbc, 0x83, 0xde, 0x9d, 0x45, 0xf3, 0x51,
	0x16, 0xc8, 0xc6, 0x90, 0xec, 0x2c, 0x9a, 0xc5, 0x55, 0xad, 0xa8, 0x11, 0xf1, 0xf4, 0xdc, 0xc5,
	0x15, 0xef, 0x56, 0x7c, 0x85, 0xc4, 0x2b, 0x48, 0xc8, 0xf7, 0xe5, 0x33, 0x3b, 0x87, 0xee, 0x62,
	0x59, 0xbb, 0xba, 0xf9, 0x92, 0x9e, 0x54, 0x6a, 0xe3, 0x6a, 0xbd, 0xaf, 0xc5, 0x02, 0x62, 0x52,
	0x5b, 0xf6, 0x0c, 0x62, 0x1a, 0x40, 0x2b, 0x87, 0xb3, 0xc1, 0xfc, 0xfc, 0xc2, 0xbf, 0x20, 0xf5,
	0x16, 0x6a, 0xad, 0xd3, 0xaa, 0xc9, 0x26, 0xd0, 0x73, 0xda, 0xc9, 0xfb, 0xeb, 0x5d, 0xe1, 0xc7,
	0x84, 0xe9, 0x01, 0x8b, 0x5f, 0x5d, 0xe8, 0x35, 0xfa, 0xa3, 0xed, 0x82, 0xe3, 0xed, 0xd8, 0x63,
	0xe8, 0x67, 0xba, 0x28, 0xa5, 0xda, 0x1f, 0x16, 0x6f, 0x09, 0xda, 0x50, 0xc9, 0x02, 0x79, 0x58,
	0x6d, 0x48, 0x35, 0x7b, 0x02, 0x40, 0xde, 0xdb, 0x7d, 0x49, 0xd3, 0xa2, 0x69, 0x30, 0x8b, 0xd3,
	0x23, 0x86, 0x09, 0x18, 0x36, 0xe8, 0x9a, 0xbc, 0xb1, 0xf7, 0xfe, 0xc1, 0xb1, 0x29, 0x0c, 0x56,
	0x68, 0x33, 0x93, 0x97, 0x2e, 0xd7, 0x8a, 0x27, 0x5e, 0x72, 0x4c, 0xb1, 0x47, 0x10, 0x67, 0xfa,
	0x3b, 0x1a, 0x7e, 0xe6, 0x7b, 0x15, 0x20, 0x9f, 0xcd, 0x0c, 0xa2, 0xb2, 0x5b, 0xed, 0x2c, 0xef,
	0x4d, 0x43, 0xf2, 0x1d, 0x51, 0xf4, 0x3e, 0xe5, 0xbd, 0xdc, 0xdf, 0xe6, 0x14, 0x5d, 0xdf, 0x27,
	0xd2, 0x12, 0x4d, 0x17, 0x0d, 0xe5, 0x05, 0x6d, 0xd7, 0x13, 0x94, 0x51, 0x6e, 0x3f, 0x18, 0x44,
	0x3e, 0xf0, 0x67, 0x52, 0x23, 0xe2, 0xb3, 0xad, 0x34, 0x1b, 0xe4, 0xc3, 0x69, 0x30, 0x0b, 0xd2,
	0x1a, 0xf9, 0x69, 0x72, 0x5f, 0x1d, 0x19, 0xff, 0xcf, 0x5b, 0x5a, 0x42, 0xac, 0x20, 0x5a, 0xca,
	0x0d, 0xb2, 0x11, 0x84, 0x6a, 0x57, 0xf8, 0xd8, 0xe3, 0x94, 0x4a, 0x4a, 0xd5, 0xe6, 0x3f, 0xd0,
	0xc7, 0x1d, 0xa7, 0xbe, 0x3e, 0x49, 0x35, 0x7c, 0x90, 0xea, 0x18, 0x12, 0x8b, 0xd2, 0x64, 0x5b,
	0x9f, 0x78, 0x3f, 0xad, 0xd1, 0xfc, 0x67, 0x17, 0x22, 0x7f, 0x96, 0x33, 0x48, 0x2e, 0x0d, 0x4a,
	0x87, 0xec, 0xe4, 0x54, 0x26, 0xc3, 0x0a, 0xd7, 0x6b, 0x75, 0xd8, 0x4b, 0x88, 0x6f, 0x9c, 0x34,
	0xae, 0x11, 0x36, 0x97, 0xdd, 0x08, 0xab, 0x6b, 0x15, 0x1d, 0xf6, 0x1c, 0xc2, 0xf7, 0x6a, 0xf5,
	0x37, 0xd9, 0x61, 0xde, 0x53, 0x88, 0x3e, 0xe5, 0xd6, 0x31, 0xa8, 0x78, 0x7a, 0xe9, 0xc9, 0xa0,
	0x1d, 0x45, 0x92, 0xb7, 0x00, 0xed, 0x77, 0xf6, 0x60, 0xe0, 0xb8, 0x36, 0x9e, 0x7c, 0x89, 0xa2,
	0xc3, 0xe6, 0x30, 0xba, 0x2b, 0x57, 0xd2, 0xe1, 0x3f, 0xdc, 0x27, 0xeb, 0x7c, 0x4d, 0xfc, 0x9f,
	0xc0, 0x9b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xc8, 0xf6, 0xc5, 0x17, 0x04, 0x00, 0x00,
}
