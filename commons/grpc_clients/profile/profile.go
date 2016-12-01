package profileClient

import (
	"context"

	pb "platform/commons/protos/profile"
	"platform/utils"

	"google.golang.org/grpc"
)

// Profile profile grpc client
type Profile struct {
	conn   *grpc.ClientConn
	client pb.ProfileClient
}

// NewProfile  create a new grpc profile client
func NewProfile(address string) *Profile {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("profile grpc server cant not connect: %v", err)
	}

	p := new(Profile)
	p.conn = conn
	p.client = pb.NewProfileClient(p.conn)

	return p
}

// Close close grpc client
func (p *Profile) Close() error {
	return p.conn.Close()
}

// Withdraw withdraw
func (p *Profile) Withdraw(in *pb.Ammount) (*pb.Status, error) {
	defer p.Close()
	return p.client.Withdraw(context.Background(), in)
}

// WithdrawRollback rollback api
func (p *Profile) WithdrawRollback(in *pb.Message) (*pb.Status, error) {
	defer p.Close()
	return p.client.WithdrawRollback(context.Background(), in)
}

// WithdrawCommit commit api
func (p *Profile) WithdrawCommit(in *pb.Message) (*pb.Status, error) {
	defer p.Close()
	return p.client.WithdrawCommit(context.Background(), in)
}

// FriendRequest friend request
func (p *Profile) FriendRequest(in *pb.Request) (*pb.RequestID, error) {
	defer p.Close()
	return p.client.FriendRequest(context.Background(), in)
}

// FriendList friend list
func (p *Profile) FriendList(in *pb.Message) (*pb.Friends, error) {
	defer p.Close()
	return p.client.FriendList(context.Background(), in)
}
