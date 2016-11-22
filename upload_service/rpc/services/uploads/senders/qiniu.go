package senders

import (
	"context"
	"platform/utils"

	"qiniupkg.com/api.v7/kodo"
)

// QiniuProvider upload with qiniu services
type QiniuProvider struct {
	provider                                    Provider
	accessKey, secretKey, bucketName, urlPrefix string
}

// NewQiniuProvider qiniu upload service
func NewQiniuProvider(provider Provider) *QiniuProvider {
	p := new(QiniuProvider)
	p.provider = provider
	p.accessKey = "iUAPz2ApOeuqBGnGLVrWi6qWFESGuLFt9BCrH9pD"
	p.secretKey = "0hdfvXGr-FisuMr-O6DpEylcfdfY4ZneWQiPkY7O"
	p.bucketName = "file"
	p.urlPrefix = "http://oaa75dzf2.bkt.clouddn.com/"
	return p
}

// Do the upload process
func (p *QiniuProvider) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("qiniu provider upload error: %v", err)
		}
	}()

	kodo.SetMac(p.accessKey, p.secretKey)
	zone := 0
	client := kodo.New(zone, nil)

	bucket := client.Bucket(p.bucketName)
	ctx := context.Background()
	visitKey := p.provider.Filename() // file name to access
	filename := p.provider.Filename() // local file path
	err = bucket.PutFile(ctx, nil, visitKey, filename, nil)
	return
}

// RemoteURL get the uploaded file remote url
func (p *QiniuProvider) RemoteURL() string {
	return p.urlPrefix + p.provider.Filename()
}
