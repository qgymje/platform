package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"tech.cloudzen/live_broadcast/rpc/models"
	"tech.cloudzen/live_broadcast/rpc/services/broadcasting"
	pb "tech.cloudzen/protos/broadcasting"
	"tech.cloudzen/utils"
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

	session := utils.ConnectMongodb()
	models.InitMongodb(session)
}

func main() {
	port := utils.GetConf().GetString("app.rpc_port")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBroadcastingServer(s, &broadcastings.BroadcastingServer{})
	broadcastings.StartToReceive()
	s.Serve(lis)
}
