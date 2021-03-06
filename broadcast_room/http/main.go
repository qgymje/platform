package main

import (
	"flag"
	"log"

	"platform/broadcast_room/http/controllers"
	"platform/commons/middlewares"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

var (
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	configPath = flag.String("conf", "./configs/", "set config path")
	port       = flag.String("port", ":3001", "service port")
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

	rr := r.Group("/v1/room")
	{
		room := new(controllers.Room)
		rr.GET("/", room.List)
		rr.POST("/", room.Create)
		rr.GET("/info", room.Info)
		rr.GET("/info/:room_id", room.Info)
		rr.GET("/types", room.Types)
		rr.POST("/follow", room.Follow)
		rr.POST("/unfollow", room.Unfollow)
	}

	if err := r.Run(getPort()); err != nil {
		log.Fatal(err)
	}
}
