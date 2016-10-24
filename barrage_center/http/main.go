package main

import (
	"flag"
	"log"

	"platform/barrage_center/http/controllers"
	"platform/commons/middlewares"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

var (
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	configPath = flag.String("conf", "./configs/", "set config path")
	port       = flag.String("port", ":3005", "service port")
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

	g := r.Group("/v1/barrage")
	{
		b := new(controllers.Barrage)
		g.GET("/", b.List)
		g.POST("/", b.Create)
	}

	if err := r.Run(getPort()); err != nil {
		log.Fatal(err)
	}
}
