package utils

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

func ConnectMongodb() *mgo.Session {
	m := GetConf().GetStringMapString("mongodb")
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{m["hosts"]},
		Timeout:  60 * time.Second,
		Database: m["dbname"],
		Username: m["username"],
		Password: m["password"],
	}
	mongoSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic("connect mongodb failed")
	}
	//defer mongoSession.Close()

	mongoSession.SetMode(mgo.Monotonic, true)
	return mongoSession
}
