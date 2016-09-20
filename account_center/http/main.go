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

	u := r.Group("/user")
	{
		user := new(controllers.User)
		u.POST("/register", user.Register)
		u.GET("/auth/:token", user.Auth)
		u.PUT("/login", user.Login)
		u.DELETE("/logout", user.Logout)
		u.GET("/info/:user_id", user.Info)

		sms := new(controllers.SMS)
		u.POST("/register/sms", sms.RegisterCode)
		u.PUT("/register/sms", sms.VerifyRegisterCode)

		email := new(controllers.Email)
		u.POST("/register/email", email.RegisterCode)
		u.PUT("/register/email", email.VerifyRegisterCode)
	}

	port := utils.GetConf().GetString("app.http_port")
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
