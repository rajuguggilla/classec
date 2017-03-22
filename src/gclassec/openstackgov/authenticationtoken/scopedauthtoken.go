package authenticationtoken

import (

	"strings"
	"encoding/json"
	"runtime"
	"os"
	"gclassec/structs/openstackInstance"
	"gclassec/errorcodes/errcode"
	"errors"
	"gclassec/loggers"
)





func ScopedAuthToken() (string, openstackInstance.OpenStackEndpoints, error){
	logger := Loggers.New()
	var sAuthToken string
	var sAuthError error
	var sApiEndpointsStruct openstackInstance.OpenStackEndpoints
	var filename string = "openstackgov/authenticationtoken/scopedauthtoken.go"
	_, filePath, _,ok := runtime.Caller(0)
	if ok == false{
		//logger.Info("Error in locating file.",errcode.ErrFileLocate)
		logger.Error("Error in locating file.",errcode.ErrFileLocate)
		return	sAuthToken, sApiEndpointsStruct, errors.New(errcode.ErrFileLocate)
	}
	logger.Info("CurrentFilePath:==", filePath)
	absPath := (strings.Replace(filePath, filename, "conf/openstackconfig.json", 1))
	logger.Info("OpenStackConfigurationFilePath:==", absPath)
	file, oerr := os.Open(absPath)
	if oerr != nil{
		if os.IsNotExist(oerr){
			//logger.Info("File does not exist.",errcode.ErrFileNotExist)
			logger.Error("File does not exist.",errcode.ErrFileNotExist)
			return sAuthToken, sApiEndpointsStruct, oerr
		}else {
			//logger.Info("Error in opening file.",errcode.ErrFileOpen)
			logger.Error("Error in opening file.",errcode.ErrFileOpen)
			return sAuthToken, sApiEndpointsStruct, oerr
		}
	}
	tempConfig := new(openstackInstance.OpenStackUserConfig)
	decoder := json.NewDecoder(file)
	decErr := decoder.Decode(&tempConfig)
	if decErr != nil {
		//logger.Info("Error in reading configuration:==",decErr)
		logger.Error("Error in reading configuration:==",decErr)
		return sAuthToken, sApiEndpointsStruct, decErr
	}
	logger.Info("TempConfig:===")
	logger.Info("IdentityEndPoint: ", tempConfig.IdentityEndpoint)
	logger.Info("Container: ", tempConfig.Container)
	logger.Info("Password: ", tempConfig.Password)
	logger.Info("Tenanat_id: ", tempConfig.TenantId)
	logger.Info("TenantName: ", tempConfig.TenantName)
	logger.Info("Project_id: ", tempConfig.ProjectId)
	logger.Info("ProjectName: ", tempConfig.ProjectName)
	logger.Info("Region: ", tempConfig.Region)
	logger.Info("UserName: ", tempConfig.UserName)
	logger.Info("Domain: ", tempConfig.Domain)
	logger.Info("Controller: ", tempConfig.Controller)

	var reqBody, reqURL string
	tempConfig.IdentityEndpoint = (strings.Replace(tempConfig.IdentityEndpoint, "controller", tempConfig.Controller, 1))
	if (strings.Contains(tempConfig.IdentityEndpoint, "v3")){
		//Scoped with Project ID
		reqURL = tempConfig.IdentityEndpoint + "auth/tokens"
		reqBody = `{"auth":{"identity":{"methods": ["password"], "password": {"user":{"name": "` + tempConfig.UserName + `","domain": { "name": "` + tempConfig.Domain + `"}, "password": "` + tempConfig.Password + `"}}},"scope" :{"project":{"id": "`+ tempConfig.ProjectId+`"}}}}`
		logger.Info("Request Body:==", reqBody)
		logger.Info("\nRequest URL:==", reqURL)
		sAuthToken ,sApiEndpointsStruct, sAuthError = GetOpenStackAuthToken_v3(reqURL, reqBody)
	}else if (strings.Contains(tempConfig.IdentityEndpoint, "v2.0")){
		//Scoped with Tenant Name
		reqURL = tempConfig.IdentityEndpoint + "/tokens"
		reqBody = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}, "tenantName": "`+ tempConfig.TenantName+`"}}`
		logger.Info("Request Body:==", reqBody)
		logger.Info("\nRequest URL:==", reqURL)
		sAuthToken ,sApiEndpointsStruct, sAuthError = GetOpenStackAuthToken_v2(reqURL, reqBody)
	}else{
		return sAuthToken, sApiEndpointsStruct, errors.New(errcode.ErrAuthEndpoint)
	}
	//logger.Info("Len Of sAuthError:===",len(sAuthError))
	//logger.Info("Len Of sApiEndpointsStruct:===",len(sApiEndpointsStruct.ApiEndpoints))
	if sAuthError == nil{
		for i := 0; i< len(sApiEndpointsStruct.ApiEndpoints); i++ {
			//logger.Info("Before Replacing sApiEndpointsStruct.ApiEndpoints[i].EndpointURL:====",sApiEndpointsStruct.ApiEndpoints[i].EndpointURL)
			sApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(sApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "controller", tempConfig.Controller, -1)
			sApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(sApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "monasca", tempConfig.Controller, -1)
			//logger.Info("After Replacing sApiEndpointsStruct.ApiEndpoints[i].EndpointURL:====",sApiEndpointsStruct.ApiEndpoints[i].EndpointURL)
		}
	}
	//logger.Info("sAuthToken:====", sAuthToken)
	//logger.Info("sAuthError:====",sAuthError)
	//logger.Info("sApiEndpointsStruct:====",sApiEndpointsStruct)
	//logger.Info("returning form ScopedAuthToken()")
	return  sAuthToken, sApiEndpointsStruct, sAuthError

}