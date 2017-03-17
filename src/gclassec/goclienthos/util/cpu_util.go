package util


import (
	"net/http"
	"io/ioutil"
	"gclassec/goclienthos/authtoken"
	"gclassec/loggers"
	"gclassec/structs/hosstruct"
	"encoding/json"
	"fmt"
	"gclassec/errorcodes/errcode"
)

func GetCpuUtilDetails(id string) (hosstruct.DynamicData,error) {

	//fmt.Println("This to get Nothing")
	//var auth,_ = authtoken.GetHOSAuthToken()
	//fmt.Println("Auth Token in Compute.go:=====\n", auth)
	logger := Loggers.New()
	var meteringEndpoint string
	var auth, hosConfig, err = authtoken.GetHOSAuthToken()
		if err != nil{
			fmt.Println("HOS: ", errcode.ErrAuth)
			logger.Error("HOS : ", errcode.ErrAuth)
			return hosstruct.DynamicData{}, err
		}

	logger.Debug("HOS AuthToken:=====\n", auth)
	logger.Debug("HOS Configuration:=====\n %+v", hosConfig)
	for i := 0; i < len(hosConfig.Access.ServiceCatalog); i++ {
		if hosConfig.Access.ServiceCatalog[i].EndpointType =="metering"{
			//for j:= 0; j< len(hosConfig.Access.ServiceCatalog[i].Endpoints); j++ {
			meteringEndpoint = hosConfig.Access.ServiceCatalog[i].Endpoints[0].PublicURL
			logger.Info("MeetringEndPoint:====",meteringEndpoint)
			//https://120.120.120.4:8777/
					//}
			}
		}

	var reqURL string =  meteringEndpoint + "v2/meters/cpu_util/statistics?q.field=resource_id&q.field=timestamp&q.op=eq&q.op=gt&q.type=&q.type=&q.value="+id+"&q.value=2017-01-18T08%3A55%3A00"

	//var reqURL string =  "https://120.120.120.4:8777/v2/meters/cpu_util/statistics?q.field=resource_id&q.field=timestamp&q.op=eq&q.op=gt&q.type=&q.type=&q.value="+id+"&q.value=2017-01-18T08%3A55%3A00"


	req, errReq := http.NewRequest("GET", reqURL, nil)
		if errReq != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
		return hosstruct.DynamicData{}, err
		}


	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
		if errClient != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
		return hosstruct.DynamicData{}, err
		}

	logger.Info("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
		if errResp != nil{
		fmt.Println("HOS: ", errcode.ErrResp)
		logger.Error("HOS : ", errcode.ErrResp)
		return hosstruct.DynamicData{}, err
		}


	logger.Info("respBody:==\n",respBody)
	//respBodyInString:= string(respBody)

	//fmt.Println("\nrespBodyInString:==\n",respBodyInString)
	//return respBodyInString

	var jsonComputeResponse hosstruct.DynamicData
        if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
                fmt.Println("Error in Unmarshing:==", err)
        }


	return jsonComputeResponse, nil


}
