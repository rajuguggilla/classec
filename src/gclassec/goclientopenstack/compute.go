// +build !unit

// Copyright (c) 2014 Hewlett-Packard Development Company, L.P.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package goclientcompute
import (
	"fmt"
	"time"
	"gclassec/goclientopenstack/openstack"
	"encoding/json"
	"net/http"
	"gclassec/goclientopenstack/flavor"
	"gclassec/goclientopenstack/compute"
	"strings"
	"runtime"
	"os"
	"gclassec/loggers"
	"gclassec/errorcodes/errcode"
)
type Configuration struct {
    Host    string
    Username   string
    Password   string
    ProjectID   string
    ProjectName   string
    Container   string
    ImageRegion string
	Controller string


}

func Compute() []compute.DetailResponse {
	//config := getConfig()
	logger := Loggers.New()
	filename := "goclientopenstack/compute.go"
       _, filePath, _, _ := runtime.Caller(0)
       logger.Info("CurrentFilePath:==",filePath)
       ConfigFilePath :=(strings.Replace(filePath, filename, "conf/computeVM.json", 1))
       logger.Info("ABSPATH:==",ConfigFilePath)
	file, errOpen := os.Open(ConfigFilePath)

	if errOpen != nil{
		fmt.Println("Error : ", errcode.ErrFileOpen)
		logger.Error("Error : ", errcode.ErrFileOpen)
		return []compute.DetailResponse{}
	}

	//dir, _ := os.Getwd()
	//file, _ := os.Open(dir + "/src/gclassec/conf/computeVM.json")
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		logger.Error("error:", err)
		return []compute.DetailResponse{}
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
	}
	if !auth.GetExpiration().After(time.Now()) {
		logger.Error("There was an error. The auth token has an invalid expiration.")
	}
	logger.Debug("OpenStack: ",auth)
	// Find the endpoint for the Nova Compute service.
	url, err := auth.GetEndpoint("compute", "")
	url = strings.Replace(url,"compute", creds.Controller ,1)
	if url == "" || err != nil {
		logger.Error("EndPoint Not Found.")
		logger.Error(err)
	}
	// Make a new client with these creds
	sess, err := openstack.NewSession(nil, auth, nil)
	if err != nil {
		panicString := fmt.Sprint("Error creating new Session:", err)
		logger.Error(panicString)
	}
	logger.Info(url)
	computeService := compute.Service{
		Session: *sess,
		Client:  *http.DefaultClient,
		URL:     url, // We're forcing Volume v2 for now
	}
	computeDetails, err := computeService.InstancesDetail()
	if err != nil {
		panicString := fmt.Sprint("Cannot access Compute:", err)
		logger.Error(panicString)
	}
	logger.Info("computedetails printing..")
	logger.Info(computeDetails)
	var computeIDs = make([]string, 0)
	for _, element := range computeDetails {
		computeIDs = append(computeIDs, element.ID)

	}
	logger.Info(computeIDs)
	if len(computeIDs) == 0 {
		panicString := fmt.Sprint("No instances found, check to make sure access is correct")
		logger.Error(panicString)
	}
	return computeDetails
}

func FinalCompute() ([]compute.DetailResponse, error) {
	logger := Loggers.New()
	var flvObj []flavor.DetailResponse
	flvObj, err := flavor.Flavor()

		if err !=nil{
			fmt.Println("OpenStack : ", errcode.ErrAuth)
			logger.Error("OpenStack : ", errcode.ErrAuth)
			return []compute.DetailResponse{}, err
		}



	logger.Info("&**********Showing FLVOBJ&************")
	logger.Info(flvObj)
	logger.Info("*********************")
	logger.Info("flvObj.FlavorID::", flvObj[1].FlavorID)

	var obj []compute.DetailResponse
	obj = Compute()
	fmt.Println("77778888899999", obj[1].Flavor.FlavorID)
	for i:=0; i<len(obj); i++{
		tempFID :=obj[i].Flavor.FlavorID
		for j:=0; j<len(flvObj); j++{
			if tempFID==flvObj[j].FlavorID{
				obj[i].Flavor.Name=flvObj[j].Name
				obj[i].Flavor.Disk=flvObj[j].Disk
				obj[i].Flavor.RAM=flvObj[j].RAM
				obj[i].Flavor.VCPU=flvObj[j].VCPU
			}
		}
	}
	out, err := json.Marshal(obj)
	if err != nil {
        	logger.Error("Error : ", err)
		return []compute.DetailResponse{}, err
    	}
	logger.Info("Out Sritng")
	logger.Info(string(out))
	temp := string(out)
	temp1 := strings.TrimPrefix(temp, "[{")
	tempstr:= strings.TrimSuffix(temp1, "}]")
	tempVar := strings.Split(tempstr,"},{")
	logger.Info("-----------TempVar----------")
	for i:=0; i<len(tempVar);i++{
		logger.Info(tempVar[i])
	}
	for i:=0; i<len(tempVar);i++{
		nevVar := string(tempVar[i])
		tempVar1 := strings.Split(nevVar,",")
		logger.Info("-----------TempVar1.----------",i)
		for j:=0; j<len(tempVar1);j++{
			logger.Info(tempVar1[j])
		}
	}
	return obj, nil
}