package uploads

import (
	"io"
	pb "platform/commons/protos/upload"
)

// Server server is an grpc server object
type Server struct {
}

// Send implement the Send rpc
func (s *Server) Send(stream pb.Upload_SendServer) error {
	for {
		buf, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Status{
				Success: true,
			})
		}
		if err != nil {
			return err
		}

	}
}
