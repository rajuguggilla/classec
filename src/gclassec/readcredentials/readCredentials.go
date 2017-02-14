package readazurecreds

import (
	"encoding/json"
	"runtime"
	"strings"
	"os"
	"gclassec/loggers"
)

type Configuration struct {
    ClientId    string
    ClientSecret   string
    SubscriptionId   string
    TenantId   string

}

func Configurtion() Configuration{
	logger := Loggers.New()
	filename := "readcredentials/readCredentials.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Debug("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/azurecred.json", 1))
       logger.Debug("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/azurecred.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		logger.Error("error:", err)
	}

	return configuration
}
