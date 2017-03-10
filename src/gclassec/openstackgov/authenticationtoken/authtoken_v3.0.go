package authenticationtoken

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"

	"gclassec/structs/openstackInstance"
)



func GetOpenStackAuthToken_v3(reqURL string, reqBody string) (string, openstackInstance.OpenStackEndpoints, string){

	var authToken string
	var Endpoints	openstackInstance.OpenStackEndpoints
	var jsonAuthTokenBody openstackInstance.OpenStackAutToken_v3

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
	respBodyInString := string(respBody)
	fmt.Println("\nrespBodyInString:===", respBodyInString)
	authToken = res.Header.Get("X-Subject-Token")
	fmt.Println("X-Subject-Token:===", authToken)

	unmErr := json.Unmarshal(respBody, &jsonAuthTokenBody)
	Templen :=len(jsonAuthTokenBody.Token.Catalog)
	fmt.Println("Length of jsonAuthTokenBody.Token.Catalog:==", Templen)
	if unmErr != nil{
		fmt.Println("Error in unmarshing:==",unmErr)
		return authToken, Endpoints, "Error in unmarshing."
	}else{
		for i:=0;i<Templen;i++ {
			fmt.Println("Value of I:==",i)
			Endpoints.ApiEndpoints = append(Endpoints.ApiEndpoints,openstackInstance.EndpointStruct{})
			Endpoints.ApiEndpoints[i].EndpointName = jsonAuthTokenBody.Token.Catalog[i].Name
			Endpoints.ApiEndpoints[i].EndpointType = jsonAuthTokenBody.Token.Catalog[i].Type
			temp2:= len(jsonAuthTokenBody.Token.Catalog[i].Endpoints)
			fmt.Println("Length of jsonAuthTokenBody.Token.Catalog:==", temp2)
			for j:=0; j<len(jsonAuthTokenBody.Token.Catalog[i].Endpoints); j++{
				fmt.Println("Value of J:==",j)
				if jsonAuthTokenBody.Token.Catalog[i].Endpoints[j].Interface == "public" {
					Endpoints.ApiEndpoints[i].EndpointURL = jsonAuthTokenBody.Token.Catalog[i].Endpoints[j].URL
				}
			}
			//if i < (Templen-1){
			//Endpoints.ApiEndpoints = append(Endpoints.ApiEndpoints,EndpointStruct{})
			//}
		}
	}
	fmt.Println("OSResponseBody in string:===")
	fmt.Printf("%+v\n\n", jsonAuthTokenBody)
	fmt.Println("Endpoints in string:===")
	fmt.Printf("%+v\n\n", Endpoints)
	fmt.Println("AuthToken:==",authToken)
	return  authToken, Endpoints, ""
}