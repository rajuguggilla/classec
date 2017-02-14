package Loggers

import (
	"log"
	"os"
	"runtime"
	"strings"
)

//MyLogger custom logger
type MyLogger struct {
	ErrorL *log.Logger
	InfoL  *log.Logger
	DebugL *log.Logger
	WarnL	*log.Logger
}

//New returns logger
//TODO : deal with error
func New() *MyLogger {
	filename := "newlogs.go"
	_, filePath, _, _ := runtime.Caller(0)
	//fmt.Println("CurrentFilePath:==",filePath)
	ErrorFilePath :=(strings.Replace(filePath, filename, "log_error.txt", 1))
	//fmt.Println("Error File Path:==",ErrorFilePath)
	InfoFilePath :=(strings.Replace(filePath, filename, "log_info.txt", 1))
	//fmt.Println("Info File Path:==",InfoFilePath)
	DebugFilePath :=(strings.Replace(filePath, filename, "log_debug.txt", 1))
	//fmt.Println("Debug File Path:==",DebugFilePath)
	WarnFilePath :=(strings.Replace(filePath, filename, "log_warn.txt", 1))
	//fmt.Println("Warning File Path:==",WarnFilePath)
	ml := new(MyLogger)
	if logFile, err := os.OpenFile(ErrorFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err == nil {
		log.SetOutput(logFile)
		ml.ErrorL = log.New(logFile, "ERROR ", log.LUTC)
	}
	if logFile, err := os.OpenFile(InfoFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err == nil {
		log.SetOutput(logFile)
		ml.InfoL = log.New(logFile, "INFO ", log.LUTC)
	}
	if logFile, err := os.OpenFile(DebugFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err == nil {
		log.SetOutput(logFile)
		ml.DebugL = log.New(logFile, "DEBUG ", log.LUTC)
	}
	if logFile, err := os.OpenFile(WarnFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); err == nil {
		log.SetOutput(logFile)
		ml.WarnL = log.New(logFile, "WARN ", log.LUTC)
	}
	return ml
}

func (ml MyLogger) Error(data ...interface{}) {
	ml.ErrorL.Println(data...)
}

func (ml MyLogger) Info(data ...interface{}) {
	ml.InfoL.Println(data...)
}

func (ml MyLogger) Debug(data ...interface{}) {
	ml.DebugL.Println(data...)
}

func (ml MyLogger) Warn(data ...interface{}) {
	ml.WarnL.Println(data...)
}
