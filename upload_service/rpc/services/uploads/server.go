package uploads

import (
	pb "platform/commons/protos/upload"
	"platform/utils"

	"golang.org/x/net/context"
)

// Server server is an grpc server object
type Server struct {
}

// Send implement the Send rpc
func (s *Server) Send(ctx context.Context, in *pb.FileInfo) (*pb.Status, error) {
	var err error
	defer func() {
		if err != nil {
			utils.GetLog().Error("upload service server send error: ", err)
		}
	}()
	/*
		u := NewUploader(in.Filename, in.Content)
		if err := u.Do(); err != nil {
			return nil, errors.New(u.ErrorCode().String())
		}
	*/
	return &pb.Status{Success: true}, nil
}
