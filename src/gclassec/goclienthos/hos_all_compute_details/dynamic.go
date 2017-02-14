package hos_all_compute_details
import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	//"hos/HOSAuthToken"
//	"hos/HOS_API_Function/comp"
	"gclassec/goclienthos/authtoken"
)

type AvgData struct {
	Avg float64 `json:"avg"`
	
}
type LatestDynamicData struct {
        Count int `json:"count"`
        DurationStart string `json:"duration_start"`
        Min float64 `json:"min"`
        DurationEnd string `json:"duration_end"`
        Max float64 `json:"max"`
        Sum float64 `json:"sum"`
        Period int `json:"period"`
        PeriodEnd string `json:"period_end"`
        Duration float64 `json:"duration"`
        PeriodStart string `json:"period_start"`
        Avg float64 `json:"avg"`
        Groupby Group `json:"groupby"`
        Unit string `json:"unit"`
 }



type DynamicData []struct {
        Count int `json:"count"`
        DurationStart string `json:"duration_start"`
        Min float64 `json:"min"`
        DurationEnd string `json:"duration_end"`
        Max float64 `json:"max"`
        Sum float64 `json:"sum"`
        Period int `json:"period"`
        PeriodEnd string `json:"period_end"`
        Duration float64 `json:"duration"`
        PeriodStart string `json:"period_start"`
        Avg float64 `json:"avg"`
        Groupby Group `json:"groupby"`
        Unit string `json:"unit"`
 }


type Group struct{
        GroupBy string `json:"groupby"`
}

type DynamicDataResponse struct {
        Data []DynamicData
}



func AvgCpuUtil(id string) float64{

	fmt.Println("\n\n================================================== for Id ::::::::::::    %s ======================",id)
	//fmt.Println("This to get Nothing")
	//var auth,_ = HOSAuthToken.GetHOSAuthToken()
	//fmt.Println("Auth Token in Compute.go:=====\n", auth)
	var meteringEndpoint string
	var auth, hosConfig = authtoken.GetHOSAuthToken()
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


        var jsonComputeResponse DynamicData
        if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
                fmt.Println("Error in Unmarshing:==", err)
        }
  //      fmt.Println("Printing Initial jsonComputeResponse ")
//        fmt.Printf("%+v\n\n", jsonComputeResponse)

	//fmt.Println("Printing Initial jsonComputeResponse ")
        //fmt.Printf("%+v\n\n", jsonComputeResponse)
        var cpu_util LatestDynamicData
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
