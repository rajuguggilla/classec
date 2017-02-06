package OpenStack

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gclassec/goclientopenstack/GetAuthToken"
)
type ComputeResponse struct {
       Servers			[]ServersResponse	`json:"servers"`
}
type ServersResponse struct {
	Name			string			`json:"name"`
	ID			string			`json:"id"`
	Status			string			`json:"status"`
	Availability_zone	string			`json:"OS-EXT-AZ:availability_zone"`
	Created			string			`json:"created"`
	Flavor			FlavorStruct		`json:"flavor"`
	Addresses		address			`json:"addresses"`
	Security_groups		[]security_groups	`json:"security_groups"`
	Key_name		string 			`json:"key_name"`
	Image			Image			`json:"image"`
	Tenant_id		string			`json:"tenant_id"`
	Updated			string			`json:"updated"`
	User_id			string			`json:"user_id"`
	HostId			string			`json:"hostId"`
	Task_state		string			`json:"OS-EXT-STS:task_state"`
	Vm_state		string			`json:"OS-EXT-STS:vm_state"`
	Launched_at		string			`json:"OS-SRV-USG:launched_at"`
	Volumes_attached	[]Volume		`json:"os-extended-volumes:volumes_attached"`
	Progress		int64			`json:"progress"`
	IPV4			string			`json:"accessIPv4"`
	IPV6			string			`json:"accessIPv6"`
	Power_State		int64			`json:"OS-EXT-STS:power_state"`
	ConfigDrive		string			`json:"config_drive"`
	DiskConfig		string			`json:"OS-DCF:diskConfig"`
}

type Volume struct{
	Vol			string			`json:"Os-extended-volumes:volumes_attached"`
}

type Image struct{
	ID			string			`json:"Id"`
}

type security_groups struct{
	Name			string			`json:"name"`
}

type address struct{
	Mac_addr		string			`json:"OS-EXT-IPS-MAC:mac_addr"`
	Version			string			`json:"version"`
	Addr			string			`json:"addr"`
	Type			string			`json:"OS-EXT-IPS:type"`
}

func Compute() ComputeResponse {
       flavor := Flavor()
       fmt.Println("Flavor Output", flavor)
       var auth = GetAuthToken.GetOpenStackAuthToken()

       var reqURL string = "http://110.110.110.5:8774/v2/99e1e2d5093446f1b5ae11939272c2df/servers/detail"
       req, _ := http.NewRequest("GET", reqURL, nil)
       req.Header.Add("x-auth-token", auth)
       req.Header.Add("content-type", "application/json")

       res, _ := http.DefaultClient.Do(req)
       defer res.Body.Close()
       respBody, _ := ioutil.ReadAll(res.Body)

       respBodyInString:= string(respBody)
       fmt.Println("\nrespBodyInString:==\n",respBodyInString)
       var jsonComputeResponse ComputeResponse
       if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
              fmt.Println("Error in unmarshing:==",err)
       }

       fmt.Printf("%+v\n\n", jsonComputeResponse)
       return jsonComputeResponse
}