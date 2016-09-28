package utils

import "os"

type Env int

const (
	ENV_NONE = 0
	ENV_DEV  = 1
	ENV_TEST = 2
	ENV_PROD = 3
)

var envString = [...]string{
	"none",
	"dev",
	"test",
	"prod",
}

func (e Env) String() string {
	return envString[int(e)]
}

var env Env

func SetEnv(k string) {
	found := false
	for i, v := range envString {
		if k == v {
			found = true
			env = Env(i)
		}
	}

	if !found {
		panic("env mode error")
		os.Exit(2)
	}

}

func GetEnv() Env {
	return env
}

func IsDev() bool {
	return GetEnv() == ENV_DEV
}

func IsTest() bool {
	return GetEnv() == ENV_TEST
}

func IsProd() bool {
	return GetEnv() == ENV_PROD
}
