package readstructconf

import (
	"runtime"
	"fmt"
	"strings"
	"os"
	"gclassec/errorcodes/errcode"
	"encoding/json"
	"gclassec/structs/configurationstruct"
	"gclassec/loggers"
)

var logger = Loggers.New()

func ReadStructConfigFile() int64{
    filename := "confmanagement/readstructconf/struct_conf.go"
    _, filePath, _, _ := runtime.Caller(0)
    //logger.Debug("CurrentFilePath:==",filePath)
	fmt.Println("CurrentFilePath:==",filePath)
    ConfigFilePath :=(strings.Replace(filePath, filename, "conf/jobconf.json", 1))
    //logger.Debug("ABSPATH:==",ConfigFilePath)
	fmt.Println("ABSPATH:==",ConfigFilePath)
    file, errOpen := os.Open(ConfigFilePath)
    if errOpen != nil{
        fmt.Println("Error : ", errcode.ErrFileOpen)
        logger.Error("Error : ", errcode.ErrFileOpen)
	return 0
    }
    decoder := json.NewDecoder(file)
    configuration := configurationstruct.Configuration{}
    errDecode := decoder.Decode(&configuration)
    if errDecode != nil {
        fmt.Println("Error : ", errcode.ErrDecode)
        logger.Error("Error : ",errcode.ErrDecode)
        return 0
    }
    return configuration.StandardizedStruct
}