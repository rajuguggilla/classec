package readhosconf

import (
	"gclassec/loggers"
	"runtime"
	"strings"
	"os"
	"encoding/json"
)

type Configuration struct {
    IdentityEndpoint    string
    UserName   string
    Password   string
    ProjectID   string
    ProjectName   string
    Container   string
    Region string
    TenantName string
    TenantID string

}

func Configurtion() Configuration{
	logger := Loggers.New()
	filename := "confmanagement/readhosconf/hosconf.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Debug("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/hosconfiguration.json", 1))
       logger.Debug("ABSPATH:==",ConfigFilePath)
	file, err1:= os.Open(ConfigFilePath)
	if err1 != nil {
			println(" conf File is not present")
			}
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/awsconf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		logger.Error("error:", err)

	}
	return configuration

}

