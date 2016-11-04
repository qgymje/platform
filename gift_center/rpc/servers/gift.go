package servers

import (
	pb "platform/commons/protos/gift"
	"platform/utils"

	"golang.org/x/net/context"
)

// Gift server
type Gift struct {
}

// List gift list
func (g *Gift) List(ctx context.Context, in *pb.Page) (*pb.Gifts, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.Gift.List error: %+v", err)
		}
	}()

	return nil, nil
}

// Send send a gift
func (g *Gift) Send(ctx context.Context, in *pb.SendGift) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("servers.Gift.Send error: %+v", err)
		}
	}()

	return nil, nil
}
