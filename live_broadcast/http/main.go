package main

import (
	"flag"
	"log"

	"platform/live_broadcast/http/controllers"
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

	bro := r.Group("/broadcasting")
	{
		b := new(controllers.Broadcasting)
		bro.GET("/join/:id", b.Join)
		bro.GET("/leave/:id", b.Leave)
	}

	port := utils.GetConf().GetString("app.http_port")
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
