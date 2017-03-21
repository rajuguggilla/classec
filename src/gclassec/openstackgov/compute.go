package openstackgov


import (
	"io/ioutil"
	"gclassec/errorcodes/errcode"
	"net/http"
	"encoding/json"
	"gclassec/structs/openstackInstance"
	"gclassec/openstackgov/authenticationtoken"
	"errors"
	"gclassec/loggers"
)

func ListComputeInstances() (openstackInstance.ComputeListStruct,error){
	logger := Loggers.New()
	var tempFlvStruct openstackInstance.FlavorsListStruct
	var flvError error
	//logger.Info("---------------------------------------Flavours Details in Compute Start------------------------------------------------------")
	tempFlvStruct,flvError=ListFlavors()
	if flvError != nil{
		logger.Error("Error In getting Flavours Details:==",flvError)
	}else{

		logger.Info(tempFlvStruct)
	}
	//logger.Info("---------------------------------------Flavours Details in Compute End------------------------------------------------------")
	var computeEndpoint string
	var authToken string
	var authError error
	var endpointsStruct openstackInstance.OpenStackEndpoints
	var jsonCompStruct openstackInstance.ComputeListStruct

	//logger.Info("=====================Scoped Authentication Token====================")
	authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(false)
	logger.Info("authToken:==", authToken)
	logger.Info("endpointsStruct:==",endpointsStruct)
	logger.Info("authError:==",authError)

	if authError != nil {
		logger.Error("authError:==",authError)
		return jsonCompStruct,authError
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

	var reqURL string =  computeEndpoint + "/servers/detail"
	logger.Info("reqURL:====",reqURL)
	//var reqURL string = "https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c/flavors/detail"
	req, errReq := http.NewRequest("GET", reqURL, nil)
	if errReq != nil{
		logger.Error("HOS: ", errcode.ErrReq)
		return jsonCompStruct,errors.New(errcode.ErrReq)
	}

	req.Header.Add("x-auth-token", authToken)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
	if errClient != nil{
		logger.Error("HOS: ", errcode.ErrReq)
		return jsonCompStruct, errors.New(errcode.ErrReq)
	}

	logger.Info("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
	if errResp != nil{
		logger.Info("HOS: ", errcode.ErrResp)
		return jsonCompStruct, errors.New(errcode.ErrResp)
	}

	respBodyInString:= string(respBody)
	logger.Info("\nrespBodyInString:==\n",respBodyInString)
	unmError := json.Unmarshal(respBody, &jsonCompStruct)
	if unmError != nil {
		logger.Error("Error in Unmarshing:==", unmError)
	}

	logger.Info("jsonCompStruct Before Updating Flavors:===", jsonCompStruct)

	for i:=0; i<len(jsonCompStruct.Servers);i++{
		tempFID := jsonCompStruct.Servers[i].Flavor.FlavorID
		for j:=0; j<len(tempFlvStruct.Flavors);j++ {
			if tempFID == tempFlvStruct.Flavors[j].FlavorId{
				jsonCompStruct.Servers[i].Flavor.FlavorName=tempFlvStruct.Flavors[j].FlavorName
				jsonCompStruct.Servers[i].Flavor.Disk=tempFlvStruct.Flavors[j].Disk
				jsonCompStruct.Servers[i].Flavor.Ram=tempFlvStruct.Flavors[j].RAM
				jsonCompStruct.Servers[i].Flavor.VCPUS=tempFlvStruct.Flavors[j].VCPUS
			}
		}
	}
	logger.Info("jsonCompStruct After Updating Flavors:===", jsonCompStruct)
	return jsonCompStruct,nil
}








