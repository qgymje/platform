package main

import (
	"qiniupkg.com/api.v7/kodo"
	//"qiniupkg.com/api.v7/conf"
	"fmt"

	"qiniupkg.com/api.v7/kodocli"
)

var (
	//设置上传到的空间
	bucket = "file"
	key    = "yourdefinekey_xx" //覆盖上传需要使用该key
)

// PutRet 构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

//reffernce : http://developer.qiniu.com/code/v7/sdk/go.html#upload-form

func main() {
	//初始化AK，SK
	//conf.ACCESS_KEY = "iUAPz2ApOeuqBGnGLVrWi6qWFESGuLFt9BCrH9pD"
	//conf.SECRET_KEY = "0hdfvXGr-FisuMr-O6DpEylcfdfY4ZneWQiPkY7O"

	cfg := kodo.Config{}
	cfg.AccessKey = "iUAPz2ApOeuqBGnGLVrWi6qWFESGuLFt9BCrH9pD"
	cfg.SecretKey = "0hdfvXGr-FisuMr-O6DpEylcfdfY4ZneWQiPkY7O"

	//创建一个Client
	c := kodo.New(0, &cfg)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		//Scope:   bucket ,  	    //非覆盖上传
		Scope: bucket + ":" + key, //覆盖上传
		//设置Token过期时间
		Expires: 3600,
	}

	//生成一个上传token
	token := c.MakeUptoken(policy)
	fmt.Println("token=", token)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	//设置上传文件的路径
	filepath := "./test2.jpg"
	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名

	//res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil) //非覆盖上传
	res := uploader.PutFile(nil, &ret, token, key, filepath, nil) //覆盖上传

	//打印返回的信息
	fmt.Println(ret.Hash, ret.Key)

	//打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}
	// check the file on qihiu :  https://portal.qiniu.com/bucket/file/resource
	domain := "http://oaa75dzf2.bkt.clouddn.com/"
	url := domain + ret.Key

	fmt.Println("the file's url is :" + url)

}
