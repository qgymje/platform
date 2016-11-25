package main

import (
	"flag"
	"log"

	"platform/chat_center/http/controllers"
	"platform/commons/middlewares"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

var (
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	configPath = flag.String("conf", "./configs/", "set config path")
	port       = flag.String("port", ":3011", "service port")
)

func initEnv() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
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
	if utils.IsDev() {
		r.Use(middlewares.FakedLogin())
	}

	uploadPath := "./uploads"
	utils.EnsurePath(uploadPath)
	controllers.SetUploadPath(uploadPath)
	r.Static("/v1/chat/uploads", uploadPath)

	rr := r.Group("/v1/chat")
	{
		chat := new(controllers.Chat)
		rr.GET("/", chat.List)
		rr.POST("/", chat.Create)
		rr.POST("/send", chat.Send)
	}

	if err := r.Run(getPort()); err != nil {
		log.Fatal(err)
	}
}
