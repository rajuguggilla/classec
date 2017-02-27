package instancetags

import (
	"net/http"
	"gclassec/loggers"
	"fmt"
	"encoding/json"
	"gclassec/structs/tagstruct"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/dbmanagement"
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

type InstanceTag struct {
	InstanceId string
	Cloud string
	Tagname string
}

func InstanceProvider(w http.ResponseWriter, r *http.Request) {
	inst := []InstanceTag{}

	/*azure_struct := []azurestruct.AzureInstances{}

	ac := goclientazure.NewVirtualMachinesClient(c["AZURE_SUBSCRIPTION_ID"])
	ac.Authorizer = spt
	ls, _ := ac.ListAll()*/
	res := json.NewDecoder(r.Body)
	fmt.Println("------Response Body-------",res)
	fmt.Printf("Type of res : %T", res)
	res.Decode(&inst)
	fmt.Printf("instance",inst)
	for _,r := range inst{
		fmt.Println(r.InstanceId)
		user := tagstruct.Providers{InstanceId:r.InstanceId, Cloud:r.Cloud, Tagname:r.Tagname}
		db.Create(&user)
		db.Model(&user).Updates(&user)
	}
}