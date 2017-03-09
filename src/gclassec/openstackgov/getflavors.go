package openstackgov

import (
	"net/http"
	"gclassec/confmanagement/readcomputeVM"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"gclassec/structs/openstackInstance"
)

func Getflavors() openstackInstance.FlvRespStruct{
	var openstackcreds = readcomputeVM.Configurtion()

	url := openstackcreds.ComputeHost + "/flavors/detail"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-auth-token",GetAuth() )
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil{
		fmt.Errorf("Error : ", err)
	}
	//fmt.Println("res : ", res)
	defer res.Body.Close()
	body, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil{
		fmt.Errorf("Error : ", err1)
	}

	//fmt.Println("res : ", res)
	fmt.Println("string(body) : ", string(body))
	//_ = json.NewEncoder(w).Encode(res.Body)

	var jsonFlavorResponse openstackInstance.FlvRespStruct
	if err := json.Unmarshal(body, &jsonFlavorResponse); err != nil {
		fmt.Errorf("Error in Unmarshing:==", err)
	}

	_ = json.NewEncoder(os.Stdout).Encode(&jsonFlavorResponse)

	return jsonFlavorResponse

}
