package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
)

type Logger struct {
	*logs.BeeLogger
}

var logger *logs.BeeLogger

func InitLogger() {
	logger = logs.NewLogger(10000)
	defer logger.Flush()

	fileName := GetConf().GetString("log.filename")
	log.Println("logger filename = ", fileName)
	_, err := os.Stat(fileName)
	if err != nil {
		err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
		if err != nil {
			panic("无法创建日志目录")
		}
	}

	enableCallDepth := GetConf().GetBool("log.enable_func_call_depth")
	logger.EnableFuncCallDepth(enableCallDepth)

	logConf := fmt.Sprintf(`{"filename":"%s"}`, fileName)
	logger.SetLogger("file", logConf)
	logger.SetLevel(getLogLevel())
}

func getLogLevel() int {
	env := GetEnv()
	switch env {
	case ENV_DEV:
		return logs.LevelDebug
	case ENV_TEST:
		return logs.LevelDebug
	case ENV_PROD:
		return logs.LevelError
	default:
		return logs.LevelDebug
	}
}

func GetLog() *Logger {
	return &Logger{
		logger,
	}
}

func (l *Logger) Error(fmt string, err error) {
	err = errors.Wrap(err, "")
	l.BeeLogger.Error(fmt, err)
}
