package compute

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gclassec/goclienthos/authtoken"
	"gclassec/structs/hosstruct"
	"gclassec/loggers"
	"fmt"
	"gclassec/errorcodes/errcode"
)




func Compute() (hosstruct.ComputeResponse, error) {


	//fmt.Println("This to get Nothing")
	logger := Loggers.New()
	var computeEndpoint string
	var auth, hosConfig, err = authtoken.GetHOSAuthToken()
		if err != nil{
			fmt.Println("HOS : ", errcode.ErrAuth)
			logger.Error("HOS :", errcode.ErrAuth)
			return hosstruct.ComputeResponse{}, err
		}
	logger.Debug("HOS AuthToken:=====\n", auth)
	logger.Debug("HOS Configuration:=====\n %+v", hosConfig)
	for i := 0; i < len(hosConfig.Access.ServiceCatalog); i++ {
		if hosConfig.Access.ServiceCatalog[i].EndpointType =="compute"{
			//for j:= 0; j< len(hosConfig.Access.ServiceCatalog[i].Endpoints); j++ {
			computeEndpoint = hosConfig.Access.ServiceCatalog[i].Endpoints[0].PublicURL
			logger.Info("ComputeEndPoint:====",computeEndpoint)
			//https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c
					//}
			}
		}

	var reqURL string =  computeEndpoint + "/servers/detail"
	//var reqURL string = "http://" + hosConfiguration.KeystoneEndpointIP + ":8774/v2.1/" + hosConfiguration.TenantId + "/servers/detail"
	logger.Info("Request Body:==",reqURL)
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	logger.Info("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)

	//fmt.Print("respBody:==\n",respBody)
	respBodyInString:= string(respBody)
	logger.Info("\nrespBodyInString:==\n",respBodyInString)
	//return respBodyInString
	var jsonComputeResponse hosstruct.ComputeResponse
	if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
		logger.Error("Error in Unmarshing:==", err)
	}
	logger.Info("Printing Initial jsonComputeResponse ")
	logger.Info("%+v\n\n", jsonComputeResponse)
	//return jsonComputeResponse

	var FlavorsList hosstruct.FlvRespStruct
	FlavorsList = Flavors()
	//var jsonFlavorList FlvRespStruct
	//if err := json.Unmarshal([]byte(FlavorsStringList), &jsonFlavorList); err != nil {
	//	fmt.Println("Error in Unmarshing:==", err)
	//}
	//
	logger.Info("%+v\n\n", FlavorsList)


	for i:=0; i<len(jsonComputeResponse.Servers);i++{
		tempFID := jsonComputeResponse.Servers[i].Flavor.FlavorID
		for j:=0; j<len(FlavorsList.Flavors);j++ {
			if tempFID == FlavorsList.Flavors[j].FlavorID{
				jsonComputeResponse.Servers[i].Flavor.FlavorName=FlavorsList.Flavors[j].FlavorName
				jsonComputeResponse.Servers[i].Flavor.Disk=FlavorsList.Flavors[j].Disk
				jsonComputeResponse.Servers[i].Flavor.Ram=FlavorsList.Flavors[j].Ram
				jsonComputeResponse.Servers[i].Flavor.VCPUS=FlavorsList.Flavors[j].VCPUS

			}

		}

	}
	logger.Info("Printing Final jsonComputeResponse ")
	logger.Info("%+v\n\n", jsonComputeResponse)
	TempStr, _ := json.Marshal(&jsonComputeResponse)
	logger.Info("Printing Final jsonComputeResponse in string:===\n\n ",string(TempStr))
	return jsonComputeResponse, nil
	//return string(TempStr)

}
