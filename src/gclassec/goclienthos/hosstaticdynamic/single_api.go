package hosstaticdynamic

import (
	"fmt"
	"net/http"
	"io/ioutil"
//	"hos/HOSAuthToken"
//	"hos/HOS_API_Function/comp/flavor"
	"encoding/json"

//	"HOS_API_Function"
	//"hos/HOS_API_Function/comp/flavor"
//"hos/HOSAuthToken"

	"gclassec/goclienthos/authtoken"
	"gclassec/goclienthos/compute"
	"gclassec/structs/hosstruct"
	"gclassec/errorcodes/errcode"
	"gclassec/loggers"
)






func ComputeWithCPU() hosstruct.CompleteComputeResponse{
	logger := Loggers.New()
	//fmt.Println("This to get Nothing")
	var computeEndpoint string
	var auth, hosConfig, err = authtoken.GetHOSAuthToken()

		if err != nil{
			fmt.Println("HOS : ", errcode.ErrAuth)
			logger.Error("HOS : ", errcode.ErrAuth)
			return hosstruct.CompleteComputeResponse{}
		}
	fmt.Println("HOS AuthToken:=====\n", auth)
	fmt.Println("HOS Configuration:=====\n %+v", hosConfig)
	for i := 0; i < len(hosConfig.Access.ServiceCatalog); i++ {
		if hosConfig.Access.ServiceCatalog[i].EndpointType =="compute"{
			//for j:= 0; j< len(hosConfig.Access.ServiceCatalog[i].Endpoints); j++ {
			computeEndpoint = hosConfig.Access.ServiceCatalog[i].Endpoints[0].PublicURL
			fmt.Println("ComputeeNDpOINT:====",computeEndpoint)
			//https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c
					//}
			}
		}

	var reqURL string =  computeEndpoint + "/servers/detail"
	//var reqURL string = "http://" + hosConfiguration.KeystoneEndpointIP + ":8774/v2.1/" + hosConfiguration.TenantId + "/servers/detail"
	fmt.Println("Request Body:==",reqURL)
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)

	var jsonComputeResponse hosstruct.CompleteComputeResponse

	if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
		fmt.Println("Error in Unmarshing:==", err)
	}



	var FlavorsList hosstruct.FlvRespStruct
	FlavorsList = compute.Flavors()

	fmt.Printf("%+v\n\n", FlavorsList)


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

	fmt.Println("Printing Final jsonComputeResponse ")
	fmt.Printf("%+v\n\n", jsonComputeResponse)



	fmt.Println("\n\n ++++++++++++++++++++++++++++++ End of Static +++++++++++++++++++++++++++++\n\n")

	for k:=0; k<len(jsonComputeResponse.Servers);k++{
		temID := jsonComputeResponse.Servers[k].InstanceID
		dynamicData := AvgCpuUtil(temID)
		//fmt.Println("dynamicData)
		fmt.Println("\n\n--------------------------------------------\n\n")
		fmt.Println("@@@@@@@@@@@@@@@@@@---  DYNAMIC @@@@@@@@@@@@@",dynamicData)
		//var jsonStaticResponse MyStruct
		//jsonStaticResponse.ServersResp = jsonStaticResponse1.Servers[k]
		//jsonStaticResponse.Cpu_Util = dynamicData
		//jsonStaticResponses = append(jsonStaticResponses,jsonStaticResponse)
		jsonComputeResponse.Servers[k].Cpu_Util=dynamicData
	}



	fmt.Println("\n\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")

	fmt.Printf("%+v",jsonComputeResponse)
	return jsonComputeResponse


}
