package main

import (
	"flag"
	"log"
	"net"

	pb "platform/commons/protos/coupon"
	"platform/coupon_center/rpc/models"
	"platform/coupon_center/rpc/services"
	"platform/coupon_center/rpc/services/coupons"
	"platform/utils"

	"google.golang.org/grpc"
)

var (
	configPath = flag.String("conf", "./configs/", "set config path")
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	port       = flag.String("port", ":4004", "service port")
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

	models.InitModels()
}

func getPort() string {
	if *port == "" {
		return utils.GetConf().GetString("app.rpc_port")
	}
	return *port
}

func main() {
	lis, err := net.Listen("tcp", getPort())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCouponServer(s, &services.Server{})
	go coupons.Sync()
	err = s.Serve(lis)
	if err != nil {
		log.Println("server start failed: ", err)
	}
}
