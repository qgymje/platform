package userClient

import (
	"context"
	"log"

	pb "platform/commons/protos/user"
	"platform/utils"

	"google.golang.org/grpc"
)

type User struct {
	conn   *grpc.ClientConn
	client pb.UserClient // why not pointer?
}

func NewUser(address string) *User {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("user grpc server cant not connect: %v", err)
	}

	u := new(User)
	u.conn = conn
	u.client = pb.NewUserClient(u.conn)

	return u
}

func (u *User) Close() error {
	return u.conn.Close()
}

func (u *User) ValidCode(in *pb.Phone) (*pb.Code, error) {
	defer u.Close()
	return u.client.ValidCode(context.Background(), in)
}

func (u *User) Register(in *pb.RegisterInfo) (*pb.UserInfo, error) {
	defer u.Close()
	upb, err := u.client.Register(context.Background(), in)
	log.Println(err)
	return upb, err
}

func (u *User) Login(in *pb.LoginInfo) (*pb.UserInfo, error) {
	defer u.Close()
	return u.client.Login(context.Background(), in)
}

func (u *User) Logout(in *pb.Token) (*pb.Status, error) {
	defer u.Close()
	return u.client.Logout(context.Background(), in)
}

func (u *User) Auth(in *pb.Token) (*pb.UserInfo, error) {
	defer u.Close()
	return u.client.Auth(context.Background(), in)
}

func (u *User) Info(in *pb.UserID) (*pb.UserInfo, error) {
	defer u.Close()
	return u.client.Info(context.Background(), in)
}
