package uploadClient

import (
	"context"
	"io"
	"log"
	pb "platform/commons/protos/upload"
	"platform/utils"

	"google.golang.org/grpc"
)

// Upload represent a upload grpc client object
type Upload struct {
	conn   *grpc.ClientConn
	client pb.UploadClient
}

// NewUpload create a grpc client
func NewUpload(address string) *Upload {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("upload grpc server cant not connect: %v", err)
	}

	u := new(Upload)
	u.conn = conn
	u.client = pb.NewUploadClient(u.conn)

	return u
}

// Close close the connection
func (u *Upload) Close() error {
	return u.conn.Close()
}

// Send the uploaded file
func (u *Upload) Send(filename string, reader io.Reader) (*pb.Status, error) {
	defer u.Close()
	stream, err := u.client.Send(context.Background())
	if err != nil {
		log.Fatalf("grpc client upload send error: %v", err)
	}
	// streaming send
	stream.Send()

	return stream.CloseAndRecv()
}
