package HOS_API_Function

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"gclassec/goclienthos/HOSAuthToken"
	"encoding/json"
)

type FlavorsStruct struct{

	FlavorID 	string 		`json:"id"`
	FlavorName 	string		`json:"name"`
	Ram		int32		`json:"ram"`
	VCPUS		int32		`json:"vcpus"`
	Disk		int32		`json:"disk"`
	//Links 		SubLinks	`json:"links"`

}
type FlvRespStruct struct {
	Flavors []FlavorsStruct		`json:"flavors"`
}


//func Flavors() string {
func Flavors() FlvRespStruct {
	//fmt.Println("This to get Nothing")
	var computeEndpoint string
	var auth, hosConfig = HOSAuthToken.GetHOSAuthToken()
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

	var reqURL string =  computeEndpoint + "/flavors/detail"

	//var reqURL string = "https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c/flavors/detail"
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
	var jsonFlavorResponse FlvRespStruct
	if err := json.Unmarshal(respBody, &jsonFlavorResponse); err != nil {
		fmt.Println("Error in Unmarshing:==", err)
	}

	fmt.Printf("%+v\n\n", jsonFlavorResponse)
	return jsonFlavorResponse

}


