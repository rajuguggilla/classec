package openstackgov

import (

	"gclassec/confmanagement/readcomputeVM"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"net/http"
	"strings"
	"io/ioutil"

	"github.com/gorilla/mux"
)

func Createserver(w http.ResponseWriter, r *http.Request)  {
	var openstackcreds = readcomputeVM.Configurtion()


	imageref := "957edc17-b1f1-4ec3-9d24-988a438d9de9"
       	flavorref := "48d0ccdc-c80e-4752-bcdf-e44873528a4d"
       	availability := "nova"
	reqURL := openstackcreds.ComputeHost + "/servers"
	vars := mux.Vars(r)
      	vmName := vars["instancename"]

	var reqBody string = `{"server":{"name": "`+vmName+`", "imageRef": "`+ imageref +`", "flavorRef": "`+flavorref +`", "availability_zone": "`+availability+`"}}`


	fmt.Println("Request Body:==", reqBody)
	fmt.Println("\nRequest URL:==", reqURL)
	req, _ := http.NewRequest("POST", reqURL, strings.NewReader(reqBody))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("x-auth-token", GetAuth() )


	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error in Request Response:==", err)
	} else {
		fmt.Println("Status:==", res.Status)
		defer res.Body.Close()
	}
	respBody, _ := ioutil.ReadAll(res.Body)

	respBodyInString:= string(respBody)
	fmt.Println("\nrespBodyInString:==\n",respBodyInString)
	rBodyInByte := []byte(respBody)
	fmt.Println("respBodyInByte",rBodyInByte)

}