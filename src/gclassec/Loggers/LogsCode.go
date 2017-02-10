package Loggers

import (
	"io"
	"os"
	"log"
)

var logCloser io.Closer
func MyLogger() {
	logFile, err := os.OpenFile("C:/Git/goclassec/src/gclassec/Logs/log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
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
