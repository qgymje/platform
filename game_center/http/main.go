package main

import (
	"flag"
	"log"

	"platform/account_center/http/controllers"
	"platform/commons/middlewares"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

var (
	configPath = flag.String("conf", "./configs/", "set config path")
	env        = flag.String("env", "dev", "set env: dev, test, prod")
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

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.APIVersion())
	r.Use(middlewares.RecordRequestBegin())

	g := r.Group("/game")
	{
		game := new(controllers.Game)
		g.GET("", game.List)
		g.PUT("/preference", game.UpdatePreference)
		g.GET("/preference", game.Preference)
		g.POST("/start", game.Start)
	}

	port := utils.GetConf().GetString("app.http_port")
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
