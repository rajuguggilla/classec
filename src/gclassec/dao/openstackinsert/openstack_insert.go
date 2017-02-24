package openstackinsert

import (

	"gclassec/confmanagement/readopenstackconf"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/structs/openstackInstance"
	"gclassec/goclientopenstack"
	"gclassec/loggers"
	"regexp"
	"fmt"
	"gclassec/structs/tagstruct"
	"gclassec/errorcodes/errcode"
)
type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}
var logger = Loggers.New()
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

func InsertInstances(){
	//println(examples.ComputeFunc())
	computeDetails, err:= goclientcompute.FinalCompute()
	if err != nil{
		return
	}
	println(computeDetails)
	logger.Info(computeDetails)
	tx := db.Begin()
	db.SingularTable(true)

	tag := []tagstruct.Providers{}

	//create a regex `(?i)openstack` will match string contains "openstack" case insensitive
	reg := regexp.MustCompile("(?i)openstack")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("Openstack")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("Openstack")).Find(&tag)

	openstack_struct := []openstackInstance.Instances{}

	er := db.Find(&openstack_struct).Error

	if er != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}

	db.Find(&openstack_struct)

	for _, element := range openstack_struct {
       db.Table("instances").Where("name = ?",element.Name).Update("deleted", true)
}

	for _, element := range computeDetails {
       for _, ele := range openstack_struct {
              if element.Name != ele.Name {
                     continue
              }else{
                     user := openstackInstance.Instances{Name:element.Name, InstanceID:element.ID, Status:element.Status, RAM:element.Flavor.RAM, VCPU:element.Flavor.VCPU, Flavor:element.Flavor.Name, Storage:element.Flavor.Disk, AvailabilityZone:element.Availability_zone, CreationTime:element.Created,
                            FlavorID:element.Flavor.FlavorID, IPAddress:element.IPV4, KeyPairName:element.Key_name, ImageName:element.Image.ID, Tagname:"Nil", Deleted:true}
                     db.Model(&user).Where("name = ?",element.Name).Updates(user)
              }
       }
}

	for _, i := range openstack_struct{
		if len(tag) == 0 {
			fmt.Println("----Nothing in Tag----")
			db.Table("instances").Where("instance_id = ?", i.InstanceID).Update("tagname","Nil")
		}else {
			for _, el := range tag {
					if i.InstanceID != el.InstanceId{
						fmt.Println("----No Tag for this instance----")
						db.Table("instances").Where("instance_id = ?", i.InstanceID).Update("tagname","Nil")
					}else {
						fmt.Println("----Update Tag for this instance----")
						fmt.Println("el.Tagname : ", el.Tagname)
						db.Table("instances").Where("instance_id = ?", i.InstanceID).Update("tagname",el.Tagname)
					}
				}
		}
	}

	for _, element := range openstack_struct {
              for _, ele := range computeDetails{
                     if element.Name != ele.Name {
                     continue
                     fmt.Println("insdie  continue")
                     }else{
                            db.Table("instances").Where("name = ?",element.Name).Update("deleted", false)
              }
              }
              }
	logger.Info("Successful in InsertInstances.")
	tx.Commit()
}