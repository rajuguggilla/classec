package hosstaticdynamic

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gclassec/goclienthos/authtoken"
	"gclassec/goclienthos/compute"
	"gclassec/structs/hosstruct"
	"gclassec/errorcodes/errcode"
	"gclassec/loggers"
	"gclassec/structs/tagstruct"
	"regexp"
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/dbmanagement"
	"gclassec/goclienthos/util"
)



var dbtype string = dbmanagement.ENVdbtype
var dbname  string = dbmanagement.ENVdbnamegodb
var dbusername string = dbmanagement.ENVdbusername
var dbpassword string = dbmanagement.ENVdbpassword
var dbhostname string = dbmanagement.ENVdbhostname
var dbport string = dbmanagement.ENVdbport

var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))

var db,err  = gorm.Open(dbtype, c)


func ComputeWithCPU() hosstruct.CompleteComputeResponse{
	logger := Loggers.New()
	//fmt.Println("This to get Nothing")
	tx := db.Begin()
	db.SingularTable(true)

	tag := []tagstruct.Providers{}

	//create a regex `(?i)hos` will match string contains "hos" case insensitive
	reg := regexp.MustCompile("(?i)hos")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("hos")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("hos")).Find(&tag)

	var computeEndpoint string
	var auth, hosConfig, err = authtoken.GetHOSAuthToken()

		if err != nil{
			fmt.Println("HOS : ", errcode.ErrAuth)
			logger.Error("HOS : ", errcode.ErrAuth)
			return hosstruct.CompleteComputeResponse{}
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
		return hosstruct.CompleteComputeResponse{}
		}

	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
		if errClient != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
		return hosstruct.CompleteComputeResponse{}
		}

	fmt.Println("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
		if errResp != nil{
		fmt.Println("HOS: ", errcode.ErrResp)
		logger.Error("HOS : ", errcode.ErrResp)
		return hosstruct.CompleteComputeResponse{}
		}

	var jsonComputeResponse hosstruct.CompleteComputeResponse

	if err := json.Unmarshal(respBody, &jsonComputeResponse); err != nil {
		fmt.Println("Error in Unmarshing:==", err)
	}



	var FlavorsList hosstruct.FlvRespStruct
	FlavorsList = compute.Flavors()

	fmt.Printf("%+v\n\n", FlavorsList)


	for i:=0; i<len(jsonComputeResponse.Servers);i++{
		tempFID := jsonComputeResponse.Servers[i].Flavor.FlavorID
		for j:=0; j<len(FlavorsList.Flavors);j++ {
			if tempFID == FlavorsList.Flavors[j].FlavorID{
				jsonComputeResponse.Servers[i].Flavor.FlavorName=FlavorsList.Flavors[j].FlavorName
				jsonComputeResponse.Servers[i].Flavor.Disk=FlavorsList.Flavors[j].Disk
				jsonComputeResponse.Servers[i].Flavor.Ram=FlavorsList.Flavors[j].Ram
				jsonComputeResponse.Servers[i].Flavor.VCPUS=FlavorsList.Flavors[j].VCPUS

			}

		}

	}

	fmt.Println("Printing Final jsonComputeResponse ")
	fmt.Printf("%+v\n\n", jsonComputeResponse)



	fmt.Println("\n\n ++++++++++++++++++++++++++++++ End of Static +++++++++++++++++++++++++++++\n\n")
/*
	for k:=0; k<len(jsonComputeResponse.Servers);k++{
		temID := jsonComputeResponse.Servers[k].InstanceID
		dynamicData := AvgCpuUtil(temID)
		//fmt.Println("dynamicData)
		fmt.Println("\n\n--------------------------------------------\n\n")
		fmt.Println("@@@@@@@@@@@@@@@@@@---  DYNAMIC @@@@@@@@@@@@@",dynamicData)
		//var jsonStaticResponse MyStruct
		//jsonStaticResponse.ServersResp = jsonStaticResponse1.Servers[k]
		//jsonStaticResponse.Cpu_Util = dynamicData
		//jsonStaticResponses = append(jsonStaticResponses,jsonStaticResponse)
		jsonComputeResponse.Servers[k].Cpu_Util=dynamicData
	}*/


	for k:=0; k<len(jsonComputeResponse.Servers);k++ {
		temID := jsonComputeResponse.Servers[k].InstanceID
		DynamicData, err := util.GetCpuUtilDetails(temID)
		if err != nil {
			fmt.Println("Error", err)
		}
		for element1:=0;element1<len(DynamicData);element1++ {
			if element1 == len(DynamicData) - 1 {
				element := DynamicData[element1]


				jsonComputeResponse.Servers[k].Cpu_Util = element.Avg

			}
		}
	}

	for j := 0; j < len(jsonComputeResponse.Servers); j++{
		if len(tag) == 0 {
			jsonComputeResponse.Servers[j].Tagname = "Nil"
		}else {
			for _, el := range tag {
				if jsonComputeResponse.Servers[j].InstanceID != el.InstanceId{
					jsonComputeResponse.Servers[j].Tagname = "Nil"
				}else {
					jsonComputeResponse.Servers[j].Tagname = el.Tagname
				}
			}
		}
	}


	fmt.Println("\n\n++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n")

	fmt.Printf("%+v",jsonComputeResponse)
	tx.Commit()
	return jsonComputeResponse


}
