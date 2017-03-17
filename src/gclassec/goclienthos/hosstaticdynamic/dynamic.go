package hosstaticdynamic
import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	//"hos/HOSAuthToken"
//	"hos/HOS_API_Function/comp"
	"gclassec/goclienthos/authtoken"
	"gclassec/errorcodes/errcode"
	"gclassec/loggers"
	"gclassec/structs/hosstruct"
)




func AvgCpuUtil(id string) float64{
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
			return 0
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


	req, errReq := http.NewRequest("GET", reqURL, nil)
	if errReq != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
		return 0
		}

	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
		if errClient != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
		return 0
		}


//	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
		if errResp != nil{
		fmt.Println("HOS: ", errcode.ErrResp)
		logger.Error("HOS : ", errcode.ErrResp)
		return 0
		}

 
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
                        if i == len(jsonComputeResponse)-1{

//                                fmt.Println("\n----------------Latest Record",jsonComputeResponse[i])
			          cpu_util = jsonComputeResponse[i]
			//	return jsonComputeResponse[i]

                        }

                }

	return cpu_util.Avg

//       res1 := HOS.Compute()
  //    fmt.Println("\n\n===Compute =====",res1)

// return jsonComputeResponse

}
