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
	"gclassec/confmanagement/readazureconf"
)

var logger = Loggers.New()

var dbcredentials = readazureconf.Configurtion()
var dbtype string = dbcredentials.Dbtype
var dbname  string = dbcredentials.Dbname
var dbusername string = dbcredentials.Dbusername
var dbpassword string = dbcredentials.Dbpassword
var dbhostname string = dbcredentials.Dbhostname
var dbport string = dbcredentials.Dbport
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
	}
}