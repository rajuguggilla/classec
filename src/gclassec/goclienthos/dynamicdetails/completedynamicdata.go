package dynamicdetails

import (
	"encoding/json"
	"fmt"
	"gclassec/loggers"
	"gclassec/goclienthos/authtoken"
	"gclassec/errorcodes/errcode"
	"net/http"
	"io/ioutil"
	"gclassec/structs/hosstruct"
)

func CpuUtil(id string) hosstruct.LatestDynamicData{
	logger := Loggers.New()
	fmt.Println("\n\n================================================== for Id ::::::::::::    %s ======================",id)
	//fmt.Println("This to get Nothing")
	//var auth,_ = HOSAuthToken.GetHOSAuthToken()
	//fmt.Println("Auth Token in Compute.go:=====\n", auth)
	var meteringEndpoint string
	var auth, hosConfig, err = authtoken.GetHOSAuthToken()

		if err != nil{
			fmt.Println("HOS : ", errcode.ErrAuth)
			logger.Error("HOS : ", errcode.ErrAuth)
			return hosstruct.LatestDynamicData{}
		}
//	fmt.Println("HOS AuthToken:=====\n", auth)
//	fmt.Println("HOS Configuration:=====\n %+v", hosConfig)
	for i := 0; i < len(hosConfig.Access.ServiceCatalog); i++ {
		if hosConfig.Access.ServiceCatalog[i].EndpointType =="metering"{
			//for j:= 0; j< len(hosConfig.Access.ServiceCatalog[i].Endpoints); j++ {
			meteringEndpoint = hosConfig.Access.ServiceCatalog[i].Endpoints[0].PublicURL
//			fmt.Println("ComputeeNDpOINT:====",meteringEndpoint)
			//https://120.120.120.4:8777/
					//}
			}
		}

	var reqURL string =  meteringEndpoint + "v2/meters/util/statistics?q.field=resource_id&q.field=timestamp&q.op=eq&q.op=gt&q.type=&q.type=&q.value="+id+"&q.value=2017-02-09"

	//var reqURL string =  "https://120.120.120.4:8777/v2/meters/cpu_util/statistics?q.field=resource_id&q.field=timestamp&q.op=eq&q.op=gt&q.type=&q.type=&q.value="+id+"&q.value=2017-01-18T08%3A55%3A00"


	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
//	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)

//	fmt.Print("respBody:==\n",respBody)
//	respBodyInString:= string(respBody)
//	fmt.Println("\nrespBodyInString:==\n",respBodyInString)
//	return respBodyInString


        var jsonComputeResponse hosstruct.DynamicData
        if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
                fmt.Println("Error in Unmarshing:==", err)
        }
  //      fmt.Println("Printing Initial jsonComputeResponse ")
//        fmt.Printf("%+v\n\n", jsonComputeResponse)

	//fmt.Println("Printing Initial jsonComputeResponse ")
        //fmt.Printf("%+v\n\n", jsonComputeResponse)
        var cpu_util hosstruct.LatestDynamicData
        for i:=0; i<len(jsonComputeResponse); i++{
		// To get latest data (last)
                        if i == len(jsonComputeResponse)-1{

//                                fmt.Println("\n----------------Latest Record",jsonComputeResponse[i])
			          cpu_util = jsonComputeResponse[i]
			//	return jsonComputeResponse[i]

                        }

                }

	return cpu_util

//       res1 := HOS.Compute()
  //    fmt.Println("\n\n===Compute =====",res1)

// return jsonComputeResponse

}