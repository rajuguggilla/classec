package hosinsert

import (
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/goclienthos/compute"
	"gclassec/confmanagement/readopenstackconf"
	_ "github.com/go-sql-driver/mysql"
	"gclassec/structs/hosstruct"

)

var dbcredentials = readopenstackconf.Configurtion()
var dbtype string = dbcredentials.Dbtype
var dbname  string = dbcredentials.Dbname
var dbusername string = dbcredentials.Dbusername
var dbpassword string = dbcredentials.Dbpassword
var dbhostname string = dbcredentials.Dbhostname
var dbport string = dbcredentials.Dbport
var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))
var db,err  = gorm.Open(dbtype, c)

func InsertHOSInstances(){
	//println(examples.ComputeFunc())
	computeDetails:= compute.Compute()
	//println(computeDetails)
	//user := openstackInstance.Instances{}
	//db.Delete(&user)
	/*for _, element := range computeDetails {
		//println(element.Name,element.ID,element.Status,element.Progress)
		*//*user :=	openstackInstance.Instances{Name:element.Name,InstanceID:element.ID,Status:element.Status,AvailabilityZone:element.Availability_zone,CreationTime:element.Created,
		Volumes:element.Volumes_attached,KeyPairName:element.Key_name}*//*
		user := hosstruct.HosInstances{Vm_Name:element.Vm_Name,InstanceID:element.InstanceID,Flavor:element.Flavor.FlavorID,Status:element.Status,Image:element.Image,Security_Groups:element.Security_Groups,Availability_Zone:element.Availability_Zone}
		db.Create(&user)
		db.Model(&user).Updates(&user)
	}*/
	for _, element := range computeDetails.Servers{
		user := hosstruct.HosInstances{Vm_Name:element.Vm_Name,InstanceID:element.InstanceID,FlavorID:element.Flavor.FlavorID,FlavorName:element.Flavor.FlavorName,Status:element.Status,Image:element.Image.ImageID,SecurityGroups:element.Security_Groups.Name,AvailabilityZone:element.Availability_Zone,KeypairName:element.Key_name,Ram:element.Flavor.Ram,VCPU:element.Flavor.VCPUS,Disk:element.Flavor.Disk}
		db.Create(&user)
		db.Model(&user).Updates(&user)
	}

}