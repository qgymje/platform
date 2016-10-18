package utils

import (
	"math/rand"
	"net"
	"time"
)

var rander *rand.Rand

// InitRander called in system initialized
func InitRander() {
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GetRand get the global rand seed
func GetRand() *rand.Rand {
	return rander
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return min + GetRand().Intn(max-min)
}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generates a random string, length by n
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[GetRand().Intn(len(letterBytes))]
	}
	return string(b)
}

// RandPort asks OS a random address
func RandPort() int {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		GetLog().Error("utils.RandPort error: %+v", err)
	}
	defer lis.Close()
	return lis.Addr().(*net.TCPAddr).Port
}
