package openstackgov


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"gclassec/confmanagement/readcomputeVM"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

type ComputeResponse struct {
	Servers       []ServerResponse      `json:"Servers"`
}

type ServerResponse struct {
	Id	string		`json:"Id"`
	Name	string		`json:"Name"`
}

func Getserver(w http.ResponseWriter, r *http.Request)  {
	var openstackcreds = readcomputeVM.Configurtion()

	url := openstackcreds.ComputeHost

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-auth-token",GetAuth() )
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil{
		fmt.Errorf("Error : ", err)
	}
	fmt.Println("res : ", res)
	defer res.Body.Close()
	body, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil{
		fmt.Errorf("Error : ", err1)
	}

	fmt.Println("res : ", res)
	fmt.Println("string(body) : ", string(body))
	//_ = json.NewEncoder(w).Encode(res.Body)

	var jsonComputeResponse ComputeResponse
	if err := json.Unmarshal(body, &jsonComputeResponse); err != nil {
		fmt.Errorf("Error in Unmarshing:==", err)
	}

	_ = json.NewEncoder(w).Encode(&jsonComputeResponse)

}