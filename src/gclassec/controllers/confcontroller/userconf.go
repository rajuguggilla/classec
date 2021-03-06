package confcontroller

import (
	"net/http"
	"fmt"
	"os"
	"bufio"
	"runtime"
	"strings"
	"io/ioutil"
	"bytes"
	"gclassec/authmanagment"
	"gclassec/loggers"
	"encoding/json"
	"gclassec/errorcodes/errcode"
	"gclassec/structs/configurationstruct"
)

var redirectTarget string
var logger = Loggers.New()
 const indexPage = `
  <h1>Select Provider</h1>
  <form method="post" action="/providers">
      <label for="provider">Provider</label>
      <input type="text" id="provider" name="provider">
      </br></br>
      <button type="submit">Select</button>
  </form>`

 const osPage = `
  <h1>Openstack Credentials</h1>
  <form method="post" action="/providers/openstack">
      <label for="host">Host</label>
      <input type="text" id="host" name="host"></br></br>

      <label for="username">Username</label>
      <input type="text" id="username" name="username"></br></br>

      <label for="password">Password</label>
      <input type="text" id="password" name="password"></br></br>

      <label for="projectid">ProjectID</label>
      <input type="text" id="projectid" name="projectid"></br></br>

      <label for="projectname">ProjectName</label>
      <input type="text" id="projectname" name="projectname"></br></br>

      <label for="container">Container</label>
      <input type="text" id="container" name="container"></br></br>

      <label for="imageregion">ImageRegion</label>
      <input type="text" id="imageregion" name="imageregion"></br></br>

      <label for="controller">Controller</label>
      <input type="text" id="controller" name="controller"></br></br>

      <button type="submit">Submit</button>
  </form>`

const azurePage = `
  <h1>Azure Credentials</h1>
  <form method="post" action="/providers/azure">
      <label for="clientid">Client ID</label>
      <input type="text" id="clientid" name="clientid"></br></br>

      <label for="clientsecret">Client Secret</label>
      <input type="text" id="clientsecret" name="clientsecret"></br></br>

      <label for="subscriptionid">Subscription ID</label>
      <input type="text" id="subscriptionid" name="subscriptionid"></br></br>

      <label for="tenantid">Tenant ID</label>
      <input type="text" id="tenantid" name="tenantid"></br></br>

      <button type="submit">Submit</button>
  </form>`

type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)

func NewUserController() *UserController {
    return &UserController{}
}

func (uc UserController) SelectProvider(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexPage)
}

func (uc UserController) OpenstackCreds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, osPage)
}

func (uc UserController) AzureCreds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, azurePage)
}

