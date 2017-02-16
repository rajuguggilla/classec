package hos_all_compute_details

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
)


//type MyStruct struct{
//	ServersResp	ServersResponse `json:"Server"`
//	Cpu_Util	float64
//}


type CompleteComputeResponse struct {
      Servers       []CompleteServersResponse      `json:"servers"`
}


type CompleteServersResponse struct {
       Status     	      string `json:"status"`
       Updated                string `json:"updated"`
       HostId                 string `json:"hostId"`
       HostName		      string `json:"OS-EXT-SRV-ATTR:host"`
       //Addresses	      AddressesStruct `json:"addresses"`
       //Links		      []SubLinks	`json:"links"`
       Image 		      ImageStruct	`json:"image"`
       Key_name               string `json:"key_name"`
       Task_State	      string	`json:"OS-EXT-STS:task_state"`
       Vm_State		string		`json:"OS-EXT-STS:vm_state"`
       //InstanceName           string `json:"OS-EXT-SRV-ATTR:instance_name"`
       //Launched_At           string `json:"OS-SRV-USG:launched_at"`
       //Hypervisor_Hostname           string `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
       InstanceID             string `json:"id"`
       Flavor                hosstruct.FlavorsStruct `json:"flavor"`
       Security_Groups       SubSecurityGroup `json:"security_groups"`
       //Terminated_At               string `json:"OS-SRV-USG:terminated_at"`
       Availability_Zone               string `json:"OS-EXT-AZ:availability_zone"`
       User_Id               string `json:"user_id"`
       Vm_Name               string `json:"name"`
       //Created_At               string `json:"created"`
       Tenant_Id               string `json:"tenant_id"`
       //DiskConfig               string `json:"OS-DCF:diskConfig"`
       //volumes_attached               string `json:"os-extended-volumes:volumes_attached"`
       //AccessIPv4               string `json:"accessIPv4"`
       //AccessIPv6               string `json:"accessIPv6"`
       //Progress               int32 `json:"progress"`
       Power_State               int32 `json:"OS-EXT-STS:power_state"`
       //Config_Drive               string `json:"config_drive"`
       //Metadata               string `json:"metadata"`
	Cpu_Util	float64

}

type SubAddress struct {
	MacAddr		string	`json:"OS-EXT-IPS-MAC:mac_addr"`
	Version		string	`json:"version"`
	IpAddress	string	`json:"addr"`
	Type		string	`json:"OS-EXT-IPS:type"`
}

type AddressesStruct struct{
	Lbpvtnet 	[]SubAddress	`json:"lbpvtnet"`
}
type SubLinks struct {
	Href 	string	`json:"href"`
	Rel 	string 	`json:"rel"`
}

type ImageStruct struct {
	ImageID		string		`json:"id"`
	//ImageLinks	SubLinks	`json:"links"`
}

type SubSecurityGroup struct {
	Name 	string		`json:"name"`
}
//func Compute() string {
//func Compute() ComputeResponse {







//func main() {

func ComputeWithCPU() CompleteComputeResponse{
	fmt.Println("Hello .......:")
	//fmt.Println("This to get Nothing")
	var computeEndpoint string
	var auth, hosConfig, err = authtoken.GetHOSAuthToken()

		if err != nil{
			return CompleteComputeResponse{}
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

	//fmt.Print("respBody:==\n",respBody)
//	respBodyInString:= string(respBody)
//	fmt.Println("\nrespBodyInString:==\n",respBodyInString)
	//return respBodyInString
	var jsonComputeResponse CompleteComputeResponse
	//var staticData ServersResponse
//	var allInstance ComputeResponse
	if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
		fmt.Println("Error in Unmarshing:==", err)
	}
//	fmt.Println("Printing Initial jsonComputeResponse ")
//	fmt.Printf("%+v\n\n", jsonComputeResponse)
	//return jsonComputeResponse
//	var jsonComputeResponse ComputeResponse

//	if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
//		fmt.Println("Error in Unmarshing:==", err)
//	}



	var FlavorsList hosstruct.FlvRespStruct
	FlavorsList = compute.Flavors()
	//var jsonFlavorList FlvRespStruct
	//if err := json.Unmarshal([]byte(FlavorsStringList), &jsonFlavorList); err != nil {
	//	fmt.Println("Error in Unmarshing:==", err)
	//}
	//
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
/*	TempStr, _ := json.Marshal(&jsonComputeResponse)
	fmt.Println("Printing Final jsonComputeResponse in string:===\n\n ",string(TempStr))*/
	//fmt.Println("compute=",jsonComputeResponse)
//	return jsonComputeResponse
	//return string(TempStr)

	/*for k:=0; k<len(allInstance.Servers);k++{
		allInstance.Servers[k].InstanceID
	}*/

	/*for k:=0; k<len(allInstance.Servers);k++{
		ID := jsonComputeResponse.Servers[k].InstanceID
		if ID == allInstance.Servers[k].InstanceID {
			allInstance.Servers[k].InstanceID = ID
		}
	}*/

//	fmt.Println("********* Static Data **********",jsonComputeResponse)



//	dynamicData := HOS.Dynamic()
//	fmt.Println("@@@@@@@@@@@@@@@@@@  DYNAMIC @@@@@@@@@@@@@",dynamicData)

/*	var jsonStaticResponse MyStruct
	&jsonStaticResponse.ServersResp=&jsonComputeResponse
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAA", &jsonStaticResponse.SercersResp)
*/

/*	if err := json.Unmarshal(respBody, &jsonStaticResponse.ServersResp); err != nil {
		fmt.Println("Error in Unmarshing:==", err)
	}

fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$", jsonStaticResponse.ServersResp)
        TempStr1, _ := json.Marshal(&jsonStaticResponse.ServersResp)
fmt.Println("++++++++++++++++===============++++++++++++++++==========",TempStr1)*/



	//var jsonStaticResponse1 ComputeResponse
	//var jsonStaticResponses []MyStruct
	////jsonStaticResponse.ServersResp=jsonComputeResponse
	//if err := json.Unmarshal(respBody, &jsonStaticResponse1); err != nil {
	//	fmt.Println("Error in Unmarshing:==", err)
	//}

	/*if err := json.Unmarshal(jsonComputeResponse, &jsonStaticResponse.ServersResp); err != nil {
		fmt.Println("Error in Unmarshing:==", err)
		}*/
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
