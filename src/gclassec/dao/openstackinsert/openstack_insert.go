package openstackinsert

import (
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/structs/openstackInstance"
	"gclassec/loggers"
	"fmt"
	"gclassec/errorcodes/errcode"
	"gclassec/dbmanagement"
	"gclassec/openstackgov"
	"gclassec/confmanagement/readopenstackconfig"
	"gclassec/openstackgov/ceilometer"
//	"github.com/vmware/govmomi/govc/ls"
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

	openstack_struct := []openstackInstance.Instances{}

	er := db.Find(&openstack_struct).Error

	if er != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		//tx.Rollback()
		return
	}

	db.Find(&openstack_struct)

	/*for _, element := range computeDetails.Servers {
			user := openstackInstance.Instances{Name:element.ServerName, InstanceID:element.ServerId, Status:element.Status, RAM:element.Flavor.Ram, VCPU:element.Flavor.VCPUS, Flavor:element.Flavor.FlavorName, Storage:element.Flavor.Disk, AvailabilityZone:element.AvailabilityZone, CreationTime:element.CreatedAt,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.AccessIPv4, KeyPairName:element.KeyName, ImageName:element.Image.ImageID, Deleted:false, Classifier: temp.ProjectName }
                    db.Create(&user)
		}*/



	if (len(openstack_struct)==0){
		for _, element := range computeDetails.Servers {
			user := openstackInstance.Instances{Name:element.ServerName, InstanceID:element.ServerId, Status:element.Status, RAM:element.Flavor.Ram, VCPU:element.Flavor.VCPUS, Flavor:element.Flavor.FlavorName, Storage:element.Flavor.Disk, AvailabilityZone:element.AvailabilityZone, CreationTime:element.CreatedAt,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.AccessIPv4, KeyPairName:element.KeyName, ImageName:element.Image.ImageID, Deleted:false, Classifier: temp.ProjectName }
                    db.Create(&user)
		}
	}else{
		for _, element := range computeDetails.Servers {
		db.Where("name =?",element.ServerName).Find(&openstack_struct)
		if(len(openstack_struct)==0){
			 user := openstackInstance.Instances{Name:element.ServerName, InstanceID:element.ServerId, Status:element.Status, RAM:element.Flavor.Ram, VCPU:element.Flavor.VCPUS, Flavor:element.Flavor.FlavorName, Storage:element.Flavor.Disk, AvailabilityZone:element.AvailabilityZone, CreationTime:element.CreatedAt,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.AccessIPv4, KeyPairName:element.KeyName, ImageName:element.Image.ImageID, Deleted:false, Classifier: temp.ProjectName }
                    db.Create(&user)
		}else{
			user := openstackInstance.Instances{Name:element.ServerName, InstanceID:element.ServerId, Status:element.Status, RAM:element.Flavor.Ram, VCPU:element.Flavor.VCPUS, Flavor:element.Flavor.FlavorName, Storage:element.Flavor.Disk, AvailabilityZone:element.AvailabilityZone, CreationTime:element.CreatedAt,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.AccessIPv4, KeyPairName:element.KeyName, ImageName:element.Image.ImageID, Deleted:true, Classifier: temp.ProjectName }
                     db.Model(&user).Where("name = ?",element.ServerName).Updates(user)
		}
}
	}
	fmt.Println("hellloooooooooooooooooooooooo")
	//ENVcount:= 0


	/*for _, element := range openstack_struct {
       db.Table("instances").Where("name = ?",element.Name).Update("deleted", true)
}*/
	db.Find(&openstack_struct)
	fmt.Println(openstack_struct)
	for _, element := range openstack_struct {
              fmt.Println("inside openstack delete")
             for i:=0;i<len(computeDetails.Servers);i++{
                     if element.Name != computeDetails.Servers[i].ServerName {
			     if(i == len(computeDetails.Servers)-1) {
                                   fmt.Println("hello")
                                   db.Table("instances").Where("name = ?",element.Name ).Update("deleted", true)
                                   fmt.Println("insdie  continue")
                            }
                            continue
                     }else{
                            db.Table("instances").Where("name = ?",element.Name ).Update("deleted", false)
                            break
              }

             }
       }
	logger.Info("Successful in InsertInstances.")
	tx.Commit()
}



func OSDynamicInsert() error{

	dynamicDetails, err := ceilometer.DynamicDetails()

	if err != nil{
		return err
	}
	logger.Info(dynamicDetails)


	// Inserting Dynamic Data into Database
	for _, element := range dynamicDetails.Servers{
		user := openstackInstance.DynamicInstances{Vm_Name:element.Vm_Name, InstanceID:element.InstanceID, Count:element.Count, DurationStart:element.DurationStart, Min:element.Min,DurationEnd:element.DurationEnd, Max:element.Max, Sum:element.Sum, Period:element.Period, PeriodEnd:element.PeriodEnd, Duration:element.Duration, PeriodStart:element.PeriodStart, Avg:element.Avg, Unit:element.Unit}
		db.Create(&user)
		//db.Model(&user).Updates(&user)
	}

	logger.Info("Successful in InsertHOSDynamicInstance")
	return nil
}