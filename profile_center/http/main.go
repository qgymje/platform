package main

import (
	"flag"
	"log"

	"platform/commons/middlewares"
	"platform/profile_center/http/controllers"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

var (
	configPath = flag.String("conf", "./configs/", "set config path")
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	port       = flag.String("port", ":3010", "service port")
)

func initEnv() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	flag.Parse()
	log.Println("current env is: ", *env)
	utils.SetEnv(*env)
}

func init() {
	initEnv()
	utils.InitConfig(*configPath)
	utils.InitLogger()
	utils.InitRander()
}

func getPort() string {
	if *port == "" {
		return utils.GetConf().GetString("app.http_port")
	}
	return *port
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.APILang())
	r.Use(middlewares.RecordRequestBegin())
	if utils.IsDev() {
		r.Use(middlewares.FakedLogin())
	}

	uploadPath := "./uploads"
	utils.EnsurePath(uploadPath)
	controllers.SetUploadPath(uploadPath)
	r.Static("/uploads", uploadPath)

	c := r.Group("/v1/friend")
	{
		friend := new(controllers.Friend)
		c.GET("/", friend.List)
		c.POST("/", friend.Request)
		c.PUT("/agree", friend.Agree)
		c.PUT("/refuse", friend.Refuse)
		c.PUT("/recommend", friend.Recommend)
	}

	if err := r.Run(getPort()); err != nil {
		log.Fatal(err)
	}
}