func (uc UserController) ProviderHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, indexPage)
	logger.Info("--------In Provider Handler--------")
	provider := r.FormValue("provider")
	logger.Info("Provider : ")
	logger.Info(provider)

	if provider == "openstack"{
		//setSession(provider, w)
		redirectTarget = "/selectedOs"
	}

	if provider == "azure"{
		//setSession(provider, w)
		redirectTarget = "/selectedAzure"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func (uc UserController) ProviderOpenstack(w http.ResponseWriter, r *http.Request) {
	//host := r.FormValue("host")

	logger.Info("-------Response Body---------")

	// Read the content
	var bodyBytes []byte
	if r.Body != nil {
  		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	logger.Info(bodyString)

	/*c := map[string]string{
		"host":       r.FormValue("host"),
		"username":   r.FormValue("username"),
		"password": r.FormValue("password"),
		"projectid": r.FormValue("projectid"),
		"projectname": r.FormValue("projectname"),
		"container": r.FormValue("container"),
		"imageregion": r.FormValue("imageregion"),
		"controller": r.FormValue("controller")}

  	outputjson,_:=json.Marshal(c)*/

	filename := "controllers/confcontroller/userconf.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Debug("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/computeVM.json", 1))
       logger.Debug("ABSPATH:==",ConfigFilePath)
	f, err := os.Create(ConfigFilePath)

	//f, err := os.OpenFile("C:/goclassec/src/gclassec/conf/dependencies.env", os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		logger.Error("Error: ", err)
		panic(err)
	}

	defer f.Close()

	/*for _, line := range c {
		if _, err = f.WriteString(line); err != nil {
			panic(err)
		}
	}*/

	//define the 'string writer'
  	filewriter:=bufio.NewWriter(f)

  	//write the JSON string. First we need to convert the outputjson to string, and then write it
  	//filewriter.WriteString(string(outputjson))
	filewriter.WriteString(bodyString)
  	filewriter.Flush()
}

func (uc UserController) ProviderAzure(w http.ResponseWriter, r *http.Request) {
	/*c := map[string]string{
		"clientid": r.FormValue("clientid"),
		"clientsecret": r.FormValue("clientsecret"),
		"subscriptionid": r.FormValue("subscriptionid"),
		"tenantid": r.FormValue("tenantid")}

	outputjson,_:=json.Marshal(c)*/

	// Read the content
	var bodyBytes []byte
	if r.Body != nil {
  		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	logger.Info(bodyString)

	filename := "controllers/confcontroller/userconf.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Debug("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/azurecred.json", 1))
       logger.Debug("ABSPATH:==",ConfigFilePath)
	f, err := os.Create(ConfigFilePath)

	//f, err := os.OpenFile("C:/goclassec/src/gclassec/conf/dependencies.env", os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		logger.Error("Error: ", err)
		panic(err)
	}

	defer f.Close()

	//define the 'string writer'
  	filewriter:=bufio.NewWriter(f)

  	//write the JSON string. First we need to convert the outputjson to string, and then write it
  	//filewriter.WriteString(string(outputjson))
	filewriter.WriteString(bodyString)
  	filewriter.Flush()
}


func (uc UserController) UpdateAwsCredentials(w http.ResponseWriter, r *http.Request){

	if ReadJobConfigFile()!=0{
		authmanagment.AwsCredentials(w, r)
	}else{
		fmt.Fprintf(w, "You Were Not Allowed To Change User Configuration Through API")
	}
}

func (uc UserController) GetAwsCredentials(w http.ResponseWriter, r *http.Request){

	resp := authmanagment.ReadAwsCredentials()
	_ = json.NewEncoder(w).Encode(&resp)


}

func (uc UserController) UpdateAzureCredentials(w http.ResponseWriter, r *http.Request){

	if ReadJobConfigFile()!=0{
		authmanagment.AzureCredentials(w, r)
	}else{
		fmt.Fprintf(w, "Not Allowed To Change User Configuration Through API")
	}

}

func (uc UserController) GetAzureCredentials(w http.ResponseWriter, r *http.Request){

	resp := authmanagment.ReadAzureCredentials()
	_ = json.NewEncoder(w).Encode(&resp)
}

func (uc UserController) UpdateOsCredentials(w http.ResponseWriter, r *http.Request){

	if ReadJobConfigFile()!=0{
		authmanagment.OpenstackCredentials(w, r)
	}else{
		fmt.Fprintf(w, "Not Allowed To Change User Configuration Through API")
	}

}

func (uc UserController) GetOsCredentials(w http.ResponseWriter, r *http.Request){


	resp := authmanagment.ReadOpenstackCredentials()
	_ = json.NewEncoder(w).Encode(&resp)

}

func (uc UserController) UpdateHosCredentials(w http.ResponseWriter, r *http.Request){

	if ReadJobConfigFile()!=0{
		authmanagment.HosCredentials(w, r)
	}else{
		fmt.Fprintf(w, "Not Allowed To Change User Configuration Through API")
	}

}

func (uc UserController) GetHosCredentials(w http.ResponseWriter, r *http.Request){

	resp := authmanagment.ReadHosCredentials()
	_ = json.NewEncoder(w).Encode(&resp)

}

func (uc UserController) UpdateVmwareCredentials(w http.ResponseWriter, r *http.Request){

	if ReadJobConfigFile()!=0{
		authmanagment.VmwareCredentials(w, r)
	}else{
		fmt.Fprintf(w, "Not Allowed To Change User Configuration Through API")
	}

}

func (uc UserController) GetVmwareCredentials(w http.ResponseWriter, r *http.Request){

	resp := authmanagment.ReadVmwareCredentials()
	_ = json.NewEncoder(w).Encode(&resp)
}

func ReadJobConfigFile() int64{
    filename := "controllers/confcontroller/userconf.go"
    _, filePath, _, _ := runtime.Caller(0)
    //logger.Debug("CurrentFilePath:==",filePath)
	fmt.Println("CurrentFilePath:==",filePath)
    ConfigFilePath :=(strings.Replace(filePath, filename, "conf/jobconf.json", 1))
    //logger.Debug("ABSPATH:==",ConfigFilePath)
	fmt.Println("ABSPATH:==",ConfigFilePath)
    file, errOpen := os.Open(ConfigFilePath)
    if errOpen != nil{
        fmt.Println("Error : ", errcode.ErrFileOpen)
        logger.Error("Error : ", errcode.ErrFileOpen)
	return 0
    }
    decoder := json.NewDecoder(file)
    configuration := configurationstruct.Configuration{}
    errDecode := decoder.Decode(&configuration)
    if errDecode != nil {
        fmt.Println("Error : ", errcode.ErrDecode)
        logger.Error("Error : ",errcode.ErrDecode)
        return 0
    }
    return configuration.UpdateUsingAPI

}