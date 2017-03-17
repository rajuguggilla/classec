package dynamicdetails


import (
	"fmt"
	"gclassec/loggers"
	//"gclassec/structs/tagstruct"
	"gclassec/goclienthos/authtoken"
	"gclassec/errorcodes/errcode"
	"net/http"
	"io/ioutil"
	"encoding/json"
//	"gclassec/goclienthos/compute"
	"gclassec/structs/hosstruct"
	"gclassec/goclienthos/util"
)

func DynamicDetails()(hosstruct.CompleteDynamicResponse,error) {
	logger := Loggers.New()


	var computeEndpoint string
	var auth, hosConfig, err = authtoken.GetHOSAuthToken()

		if err != nil{
			fmt.Println("HOS : ", errcode.ErrAuth)
			logger.Error("HOS : ", errcode.ErrAuth)
			return hosstruct.CompleteDynamicResponse{}, err
		}
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

	var reqURL string =  computeEndpoint + "/servers/detail"
	//var reqURL string = "http://" + hosConfiguration.KeystoneEndpointIP + ":8774/v2.1/" + hosConfiguration.TenantId + "/servers/detail"
	fmt.Println("Request Body:==",reqURL)

	req, errReq := http.NewRequest("GET", reqURL, nil)
	if errReq != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
		return hosstruct.CompleteDynamicResponse{}, err
		}


	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
	if errClient != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
		return hosstruct.CompleteDynamicResponse{}, err
		}


	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
		if errResp != nil{
		fmt.Println("HOS: ", errcode.ErrResp)
		logger.Error("HOS : ", errcode.ErrResp)
		return hosstruct.CompleteDynamicResponse{}, err
		}



	var jsonDynamicResponse hosstruct.CompleteDynamicResponse

	if err := json.Unmarshal(respBody, &jsonDynamicResponse); err != nil {
		fmt.Println("Error in Unmarshing:==", err)
	}


/*	staticDetail := compute.Compute()
	tempStruct := new(hosstruct.CompleteDynamicResponse)



	for k := 0; k < len(staticDetail.Servers); k++{
		tempStruct.Servers[k].InstanceID = append(staticDetail.Servers[k].InstanceID)
		tempStruct.Servers[k].Vm_Name = append(staticDetail.Servers[k].Vm_Name)

	}*/


	fmt.Println("Printing Final jsonComputeResponse ")
	fmt.Printf("%+v\n\n", jsonDynamicResponse)



	fmt.Println("\n\n ++++++++++++++++++++++++++++++ End of Static +++++++++++++++++++++++++++++\n\n")

	for k:=0; k<len(jsonDynamicResponse.Servers);k++ {
		temID := jsonDynamicResponse.Servers[k].InstanceID
		DynamicData, err := util.GetCpuUtilDetails(temID)
		if err != nil {
			fmt.Println("Error", err)
		}
		for element1:=0;element1<len(DynamicData);element1++ {
			if element1 == len(DynamicData) - 1 {
				element := DynamicData[element1]


				jsonDynamicResponse.Servers[k].Count = element.Count
				jsonDynamicResponse.Servers[k].DurationStart = element.DurationStart
				jsonDynamicResponse.Servers[k].Min = element.Min
				jsonDynamicResponse.Servers[k].DurationEnd = element.DurationEnd
				jsonDynamicResponse.Servers[k].Max = element.Max
				jsonDynamicResponse.Servers[k].Sum = element.Sum
				jsonDynamicResponse.Servers[k].Period = element.Period
				jsonDynamicResponse.Servers[k].PeriodEnd = element.PeriodEnd
				jsonDynamicResponse.Servers[k].Duration = element.Duration
				jsonDynamicResponse.Servers[k].PeriodStart = element.PeriodStart
				jsonDynamicResponse.Servers[k].Avg = element.Avg
				jsonDynamicResponse.Servers[k].Groupby = element.Groupby
				jsonDynamicResponse.Servers[k].Unit = element.Unit
			}
		}
	}




		//fmt.Println("dynamicData)
		//fmt.Println("\n\n--------------------------------------------\n\n")
		//fmt.Println("@@@@@@@@@@@@@@@@@@---  DYNAMIC @@@@@@@@@@@@@",DynamicData)
		//var jsonStaticResponse MyStruct
		//jsonStaticResponse.ServersResp = jsonStaticResponse1.Servers[k]
		//jsonStaticResponse.Cpu_Util = dynamicData
		/*//jsonStaticResponses = append(jsonStaticResponses,jsonStaticResponse)
		jsonDynamicResponse.Servers[k].Count = DynamicData.Count
		jsonDynamicResponse.Servers[k].DurationStart = DynamicData.DurationStart
		jsonDynamicResponse.Servers[k].Min = DynamicData.Min
		jsonDynamicResponse.Servers[k].DurationEnd = DynamicData.DurationEnd
		jsonDynamicResponse.Servers[k].Max = DynamicData.Max
		jsonDynamicResponse.Servers[k].Sum = DynamicData.Sum
		jsonDynamicResponse.Servers[k].Period = DynamicData.Period
		jsonDynamicResponse.Servers[k].PeriodEnd = DynamicData.PeriodEnd
		jsonDynamicResponse.Servers[k].Duration = DynamicData.Duration
		jsonDynamicResponse.Servers[k].PeriodStart = DynamicData.PeriodStart
		jsonDynamicResponse.Servers[k].Avg = DynamicData.Avg
		jsonDynamicResponse.Servers[k].Groupby = DynamicData.Groupby
		jsonDynamicResponse.Servers[k].Unit = DynamicData.Unit*/




	//}


	fmt.Println("\n\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
	fmt.Printf("%+v",jsonDynamicResponse)
	return jsonDynamicResponse, nil
}
