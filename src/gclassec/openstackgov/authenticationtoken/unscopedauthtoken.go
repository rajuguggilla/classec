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




func UnscopedAuthToken() (string, openstackInstance.OpenStackEndpoints, error){
	logger := Loggers.New()
	var usAuthToken string
	var usAuthError error
	var usApiEndpointsStruct openstackInstance.OpenStackEndpoints
	var filename string = "openstackgov/authenticationtoken/unscopedauthtoken.go"
	_, filePath, _,ok := runtime.Caller(0)
	if ok == false{
		logger.Error("Error in locating file.")
		return	usAuthToken, usApiEndpointsStruct, errors.New(errcode.ErrFileLocate)
	}
	logger.Info("CurrentFilePath:==", filePath)
	absPath := (strings.Replace(filePath, filename, "conf/openstackconfig.json", 1))
	logger.Info("OpenStackConfigurationFilePath:==", absPath)
	file, oerr := os.Open(absPath)
	if oerr != nil{
		if os.IsNotExist(oerr){
			logger.Error(errcode.ErrFileNotExist)
			return usAuthToken, usApiEndpointsStruct, oerr
		}else {
			logger.Error(errcode.ErrFileOpen)
			return usAuthToken, usApiEndpointsStruct, oerr
		}
	}
	tempConfig := new(openstackInstance.OpenStackUserConfig)
	decoder := json.NewDecoder(file)
	decErr := decoder.Decode(&tempConfig)
	if decErr != nil {
		logger.Error("Error in reading configuration:==", decErr)
		return usAuthToken, usApiEndpointsStruct, decErr
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
		//Un-scoped
		reqURL = tempConfig.IdentityEndpoint + "auth/tokens"
		reqBody = `{"auth":{"identity":{"methods": ["password"], "password": {"user":{"name": "` + tempConfig.UserName + `","domain": { "name": "` + tempConfig.Domain + `"}, "password": "` + tempConfig.Password + `"}}}}}`
		logger.Info("Request Body:==", reqBody)
		logger.Info("\nRequest URL:==", reqURL)
		usAuthToken ,usApiEndpointsStruct, usAuthError = GetOpenStackAuthToken_v3(reqURL, reqBody)
	}else if (strings.Contains(tempConfig.IdentityEndpoint, "v2.0")){
		//Un-Scoped
		reqURL = tempConfig.IdentityEndpoint + "/tokens"
		reqBody = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}}}`
		logger.Info("Request Body:==", reqBody)
		logger.Info("\nRequest URL:==", reqURL)
		usAuthToken ,usApiEndpointsStruct, usAuthError = GetOpenStackAuthToken_v2(reqURL, reqBody)
	}else{
		return usAuthToken, usApiEndpointsStruct, errors.New(errcode.ErrAuthEndpoint)
	}

	if usAuthError != nil {
		for i := 0; i < len(usApiEndpointsStruct.ApiEndpoints); i++ {
			usApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(usApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "controller", tempConfig.Controller, -1)
			usApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(usApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "monasca", tempConfig.Controller, -1)
		}
	}
	return  usAuthToken, usApiEndpointsStruct, usAuthError

}