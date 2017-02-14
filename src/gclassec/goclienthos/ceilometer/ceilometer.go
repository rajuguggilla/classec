package ceilometer

import (
	"net/http"
	"io/ioutil"
	"gclassec/goclienthos/authtoken"
	"gclassec/loggers"
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
	var auth, hosConfig = authtoken.GetHOSAuthToken()
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
	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("x-auth-token", auth)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	logger.Info("Status:======== ", res.Status)
	defer res.Body.Close()
	respBody, _ := ioutil.ReadAll(res.Body)

	logger.Info("respBody:==\n",respBody)
	respBodyInString:= string(respBody)
//	fmt.Println("\nrespBodyInString:==\n",respBodyInString)

	return respBodyInString
}
