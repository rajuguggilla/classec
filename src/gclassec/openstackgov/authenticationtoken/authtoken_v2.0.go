package authenticationtoken

import (
	"strings"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"gclassec/structs/openstackInstance"
)



func GetOpenStackAuthToken_v2(reqURL string, reqBody string ) (string,openstackInstance.OpenStackEndpoints,string){

	var authToken string
	var Endpoints	openstackInstance.OpenStackEndpoints
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
	var jsonAuthTokenBody openstackInstance.OpenStackAutToken_v2

	unmErr := json.Unmarshal(respBody, &jsonAuthTokenBody);
	if unmErr != nil{
		fmt.Println("Error in unmarshing:==",unmErr)
		return authToken, Endpoints, "Error in unmarshing."
	}else {
		authToken= jsonAuthTokenBody.Access.Token.AuthToken
		for i:=0;i<len(jsonAuthTokenBody.Access.ServiceCatalog);i++ {
			Endpoints.ApiEndpoints = append(Endpoints.ApiEndpoints,openstackInstance.EndpointStruct{})
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
