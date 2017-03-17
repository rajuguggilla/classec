package instancestatus

import (
	_ "github.com/go-sql-driver/mysql"
	//"gclassec/structs/vmwarestructs"
	//"gclassec/errorcodes/errcode"
	"gclassec/dbmanagement"
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/loggers"
	//"fmt"
	"encoding/json"
	"os"
	"gclassec/structs/vmwarestructs"
	//"fmt"
	//"gclassec/dao/openstackinsert"

	"fmt"
	"gclassec/dao/vmwareinsert"
	"gclassec/dao/hosinsert"
	"gclassec/openstackgov"
	"gclassec/dao/azureinsert"
	"net/http"
	"gclassec/readcredentials"
	"gclassec/authmanagment"
)
var azurecreds = readazurecreds.Configurtion()
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
type GetstatusInstances struct {
	Provider string
	Type string
	RunningInstances int
	StoppedInstances int
	Classifier string
}

func Getinstancestatus(w http.ResponseWriter, r *http.Request) {

	//getstatusinstances := []GetstatusInstances{}

//tx := db.Begin()
	//var count int
	db.SingularTable(true)
	vmware_struct := []vmwarestructs.VmwareInstances{}
	//errFind := db.Model(&vmware_struct).Count(&vmware_struct).Where("PowerState = ?","poweredOn")
	//db.Model(&vmware_struct).Where("name = ?", "jinzhu").Count(&count)
	//count := SELECT  count(*) FROM `vmware_instances` WHERE PowerState="poweredOn";
	//  _ = json.NewEncoder(os.Stdout).Encode(db.Table("vmware_instances").Select("COUNT(*)").Where("PowerState = ?","poweredOn"))
	/*_ = json.NewEncoder(os.Stdout).Encode(db.Table("vmware_instances").Select("*"))
	rows, err := db.Query("SELECT COUNT(*) as count FROM  table_name")
        fmt.Println("Total count:",checkCount(rows))*/
	_ = json.NewEncoder(os.Stdout).Encode(db.Where("PowerState = ?", "poweredOn").Find(&vmware_struct))
	/*if errFind != nil {
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}*/
	//openstackinsert.ENVcount
	//fmt.Print(*(&count))
	err,count1,count2 :=vmwareinsert.VmwareInsert()
fmt.Println(count1)
	fmt.Print(count2)
	println(err)

	poweredoncount,poweredoffcount := hosinsert.HosInsert()
	fmt.Println("HOS poweredoncount",poweredoncount)
	fmt.Println("HOS poweredoffcount",poweredoffcount)

	poweredoncount1,poweredoffcount1 := openstackgov.GetServer()
	fmt.Println("Openstack on count",poweredoncount1)
	fmt.Println("Openstack off count",poweredoffcount1)

	err,poweredoncount2,poweredoffcount2 := azureinsert.AzureInsert()
	fmt.Println("Azure poweron count",poweredoncount2)
	fmt.Println("Azure poweroff count",poweredoffcount2)

	var basenameOpts = []GetstatusInstances {
    GetstatusInstances {
        Provider: "VMware",
        Type: "Public",
        RunningInstances: count1,
        StoppedInstances: count2,
	    Classifier:vmwareinsert.EnvUserName,
    },
    GetstatusInstances {
        Provider: "Azure",
        Type: "Public",
        RunningInstances: poweredoncount2,
        StoppedInstances: poweredoffcount2,
	    Classifier:azurecreds.SubscriptionId,
    },
		GetstatusInstances {
        Provider: "HOS",
        Type: "Private",
        RunningInstances: poweredoncount,
        StoppedInstances: poweredoffcount,
			Classifier:authmanagment.ReadHosCredentials().ProjectName,
    },
		GetstatusInstances {
        Provider: "Openstack",
        Type: "Private",
        RunningInstances: poweredoncount1,
        StoppedInstances: poweredoffcount1,
			Classifier:authmanagment.ReadOpenstackCredentials().ProjectName,
    },
}

	/*user := &GetstatusInstances{
  Provider: "VMware",
  Type: "Public",
      RunningInstances:count1,
		StoppedInstances:count2,

}
	b,err := json.Marshal(user)*/
	_ = json.NewEncoder(w).Encode(&basenameOpts)

	/*for i:=0;i<=4;i++{
		user
	}*/

	//_ = json.NewEncoder(os.Stdout).Encode(db.Model(&vmware_struct).Where("PowerState = ?", "poweredOn"))
return

}

