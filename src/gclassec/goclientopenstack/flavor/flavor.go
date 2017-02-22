package flavor
import (
	"fmt"
	"time"
	"gclassec/goclientopenstack/openstack"
	"encoding/json"
	"net/http"
	"runtime"
	"strings"
	"os"
	"gclassec/loggers"
)
type Configuration struct {
    Host	string
    Username	string
    Password	string
    ProjectID	string
    ProjectName	string
    Container	string
    ImageRegion	string
	Controller string
}

func Flavor() ([]DetailResponse, error){
	//config := getConfig()
	logger := Loggers.New()
	filename := "goclientopenstack/flavor/flavor.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Debug("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/computeVM.json", 1))
       logger.Debug("ABSPATH:==",ConfigFilePath)
	file, _ := os.Open(ConfigFilePath)
	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/computeVM.json")
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		logger.Error("error:", err)
		//return []DetailResponse{},err
	}

	// Authenticate with a username, password, tenant id.
	creds := openstack.AuthOpts{
		AuthUrl:     config.Host,
		ProjectName: config.ProjectName,
		Username:    config.Username,
		Password:    config.Password,
		Controller:	config.Controller,
	}
	auth, err := openstack.DoAuthRequest(creds)
	if err != nil {
		panicString := fmt.Sprint("There was an error authenticating:", err)
		logger.Error(panicString)
//		panic(panicString)
		return []DetailResponse{}, err

	}
	if !auth.GetExpiration().After(time.Now()) {
		logger.Error("There was an error. The auth token has an invalid expiration.")
		//panic("There was an error. The auth token has an invalid expiration.")
	}
	logger.Debug(auth)
	// Find the endpoint for the Nova Compute service.
	url, err := auth.GetEndpoint("compute", "")
	url = strings.Replace(url,"compute", creds.Controller ,1)
	if url == "" || err != nil {
		logger.Error("EndPoint Not Found.")
	//	panic("EndPoint Not Found.")
		logger.Error(err)
	//	panic(err)

		return []DetailResponse{},err
	}
	// Make a new client with these creds
	sess, err := openstack.NewSession(nil, auth, nil)
	if err != nil {
		panicString := fmt.Sprint("Error creating new Session:", err)
		logger.Error(panicString)
	//	panic(panicString)

		return []DetailResponse{}, err
	}
	logger.Info(url)
	flavorService := Service{
		Session: *sess,
		Client:  *http.DefaultClient,
		URL:     url, // We're forcing Volume v2 for now
	}
	flavorDetails, err := flavorService.FlavorsDetail()

	if err != nil{
		return []DetailResponse{}, err
	}


	logger.Info(flavorDetails,"00000000000000000000000000")
	if err != nil {
		panicString := fmt.Sprint("Cannot access Compute:", err)
		logger.Error(panicString)
	//	panic(panicString)
		return []DetailResponse{},err
	}
	logger.Info("computedetails printing..")
	logger.Info(flavorDetails)
	var flavorIDs = make([]string, 0)
	for _, element := range flavorDetails {
		flavorIDs = append(flavorIDs, element.FlavorID)
	}
	logger.Info(flavorIDs)
	if len(flavorIDs) == 0 {
		panicString := fmt.Sprint("No instances found, check to make sure access is correct")
		logger.Error(panicString)
	//	panic(panicString)
	}
	return flavorDetails, nil
}