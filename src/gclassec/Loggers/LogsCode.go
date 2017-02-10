package Loggers

import (
	"io"
	"os"
	"log"
	"runtime"
	"fmt"
	"strings"
)

var logCloser io.Closer
func MyLogger() {
	filename := "LogsCode.go"
       _, filePath, _, _ := runtime.Caller(0)
       fmt.Println("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "log.txt", 1))
       fmt.Println("ABSPATH:==",ConfigFilePath)

	logFile, err := os.OpenFile(ConfigFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	logCloser = logFile
	log.SetOutput(logFile)
	//log.Set
}
func CloseMyLogger() {
	logCloser.Close()
}
