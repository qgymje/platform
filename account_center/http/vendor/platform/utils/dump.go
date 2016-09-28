package utils

import "github.com/davecgh/go-spew/spew"

func Dump(x ...interface{}) {
	spew.Dump(x)
}

func Sdump(x ...interface{}) string {
	return spew.Sdump(x)
}
