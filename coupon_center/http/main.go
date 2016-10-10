package main

import (
	"flag"
	"log"

	"platform/commons/middlewares"
	"platform/coupon_center/http/controllers"
	"platform/utils"

	"github.com/gin-gonic/gin"
)

var (
	configPath = flag.String("conf", "./configs/", "set config path")
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	port       = flag.String("port", ":3004", "service port")
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

	uploadPath := "./uploads"
	utils.EnsurePath(uploadPath)
	controllers.SetUploadPath(uploadPath)
	r.Static("/uploads", uploadPath)

	c := r.Group("/coupon")
	{
		// web api
		account := new(controllers.Account)
		c.POST("/account", account.Create)
		c.GET("/account", account.List)
		c.PUT("/account/perm", account.UpdatePermission)
		c.DELETE("/account", account.Delete)

		store := new(controllers.Store)
		c.GET("/store", store.List)
		c.GET("/store/:store_id", store.Show)
		c.POST("/store", store.Create)
		c.PUT("/store", store.Update)

		coupon := new(controllers.Coupon)
		c.GET("/", coupon.List)
		c.GET("/bystore/:store_id", coupon.ListByStore)
		c.GET("/show/:coupon_id", coupon.Show)
		c.POST("/", coupon.Create)
		c.PUT("/", coupon.Update)

		// mobile api
	}

	if err := r.Run(getPort()); err != nil {
		log.Fatal(err)
	}
}
