package ceilometer

import (
	"gclassec/structs/openstackInstance"
	"fmt"
	"gclassec/openstackgov/authenticationtoken"
	"gclassec/openstackgov"
)

func DynamicDetails()(openstackInstance.CompleteDynamicResponse,error) {


      //  var computeEndpoint string
	var authToken string
	var authError error
	var endpointsStruct openstackInstance.OpenStackEndpoints
	//var jsonCompStruct openstackInstance.ComputeListStruct

	var tempCompStruct openstackInstance.ComputeListStruct
	var compError error
	//logger.Info("---------------------------------------Flavours Details in Compute Start------------------------------------------------------")
	tempCompStruct,compError = openstackgov.ListComputeInstances()
	if compError != nil{
		fmt.Println("Error In getting Flavours Details:==",compError)
	}else{

		fmt.Println(tempCompStruct)
	}


	fmt.Println("=====================Scoped Authentication Token====================")
	authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(false)
	fmt.Println("authToken:==", authToken)
	fmt.Println("endpointsStruct:==",endpointsStruct)
	fmt.Println("authError:==",authError)

	if authError!= nil {
		fmt.Println("authError:==",authError)
		return openstackInstance.CompleteDynamicResponse{},authError
	}else{
		fmt.Println("authToken:==", authToken)
		fmt.Println("endpointsStruct:==", endpointsStruct)
	}


	var jsonDynamicResponse openstackInstance.CompleteDynamicResponse

	//if err := json.Unmarshal(respBody, &jsonDynamicResponse); err != nil {
	//	fmt.Println("Error in Unmarshing:==", err)
	//}

	Templen :=len(tempCompStruct.Servers)
	for i:=0;i<Templen;i++ {
		//logger.Info("Value of I:==",i)
		jsonDynamicResponse.Servers = append(jsonDynamicResponse.Servers, openstackInstance.DynamicInstances_db{})
		jsonDynamicResponse.Servers[i].Vm_Name = tempCompStruct.Servers[i].ServerName
		jsonDynamicResponse.Servers[i].InstanceID = tempCompStruct.Servers[i].ServerId
		//temp2:= len(jsonAuthTokenBody.Token.Catalog[i].Endpoints)
		//logger.Info("Length of jsonAuthTokenBody.Token.Catalog:==", temp2)

	}

	fmt.Println("Printing Final jsonComputeResponse ")
	fmt.Printf("%+v\n\n", jsonDynamicResponse)



	fmt.Println("\n\n ++++++++++++++++++++++++++++++ End of Static +++++++++++++++++++++++++++++\n\n")

	for k:=0; k<len(jsonDynamicResponse.Servers);k++ {
		temID := jsonDynamicResponse.Servers[k].InstanceID
		DynamicData, err := GetCeilometerDetail(temID)
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



























//	var jsonDynamicResponse hosstruct.CompleteDynamicResponse
//
//	if err := json.Unmarshal(respBody, &jsonDynamicResponse); err != nil {
//		fmt.Println("Error in Unmarshing:==", err)
//	}
//
//
///*	staticDetail := compute.Compute()
//	tempStruct := new(hosstruct.CompleteDynamicResponse)
//
//
//
//	for k := 0; k < len(staticDetail.Servers); k++{
//		tempStruct.Servers[k].InstanceID = append(staticDetail.Servers[k].InstanceID)
//		tempStruct.Servers[k].Vm_Name = append(staticDetail.Servers[k].Vm_Name)
//
//	}*/
//
//
//	fmt.Println("Printing Final jsonComputeResponse ")
//	fmt.Printf("%+v\n\n", jsonDynamicResponse)
//
//
//
//	fmt.Println("\n\n ++++++++++++++++++++++++++++++ End of Static +++++++++++++++++++++++++++++\n\n")
//
//	for k:=0; k<len(jsonDynamicResponse.Servers);k++ {
//		temID := jsonDynamicResponse.Servers[k].InstanceID
//		DynamicData, err := util.GetCpuUtilDetails(temID)
//		if err != nil {
//			fmt.Println("Error", err)
//		}
//		for element1:=0;element1<len(DynamicData);element1++ {
//			if element1 == len(DynamicData) - 1 {
//				element := DynamicData[element1]
//
//
//				jsonDynamicResponse.Servers[k].Count = element.Count
//				jsonDynamicResponse.Servers[k].DurationStart = element.DurationStart
//				jsonDynamicResponse.Servers[k].Min = element.Min
//				jsonDynamicResponse.Servers[k].DurationEnd = element.DurationEnd
//				jsonDynamicResponse.Servers[k].Max = element.Max
//				jsonDynamicResponse.Servers[k].Sum = element.Sum
//				jsonDynamicResponse.Servers[k].Period = element.Period
//				jsonDynamicResponse.Servers[k].PeriodEnd = element.PeriodEnd
//				jsonDynamicResponse.Servers[k].Duration = element.Duration
//				jsonDynamicResponse.Servers[k].PeriodStart = element.PeriodStart
//				jsonDynamicResponse.Servers[k].Avg = element.Avg
//				jsonDynamicResponse.Servers[k].Groupby = element.Groupby
//				jsonDynamicResponse.Servers[k].Unit = element.Unit
//			}
//		}
//	}
//
//
//
//
//		//fmt.Println("dynamicData)
//		//fmt.Println("\n\n--------------------------------------------\n\n")
//		//fmt.Println("@@@@@@@@@@@@@@@@@@---  DYNAMIC @@@@@@@@@@@@@",DynamicData)
//		//var jsonStaticResponse MyStruct
//		//jsonStaticResponse.ServersResp = jsonStaticResponse1.Servers[k]
//		//jsonStaticResponse.Cpu_Util = dynamicData
//		/*//jsonStaticResponses = append(jsonStaticResponses,jsonStaticResponse)
//		jsonDynamicResponse.Servers[k].Count = DynamicData.Count
//		jsonDynamicResponse.Servers[k].DurationStart = DynamicData.DurationStart
//		jsonDynamicResponse.Servers[k].Min = DynamicData.Min
//		jsonDynamicResponse.Servers[k].DurationEnd = DynamicData.DurationEnd
//		jsonDynamicResponse.Servers[k].Max = DynamicData.Max
//		jsonDynamicResponse.Servers[k].Sum = DynamicData.Sum
//		jsonDynamicResponse.Servers[k].Period = DynamicData.Period
//		jsonDynamicResponse.Servers[k].PeriodEnd = DynamicData.PeriodEnd
//		jsonDynamicResponse.Servers[k].Duration = DynamicData.Duration
//		jsonDynamicResponse.Servers[k].PeriodStart = DynamicData.PeriodStart
//		jsonDynamicResponse.Servers[k].Avg = DynamicData.Avg
//		jsonDynamicResponse.Servers[k].Groupby = DynamicData.Groupby
//		jsonDynamicResponse.Servers[k].Unit = DynamicData.Unit*/
//
//
//
//
//	//}
//
//
//	fmt.Println("\n\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")
//	fmt.Printf("%+v",jsonDynamicResponse)
//	return jsonDynamicResponse, nil
//}
//

