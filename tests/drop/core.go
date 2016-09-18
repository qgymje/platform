package main

import (
	"github.com/micro/go-micro"
	"tech.cloudzen/libs/data/dropGen"
	"fmt"
)

func main() {
	service := micro.NewService()
	service.Init(
		micro.Name("cloudzen.tests.drop"),
		micro.Version("last"),
	)
	service.String()
	for _, drop := range dropGen.Drop(3) {
		fmt.Println(drop)
	}
}
