package main

import (
	libsUtils "tech.cloudzen/libs/utils"
)

func main() {
	logConfig := map[string]interface{}{"log_path":"D:\\temp\\log\\xxb\\xx", "debug":true}
	libsUtils.LogConfig = logConfig
	Log := libsUtils.NewLogger()
	Log.Info("Info")
	Log.Debug("Debug")  //debug日志
	Log.Error("Error")
 }
