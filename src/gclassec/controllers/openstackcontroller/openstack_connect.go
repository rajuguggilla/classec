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
	"fmt"
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
		//var standardresponse azurestruct.StandardizedAzure
		standardresponse := []openstackInstance.StandardizedOpenstack{}
		/*byte, _ := json.Marshal(&hos_compute)
		fmt.Println(string(byte))
		if err := json.Unmarshal(byte, &standardresponse); err != nil {
			fmt.Println("Error in Unmarshing:==", err)
		}*/
		fmt.Println(len(openstack_struct))
		fmt.Println(&openstack_struct)
		//fmt.Println(&standardresponse)
		for i:=0; i<len(openstack_struct);i++{
			response := openstackInstance.StandardizedOpenstack{}
			response.Name = openstack_struct[i].Name
			response.InstanceID = openstack_struct[i].InstanceID
			response.FlavorID = openstack_struct[i].FlavorID
			response.Flavor = openstack_struct[i].Flavor
			response.Storage = openstack_struct[i].Storage
			response.KeyPairName = openstack_struct[i].KeyPairName
			response.ImageName = openstack_struct[i].ImageName
			response.RAM = openstack_struct[i].RAM
			response.SecurityGroup = openstack_struct[i].SecurityGroup
			response.VCPU = openstack_struct[i].VCPU
			response.Deleted = openstack_struct[i].Deleted
			response.Tagname = openstack_struct[i].Tagname
			response.AvailabilityZone = openstack_struct[i].AvailabilityZone
			response.Status = openstack_struct[i].Status
			response.CreationTime = openstack_struct[i].CreationTime
			response.IPAddress = openstack_struct[i].IPAddress
			response.Volumes = openstack_struct[i].Volumes
			response.InsertionDate = openstack_struct[i].InsertionDate

			standardresponse = append(standardresponse, response)
		}

		fmt.Println(&standardresponse)
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