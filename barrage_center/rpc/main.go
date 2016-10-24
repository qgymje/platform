package main

import (
	"flag"
	"log"
	"net"

	"platform/barrage_center/rpc/models"
	"platform/barrage_center/rpc/services"
	pb "platform/commons/protos/barrage"
	"platform/utils"

	"google.golang.org/grpc"
)

var (
	configPath = flag.String("conf", "./configs/", "set config path")
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	port       = flag.String("port", ":4007", "service port")
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
	session := utils.ConnectMongodb()
	models.InitMongodb(session)
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
	pb.RegisterBarrageServer(s, &services.Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Println("server start failed: ", err)
	}
}
