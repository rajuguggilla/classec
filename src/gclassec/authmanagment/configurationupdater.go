package authmanagment

import(
	"net/http"
	"io/ioutil"
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"os"
	"gclassec/loggers"
)
const updatingFileName = "authmanagment/configurationupdater.go"
const awsFileName = "conf/awscred.json"
const azureFileName = "conf/azurecred.json"
const hosFileName = "conf/hosconfiguration.json"
const osFileName = "conf/computeVM.json"
const vmwareFileName = "conf/vmwareconf.json"


var logger = Loggers.New()

func MyFileWriter(data string, configFile string)(string){
	logger.Debug("requestBody:==",data)
	// Split on NewLine.
    	tempVariableString := strings.Split(data, "&")
    	 //Display all elements.
	logger.Info("TempVariable Length:",len(tempVariableString))
	for i:= range tempVariableString {
		logger.Info("\nTempVariable %d:%s", i,tempVariableString[i])
	}
	//f, err := os.Create(configFile)
	//var err = os.Remove(path)
	//checkError(err)
	f,err := os.OpenFile(configFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		logger.Error("Error in creating File:==", err)
		return("Failed")
	}
	defer f.Close()
	f.WriteString("{")
	for i:=0; i<len(tempVariableString); i++ {
		f.WriteString("\n\"")
		temp := (strings.Replace(tempVariableString[i],"=",  "\":\"", 1))
		f.WriteString(temp)
		if i==len(tempVariableString)-1 {
			f.WriteString("\"")
		}else{
			f.WriteString("\",")
		}
	}
	f.WriteString("\n}")
	logger.Debug("Ok Successful in MyFileWriter")
	return "Ok Sucessfull"
}


func AwsCredentials(w http.ResponseWriter, r *http.Request){
	var bodyBytes []byte
	if r.Body != nil {
  		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	//fmt.Println(bodyString)
	//filename := "authmanagment/configurationupdater.go"
	_, filePath, _, _ := runtime.Caller(0)
        logger.Debug("\nCurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, updatingFileName, awsFileName, 1))
        logger.Debug("\nABSPATH:==",ConfigFilePath)
	resp:= MyFileWriter(bodyString, ConfigFilePath)
	fmt.Fprintf(w,resp)
}


func AzureCredentials(w http.ResponseWriter, r *http.Request){
	var bodyBytes []byte
	if r.Body != nil {
  		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	//filename := "authmanagment/configurationupdater.go"
        _, filePath, _, _ := runtime.Caller(0)
        logger.Debug("CurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, updatingFileName, azureFileName, 1))
        logger.Debug("ABSPATH:==",ConfigFilePath)
	resp:=(MyFileWriter(bodyString, ConfigFilePath))
	fmt.Fprintf(w,resp)
}

func OpenstackCredentials(w http.ResponseWriter, r *http.Request){
	var bodyBytes []byte
	if r.Body != nil {
  		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	logger.Info(bodyString)
	//filename := "authmanagment/configurationupdater.go"
        _, filePath, _, _ := runtime.Caller(0)
        logger.Debug("CurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, updatingFileName, osFileName, 1))
        logger.Debug("ABSPATH:==",ConfigFilePath)
	resp:=(MyFileWriter(bodyString, ConfigFilePath))
	fmt.Fprintf(w,resp)
}


func VmwareCredentials(w http.ResponseWriter, r *http.Request){
	var bodyBytes []byte
	if r.Body != nil {
  		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	//filename := "authmanagment/configurationupdater.go"
        _, filePath, _, _ := runtime.Caller(0)
        logger.Debug("CurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, updatingFileName, vmwareFileName, 1))
        logger.Debug("ABSPATH:==",ConfigFilePath)
	resp:=(MyFileWriter(bodyString, ConfigFilePath))
	fmt.Fprintf(w,resp)
}



func HosCredentials(w http.ResponseWriter, r *http.Request){
	var bodyBytes []byte
	if r.Body != nil {
  		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	//filename := "authmanagment/configurationupdater.go"
        _, filePath, _, _ := runtime.Caller(0)
        logger.Debug("CurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, updatingFileName, hosFileName, 1))
        logger.Debug("ABSPATH:==",ConfigFilePath)
	resp:=(MyFileWriter(bodyString, ConfigFilePath))
	fmt.Fprintf(w,resp)
}