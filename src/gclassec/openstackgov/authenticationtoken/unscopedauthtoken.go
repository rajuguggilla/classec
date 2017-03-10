package authenticationtoken

import (
	"fmt"
	"strings"
	"encoding/json"
	"runtime"
	"os"
	"gclassec/structs/openstackInstance"
)




func UnscopedAuthToken() (string, openstackInstance.OpenStackEndpoints, string){
	var usAuthToken string
	var usAuthError string
	var usApiEndpointsStruct openstackInstance.OpenStackEndpoints
	var filename string = "openstackgov/authenticationtoken/unscopedauthtoken.go"
	_, filePath, _,ok := runtime.Caller(0)
	if ok == false{
		fmt.Println("Error in locating file.")
		return	usAuthToken, usApiEndpointsStruct, "Error in locating file."
	}
	fmt.Println("CurrentFilePath:==", filePath)
	absPath := (strings.Replace(filePath, filename, "conf/openstackconfig.json", 1))
	fmt.Println("OpenStackConfigurationFilePath:==", absPath)
	file, oerr := os.Open(absPath)
	if oerr != nil{
		if os.IsNotExist(oerr){
			fmt.Println("File does not exist.")
			return usAuthToken, usApiEndpointsStruct, "File does not exist."
		}else {
			fmt.Println("Error in opening file.")
			return usAuthToken, usApiEndpointsStruct, "Error in opening file."
		}
	}
	tempConfig := new(openstackInstance.OpenStackUserConfig)
	decoder := json.NewDecoder(file)
	decErr := decoder.Decode(&tempConfig)
	if decErr != nil {
		fmt.Println("Error in reading configuration:==", decErr)
		return usAuthToken, usApiEndpointsStruct, "Error in reading configuration."
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

	var reqBody, reqURL string
	tempConfig.IdentityEndpoint = (strings.Replace(tempConfig.IdentityEndpoint, "controller", tempConfig.Controller, 1))
	if (strings.Contains(tempConfig.IdentityEndpoint, "v3")){
		//Un-scoped
		reqURL = tempConfig.IdentityEndpoint + "auth/tokens"
		reqBody = `{"auth":{"identity":{"methods": ["password"], "password": {"user":{"name": "` + tempConfig.UserName + `","domain": { "name": "` + tempConfig.Domain + `"}, "password": "` + tempConfig.Password + `"}}}}}`
		fmt.Println("Request Body:==", reqBody)
		fmt.Println("\nRequest URL:==", reqURL)
		usAuthToken ,usApiEndpointsStruct, usAuthError = GetOpenStackAuthToken_v3(reqURL, reqBody)
	}else if (strings.Contains(tempConfig.IdentityEndpoint, "v2.0")){
		//Un-Scoped
		reqURL = tempConfig.IdentityEndpoint + "/tokens"
		reqBody = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}}}`
		fmt.Println("Request Body:==", reqBody)
		fmt.Println("\nRequest URL:==", reqURL)
		usAuthToken ,usApiEndpointsStruct, usAuthError = GetOpenStackAuthToken_v2(reqURL, reqBody)
	}else{
		return usAuthToken, usApiEndpointsStruct, "Please provide a valid IdentityEndPoint of type v3 or v2.0"
	}

	if usAuthError != "" {
		for i := 0; i < len(usApiEndpointsStruct.ApiEndpoints); i++ {
			usApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(usApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "controller", tempConfig.Controller, -1)
			usApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(usApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "monasca", tempConfig.Controller, -1)
		}
	}
	return  usAuthToken, usApiEndpointsStruct, usAuthError




}