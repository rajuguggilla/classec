package OpenStack

import (
       "fmt"
       "net/http"
       "io/ioutil"
       "gclassec/goclientopenstack/GetAuthToken"
       "encoding/json"
)
type FlavorResponse struct{
       Flavors              []FlavorStruct              `json:"flavors"`
}
type FlavorStruct struct{
	Name                string                      `json:"name"`
	RAM                 int64                       `json:"ram"`
	VCPU                int64                       `json:"vcpus"`
	Disk                int64                       `json:"disk"`
	FlavorID            string                      `json:"id"`
}
func Flavor() FlavorResponse {
       var auth = GetAuthToken.GetOpenStackAuthToken()
       fmt.Println("Auth Token in Flavor.go:=====\n", auth)

       var reqURL string = "http://110.110.110.5:8774/v2/99e1e2d5093446f1b5ae11939272c2df/flavors/detail"
       req, _ := http.NewRequest("GET", reqURL, nil)
       req.Header.Add("x-auth-token", auth)
       req.Header.Add("content-type", "application/json")

       res, _ := http.DefaultClient.Do(req)
       fmt.Println("Status : ", res.Status)
       defer res.Body.Close()
       respBody, _ := ioutil.ReadAll(res.Body)

       respBodyInString:= string(respBody)
       fmt.Println("\nrespBodyInString:==\n",respBodyInString)
       var jsonFlavorResponse FlavorResponse
       if err := json.Unmarshal(respBody, &jsonFlavorResponse); err != nil {
              fmt.Println("Error in unmarshing:==",err)
       }

       fmt.Printf("%+v\n\n", jsonFlavorResponse)
       return jsonFlavorResponse
}