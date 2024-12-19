package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logOnce sync.Once
	l       *logrus.Logger
)

func GetLogger() *logrus.Logger {
	logOnce.Do(func() {
		initLogger()
	})
	return l
}

func initLogger() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			funcName := path.Base(frame.Function)
			arr := strings.Split(funcName, ".")
			if len(arr) > 1 {
				funcName = arr[len(arr)-1]
			}
			return funcName, fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
	})

	if len(os.Getenv("LOG_OUT")) > 0 {
		logger.SetOutput(os.Stdout)
	} else {
		logDir := "/var/log/coffeechat/"
		logName := "user.log"
		var filePerm os.FileMode = 0750
		if err := os.MkdirAll(logDir, filePerm); err != nil {
			panic(err)
		}

		logFile, err := os.OpenFile(path.Join(logDir, logName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, filePerm)
		if err != nil {
			panic(err)
		}
		logger.SetOutput(logFile)
	}

	l = logger
}
