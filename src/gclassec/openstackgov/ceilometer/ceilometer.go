package ceilometer

import (
	"gclassec/structs/openstackInstance"
	"fmt"
	"gclassec/errorcodes/errcode"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gclassec/loggers"
	"gclassec/openstackgov/authenticationtoken"
	"time"
	"strconv"
)

func GetCeilometerDetail(id string) ([]openstackInstance.MeterStruct,error) {
	logger := Loggers.New()
	var MeteringEndpoint string
	var authToken string
	var authError error
	var endpointsStruct openstackInstance.OpenStackEndpoints
	//var jsonceilometerStruct openstackInstance.MeterStruct
       fmt.Println("=====================Scoped Authentication Token====================")
	authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(false)
	fmt.Println("authToken:==", authToken)
	fmt.Println("endpointsStruct:==",endpointsStruct)
	fmt.Println("authError:==",authError)

	if authError!= nil {
		fmt.Println("authError:==",authError)
		return []openstackInstance.MeterStruct{},authError
	}else{
		fmt.Println("authToken:==", authToken)
		fmt.Println("endpointsStruct:==", endpointsStruct)
	}
	for i := 0; i < len(endpointsStruct.ApiEndpoints); i++ {
		if endpointsStruct.ApiEndpoints[i].EndpointType =="metering"{
			MeteringEndpoint = endpointsStruct.ApiEndpoints[i].EndpointURL
			fmt.Println("MeteringEndPoint:====",MeteringEndpoint)
			//https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c
			}
		}
	t := time.Now()

	var reqURL string =  MeteringEndpoint + "/v2/meters/cpu_util/statistics?q.field=resource_id&q.field=timestamp&q.op=eq&q.op=gt&q.type=&q.type=&q.value="+id+"&q.value="+strconv.Itoa(t.Year())+"-"+strconv.Itoa(int(t.Month()))+"-"+strconv.Itoa(t.Day())

	//var reqURL string = meteringEndpoint+ "/v2/meters/cpu_util/statistics?q.field=resource_id&q.field=timestamp&q.op=eq&q.op=gt&q.type=&q.type=&q.value="+id+"&q.value=2017-01-18T08%3A55%3A00"
	//http://110.110.112.10:8777/v2/meters/cpu_util/statistics?q.field=resource_id&q.field=timestamp&q.op=eq&q.op=gt&q.type=&q.type=&q.value=06b49c3e-3e80-4a22-82ac-20d43f1b9009&q.value=2017-01-18T08%3A55%3A00
	fmt.Println("reqURL:====",reqURL)

	fmt.Println("reqURL:====",reqURL)
	//var reqURL string = "https://120.120.120.4:8774/v2.1/cf5489c2c0d040c6907eeae1d7d2614c/flavors/detail"
	req, errReq := http.NewRequest("GET", reqURL, nil)
	if errReq != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		return []openstackInstance.MeterStruct{},errReq
	}

	req.Header.Add("x-auth-token", authToken)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
	if errClient != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		return []openstackInstance.MeterStruct{}, errClient
	}

	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
	if errResp != nil{
		fmt.Println("HOS: ", errcode.ErrResp)
 		return []openstackInstance.MeterStruct{}, errResp
	}


	logger.Info("respBody:==\n",respBody)
	respBodyInString:= string(respBody)

	fmt.Println("\nrespBodyInString:==\n",respBodyInString)
//	return respBodyInString, nil

	var jsonDynamicResponse []openstackInstance.MeterStruct
        if err := json.Unmarshal(respBody, &jsonDynamicResponse); err != nil {
                fmt.Println("Error in Unmarshing:==", err)
        }


	return jsonDynamicResponse, nil

}
// var cpu_util hosstruct.LatestDynamicData
//        for i:=0; i<len(jsonComputeResponse); i++{
//                        if i == len(jsonComputeResponse)-1{
//
////                                fmt.Println("\n----------------Latest Record",jsonComputeResponse[i])
//			          cpu_util = jsonComputeResponse[i]
//			//	return jsonComputeResponse[i]
//
//                        }
//
//                }
//
//	return cpu_util.Avg
//}


