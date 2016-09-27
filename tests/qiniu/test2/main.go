package main

import (
	"context"
	"flag"
	"log"

	"qiniupkg.com/api.v7/kodo"
)

var (
	// test account only have one bucket
	folder = flag.String("folder", "file", "bucket name, like avatar, cover")
	file   = flag.String("file", "../test2.jpg", "file path")
)

const (
	accessKey = "iUAPz2ApOeuqBGnGLVrWi6qWFESGuLFt9BCrH9pD"
	secretKey = "0hdfvXGr-FisuMr-O6DpEylcfdfY4ZneWQiPkY7O"
)

func init() {
	flag.Parse()
	log.SetFlags(log.Ltime | log.Llongfile)
}

func main() {
	kodo.SetMac(accessKey, secretKey)
	zone := 0
	client := kodo.New(zone, nil)

	bucket := client.Bucket(*folder)
	ctx := context.Background()
	visitKey := "hellscream.jpg"
	err := bucket.PutFile(ctx, nil, visitKey, *file, nil)
	if err != nil {
		log.Println("putfile error: ", err)
		return
	}

	fileURL := "http://oaa75dzf2.bkt.clouddn.com/" + visitKey
	log.Println(fileURL)
}
