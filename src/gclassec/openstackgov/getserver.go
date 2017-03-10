package openstackgov


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"gclassec/confmanagement/readcomputeVM"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"gclassec/structs/openstackInstance"
)



func Getserver(w http.ResponseWriter, r *http.Request)  {
	var openstackcreds = readcomputeVM.Configurtion()

	url := openstackcreds.ComputeHost + "/servers/detail"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-auth-token",GetAuth() )
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil{
		fmt.Errorf("Error : ", err)
	}
	fmt.Println("res : ", res)
	defer res.Body.Close()
	body, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil{
		fmt.Errorf("Error : ", err1)
	}

	fmt.Println("res : ", res)
	fmt.Println("string(body) : ", string(body))
	//_ = json.NewEncoder(w).Encode(res.Body)

	var jsonComputeResponse openstackInstance.ComputeResponse
	if err := json.Unmarshal(body, &jsonComputeResponse); err != nil {
		fmt.Errorf("Error in Unmarshing:==", err)
	}

	var FlavorsList openstackInstance.FlvRespStruct
	FlavorsList = Getflavors()

	fmt.Println("FlavorsList : ", FlavorsList)

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

	_ = json.NewEncoder(w).Encode(&jsonComputeResponse)

}