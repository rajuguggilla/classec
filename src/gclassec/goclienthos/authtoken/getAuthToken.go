package authtoken

import (
	"strings"
	"fmt"
	"runtime"
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"
)


type HOSAutToken struct{
	Access 	AccessStruct	`json:"access"`
}


type  AccessStruct struct {
	Token  		TokenStruct		`json:"token"`
	ServiceCatalog	[]ServiceCatalogStruct	`json:"serviceCatalog"`
	User		UserStruct		`json:"user"`
	Metadata	Metadata		`json:"metadata"`

}

type TokenStruct struct{
	Issued_at	string		`json:"issued_at"`
	Expires		string		`json:"expires"`
	AuthToken	string		`json:"id"`
	Tenant		TenantStruct	`json:"tenant"`
	Audit_ids	[]string	`json:"audit_ids"`
}
type TenantStruct struct{
	Description	string		`json:"description"`
	Enabled		bool		`json:"enabled"`
	TenanatID	string		`json:"id"`
	TenantName	string		`json:"name"`
}

type ServiceCatalogStruct struct{
	Endpoints		[]EndpointsStruct	`json:"endpoints"`
	Endpoints_links		[]string		`json:"endpoints_links"`
	EndpointType		string			`json:"type"`
	EndpointName		string			`json:"name"`
}
type EndpointsStruct struct{
	AdminURL		string	`json:"adminURL"`
	Region			string	`json:"region"`
	EndpiontID		string	`json:"id"`
	InternalURL		string	`json:"internalURL"`
	PublicURL		string	`json:"publicURL"`
}

type UserStruct struct{
	UserName	string		`json:"username"`
	Roles_links	[]string	`json:"roles_links"`
	UserID		string		`json:"id"`
	Roles		[]RolesStruct	`json:"roles"`
	Name		string		`json:"name"`
}
type RolesStruct struct{
	RoleName 	string		`json:"name"`
}

type Metadata struct{
	Is_admin	int64		`json:"is_admin"`
	Roles		[]string	`json:"roles"`
}





type Configuration struct {
	IdentityEndpoint	string	`json:"IdentityEndpoint"`
    	UserName		string	`json:"userName"`
	Password		string	`json:"password"`
    	TenantName 		string	`json:"tenantName"`
    	TenantId 		string	`json:"tenantID"`
	ProjectId		string	`json:"projectID"`
	ProjectName		string	`json:"projectName"`
    	Container 		string	`json:"container"`
    	Region	 		string	`json:"region"`
}

func GetHOSAuthToken() (string, HOSAutToken){
//func main(){
	var filename string = "goclienthos/authtoken/getAuthToken.go"
	_, filePath, _, _ := runtime.Caller(0)
	fmt.Println("CurrentFilePath:==",filePath)
	absPath :=(strings.Replace(filePath, filename, "conf/hosconfiguration.json", 1))
	//absPath :=(strings.Replace(filePath, filename, "openStackConfiguration.json", 1))
	fmt.Println("HOSConfigurationFilePath:==",absPath)
	file, _ := os.Open(absPath)
	decoder := json.NewDecoder(file)
	tempConfig := Configuration{}
	err := decoder.Decode(&tempConfig)
	if err != nil{
		fmt.Println("ConfigurationError:", err)
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

	var reqBody string = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}, "tenantName": "`+ tempConfig.TenantName+`"}}`
	//var reqBody string = `{"auth":{"passwordCredentials":{"username": "` + tempConfig.UserName +`", "password": "`+ tempConfig.Password +`"}, "tenantId": "`+ tempConfig.TenantId +`", "tenantName": "`+ tempConfig.TenantName+`", "Container": "`+ tempConfig.Container +`","ImageRegion": "`+ tempConfig.Region +`"}}`
	fmt.Println("Request Body:==",reqBody)

	var reqURL string = tempConfig.IdentityEndpoint + "/tokens"
	fmt.Println("\nRequest URL:==",reqURL)

	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(reqBody))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	fmt.Println("Printing request:==",req)
	res, _ := http.DefaultClient.Do(req)
	fmt.Println("Status:==", res.Status)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)

	//fmt.Print("In GET HOS AUTH TOKEN respBody:==",respBody)

	respBodyInString:= string(respBody)
	fmt.Println("\n\n\nIn GET HOS AUTH TOKEN respBodyInString:==\n\n",respBodyInString)
	fmt.Println("\n\n\n")
	//rBodyInByte := []byte(respBody)
	//fmt.Println("rBodyInByte",rBodyInByte)

	var jsonAuthTokenBody HOSAutToken

	//respMarshed,_ := json.Marshal(rBodyInByte)
	//fmt.Println("marshedRespBody:===",respMarshed)
	//stringRespMarshed:=string(respMarshed)
	//fmt.Println("marshedBody in string", stringRespMarshed)
	if err = json.Unmarshal(respBody, &jsonAuthTokenBody); err != nil{
		fmt.Println("Error in unmarshing:==",err)
	}

	//newDecoder := json.NewDecoder(respBody)
	//newTempConfig := Endpoint{}
	//error := newDecoder.Decode(&newTempConfig)
	//if error != nil{
	//	fmt.Println("ConfigurationError:", error)
	//}
	//
	fmt.Println("\nIn GET HOS AUTH TOKEN HOSResponseBody:===\n", jsonAuthTokenBody)
	fmt.Printf("\nIn GET HOS AUTH TOKEN jsonAuthTokenBody:===\n %+v\n\n", jsonAuthTokenBody)
	fmt.Println("\nIn GET HOS AUTH TOKEN AuthToken:==\n",jsonAuthTokenBody.Access.Token.AuthToken)
	return  jsonAuthTokenBody.Access.Token.AuthToken, jsonAuthTokenBody

}
