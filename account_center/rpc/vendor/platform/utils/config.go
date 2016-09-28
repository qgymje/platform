package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var conf *viper.Viper

func InitConfig(path string) {
	conf = viper.New()
	conf.AddConfigPath(path)
	conf.SetConfigName(fmt.Sprintf("%s", env))
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		log.Fatal("read config file error: ", err)
	}
}

func GetConf() *viper.Viper {
	return conf
}
