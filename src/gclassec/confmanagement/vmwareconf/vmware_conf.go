package vmwareconf

import (
	"runtime"
	"fmt"
	"strings"
	"os"
	"encoding/json"
)

type Vmware struct {
   EnvURL 	 string
   EnvUserName 	 string
   EnvPassword 	 string
   EnvInsecure 	 string
}

func Configurtion() Vmware{
	filename := "confmanagement/vmwareconf/vmware_conf.go"
       _, filePath, _, _ := runtime.Caller(0)
       fmt.Println("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/vmwareconf.json", 1))
       fmt.Println("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/osazureconf.json")
	decoder := json.NewDecoder(file)
	configuration := Vmware{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	return configuration
}






