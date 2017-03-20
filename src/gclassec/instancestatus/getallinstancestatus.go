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
	//"encoding/json"
	//"os"
	//"gclassec/structs/vmwarestructs"
	//"fmt"
	//"gclassec/dao/openstackinsert"

	/*"fmt"
	"gclassec/dao/vmwareinsert"
	"gclassec/dao/hosinsert"
	"gclassec/openstackgov"
	"gclassec/dao/azureinsert"*/
	//"net/http"
	"gclassec/readcredentials"
	//"gclassec/authmanagment"
	//"gclassec/goclientopenstack/openstack"
	//"gclassec/structs/hosstruct"
	//"gclassec/structs/openstackInstance"
	//"gclassec/structs/azurestruct"
	/*"encoding/json"
	"os"*/
	"fmt"
	//"gclassec/structs/vmwarestructs"
	//"gclassec/structs/azurestruct"
	//"gclassec/structs/openstackInstance"
	"gclassec/authmanagment"
	"encoding/json"
	"net/http"
	/*"os"
	"gclassec/errorcodes/errcode"
	"gclassec/dao/vmwareinsert"
	"gclassec/dao/hosinsert"
	"gclassec/openstackgov"
	"gclassec/dao/azureinsert"
*/
	"gclassec/dao/vmwareinsert"
	"gclassec/dao/hosinsert"
	"gclassec/openstackgov"
	"gclassec/dao/azureinsert"
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
var azurecreds = readazurecreds.Configurtion()

var db,err  = gorm.Open(dbtype, c)
type GetstatusInstances struct {
	Provider string
	Type string
	RunningInstances int
	StoppedInstances int
	ClassifierType string
	Classifier string

}


func Getinstancestatus(w http.ResponseWriter, r *http.Request) {

	db.SingularTable(true)
	/*count1 := 0
	count2 := 0
	hos_struct := []hosstruct.HosInstances{}
	//db.Find(&hos_struct)
	db.Where("Status = ?","Active").Find(&hos_struct).Count(&count1)
	db.Where("Status = ?","Shutoff").Find(&hos_struct).Count(&count2)
	fmt.Println(count1)
	fmt.Println(count2)

	poweredoncount := 0
	poweredoffcount := 0
	vmware_struct := []vmwarestructs.VmwareInstances{}
	//db.Find(&hos_struct)
	db.Where("PowerState = ?","poweredOn").Find(&vmware_struct).Count(&poweredoncount)
	fmt.Println(poweredoncount)
	db.Where("PowerState = ?","poweredOff").Find(&vmware_struct).Count(&poweredoffcount)
	fmt.Println(poweredoncount)
	fmt.Println(poweredoffcount)

	poweredoncount1 := 0
	poweredoffcount1 := 0
	azure_struct := []azurestruct.AzureInstances{}
	//db.Find(&hos_struct)
	db.Where("status = ?","VM running").Find(&azure_struct).Count(&poweredoncount1)
	db.Where("PowerState = ?","VM deallocated").Find(&vmware_struct).Count(&poweredoffcount1)
	fmt.Println(poweredoffcount1)
	fmt.Println(poweredoncount1)

	poweredoncount2 := 0
	poweredoffcount2 := 0
	openstack_struct := []openstackInstance.Instances{}
	db.Where("Status = ?","ACTIVE").Find(&openstack_struct).Count(&poweredoncount2)
	db.Where("Status = ?","SHUTDOWN").Find(&openstack_struct).Count(&poweredoffcount2)
	fmt.Println(poweredoffcount2)
	fmt.Println(poweredoncount2)*/

	//getstatusinstances := []GetstatusInstances{}

//tx := db.Begin()
	//var count int
	//db.SingularTable(true)
	/*vmware_struct := []vmwarestructs.VmwareInstances{}
	db.Find(&vmware_struct)
	count1,count2:=0,0
	for _,element1 := range vmware_struct{
		if(element1.PowerState == "poweredOn"){
			count1++
		}else{
			count2++
		}
	}
	hos_struct := []hosstruct.HosInstances{}
	db.Find(&hos_struct)
	poweredoncount,poweredoffcount:=0,0
	for _,element1 := range hos_struct{
		if(element1.Status == "Active"){
			poweredoncount++
		}else{
			poweredoffcount++
		}
	}
	openstack_struct := []openstackInstance.Instances{}
	db.Find(&openstack_struct)
	poweredoncount1,poweredoffcount1:=0,0
	for _,element1 := range openstack_struct{
		if(element1.Status == "ACTIVE"){
			poweredoncount1++
		}else{
			poweredoffcount1++
		}
	}
	azure_struct := []azurestruct.AzureInstances{}
	db.Find(&azure_struct)
	poweredoncount2,poweredoffcount2:=0,0
	for _,element1 := range azure_struct{
		if(element1.Status == "VM running"){
			poweredoncount2++
		}else{
			poweredoffcount2++
		}
	}*/


	//errFind := db.Model(&vmware_struct).Count(&vmware_struct).Where("PowerState = ?","poweredOn")
	//db.Model(&vmware_struct).Where("name = ?", "jinzhu").Count(&count)
	//count := SELECT  count(*) FROM `vmware_instances` WHERE PowerState="poweredOn";
	//  _ = json.NewEncoder(os.Stdout).Encode(db.Table("vmware_instances").Select("COUNT(*)").Where("PowerState = ?","poweredOn"))
/*_ = json.NewEncoder(os.Stdout).Encode(db.Table("vmware_instances").Select("*"))
	rows, err := db.Query("SELECT COUNT(*) as count FROM  table_name")
        fmt.Println("Total count:",checkCount(rows))

	_ = json.NewEncoder(os.Stdout).Encode(db.Where("PowerState = ?", "poweredOn").Find(&vmware_struct))
if errFind != nil {
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
	    ClassifierType:"Username",
	    Classifier:authmanagment.ReadVmwareCredentials().EnvUserName,
    },
    GetstatusInstances {
        Provider: "Azure",
        Type: "Public",
        RunningInstances: poweredoncount2,
        StoppedInstances: poweredoffcount2,
	    ClassifierType:"SubscriptionID",
	    Classifier:azurecreds.SubscriptionId,
    },
		GetstatusInstances {
        Provider: "HOS",
        Type: "Private",
        RunningInstances: poweredoncount,
        StoppedInstances: poweredoffcount,
			 ClassifierType:"ProjectName",
			Classifier:authmanagment.ReadHosCredentials().ProjectName,
    },
		GetstatusInstances {
        Provider: "Openstack",
        Type: "Private",
        RunningInstances: poweredoncount1,
        StoppedInstances: poweredoffcount1,
			ClassifierType:"ProjectName",
			Classifier:authmanagment.ReadOpenstackCredentials().ProjectName,

    },
}
	_ = json.NewEncoder(w).Encode(&basenameOpts)

return

}

