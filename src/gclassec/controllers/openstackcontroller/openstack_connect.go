package openstackcontroller

import(
	"strings"
	"github.com/jinzhu/gorm"
	"net/http"
	"encoding/json"

	"gclassec/structs/openstackInstance"
	"gclassec/loggers"
	"gclassec/errorcodes/errcode"
	"gclassec/dbmanagement"
	"gclassec/confmanagement/readstructconf"
)
type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}
var dbtype string = dbmanagement.ENVdbtype
var dbname  string = dbmanagement.ENVdbnamegodb
var dbusername string = dbmanagement.ENVdbusername
var dbpassword string = dbmanagement.ENVdbpassword
var dbhostname string = dbmanagement.ENVdbhostname
var dbport string = dbmanagement.ENVdbport
var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))

var db,err  = gorm.Open(dbtype, c)

func (uc UserController) GetDetailsOpenstack(w http.ResponseWriter, r *http.Request){

	tx := db.Begin()
	db.SingularTable(true)
	logger := Loggers.New()
	logger.Info("We are Fetching Static Data from Database.")
	openstack_struct := []openstackInstance.Instances{}

	errFind := db.Find(&openstack_struct).Error

	if errFind != nil{
		logger.Error("Error: ", errcode.ErrFindDB)
		tx.Rollback()
	}

	db.Find(&openstack_struct)

	if readstructconf.ReadStructConfigFile()!=0{
		standardresponse := []openstackInstance.StandardizedOpenstack{}

		for i:=0; i<len(openstack_struct);i++{
			response := openstackInstance.StandardizedOpenstack{}
			response.Name = openstack_struct[i].Name
			response.InstanceID = openstack_struct[i].InstanceID
			response.Flavor = openstack_struct[i].Flavor
			response.Storage = openstack_struct[i].Storage
			response.RAM = openstack_struct[i].RAM
			response.VCPU = openstack_struct[i].VCPU
			response.Tagname = openstack_struct[i].Tagname
			response.Status = openstack_struct[i].Status

			standardresponse = append(standardresponse, response)
		}

		_ = json.NewEncoder(w).Encode(&standardresponse)
	}else {
		_ = json.NewEncoder(w).Encode(&openstack_struct)
	}

	//_ = json.NewEncoder(w).Encode(db.Find(&openstack_struct))

		if err != nil {
			logger.Error("Error :", err)
			println(err)
		}

	tx.Commit()
	logger.Info("Successful in Fetching Data from Database.")
}