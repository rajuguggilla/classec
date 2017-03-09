package openstackgov


import (
	"fmt"
	"strings"
	"net/http"
	//"io/ioutil"
	"gclassec/confmanagement/readcomputeVM"
	"io/ioutil"
)

func GetAuth() (string) {

	var openstackcreds = readcomputeVM.Configurtion()

	UserName := openstackcreds.Username
	Password := openstackcreds.Password
	Domain := openstackcreds.ProjectName
	Tenant := openstackcreds.ProjectID

	url := openstackcreds.IdentityEndpoint
	var reqBody, reqURL string

	if (strings.Contains(url, "v3")) {
		reqURL = url + "auth/tokens"
		reqBody = `{"auth":{"identity":{"methods": ["password"], "password": {"user":{"name": "` + UserName + `","domain": { "name": "` + Domain + `"}, "password": "` + Password + `"}}}}}`
	} else {
		reqURL = url + "/tokens"
		reqBody = `{"auth":{"passwordCredentials":{"username": "` + UserName + `", "password": "` + Password + `"}, "tenantName": "` + Tenant + `"}}`
	}

	fmt.Println("Request Body:==", reqBody)
	fmt.Println("\nRequest URL:==", reqURL)

	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(reqBody))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error in Request Response:==", err)
	} else {
		fmt.Println("Status:==", res.Status)
		defer res.Body.Close()
	}
	respBody, _ := ioutil.ReadAll(res.Body)

	respBodyInString := string(respBody)
	fmt.Println("\nrespBodyInString:==\n", respBodyInString)

	var X_Subject_Token string = res.Header.Get("X-Subject-Token")
	fmt.Println("X-Subject-Token:===", X_Subject_Token)
	return X_Subject_Token
}


	/*//payload := strings.NewReader("{\r\n    \"auth\": {\r\n        \"identity\": {\r\n            \"methods\": [\r\n                \"password\"\r\n            ],\r\n            \"password\": {\r\n                \"user\": {\r\n                    \"name\": \"bhanu\",\r\n                    \"domain\": {\r\n                        \"name\": \"Default\"\r\n                    },\r\n                    \"password\": \"bhanu\"\r\n                }\r\n            }\r\n        }\r\n    }\r\n}")
//	req, _ := http.NewRequest("POST", url, payload)
//	req.Header.Add("content-type", "application/json")
	//req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("postman-token", "4760505b-881d-842c-4fa0-4fb365ab288a")
	//res, _ := http.DefaultClient.Do(req)
//	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.Header.Get("X-Subject-Token"))
	return res.Header.Get("X-Subject-Token")
	//fmt.Println(string(body))*/

