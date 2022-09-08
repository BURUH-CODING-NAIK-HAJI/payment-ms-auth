package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

func getRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(b)

	logPath := fmt.Sprintf(currentPath+"/%s", "../../log/error.log")
	return logPath
}

func CreateErrorLogger() *logrus.Logger {
	path := getRootPath()
	errorLogFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Server Shutdown, Log File Not Found")
		os.Exit(0)
	}

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetOutput(errorLogFile)

	return log
}
