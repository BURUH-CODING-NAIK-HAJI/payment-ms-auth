package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/sirupsen/logrus"
)

func getRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(b)
	regex, _ := regexp.Compile("golang-api-template.*")
	rootPath := regex.ReplaceAll([]byte(currentPath), []byte("golang-api-template/log/error.log"))
	return string(rootPath)
}

func CreateErrorLogger() *logrus.Logger {
	path := getRootPath()
	fmt.Println(path)
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
