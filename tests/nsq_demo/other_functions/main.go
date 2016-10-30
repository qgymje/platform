package main

import (
	"flag"
	"log"
	"platform/utils"
)

var (
	topic   = flag.String("topic", "", "topic name")
	channel = flag.String("channel", "", "channel name")
)

func init() {
	flag.Parse()
	utils.SetEnv("dev")
	utils.InitConfig("../../../account_center/rpc/configs/")
	utils.InitLogger()
}

func main() {
	if *topic == "" {
		log.Fatal("topic name can not be empty")
	}

	if *channel != "" {
		if err := utils.DeleteChannel(*topic, *channel); err != nil {
			log.Fatal(err)
		}
	}
	if err := utils.DeleteTopic(*topic); err != nil {
		log.Fatal(err)
	}

}
