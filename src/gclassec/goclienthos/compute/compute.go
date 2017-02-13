package compute

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gclassec/goclienthos/authtoken"
)


type ComputeResponse struct {
      Servers       []ServersResponse      `json:"servers"`
}

type ServersResponse struct {
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
       Flavor                 FlavorsStruct `json:"flavor"`
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
func Compute() ComputeResponse {


	//fmt.Println("This to get Nothing")
	var computeEndpoint string
	var auth, hosConfig = authtoken.GetHOSAuthToken()
	fmt.Println("HOS AuthToken:=====\n", auth)
	fmt.Println("HOS Configuration:=====\n %+v", hosConfig)
	for i := 0; i < len(hosConfig.Access.ServiceCatalog); i++ {
		if hosConfig.Access.ServiceCatalog[i].EndpointType =="compute"{
			//for j:= 0; j< len(hosConfig.Access.ServiceCatalog[i].Endpoints); j++ {
			computeEndpoint = hosConfig.Access.ServiceCatalog[i].Endpoints[0].PublicURL
			fmt.Println("ComputeEndPoint:====",computeEndpoint)
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
	respBodyInString:= string(respBody)
	fmt.Println("\nrespBodyInString:==\n",respBodyInString)
	//return respBodyInString
	var jsonComputeResponse ComputeResponse
	if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
		fmt.Println("Error in Unmarshing:==", err)
	}
	fmt.Println("Printing Initial jsonComputeResponse ")
	fmt.Printf("%+v\n\n", jsonComputeResponse)
	//return jsonComputeResponse

	var FlavorsList FlvRespStruct
	FlavorsList = Flavors()
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
	TempStr, _ := json.Marshal(&jsonComputeResponse)
	fmt.Println("Printing Final jsonComputeResponse in string:===\n\n ",string(TempStr))
	return jsonComputeResponse
	//return string(TempStr)

}
