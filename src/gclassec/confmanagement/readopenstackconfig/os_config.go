package readopenstackconfig

import (
	"runtime"
	"strings"
	"os"
	"encoding/json"
	"fmt"
	"gclassec/structs/openstackInstance"
)



func OpenStackConfigReader() (openstackInstance.OpenStackUserConfig){
	var tempConfig openstackInstance.OpenStackUserConfig
	var filename string = "confmanagement/readopenstackconfig/os_config.go"
	_, filePath, _,ok := runtime.Caller(0)
	if ok == false{
		fmt.Println("Error in locating file.")

	}
	fmt.Println("CurrentFilePath:==", filePath)
	absPath := (strings.Replace(filePath, filename, "conf/openstackconfig.json", 1))
	fmt.Println("OpenStackConfigurationFilePath:==", absPath)
	file, oerr := os.Open(absPath)
	if oerr != nil{
		if os.IsNotExist(oerr){
			fmt.Println("File does not exist.")
			//return sAuthToken, sApiEndpointsStruct, "File does not exist."
		}else {
			fmt.Println("Error in opening file.")
			//return sAuthToken, sApiEndpointsStruct, "Error in opening file."
		}
	}

	decoder := json.NewDecoder(file)
	decErr := decoder.Decode(&tempConfig)
	if decErr != nil {
		fmt.Println("Error in reading configuration:==", decErr)
		//return sAuthToken, sApiEndpointsStruct, "Error in reading configuration."
	}
	fmt.Println("TempConfig:===")
	fmt.Println("IdentityEndPoint: ", tempConfig.IdentityEndpoint)
	fmt.Println("Container: ", tempConfig.Container)
	fmt.Println("Password: ", tempConfig.Password)
	fmt.Println("Tenanat_id: ", tempConfig.TenantId)
	fmt.Println("TenantName: ", tempConfig.TenantName)
	fmt.Println("Project_id: ", tempConfig.ProjectId)
	fmt.Println("ProjectName: ", tempConfig.ProjectName)
	fmt.Println("Region: ", tempConfig.Region)
	fmt.Println("UserName: ", tempConfig.UserName)
	fmt.Println("Domain: ", tempConfig.Domain)
	fmt.Println("Controller: ", tempConfig.Controller)
	return  tempConfig

}
