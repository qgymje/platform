package main

import (
	"flag"
	"log"

	"platform/commons/middlewares"
	"platform/game_center/http/controllers"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

var (
	configPath = flag.String("conf", "./configs/", "set config path")
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	port       = flag.String("port", ":3002", "game center http port")
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
		return utils.GetConf().GetString("app.rpc_port")
	}
	return *port
}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.APIVersion())
	r.Use(middlewares.APILang())
	r.Use(middlewares.RecordRequestBegin())

	g := r.Group("/v1/game")
	{
		game := new(controllers.Game)
		g.POST("/", game.Create)
		g.GET("/", game.List)
		g.GET("/search/:search", game.List)
		g.GET("/types", game.Types)
		g.POST("/start", game.Start)
		g.POST("/end", game.End)
		g.GET("/preference", game.Preference)
		g.PUT("/preference", game.UpdatePreference)
	}

	if err := r.Run(getPort()); err != nil {
		log.Fatal(err)
	}
}
