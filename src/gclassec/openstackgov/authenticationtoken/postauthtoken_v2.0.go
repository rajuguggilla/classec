package AuthenticationToken
//package main
import (
	"strings"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

//------------------------------------------------------Responsebody Structure For V2 auth Token Request--------------------------------------------------------//
type OpenStackAutToken_v2 struct{
	Access 	AccessStruct_v2	`json:"access"`
}


type  AccessStruct_v2 struct {
	Token  		TokenStruct_v2		`json:"token"`
	ServiceCatalog	[]ServiceStructure_v2	`json:"serviceCatalog"`
	User		UserStruct_v2		`json:"user"`
	Metadata	Metadata_v2		`json:"metadata"`

}

type TokenStruct_v2 struct{
	Issued_at	string		`json:"issued_at"`
	Expires		string		`json:"expires"`
	AuthToken	string		`json:"id"`
	Tenant		TenantStruct_v2	`json:"tenant"`
	Audit_ids	[]string	`json:"audit_ids"`
}
type TenantStruct_v2 struct{
	Description	string		`json:"description"`
	Enabled		bool		`json:"enabled"`
	TenanatID	string		`json:"id"`
	TenantName	string		`json:"name"`
}

type ServiceStructure_v2 struct{
	Endpoints		[]EndpointsStruct_v2	`json:"endpoints"`
	Endpoints_links		[]string		`json:"endpoints_links"`
	EndpointType		string			`json:"type"`
	EndpointName		string			`json:"name"`
}
type EndpointsStruct_v2 struct{
	AdminURL		string	`json:"adminURL"`
	Region			string	`json:"region"`
	EndpiontID		string	`json:"id"`
	InternalURL		string	`json:"internalURL"`
	PublicURL		string	`json:"publicURL"`
}

type UserStruct_v2 struct{
	UserName	string		`json:"username"`
	Roles_links	[]string	`json:"roles_links"`
	UserID		string		`json:"id"`
	Roles		[]Roles_v2		`json:"roles"`
	Name		string		`json:"name"`
}
type Roles_v2 struct{
	RoleName 	string		`json:"name"`
}

type Metadata_v2 struct{
	Is_admin	int64		`json:"is_admin"`
	Roles		[]string	`json:"roles"`
}


func GetOpenStackAuthToken_v2(reqURL string, reqBody string ) (string,OpenStackEndpoints,string){
//func main(){
	var authToken string
	var Endpoints	OpenStackEndpoints
	fmt.Println("Request Body:==",reqBody)
	fmt.Println("\nRequest URL:==",reqURL)

	req, reqErr := http.NewRequest("POST", reqURL, strings.NewReader(reqBody))
	if reqErr != nil{
		fmt.Println("Error in generating http.NewRequest:==", reqErr)
		return authToken, Endpoints, "Error in generating http.NewRequest."
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	fmt.Println("Printing request:==",req)
	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		fmt.Println("Error in Request Response:==", resErr)
		return authToken, Endpoints, "Error in Request Response."
	} else {
		fmt.Println("Status:==", res.Status)
		defer res.Body.Close()
	}
	respBody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		fmt.Println("Error in reading Response Body:==", readErr)
		return authToken, Endpoints, "Error in reading Response Body."
	}

	respBodyInString:= string(respBody)
	fmt.Println("respBodyInString:==\n",respBodyInString)
	var jsonAuthTokenBody OpenStackAutToken_v2

	unmErr := json.Unmarshal(respBody, &jsonAuthTokenBody);
	if unmErr != nil{
		fmt.Println("Error in unmarshing:==",unmErr)
		return authToken, Endpoints, "Error in unmarshing."
	}else {
		authToken= jsonAuthTokenBody.Access.Token.AuthToken
		for i:=0;i<len(jsonAuthTokenBody.Access.ServiceCatalog);i++ {
			Endpoints.ApiEndpoints[i].EndpointName = jsonAuthTokenBody.Access.ServiceCatalog[i].EndpointName
			Endpoints.ApiEndpoints[i].EndpointType = jsonAuthTokenBody.Access.ServiceCatalog[i].EndpointType
			Endpoints.ApiEndpoints[i].EndpointURL = jsonAuthTokenBody.Access.ServiceCatalog[i].Endpoints[0].PublicURL
		}
	}
	fmt.Println("OSResponseBody:===\n")
	fmt.Printf("%+v\n\n", jsonAuthTokenBody)
	fmt.Println("Endpoints:===\n")
	fmt.Printf("%+v\n\n", Endpoints)
	fmt.Println("AuthToken:==",authToken)
	return  authToken, Endpoints, ""
}
