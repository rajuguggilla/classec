package openstackgov


import (
	"gclassec/structs/openstackInstance"
	"gclassec/openstackgov/authenticationtoken"
	"io/ioutil"
	"gclassec/errorcodes/errcode"
	"net/http"
	"encoding/json"
	"gclassec/loggers"
)





func ListFlavors() (openstackInstance.FlavorsListStruct, error){
	logger := Loggers.New()
	var computeEndpoint string
	var authToken string
	var authError error
	var endpointsStruct openstackInstance.OpenStackEndpoints
	var jsonFlvStruct openstackInstance.FlavorsListStruct
	//logger.Info("=====================Unscoped Authentication Token====================")
	//authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(true)
	//logger.Info("authToken:==", authToken)
	//logger.Info("endpointsStruct:==",endpointsStruct)
	//logger.Info("authError:==",authError)
	//logger.Info("=====================Scoped Authentication Token====================")
	authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(false)
	logger.Info("authToken:==", authToken)
	logger.Info("endpointsStruct:==",endpointsStruct)
	logger.Info("authError:==",authError)

	if authError!= nil{
		logger.Error("authError:==",authError)
		return jsonFlvStruct,authError
	}else{
		logger.Info("authToken:==", authToken)
		logger.Info("endpointsStruct:==", endpointsStruct)
	}
	for i := 0; i < len(endpointsStruct.ApiEndpoints); i++ {
		if endpointsStruct.ApiEndpoints[i].EndpointType =="compute"{
			computeEndpoint = endpointsStruct.ApiEndpoints[i].EndpointURL
			logger.Info("ComputeEndPoint:====",computeEndpoint)
			//https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c
			}
		}

	var reqURL string =  computeEndpoint + "/flavors/detail"
	logger.Info("reqURL:====",reqURL)
	//var reqURL string = "https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c/flavors/detail"
	req, errReq := http.NewRequest("GET", reqURL, nil)
	if errReq != nil{
		logger.Error(errcode.ErrReq,":==  ", errReq)
		return jsonFlvStruct,errReq
	}

	req.Header.Add("x-auth-token", authToken)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
	if errClient != nil{
		logger.Error(errcode.ErrReq,":==  ", errClient)
		return jsonFlvStruct, errClient
	}

	logger.Info("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
	if errResp != nil{
		logger.Error(errcode.ErrResp,":==  ",errResp)
		return jsonFlvStruct, errResp
	}

	respBodyInString:= string(respBody)
	logger.Info("\nrespBodyInString:==\n",respBodyInString)
	unmError := json.Unmarshal(respBody, &jsonFlvStruct)
	if unmError != nil {
		logger.Error("Error in Unmarshing:==", unmError)
		return	jsonFlvStruct,unmError
	}

	logger.Info("jsonFlvStruct:====", jsonFlvStruct)
	return jsonFlvStruct,nil
}




