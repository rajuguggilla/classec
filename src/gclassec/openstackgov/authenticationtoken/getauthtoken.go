package AuthenticationToken

import (
	"fmt"
	"strings"
	"encoding/json"
	"runtime"
	"os"
)

//------------------------------------------------------Structure to List All Api Endpoints --------------------------------------------------------//
type OpenStackEndpoints struct {
 	ApiEndpoints 	[]EndpointStruct
}

type EndpointStruct struct {
	EndpointName	string
	EndpointURL	string
	EndpointType	string
}

//------------------------------------------------------Structure to read Configuration file --------------------------------------------------------//
type OpenStackUserConfig struct {
	IdentityEndpoint	string	`json:"identityEndpoint"`
    	UserName		string	`json:"userName"`
	Password		string	`json:"password"`
	Domain			string	`json:"domain"`
    	TenantName 		string	`json:"tenantName"`
    	TenantId 		string	`json:"tenantID"`
	ProjectId		string	`json:"projectID"`
	ProjectName		string	`json:"projectName"`
    	Container 		string	`json:"container"`
    	Region	 		string	`json:"region"`
	Controller 		string	`json:"controller"`
}



func GetAuthToken() (string, OpenStackEndpoints, string){
	var authToken string
	var authError string
	var ApiEndpointsStruct OpenStackEndpoints
	var filename string = "openstackgov/authenticationtoken/getauthtoken.go"
	_, filePath, _,ok := runtime.Caller(0)
	if ok == false{
		fmt.Println("Error in locating file.")
		return	authToken, ApiEndpointsStruct, "Error in locating file."
	}
	fmt.Println("CurrentFilePath:==", filePath)
	absPath := (strings.Replace(filePath, filename, "conf/openStackConfig.json", 1))
	fmt.Println("OpenStackConfigurationFilePath:==", absPath)
	file, oerr := os.Open(absPath)
	if oerr != nil{
		if os.IsNotExist(oerr){
			fmt.Println("File does not exist.")
			return authToken, ApiEndpointsStruct, "File does not exist."
		}else {
			fmt.Println("Error in opening file.")
			return authToken, ApiEndpointsStruct, "Error in opening file."
		}
	}
	tempConfig := new(OpenStackUserConfig)
	decoder := json.NewDecoder(file)
	decErr := decoder.Decode(&tempConfig)
	if decErr != nil {
		fmt.Println("Error in reading configuration:==", decErr)
		return authToken, ApiEndpointsStruct, "Error in reading configuration."
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
		reqURL = tempConfig.IdentityEndpoint + "auth/tokens"
		reqBody = `{"auth":{"identity":{"methods": ["password"], "password": {"user":{"name": "` + tempConfig.UserName + `","domain": { "name": "` + tempConfig.Domain + `"}, "password": "` + tempConfig.Password + `"}}}}}`
		fmt.Println("Request Body:==", reqBody)
		fmt.Println("\nRequest URL:==", reqURL)
		authToken ,ApiEndpointsStruct, authError = GetOpenStackAuthToken_v3(reqURL, reqBody)
	}else if (strings.Contains(tempConfig.IdentityEndpoint, "v2.0")){
		reqURL = tempConfig.IdentityEndpoint + "/tokens"
		reqBody = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}, "tenantName": "`+ tempConfig.TenantName+`"}}`
		fmt.Println("Request Body:==", reqBody)
		fmt.Println("\nRequest URL:==", reqURL)
		authToken ,ApiEndpointsStruct, authError = GetOpenStackAuthToken_v2(reqURL, reqBody)
	}else{
		return authToken, ApiEndpointsStruct, "Please provide a valid IdentityEndPoint of type v3 or v2.0"
	}

	if authError != "" {
		for i := 0; i < len(ApiEndpointsStruct.ApiEndpoints); i++ {
			ApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(ApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "controller", tempConfig.Controller, -1)
			ApiEndpointsStruct.ApiEndpoints[i].EndpointURL = strings.Replace(ApiEndpointsStruct.ApiEndpoints[i].EndpointURL, "monasca", tempConfig.Controller, -1)
		}
	}
	return  authToken, ApiEndpointsStruct, authError




}