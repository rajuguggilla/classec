package authmanagment

import(
	"fmt"
	"runtime"
	"strings"
	"os"
	"encoding/json"
	"gclassec/structs/hosstruct"
	"gclassec/structs/openstackInstance"
	"gclassec/structs/vmwarestructs"
	"gclassec/structs/azurestruct"
	"gclassec/structs/awsstructs"
)


const readerFileName = "authmanagment/configurationreader.go"
//const awsFileName = "conf/awscred.json"
//const azureFileName = "conf/azurecred.json"
//const hosFileName = "conf/hosconfiguration.json"
//const osFileName = "conf/computeVM.json"
//const vmwareFileName = "conf/vmwareconf.json"



func ReadAwsCredentials() (awsstructs.Configuration){

	var tempConfig = awsstructs.Configuration{}
	_, filePath, _,ok := runtime.Caller(0)
	if ok == false{
		fmt.Println("Error in locating file.")
		return	tempConfig
	}
	fmt.Println("CurrentFilePath:==",filePath)
	absPath :=(strings.Replace(filePath, readerFileName, awsFileName, 1))
	fmt.Println("HOSConfigurationFilePath:==",absPath)
	file, oerr := os.Open(absPath)
	if oerr != nil{
		if os.IsNotExist(oerr){
			fmt.Println("File does not exist.")
			return tempConfig
		}else {
			fmt.Println("Error in opening file.")
			return tempConfig
		}
	}
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&tempConfig)
	if err != nil{
		logger.Error("ConfigurationError:", err)
	}

	fmt.Println("TempConfig:===")
	fmt.Println("accessKeyID: ",tempConfig.AccessKeyID)
    	fmt.Println("accessKeyValue: ",tempConfig.AccessKeyValue)
	//tempConfig.AccessKeyValue="XXXXXXXXXX"
	return tempConfig




}


func ReadAzureCredentials() azurestruct.Configuration {
	_, filePath, _,_ := runtime.Caller(0)
	fmt.Println("CurrentFilePath:==",filePath)
	absPath :=(strings.Replace(filePath, readerFileName, azureFileName, 1))
	fmt.Println("AzureConfigurationFilePath:==",absPath)
	file, _ := os.Open(absPath)
	decoder := json.NewDecoder(file)
	tempConfig := azurestruct.Configuration{}
	err := decoder.Decode(&tempConfig)
	if err != nil{
		logger.Error("ConfigurationError:", err)
	}

	fmt.Println("TempConfig:===")
	fmt.Println("Clientid: ",tempConfig.Clientid)
    	fmt.Println("Clientsecret: ",tempConfig.Clientsecret)
    	fmt.Println("Subscriptionid: ",tempConfig.Subscriptionid)
	fmt.Println("Tenanatid: ",tempConfig.Tenantid)
	tempConfig.Clientsecret="XXXXXXXXXX"
	return tempConfig
}

func ReadOpenstackCredentials() openstackInstance.Configuration{
	_, filePath, _, _ := runtime.Caller(0)
	fmt.Println("CurrentFilePath:==",filePath)
	absPath :=(strings.Replace(filePath, readerFileName, osFileName, 1))
	fmt.Println("HOSConfigurationFilePath:==",absPath)
	file, _ := os.Open(absPath)
	decoder := json.NewDecoder(file)
	tempConfig := openstackInstance.Configuration{}
	err := decoder.Decode(&tempConfig)
	if err != nil{
		logger.Error("ConfigurationError:", err)
	}

	fmt.Println("TempConfig:===")
	fmt.Println("IdentityEndPoint: ",tempConfig.Host)
    	fmt.Println("Container: ",tempConfig.Container)
    	fmt.Println("Password: ",tempConfig.Password)
	fmt.Println("Tenanat_id: ",tempConfig.ProjectID)
    	fmt.Println("TenantName: ",tempConfig.ProjectName)
	fmt.Println("Project_id: ",tempConfig.ProjectID)
    	fmt.Println("ProjectName: ",tempConfig.ProjectName)
	fmt.Println("Controller: ",tempConfig.Controller)
	fmt.Println("UserName: ",tempConfig.Username)
	tempConfig.Password="XXXXXXXXXX"
	return tempConfig
}


func ReadVmwareCredentials() vmwarestructs.Configuration{
	_, filePath, _, _ := runtime.Caller(0)
	fmt.Println("CurrentFilePath:==",filePath)
	absPath :=(strings.Replace(filePath, readerFileName, vmwareFileName, 1))
	fmt.Println("HOSConfigurationFilePath:==",absPath)
	file, _ := os.Open(absPath)
	decoder := json.NewDecoder(file)
	tempConfig := vmwarestructs.Configuration{}
	err := decoder.Decode(&tempConfig)
	if err != nil{
		logger.Error("ConfigurationError:", err)
	}

	fmt.Println("TempConfig:===")
	fmt.Println("IdentityEndPoint: ",tempConfig.EnvURL)
    	fmt.Println("UserName: ",tempConfig.EnvUserName)
    	fmt.Println("Password: ",tempConfig.EnvPassword)
	fmt.Println("Insecure: ",tempConfig.EnvInsecure)
	tempConfig.EnvPassword="XXXXXXXXXX"
    	return tempConfig

}



func ReadHosCredentials() hosstruct.Configuration {
	_, filePath, _, _ := runtime.Caller(0)
	fmt.Println("CurrentFilePath:==",filePath)
	absPath :=(strings.Replace(filePath, readerFileName, hosFileName, 1))
	fmt.Println("HOSConfigurationFilePath:==",absPath)
	file, _ := os.Open(absPath)
	decoder := json.NewDecoder(file)
	tempConfig := hosstruct.Configuration{}
	err := decoder.Decode(&tempConfig)
	if err != nil{
		logger.Error("ConfigurationError:", err)
	}

	fmt.Println("TempConfig:===")
	fmt.Println("IdentityEndPoint: ",tempConfig.IdentityEndpoint)
    	fmt.Println("Container: ",tempConfig.Container)
    	fmt.Println("Password: ",tempConfig.Password)
	fmt.Println("Tenanat_id: ",tempConfig.TenantId)
    	fmt.Println("TenantName: ",tempConfig.TenantName)
	fmt.Println("Project_id: ",tempConfig.ProjectId)
    	fmt.Println("ProjectName: ",tempConfig.ProjectName)
	fmt.Println("Region: ",tempConfig.Region)
	fmt.Println("UserName: ",tempConfig.UserName)
	tempConfig.Password="XXXXXXXXXX"
	return tempConfig
}
