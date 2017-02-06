package main

import (
	"gclassec/goclientopenstack/GetAuthToken"
	"fmt"
	"gclassec/goclientopenstack/OpenStack"
)
func main(){
//func openStackServer() ComputeResponse{
	auth := GetAuthToken.GetOpenStackAuthToken()
	var CompStruct OpenStack.ComputeResponse

 	CompStruct = OpenStack.Compute()
	fmt.Println("Auth Token", auth)
	fmt.Printf("Body of CompStruct in Main")
	fmt.Printf("%+v\n\n", CompStruct)
	fmt.Printf("\n\n\n\n=======================")
	fmt.Println("Compute Output", CompStruct)
	var flvStruct OpenStack.FlavorResponse
	flvStruct = OpenStack.Flavor()
	fmt.Printf("Body of flvStruct in Main")
	fmt.Printf("%+v\n\n",flvStruct)
	fmt.Printf("\n\n\n\n=======================\n\n\n\n\n\n")

	for i:=0;i<len(CompStruct.Servers); i++{
		tempFID := (CompStruct.Servers[i].Flavor.FlavorID)
		for j:=0; j<len(flvStruct.Flavors); j++{
			if tempFID == flvStruct.Flavors[j].FlavorID{
				CompStruct.Servers[i].Flavor.Name=flvStruct.Flavors[j].Name
				CompStruct.Servers[i].Flavor.Disk=flvStruct.Flavors[j].Disk
				CompStruct.Servers[i].Flavor.RAM=flvStruct.Flavors[j].RAM
				CompStruct.Servers[i].Flavor.VCPU=flvStruct.Flavors[j].VCPU
			}
		}
	}

	fmt.Printf("%+v\n\n", CompStruct)
	//return  CompStruct

}

