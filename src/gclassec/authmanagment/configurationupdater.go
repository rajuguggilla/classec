package authmanagment

import(
	"net/http"
	"io/ioutil"
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"os"
)

func MyFileWriter(data string, configFile string)(string){
	fmt.Println("requestBody:==",data)
	// Split on NewLine.
    	tempVariableString := strings.Split(data, "&")
    	 //Display all elements.
	fmt.Println("TempVariable Length:",len(tempVariableString))
	for i:= range tempVariableString {
		fmt.Printf("\nTempVariable %d:%s", i,tempVariableString[i])
	}
	f, err := os.Create(configFile)
	if err != nil {
		fmt.Println("Error in creating File:==", err)
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
	filename := "credentialeditor/configurationupdater.go"
	_, filePath, _, _ := runtime.Caller(0)
        fmt.Println("\nCurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, filename, "conf/tempawscred.json", 1))
        fmt.Println("\nABSPATH:==",ConfigFilePath)
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
	filename := "credentialeditor/configurationupdater.go"
        _, filePath, _, _ := runtime.Caller(0)
        fmt.Println("CurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, filename, "conf/tempazurecred.json", 1))
        fmt.Println("ABSPATH:==",ConfigFilePath)
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
	fmt.Println(bodyString)
	filename := "credentialeditor/configurationupdater.go"
        _, filePath, _, _ := runtime.Caller(0)
        fmt.Println("CurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, filename, "conf/temposcred.json", 1))
        fmt.Println("ABSPATH:==",ConfigFilePath)
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
	filename := "credentialeditor/configurationupdater.go"
        _, filePath, _, _ := runtime.Caller(0)
        fmt.Println("CurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, filename, "conf/tempvmwarecred.json", 1))
        fmt.Println("ABSPATH:==",ConfigFilePath)
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
	filename := "credentialeditor/configurationupdater.go"
        _, filePath, _, _ := runtime.Caller(0)
        fmt.Println("CurrentFilePath:==",filePath)
        ConfigFilePath :=(strings.Replace(filePath, filename, "conf/temphoscred.json", 1))
        fmt.Println("ABSPATH:==",ConfigFilePath)
	resp:=(MyFileWriter(bodyString, ConfigFilePath))
	fmt.Fprintf(w,resp)
}