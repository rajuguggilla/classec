package authtoken

import (
	"strings"
	"runtime"
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"gclassec/loggers"
	"gclassec/structs/hosstruct"
	"fmt"
	"gclassec/errorcodes/errcode"

)


var logger = Loggers.New()

func GetHOSAuthToken() (string, hosstruct.HOSAutToken, error){
//func main(){
	var filename string = "goclienthos/authtoken/getAuthToken.go"
	_, filePath, _, _ := runtime.Caller(0)
	logger.Debug("CurrentFilePath:==",filePath)
	absPath :=(strings.Replace(filePath, filename, "conf/hosconfiguration.json", 1))
	//absPath :=(strings.Replace(filePath, filename, "openStackConfiguration.json", 1))
	logger.Debug("HOSConfigurationFilePath:==",absPath)
	file, errOpen := os.Open(absPath)

	if errOpen != nil{
		fmt.Println("Error : ", errcode.ErrFileOpen)
		logger.Error("Error : ", errcode.ErrFileOpen)
		return "", hosstruct.HOSAutToken{}, errOpen
	}

	decoder := json.NewDecoder(file)
	tempConfig := hosstruct.Configuration{}
	err := decoder.Decode(&tempConfig)
	if err != nil{
		logger.Error("ConfigurationError:", errcode.ErrDecode)
	}

	logger.Info("TempConfig:===")
	logger.Info("IdentityEndPoint: ",tempConfig.IdentityEndpoint)
    	logger.Info("Container: ",tempConfig.Container)
    	logger.Info("Password: ",tempConfig.Password)
	logger.Info("Tenanat_id: ",tempConfig.TenantId)
    	logger.Info("TenantName: ",tempConfig.TenantName)
	logger.Info("Project_id: ",tempConfig.ProjectId)
    	logger.Info("ProjectName: ",tempConfig.ProjectName)
	logger.Info("Region: ",tempConfig.Region)
	logger.Info("UserName: ",tempConfig.UserName)

	var reqBody string = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}, "tenantName": "`+ tempConfig.TenantName+`"}}`
	//var reqBody string = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}, "tenantId": "`+ tempConfig.TenantId +`", "tenantName": "`+ tempConfig.TenantName+`", "Container": "`+ tempConfig.Container +`","ImageRegion": "`+ tempConfig.Region +`"}}`
	logger.Info("Request Body:==",reqBody)

	var reqURL string = tempConfig.IdentityEndpoint + "/tokens"
	logger.Info("\nRequest URL:==",reqURL)

	req, err := http.NewRequest("POST", reqURL, strings.NewReader(reqBody))
	if err != nil{
			return "", hosstruct.HOSAutToken{}, err
		}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	logger.Info("Printing request:==",req)
	res, err := http.DefaultClient.Do(req)


		if err != nil{
			return "", hosstruct.HOSAutToken{}, err
		}

//	logger.Info("Status:==", res.Status)
	defer res.Body.Close()
	respBody, err := ioutil.ReadAll(res.Body)

	var emptyStruct = hosstruct.HOSAutToken{}
		if err != nil{
			return "", emptyStruct, err
		}

	//fmt.Print("In GET HOS AUTH TOKEN respBody:==",respBody)

	respBodyInString:= string(respBody)
	logger.Info("\n\n\nIn GET HOS AUTH TOKEN respBodyInString:==\n\n",respBodyInString)
	logger.Info("\n\n\n")
	//rBodyInByte := []byte(respBody)
	//fmt.Println("rBodyInByte",rBodyInByte)

	var jsonAuthTokenBody hosstruct.HOSAutToken

	//respMarshed,_ := json.Marshal(rBodyInByte)
	//fmt.Println("marshedRespBody:===",respMarshed)
	//stringRespMarshed:=string(respMarshed)
	//fmt.Println("marshedBody in string", stringRespMarshed)
	if err = json.Unmarshal(respBody, &jsonAuthTokenBody); err != nil{
		logger.Error("Error in unmarshing:==",err)
	}

	//newDecoder := json.NewDecoder(respBody)
	//newTempConfig := Endpoint{}
	//error := newDecoder.Decode(&newTempConfig)
	//if error != nil{
	//	fmt.Println("ConfigurationError:", error)
	//}
	//
	logger.Info("\nIn GET HOS AUTH TOKEN HOSResponseBody:===\n", jsonAuthTokenBody)
	logger.Info("\nIn GET HOS AUTH TOKEN jsonAuthTokenBody:===\n %+v\n\n", jsonAuthTokenBody)
	logger.Info("\nIn GET HOS AUTH TOKEN AuthToken:==\n",jsonAuthTokenBody.Access.Token.AuthToken)
	return  jsonAuthTokenBody.Access.Token.AuthToken, jsonAuthTokenBody, nil

}
