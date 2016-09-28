package utils

import (
	"math/rand"
	"time"
)

var rander *rand.Rand

func InitRander() {
	rander = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRand() *rand.Rand {
	return rander
}

func RandomInt(min, max int) int {
	return min + GetRand().Intn(max-min)
}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[GetRand().Intn(len(letterBytes))]
	}
	return string(b)
}
