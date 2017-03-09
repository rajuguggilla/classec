package AuthenticationToken

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"

)

//------------------------------------------------------Responsebody Structure For V3 auth Token Request--------------------------------------------------------//

type OpenStackAutToken_v3 struct{
	Token 	AccessStruct_v3	`json:"token"`
}

type  AccessStruct_v3 struct {
	Is_Domain  		bool			`json:"is_domain"`
	Methods			[]string		`json:"methods"`
	Roles			[]DemoStruct1_v3		`json:"roles"`
	Expires_At		string			`json:"expires_at"`
	Project  		DemoStruct2_v3		`json:"project"`
	Catalog			[]SingleCatalogStruct_v3	`json:"catalog"`
	User			DemoStruct2_v3		`json:"user"`
	Audit_Ids		[]string		`json:"audit_ids"`
	Issued_At		string			`json:"issued_at"`
}

type DemoStruct1_v3 struct{
	Id	string		`json:"id"`
	Name	string		`json:"name"`
}

type DemoStruct2_v3 struct{
	Domain 	DemoStruct1_v3	`json:"domain"`
	Id	string		`json:"id"`
	Name	string		`json:"name"`
}

type SingleEndpointStruct_v3 struct{
	Region_Id	string		`json:"region_id"`
	URL		string		`json:"url"`
	Interface	string		`json:"interface"`
	Id		string		`json:"id"`
	Region		string		`json:"region"`
}

type SingleCatalogStruct_v3 struct{
	Endpoints 	[]SingleEndpointStruct_v3	`json:"endpoints"`
	Id		string			`json:"id"`
	Type		string			`json:"type"`
	Name		string			`json:"name"`
}

//------------------------------------------------------Structure to read Configuration file --------------------------------------------------------//

type Configuration struct {
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

func GetOpenStackAuthToken_v3(reqURL string, reqBody string) (string, OpenStackEndpoints, string){

	var authToken string
	var Endpoints	OpenStackEndpoints
	var jsonAuthTokenBody OpenStackAutToken_v3

	fmt.Println("Request Body:==",reqBody)
	fmt.Println("\nRequest URL:==",reqURL)

	req, reqErr := http.NewRequest("POST", reqURL, strings.NewReader(reqBody))
	if reqErr != nil{
		fmt.Println("Error in generating http.NewRequest:==", reqErr)
		return authToken, Endpoints, "Error in generating http.NewRequest."
	}

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

	fmt.Print("respBody:==",respBody)
	authToken = res.Header.Get("X-Subject-Token")
	unmErr := json.Unmarshal(respBody, &jsonAuthTokenBody)
	if unmErr != nil{
		fmt.Println("Error in unmarshing:==",unmErr)
		return authToken, Endpoints, "Error in unmarshing."
	}else{
		for i:=0;i<len(jsonAuthTokenBody.Token.Catalog);i++ {
			Endpoints.ApiEndpoints[i].EndpointName = jsonAuthTokenBody.Token.Catalog[i].Name
			Endpoints.ApiEndpoints[i].EndpointType = jsonAuthTokenBody.Token.Catalog[i].Type
			for j:=0; j<len(jsonAuthTokenBody.Token.Catalog[i].Endpoints); j++{
				if jsonAuthTokenBody.Token.Catalog[i].Endpoints[j].Interface == "public" {
					Endpoints.ApiEndpoints[i].EndpointURL = jsonAuthTokenBody.Token.Catalog[i].Endpoints[j].URL
				}
			}
		}
	}
	fmt.Println("OSResponseBody in string:===")
	fmt.Printf("%+v\n\n", jsonAuthTokenBody)
	fmt.Println("Endpoints in string:===")
	fmt.Printf("%+v\n\n", Endpoints)
	fmt.Println("AuthToken:==",authToken)
	return  authToken, Endpoints, ""
}