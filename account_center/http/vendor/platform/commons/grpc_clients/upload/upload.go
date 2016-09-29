package uploadClient

import (
	"context"
	pb "platform/commons/protos/upload"
	"platform/utils"

	"google.golang.org/grpc"
)

// Uploader represent a upload grpc client object
type Uploader struct {
	conn   *grpc.ClientConn
	client pb.UploadClient
}

// NewUpload create a grpc client
func NewUpload(address string) *Uploader {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utils.GetLog().Error("uploader grpc server cant not connect: %v", err)
	}

	u := new(Uploader)
	u.conn = conn
	u.client = pb.NewUploadClient(u.conn)

	return u
}

// Close close the connection
func (u *Uploader) Close() error {
	return u.conn.Close()
}

// Send the uploaded file
func (u *Uploader) Send(in *pb.FileInfo) (*pb.Status, error) {
	defer u.Close()
	return u.client.Send(context.Background(), in)
}
