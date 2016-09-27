package main

import (
	"flag"
	"log"
	"net"

	pb "platform/commons/protos/broadcasting"
	"platform/live_broadcast/rpc/models"
	"platform/live_broadcast/rpc/services/broadcasting"

	"platform/utils"

	"google.golang.org/grpc"
)

var (
	env        = flag.String("env", "dev", "set env: dev, test, prod")
	configPath = flag.String("conf", "./configs/", "set config path")
	port       = flag.String("port", ":4003", "game center http port")
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
	pb.RegisterBroadcastingServer(s, &broadcastings.BroadcastingServer{})
	broadcastings.StartToReceive()
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
