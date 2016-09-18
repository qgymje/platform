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

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.APIVersion())
	if utils.IsDev() {
		r.Use(middlewares.FakedLogin())
	}

	rr := r.Group("/room")
	{
		room := new(controllers.Room)
		rr.POST("", room.Create)
		rr.PUT("/start", room.Start)
		rr.PUT("/end", room.End)
		rr.POST("/barrage", room.Barrage)
	}

	port := utils.GetConf().GetString("app.http_port")
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
