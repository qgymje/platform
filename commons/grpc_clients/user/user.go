package userClient

import (
	"context"

	pb "platform/commons/protos/user"
	"platform/utils"

	"google.golang.org/grpc"
)

// User grpc user client
type User struct {
	conn   *grpc.ClientConn
	client pb.UserClient // why not pointer?
}

// NewUser create grpc user client
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

// Close close the client
func (u *User) Close() error {
	return u.conn.Close()
}

// EmailCode get the register code by email
func (u *User) EmailCode(in *pb.Email) (*pb.Code, error) {
	defer u.Close()
	return u.client.EmailCode(context.Background(), in)
}

// SMSCode get the register code by sms
func (u *User) SMSCode(in *pb.Phone) (*pb.Code, error) {
	defer u.Close()
	return u.client.SMSCode(context.Background(), in)
}

// Register register an account
func (u *User) Register(in *pb.RegisterInfo) (*pb.UserInfo, error) {
	defer u.Close()
	return u.client.Register(context.Background(), in)
}

// Login login action
func (u *User) Login(in *pb.LoginInfo) (*pb.UserInfo, error) {
	defer u.Close()
	return u.client.Login(context.Background(), in)
}

// Logout logout
func (u *User) Logout(in *pb.Token) (*pb.Status, error) {
	defer u.Close()
	return u.client.Logout(context.Background(), in)
}

// Auth get user info by token
func (u *User) Auth(in *pb.Token) (*pb.UserInfo, error) {
	defer u.Close()
	return u.client.Auth(context.Background(), in)
}

// Info user info
func (u *User) Info(in *pb.UserID) (*pb.UserInfo, error) {
	defer u.Close()
	return u.client.Info(context.Background(), in)
}
