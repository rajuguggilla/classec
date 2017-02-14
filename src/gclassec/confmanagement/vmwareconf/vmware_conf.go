package vmwareconf

import (
	"runtime"
	"strings"
	"os"
	"encoding/json"
	"gclassec/loggers"
)

type Vmware struct {
   EnvURL 	 string
   EnvUserName 	 string
   EnvPassword 	 string
   EnvInsecure 	 string
}

func Configurtion() Vmware{
	logger := Loggers.New()
	filename := "confmanagement/vmwareconf/vmware_conf.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Debug("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/vmwareconf.json", 1))
       logger.Debug("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/osazureconf.json")
	decoder := json.NewDecoder(file)
	configuration := Vmware{}
	err := decoder.Decode(&configuration)
	if err != nil {
		logger.Error("Error:", err)
	}
	return configuration
}






