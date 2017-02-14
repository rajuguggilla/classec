package compute

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gclassec/goclienthos/authtoken"
	"gclassec/loggers"
	"gclassec/structs/hosstruct"
)

//type FlavorsStruct struct{
//
//	FlavorID 	string 		`json:"id"`
//	FlavorName 	string		`json:"name"`
//	Ram		int32		`json:"ram"`
//	VCPUS		int32		`json:"vcpus"`
//	Disk		int32		`json:"disk"`
//	//Links 		SubLinks	`json:"links"`
//
//}
//type FlvRespStruct struct {
//	Flavors []FlavorsStruct		`json:"flavors"`
//}


//func Flavors() string {
func Flavors() hosstruct.FlvRespStruct {
	//fmt.Println("This to get Nothing")
	logger := Loggers.New()
	var computeEndpoint string
	var auth, hosConfig = authtoken.GetHOSAuthToken()
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

	var reqURL string =  computeEndpoint + "/flavors/detail"

	//var reqURL string = "https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c/flavors/detail"
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
	var jsonFlavorResponse hosstruct.FlvRespStruct
	if err := json.Unmarshal(respBody, &jsonFlavorResponse); err != nil {
		logger.Error("Error in Unmarshing:==", err)
	}

	logger.Info("%+v\n\n", jsonFlavorResponse)
	return jsonFlavorResponse

}


