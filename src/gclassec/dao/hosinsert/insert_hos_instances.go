package hosinsert

import (
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/goclienthos/compute"
	"gclassec/confmanagement/readopenstackconf"
	_ "github.com/go-sql-driver/mysql"
	"gclassec/structs/hosstruct"
	"gclassec/loggers"

	"gclassec/structs/tagstruct"
	"regexp"
	"fmt"
	"gclassec/errorcodes/errcode"
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

func HosInsert(){
	logger := Loggers.New()
	//println(examples.ComputeFunc())
	computeDetails:= compute.Compute()

	tx := db.Begin()
	db.SingularTable(true)

	tag := []tagstruct.Providers{}
	hos_compute := []hosstruct.HosInstances{}

	err := db.Find(&hos_compute).Error

	if err != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}

	db.Find(&hos_compute)

	//create a regex `(?i)hos` will match string contains "hos" case insensitive
	reg := regexp.MustCompile("(?i)hos")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("hos")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("hos")).Find(&tag)

	for _, element := range hos_compute {
       db.Table("hos_instances").Where("Name = ?",element.Vm_Name).Update("deleted", true)
}

	for _, element := range computeDetails.Servers {
       for _, ele := range hos_compute {
              if element.Vm_Name != ele.Vm_Name {
                     continue
              }else{
                     user := hosstruct.HosInstances{Vm_Name:element.Vm_Name,InstanceID:element.InstanceID,FlavorID:element.Flavor.FlavorID,FlavorName:element.Flavor.FlavorName,Status:element.Status,Image:element.Image.ImageID,SecurityGroups:element.Security_Groups.Name,AvailabilityZone:element.Availability_Zone,KeypairName:element.Key_name,Ram:element.Flavor.Ram,VCPU:element.Flavor.VCPUS,Disk:element.Flavor.Disk, Tagname:"Nil", Deleted:true}
                     db.Model(&user).Where("Name =?",element.Vm_Name).Updates(user)
              }
       }
}

	for _, i := range hos_compute{
		if len(tag) == 0 {
			fmt.Println("----Nothing in Tag----")
			db.Table("hos_instances").Where("Instance_id = ?", i.InstanceID).Update("tagname","Nil")
		}else {
			for _, el := range tag {
				if i.InstanceID != el.InstanceId {
					fmt.Println("----No Tag for this instance----")
					db.Table("hos_instances").Where("Instance_id = ?", i.InstanceID).Update("tagname","Nil")
				}else {
					fmt.Println("----Update Tag for this instance----")
					fmt.Println("el.Tagname : ", el.Tagname)
					db.Table("hos_instances").Where("Instance_id = ?", i.InstanceID).Update("tagname",el.Tagname)
				}
			}
		}
	}

	for _, element := range hos_compute {
              for _, ele := range computeDetails.Servers{
                     if element.Vm_Name != ele.Vm_Name {
                     continue
                     }else{
                            db.Table("hos_instances").Where("Name = ?",element.Vm_Name).Update("deleted", false)
              }
              }
              }
	logger.Info("Successful in InsertHOSInstance.")
	tx.Commit()
}