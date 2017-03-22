package hosinsert

import (
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/goclienthos/compute"
	_ "github.com/go-sql-driver/mysql"
	"gclassec/structs/hosstruct"
	"gclassec/loggers"
	"gclassec/errorcodes/errcode"
	"gclassec/dbmanagement"
	"gclassec/goclienthos/dynamicdetails"
	"gclassec/confmanagement/readhosconf"
	"fmt"
)

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

var hoscreds = readhosconf.Configurtion()
var classifier = hoscreds.ProjectName

func HosInsert() (int,int){
	logger := Loggers.New()
	//println(examples.ComputeFunc())
	computeDetails, err1 := compute.Compute()
	if err1 != nil{
		logger.Error("Error: ",errcode.ErrAuth)
		//tx.Rollback()
		return 0,0
	}

	tx := db.Begin()
	db.SingularTable(true)

	hos_compute := []hosstruct.HosInstances{}

	err := db.Find(&hos_compute).Error

	if err != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		//tx.Rollback()
		return 0,0
	}

	db.Find(&hos_compute)

	poweredoncount := 0
	poweredoffcount := 0
	for _,element1 := range computeDetails.Servers {
		if element1.Power_State == 1 {
			poweredoncount++
		} else {
			poweredoffcount++
		}
	}
	if(len(hos_compute)==0){
		for _,element :=range computeDetails.Servers{
			 user := hosstruct.HosInstances{Vm_Name:element.Vm_Name,InstanceID:element.InstanceID,FlavorID:element.Flavor.FlavorID,FlavorName:element.Flavor.FlavorName,Status:element.Status,Image:element.Image.ImageID,SecurityGroups:element.Security_Groups.Name,AvailabilityZone:element.Availability_Zone,KeypairName:element.Key_name,Ram:element.Flavor.Ram,VCPU:element.Flavor.VCPUS,Disk:element.Flavor.Disk, Deleted:false, Classifier:classifier}
			db.Create(&user)
		}
	}else{
		for _, element := range computeDetails.Servers {
			db.Where("name =?",element.Vm_Name).Find(&hos_compute)
			if(len(hos_compute)==0){
				 user := hosstruct.HosInstances{Vm_Name:element.Vm_Name,InstanceID:element.InstanceID,FlavorID:element.Flavor.FlavorID,FlavorName:element.Flavor.FlavorName,Status:element.Status,Image:element.Image.ImageID,SecurityGroups:element.Security_Groups.Name,AvailabilityZone:element.Availability_Zone,KeypairName:element.Key_name,Ram:element.Flavor.Ram,VCPU:element.Flavor.VCPUS,Disk:element.Flavor.Disk, Deleted:false, Classifier:classifier}
			db.Create(&user)
			}else{
				  user := hosstruct.HosInstances{Vm_Name:element.Vm_Name,InstanceID:element.InstanceID,FlavorID:element.Flavor.FlavorID,FlavorName:element.Flavor.FlavorName,Status:element.Status,Image:element.Image.ImageID,SecurityGroups:element.Security_Groups.Name,AvailabilityZone:element.Availability_Zone,KeypairName:element.Key_name,Ram:element.Flavor.Ram,VCPU:element.Flavor.VCPUS,Disk:element.Flavor.Disk, Deleted:false, Classifier:classifier}
                     db.Model(&user).Where("Name =?",element.Vm_Name).Updates(user)
			}
		}
	}
	/*for _, element := range hos_compute {
       db.Table("hos_instances").Where("Name = ?",element.Vm_Name).Update("deleted", true)
}*/
	db.Find(&hos_compute)

	for _, element := range hos_compute {
              fmt.Println("inside delete")
             for i:=0;i<len(computeDetails.Servers);i++{
                     if element.Vm_Name != computeDetails.Servers[i].Vm_Name {
                            if(i == len(computeDetails.Servers)-1) {
                                   fmt.Println("hello")
                                   db.Table("hos_instances").Where("Name = ?",element.Vm_Name ).Update("deleted", true)
			    }
                            continue
                     }else{
                            db.Table("hos_instances").Where("Name = ?",element.Vm_Name ).Update("deleted", false)
                            break
              }

             }
       }
	logger.Info("Successful in InsertHOSInstance.")
	tx.Commit()
	return poweredoncount,poweredoffcount
}



// Inserting Dynamic Data into database

func HOSDynamicInsert() error{

	dynamicDetails, err := dynamicdetails.DynamicDetails()

	if err != nil{
		return err
	}
	logger.Info(dynamicDetails)


	// Inserting Dynamic Data into Database
	for _, element := range dynamicDetails.Servers{
		user := hosstruct.HosDynamicInstances{Vm_Name:element.Vm_Name, InstanceID:element.InstanceID, Count:element.Count, DurationStart:element.DurationStart, Min:element.Min,DurationEnd:element.DurationEnd, Max:element.Max, Sum:element.Sum, Period:element.Period, PeriodEnd:element.PeriodEnd, Duration:element.Duration, PeriodStart:element.PeriodStart, Avg:element.Avg, Unit:element.Unit}
		db.Create(&user)
		//db.Model(&user).Updates(&user)
	}

	logger.Info("Successful in InsertHOSDynamicInstance")
	return nil
}