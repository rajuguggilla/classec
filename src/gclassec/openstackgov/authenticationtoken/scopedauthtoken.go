package authenticationtoken

import (
	"fmt"
	"strings"
	"encoding/json"
	"runtime"
	"os"
	"gclassec/structs/openstackInstance"
)





func ScopedAuthToken() (string, openstackInstance.OpenStackEndpoints, string){
	var sAuthToken string
	var sAuthError string
	var sApiEndpointsStruct openstackInstance.OpenStackEndpoints
	var filename string = "openstackgov/authenticationtoken/scopedauthtoken.go"
	_, filePath, _,ok := runtime.Caller(0)
	if ok == false{
		fmt.Println("Error in locating file.")
		return	sAuthToken, sApiEndpointsStruct, "Error in locating file."
	}
	fmt.Println("CurrentFilePath:==", filePath)
	absPath := (strings.Replace(filePath, filename, "conf/openstackconfig.json", 1))
	fmt.Println("OpenStackConfigurationFilePath:==", absPath)
	file, oerr := os.Open(absPath)
	if oerr != nil{
		if os.IsNotExist(oerr){
			fmt.Println("File does not exist.")
			return sAuthToken, sApiEndpointsStruct, "File does not exist."
		}else {
			fmt.Println("Error in opening file.")
			return sAuthToken, sApiEndpointsStruct, "Error in opening file."
		}
	}
	tempConfig := new(openstackInstance.OpenStackUserConfig)
	decoder := json.NewDecoder(file)
	decErr := decoder.Decode(&tempConfig)
	if decErr != nil {
		fmt.Println("Error in reading configuration:==", decErr)
		return sAuthToken, sApiEndpointsStruct, "Error in reading configuration."
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
		//Scoped with Project ID
		reqURL = tempConfig.IdentityEndpoint + "auth/tokens"
		reqBody = `{"auth":{"identity":{"methods": ["password"], "password": {"user":{"name": "` + tempConfig.UserName + `","domain": { "name": "` + tempConfig.Domain + `"}, "password": "` + tempConfig.Password + `"}}},"scope" :{"project":{"id": "`+ tempConfig.ProjectId+`"}}}}`
		fmt.Println("Request Body:==", reqBody)
		fmt.Println("\nRequest URL:==", reqURL)
		sAuthToken ,sApiEndpointsStruct, sAuthError = GetOpenStackAuthToken_v3(reqURL, reqBody)
	}else if (strings.Contains(tempConfig.IdentityEndpoint, "v2.0")){
		//Scoped with Tenant Name
		reqURL = tempConfig.IdentityEndpoint + "/tokens"
		reqBody = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}, "tenantName": "`+ tempConfig.TenantName+`"}}`
		fmt.Println("Request Body:==", reqBody)
		fmt.Println("\nRequest URL:==", reqURL)
		sAuthToken ,sApiEndpointsStruct, sAuthError = GetOpenStackAuthToken_v2(reqURL, reqBody)
	}else{
		return sAuthToken, sApiEndpointsStruct, "Please provide a valid IdentityEndPoint of type v3 or v2.0"
	}

	if sAuthError != "" {
		for i := 0; i < len(sApiEndpointsStruct.ApiEndpoints); i++ {
			sApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(sApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "controller", tempConfig.Controller, -1)
			sApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(sApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "monasca", tempConfig.Controller, -1)
		}
	}
	return  sAuthToken, sApiEndpointsStruct, sAuthError

}