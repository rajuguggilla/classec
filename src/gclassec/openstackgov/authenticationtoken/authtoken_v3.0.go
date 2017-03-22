package authenticationtoken

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"gclassec/structs/openstackInstance"
	"gclassec/loggers"

)



func GetOpenStackAuthToken_v3(reqURL string, reqBody string) (string, openstackInstance.OpenStackEndpoints, error){
	logger := Loggers.New()
	var authToken string
	var Endpoints	openstackInstance.OpenStackEndpoints
	var jsonAuthTokenBody openstackInstance.OpenStackAutToken_v3

	logger.Info("Request Body:==",reqBody)
	logger.Info("\nRequest URL:==",reqURL)

	req, reqErr := http.NewRequest("POST", reqURL, strings.NewReader(reqBody))
	if reqErr != nil{
		logger.Error("Error in generating http.NewRequest:==", reqErr)
		return authToken, Endpoints, reqErr
	}

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		logger.Error("Error in Request Response:==", resErr)
		return authToken, Endpoints, resErr
	} else {
		logger.Info("Status:==", res.Status)
		defer res.Body.Close()
	}
	respBody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil{
		logger.Error("Error in reading Response Body:==", readErr)
		return authToken, Endpoints, readErr
	}

	logger.Info("respBody:==",respBody)
	respBodyInString := string(respBody)
	logger.Info("\nrespBodyInString:===", respBodyInString)
	authToken = res.Header.Get("X-Subject-Token")
	logger.Info("X-Subject-Token:===", authToken)

	unmErr := json.Unmarshal(respBody, &jsonAuthTokenBody)
	Templen :=len(jsonAuthTokenBody.Token.Catalog)
	logger.Info("Length of jsonAuthTokenBody.Token.Catalog:==", Templen)
	if unmErr != nil{
		logger.Error("Error in unmarshing:==",unmErr)
		return authToken, Endpoints, unmErr
	}else{
		for i:=0;i<Templen;i++ {
			//logger.Info("Value of I:==",i)
			Endpoints.ApiEndpoints = append(Endpoints.ApiEndpoints,openstackInstance.EndpointStruct{})
			Endpoints.ApiEndpoints[i].EndpointName = jsonAuthTokenBody.Token.Catalog[i].Name
			Endpoints.ApiEndpoints[i].EndpointType = jsonAuthTokenBody.Token.Catalog[i].Type
			//temp2:= len(jsonAuthTokenBody.Token.Catalog[i].Endpoints)
			//logger.Info("Length of jsonAuthTokenBody.Token.Catalog:==", temp2)
			for j:=0; j<len(jsonAuthTokenBody.Token.Catalog[i].Endpoints); j++{
				logger.Info("Value of J:==",j)
				if jsonAuthTokenBody.Token.Catalog[i].Endpoints[j].Interface == "public" {
					Endpoints.ApiEndpoints[i].EndpointURL = jsonAuthTokenBody.Token.Catalog[i].Endpoints[j].URL
				}
			}
			//if i < (Templen-1){
			//Endpoints.ApiEndpoints = append(Endpoints.ApiEndpoints,EndpointStruct{})
			//}
		}
	}
	logger.Info("OSResponseBody in string:===")
	logger.Info("jsonAuthTokenBody:===", jsonAuthTokenBody)
	logger.Info("Endpoints in string:===")
	logger.Info("Endpoints:==", Endpoints)
	logger.Info("AuthToken:==",authToken)
	return  authToken, Endpoints, nil
}