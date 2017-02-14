package readazureconf


import (
	"encoding/json"
	"runtime"
	"strings"
	"os"
	"gclassec/loggers"
)

type Configuration struct {
    Dbtype    string
    Dbname   string
	Dbusername   string
    Dbpassword   string
	Dbhostname   string
	Dbport   string
}

func Configurtion() Configuration{
	logger := Loggers.New()
	filename := "confmanagement/readazureconf/azure_conf.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Debug("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/osazureconf.json", 1))
       logger.Debug("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/osazureconf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		logger.Error("error:", err)
	}

	return configuration
}

