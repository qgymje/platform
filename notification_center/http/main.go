package main

import (
	"flag"
	"log"

	"platform/commons/middlewares"
	"platform/notification_center/http/controllers"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

var (
	configPath = flag.String("conf", "./configs/", "set config path")
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	port       = flag.String("port", ":3012", "service port")
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
	r.Use(middlewares.RecordRequestBegin())
	if utils.IsDev() {
		r.Use(middlewares.FakedLogin())
	}

	uploadPath := "./uploads"
	utils.EnsurePath(uploadPath)
	controllers.SetUploadPath(uploadPath)
	r.Static("/uploads", uploadPath)

	c := r.Group("/v1/notify")
	{
		// web api
		notification := new(controllers.Notification)
		c.GET("/:token", notification.Notify)
	}

	if err := r.Run(getPort()); err != nil {
		log.Fatal(err)
	}
}
