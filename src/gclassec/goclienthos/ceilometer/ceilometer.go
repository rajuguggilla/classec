package ceilometer

import (
	"net/http"
	"io/ioutil"
	"gclassec/goclienthos/authtoken"
	"gclassec/loggers"
	"fmt"
	"gclassec/errorcodes/errcode"
)



type Response struct {
	Counter_name		string	`json:"counter_name"`
	User_id			string	`json:"user_id"`
	Resource_id		string	`json:"resource_id"`
	Timestamp		string	`json:"timestamp"`
	Recorded_at		string	`json:"recorded_at"`
	Message_id		string	`json:"message_id"`
	Source			string	`json:"source"`
	Counter_unit		string	`json:"counter_unit"`
	Counter_volume		string	`json:"counter_volume"`
	Project_id		string 	`json:"project_id"`
	Resource_metadata	string	`json:"resource_metadata"`
	Counter_type		string	`json:"counter_type"`

}


func GetCeilometerDetail() string{
	logger := Loggers.New()
	//fmt.Println("This to get Nothing")
	//var auth,_ = authtoken.GetHOSAuthToken()
	//fmt.Println("Auth Token in Compute.go:=====\n", auth)
	var meteringEndpoint string
	var auth, hosConfig, err = authtoken.GetHOSAuthToken()

	if err != nil{
		fmt.Println("HOS: ", errcode.ErrAuth)
		logger.Error("HOS : ", errcode.ErrAuth)
			return ""
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

	var reqURL string =  meteringEndpoint + "v2/meters"
	//var reqURL string =  "https://120.120.120.4:8777/v2/meters"
	req, errReq := http.NewRequest("GET", reqURL, nil)
		if errReq != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
			return ""
		}

	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, errClient := http.DefaultClient.Do(req)
		if errClient != nil{
		fmt.Println("HOS: ", errcode.ErrReq)
		logger.Error("HOS : ", errcode.ErrReq)
			return ""
		}
	logger.Info("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, errResp := ioutil.ReadAll(res.Body)
		if errResp != nil{
		fmt.Println("HOS: ", errcode.ErrResp)
		logger.Error("HOS : ", errcode.ErrResp)
			return ""
		}

	logger.Info("respBody:==\n",respBody)
	respBodyInString:= string(respBody)
//	fmt.Println("\nrespBodyInString:==\n",respBodyInString)

	return respBodyInString
}
