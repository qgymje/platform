package controllers

import (
	"io"
	"log"
	"net/http"
	"os"
	"platform/commons/codes"
	"platform/commons/grpc_clients/upload"
	pb "platform/commons/protos/upload"
	"platform/utils"

	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

// Profile user profile controller
type Profile struct {
	Base
}

// Avatar upload a avatar image file
func (p *Profile) Avatar(c *gin.Context) {
	file, header, err := c.Request.FormFile("avatar")
	log.Println(header.Header)
	filename := uuid.NewV4().String()
	local, _ := os.Create(getUploadPath() + filename + ".jpg")
	defer local.Close()
	_, err = io.Copy(local, file)
	if err != nil {
		respformat := p.Response(c, codes.ErrorCodeUpload, nil)
		c.JSON(http.StatusOK, respformat)
		return
	}

	client := uploadClient.NewUpload(p.getUserRPCAddress())
	fileHTTPPath := utils.BaseURL() + "/uploads/" + filename + ".jpg"
	pbFile := pb.FileInfo{Filename: filename, FilePath: fileHTTPPath}
	log.Println(pbFile)
	reply, err := client.Send(&pbFile)
	if err != nil {
		respformat := p.Response(c, rpcErrorFormat(err.Error()), nil)
		c.JSON(http.StatusOK, respformat)
		return
	}
	respformat := p.Response(c, codes.ErrorCodeSuccess, reply)
	c.JSON(http.StatusOK, respformat)
	return
}
