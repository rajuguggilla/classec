package openstackinsert

import (
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/structs/openstackInstance"
	"gclassec/loggers"
	"regexp"
	"fmt"
	"gclassec/structs/tagstruct"
	"gclassec/errorcodes/errcode"
	"gclassec/dbmanagement"
	"gclassec/openstackgov"
	"gclassec/confmanagement/readopenstackconfig"
)
type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}
var ENVcount = 0
var logger = Loggers.New()
var dbtype string = dbmanagement.ENVdbtype
var dbname  string = dbmanagement.ENVdbnamegodb
var dbusername string = dbmanagement.ENVdbusername
var dbpassword string = dbmanagement.ENVdbpassword
var dbhostname string = dbmanagement.ENVdbhostname
var dbport string = dbmanagement.ENVdbport
var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))
var db,err  = gorm.Open(dbtype, c)

func InsertInstances(){
	//println(examples.ComputeFunc())
	//computeDetails, err:= goclientcompute.FinalCompute()
	temp := readopenstackconfig.OpenStackConfigReader()
	var flvError error
	var computeDetails openstackInstance.ComputeListStruct
	computeDetails,flvError=openstackgov.ListComputeInstances()
	if flvError != nil{
		fmt.Println("Error In getting Compute Details:==",flvError)
	}else{
		fmt.Println("-----------------------------------------Compute Details-----------------------------------------------------")
		fmt.Println(computeDetails)
	}
	if err != nil{
		return
	}
	fmt.Printf("\n%v\n",computeDetails)
	logger.Info(computeDetails)
	tx := db.Begin()
	db.SingularTable(true)

	tag := []tagstruct.Tags{}

	//create a regex `(?i)openstack` will match string contains "openstack" case insensitive
	reg := regexp.MustCompile("(?i)openstack")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("Openstack")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		//tx.Rollback()
		return
	}
	db.Where("Cloud = ?", reg.FindString("Openstack")).Find(&tag)

	openstack_struct := []openstackInstance.Instances{}

	er := db.Find(&openstack_struct).Error

	if er != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		//tx.Rollback()
		return
	}

	db.Find(&openstack_struct)

	for _, element := range computeDetails.Servers {
			user := openstackInstance.Instances{Name:element.ServerName, InstanceID:element.ServerId, Status:element.Status, RAM:element.Flavor.Ram, VCPU:element.Flavor.VCPUS, Flavor:element.Flavor.FlavorName, Storage:element.Flavor.Disk, AvailabilityZone:element.AvailabilityZone, CreationTime:element.CreatedAt,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.AccessIPv4, KeyPairName:element.KeyName, ImageName:element.Image.ImageID, Tagname:"Nil", Deleted:false, Classifier: temp.ProjectName }
                    db.Create(&user)
		}



	if (len(openstack_struct)==0){
		for _, element := range computeDetails.Servers {
			user := openstackInstance.Instances{Name:element.ServerName, InstanceID:element.ServerId, Status:element.Status, RAM:element.Flavor.Ram, VCPU:element.Flavor.VCPUS, Flavor:element.Flavor.FlavorName, Storage:element.Flavor.Disk, AvailabilityZone:element.AvailabilityZone, CreationTime:element.CreatedAt,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.AccessIPv4, KeyPairName:element.KeyName, ImageName:element.Image.ImageID, Tagname:"Nil", Deleted:false, Classifier: temp.ProjectName }
                    db.Create(&user)
		}
	}else{
		for _, element := range computeDetails.Servers {
		db.Where("name =?",element.ServerName).Find(&openstack_struct)
		if(len(openstack_struct)==0){
			 user := openstackInstance.Instances{Name:element.ServerName, InstanceID:element.ServerId, Status:element.Status, RAM:element.Flavor.Ram, VCPU:element.Flavor.VCPUS, Flavor:element.Flavor.FlavorName, Storage:element.Flavor.Disk, AvailabilityZone:element.AvailabilityZone, CreationTime:element.CreatedAt,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.AccessIPv4, KeyPairName:element.KeyName, ImageName:element.Image.ImageID, Tagname:"Nil", Deleted:false, Classifier: temp.ProjectName }
                    db.Create(&user)
		}else{
			user := openstackInstance.Instances{Name:element.ServerName, InstanceID:element.ServerId, Status:element.Status, RAM:element.Flavor.Ram, VCPU:element.Flavor.VCPUS, Flavor:element.Flavor.FlavorName, Storage:element.Flavor.Disk, AvailabilityZone:element.AvailabilityZone, CreationTime:element.CreatedAt,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.AccessIPv4, KeyPairName:element.KeyName, ImageName:element.Image.ImageID, Tagname:"Nil", Deleted:true, Classifier: temp.ProjectName }
                     db.Model(&user).Where("name = ?",element.ServerName).Updates(user)
		}
}
	}
	//ENVcount:= 0
	for _,element1 := range computeDetails.Servers{
		if element1.Status =="ACTIVE"{
			ENVcount++
		}
	}

	/*for _, element := range openstack_struct {
       db.Table("instances").Where("name = ?",element.Name).Update("deleted", true)
}*/
	db.Find(&openstack_struct)
	for _, i := range openstack_struct{
		if len(tag) != 0 {
			for _, el := range tag {
				if i.InstanceID == el.InstanceId{
					fmt.Println("----Update Tag for this instance----")
					fmt.Println("el.Tagname : ", el.Tagname)
					db.Model(openstackInstance.Instances{}).Where("instance_id = ?", i.InstanceID).Update("tagname",el.Tagname)
				}

			}
		}
	}
/*
	for _, element := range openstack_struct {
              for _, ele := range computeDetails{
                     if element.Name != ele.Name {
                     continue
                     fmt.Println("insdie  continue")
                     }else{
                            db.Table("instances").Where("name = ?",element.Name).Update("deleted", false)
              }
              }
              }*/
	logger.Info("Successful in InsertInstances.")
	tx.Commit()
}