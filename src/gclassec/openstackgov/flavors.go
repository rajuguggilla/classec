package openstackgov


import (
	"gclassec/structs/openstackInstance"
	"gclassec/openstackgov/authenticationtoken"
	"fmt"
	"io/ioutil"
	"gclassec/errorcodes/errcode"
	"net/http"
	"encoding/json"
)





func ListFlavors() (openstackInstance.FlavorsListStruct, string){

	var computeEndpoint string
	var authToken string
	var authError string
	var endpointsStruct openstackInstance.OpenStackEndpoints
	var jsonFlvStruct openstackInstance.FlavorsListStruct
	//fmt.Println("=====================Unscoped Authentication Token====================")
	//authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(true)
	//fmt.Println("authToken:==", authToken)
	//fmt.Println("endpointsStruct:==",endpointsStruct)
	//fmt.Println("authError:==",authError)
	fmt.Println("=====================Scoped Authentication Token====================")
	authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(false)
	fmt.Println("authToken:==", authToken)
	fmt.Println("endpointsStruct:==",endpointsStruct)
	fmt.Println("authError:==",authError)

	if authError!="" {
		fmt.Println("authError:==",authError)
		return jsonFlvStruct,authError
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

	var reqURL string =  computeEndpoint + "/flavors/detail"
	fmt.Println("reqURL:====",reqURL)
	//var reqURL string = "https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c/flavors/detail"
	req, errReq := http.NewRequest("GET", reqURL, nil)
	if errReq != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		return jsonFlvStruct,errcode.ErrReq
	}

	req.Header.Add("x-auth-token", authToken)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
	if errClient != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		return jsonFlvStruct, errcode.ErrReq
	}

	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
	if errResp != nil{
		fmt.Println("HOS: ", errcode.ErrResp)
		return jsonFlvStruct, errcode.ErrResp
	}

	respBodyInString:= string(respBody)
	fmt.Println("\nrespBodyInString:==\n",respBodyInString)
	unmError := json.Unmarshal(respBody, &jsonFlvStruct)
	if unmError != nil {
		fmt.Println("Error in Unmarshing:==", unmError)
	}

	fmt.Println("%+v\n\n", jsonFlvStruct)
	return jsonFlvStruct,""
}




