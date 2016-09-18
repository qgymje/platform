package main

import (
	"flag"
	"log"
	"net"

	"platform/broadcast_room/rpc/models"
	"platform/broadcast_room/rpc/services/room"
	pb "platform/commons/protos/room"
	"platform/utils"

	"google.golang.org/grpc"
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
	pb.RegisterRoomServer(s, &rooms.RoomServer{})
	log.Println("broadcast room rpc server running on ", port)
	s.Serve(lis)
}
