package authenticationtoken

import (
	"strings"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"gclassec/structs/openstackInstance"
	"errors"
	"gclassec/loggers"
)



func GetOpenStackAuthToken_v2(reqURL string, reqBody string ) (string,openstackInstance.OpenStackEndpoints, error){
	logger := Loggers.New()
	var authToken string
	var Endpoints	openstackInstance.OpenStackEndpoints
	logger.Info("Request Body:==",reqBody)
	logger.Info("\nRequest URL:==",reqURL)

	req, reqErr := http.NewRequest("POST", reqURL, strings.NewReader(reqBody))
	if reqErr != nil{
		logger.Error("Error in generating http.NewRequest:==", reqErr)
		return authToken, Endpoints, reqErr
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	logger.Info("Printing request:==",req)
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

	respBodyInString:= string(respBody)
	logger.Info("respBodyInString:==\n",respBodyInString)
	var jsonAuthTokenBody openstackInstance.OpenStackAutToken_v2

	unmErr := json.Unmarshal(respBody, &jsonAuthTokenBody);
	if unmErr != nil{
		logger.Error("Error in unmarshing:==",unmErr)
		return authToken, Endpoints, errors.New("Error in unmarshing.")
	}else {
		authToken= jsonAuthTokenBody.Access.Token.AuthToken
		for i:=0;i<len(jsonAuthTokenBody.Access.ServiceCatalog);i++ {
			Endpoints.ApiEndpoints = append(Endpoints.ApiEndpoints,openstackInstance.EndpointStruct{})
			Endpoints.ApiEndpoints[i].EndpointName = jsonAuthTokenBody.Access.ServiceCatalog[i].EndpointName
			Endpoints.ApiEndpoints[i].EndpointType = jsonAuthTokenBody.Access.ServiceCatalog[i].EndpointType
			Endpoints.ApiEndpoints[i].EndpointURL = jsonAuthTokenBody.Access.ServiceCatalog[i].Endpoints[0].PublicURL
		}
	}
	logger.Info("OSResponseBody:===\n")
	logger.Info("jsonAuthTokenBody:===", jsonAuthTokenBody)
	logger.Info("Endpoints:===\n")
	logger.Info("Endpoints:===", Endpoints)
	logger.Info("AuthToken:==",authToken)
	return  authToken, Endpoints, nil
}
