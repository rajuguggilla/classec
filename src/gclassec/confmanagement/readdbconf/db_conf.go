package readdbconf
import (
	"encoding/json"
	"runtime"
	"strings"
	"os"
	"gclassec/loggers"

	"gclassec/errorcodes/errcode"
)

type Configuration struct {
    Dbtype    string
    Dbname   string
    Dbusername   string
    Dbpassword   string
    Dbhostname   string
    Dbport   string
    Dbnameforaws string
}

func Configurtion() Configuration{
	logger := Loggers.New()
	filename := "confmanagement/readdbconf/db_conf.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Debug("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/dbconf.json", 1))
       logger.Debug("ABSPATH:==",ConfigFilePath)
	file, err1:= os.Open(ConfigFilePath)
	if err1 != nil {
			println(errcode.CLAERR0001)
			}
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/awsconf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		logger.Error("error:", err)
		println(errcode.CLAERR0001)
	}
	return configuration
}