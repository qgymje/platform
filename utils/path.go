package utils

import "os"

func EnsurePath(path string) (err error) {
	_, err = os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, 0744)
	}
	if err != nil {
		GetLog().Error("helpers.MakeSurePath: MkdirAll error=%v, path=%v", err, path)
	}

	return
}

func BaseURL() string {
	baseURL := GetConf().GetString("app.base_url")
	httpport := GetConf().GetString("app.http_port")
	if httpport != "80" {
		baseURL += ":" + httpport
	}
	return baseURL
}
