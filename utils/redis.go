package utils

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

// GetRedis get a redis connection
func GetRedis() (redis.Conn, error) {
	redisAddr := GetConf().GetString("redis.address")
	conn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatalf("Couldn't connect to Redis: %v\n", err)

	}
	//defer conn.Close()
	return conn, err
}
