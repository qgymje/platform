package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
 	"os"
	"runtime"
	"time"
	libsUtils "tech.cloudzen/libs/utils"
 )

var Log *log.Logger
var DEFAULT_LOG_PATH_WIN =  "D:\\temp\\log"
var DEFAULT_LOG_PATH_UNIX =  "/var/log/tech.cloudzen"
var separator string

func NewLogger(config map[string]interface{}) *log.Logger {
	if Log != nil {
		return Log
	}
	var logPath = ""
	Log = log.New()
	if(config["debug"].(bool)==true) {
		Log.Level = log.DebugLevel
	}
	if(config["log_path"]!="") {
		logPath = config["log_path"].(string)
	}
	if(runtime.GOOS=="windows") {
		if(logPath=="") {
			logPath = DEFAULT_LOG_PATH_WIN
		}
		separator = "\\"
	}else{
		if(logPath=="") {
			logPath = DEFAULT_LOG_PATH_UNIX
		}
		separator = "/"
	}
 	os.MkdirAll(logPath, 0755)
	Log.Formatter = new(log.JSONFormatter)
  	Log.Hooks.Add(lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel : logPath + separator +"info_"+ libsUtils.TimeFormat(time.Now(), "yyyy-MM-dd") + ".log",
		log.ErrorLevel :  logPath + separator +"error_"+ libsUtils.TimeFormat(time.Now(), "yyyy-MM-dd") + ".log",
		log.DebugLevel :  logPath + separator +"debug_"+ libsUtils.TimeFormat(time.Now(), "yyyy-MM-dd") + ".log",
		log.PanicLevel :  logPath + separator +"panic_"+ libsUtils.TimeFormat(time.Now(), "yyyy-MM-dd") + ".log",
	}))
	return Log
}

func main() {
	logConfig := map[string]interface{}{"log_path":"D:\\temp\\log\\xxx\\ccc", "debug":true}
	Log = NewLogger(logConfig)
	Log.Info("Info")
	Log.Debug("Debug")
	Log.Error("Error")
	Log.Panic("Panic")
}