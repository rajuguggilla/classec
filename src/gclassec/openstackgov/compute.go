package openstackgov

import (
	"gclassec/structs/openstackInstance"
	"fmt"
	"gclassec/openstackgov/authenticationtoken"
)



import (
	"io/ioutil"
	"gclassec/errorcodes/errcode"
	"net/http"
	"encoding/json"

)

func ListComputeInstances() (openstackInstance.ComputeListStruct,string){
	var tempFlvStruct openstackInstance.FlavorsListStruct
	var flvError string
	tempFlvStruct,flvError=ListFlavors()
	if flvError != ""{
		fmt.Println("Error In getting Flavours Details:==",flvError)
	}else{
		fmt.Println("---------------------------------------Flavours Details------------------------------------------------------")
		fmt.Println(tempFlvStruct)
	}

	var computeEndpoint string
	var authToken string
	var authError string
	var endpointsStruct openstackInstance.OpenStackEndpoints
	var jsonCompStruct openstackInstance.ComputeListStruct

	fmt.Println("=====================Scoped Authentication Token====================")
	authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(false)
	fmt.Println("authToken:==", authToken)
	fmt.Println("endpointsStruct:==",endpointsStruct)
	fmt.Println("authError:==",authError)

	if authError!="" {
		fmt.Println("authError:==",authError)
		return jsonCompStruct,authError
	}else{
		fmt.Println("authToken:==", authToken)
		fmt.Println("endpointsStruct:==", endpointsStruct)
	}
	for i := 0; i < len(endpointsStruct.ApiEndpoints); i++ {
		if endpointsStruct.ApiEndpoints[i].EndpointType =="compute"{
			computeEndpoint = endpointsStruct.ApiEndpoints[i].EndpointURL
			fmt.Println("ComputeEndPoint:====",computeEndpoint)
			//https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c
			}
		}

	var reqURL string =  computeEndpoint + "/servers/detail"
	fmt.Println("reqURL:====",reqURL)
	//var reqURL string = "https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c/flavors/detail"
	req, errReq := http.NewRequest("GET", reqURL, nil)
	if errReq != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		return jsonCompStruct,errcode.ErrReq
	}

	req.Header.Add("x-auth-token", authToken)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
	if errClient != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		return jsonCompStruct, errcode.ErrReq
	}

	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
	if errResp != nil{
		fmt.Println("HOS: ", errcode.ErrResp)
		return jsonCompStruct, errcode.ErrResp
	}

	respBodyInString:= string(respBody)
	fmt.Println("\nrespBodyInString:==\n",respBodyInString)
	unmError := json.Unmarshal(respBody, &jsonCompStruct)
	if unmError != nil {
		fmt.Println("Error in Unmarshing:==", unmError)
	}

	fmt.Printf("\n%+v\n", jsonCompStruct)

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
	fmt.Printf("\n%+v\n", jsonCompStruct)
	return jsonCompStruct,""
}








