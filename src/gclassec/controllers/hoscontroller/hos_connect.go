package hoscontroller

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"gclassec/goclienthos/compute"
	"encoding/json"
	"strings"
	"github.com/jinzhu/gorm"
	"gclassec/structs/hosstruct"
	"gclassec/goclienthos/util"
	"gclassec/goclienthos/ceilometer"
//	"sid/goclassec/src/github.com/Azure/azure-sdk-for-go/arm/recoveryservices"

	"gclassec/loggers"
	"gclassec/goclienthos/hosstaticdynamic"
	"gclassec/dbmanagement"
	"gclassec/errorcodes/errcode"
	"gclassec/confmanagement/readstructconf"
	"gclassec/confmanagement/readhosconf"
	"gclassec/structs/tagstruct"
	"regexp"
)

type (

    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}
var logger = Loggers.New()
//var counter = 0
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

func (uc UserController) CpuUtilDetails(w http.ResponseWriter, r *http.Request){
        vars := mux.Vars(r)
        id := vars["id"]
        res, err := util.GetCpuUtilDetails(id)
	if err != nil{
		fmt.Println("Error:", err)
		return
	}

	logger.Info(res)
	  _ = json.NewEncoder(w).Encode(&res)
     //   fmt.Fprintf(w,res)

}

func (uc UserController) GetComputeDetails(w http.ResponseWriter, r *http.Request){

	res, err := compute.Compute()
	if err != nil{
		logger.Error("Error: ", errcode.ErrFindDB)
	}
        _ = json.NewEncoder(w).Encode(&res)
	//fmt.Fprintf(w,res)

}

func (uc UserController) Compute(w http.ResponseWriter, r *http.Request){
	tx := db.Begin()
	db.SingularTable(true)

	hos_compute := []hosstruct.HosInstances{}
	response_struct := []hosstruct.HosResponse{}

	tag := []tagstruct.Tags{}
	//create a regex `(?i)hos` will match string contains "hos" case insensitive
	reg := regexp.MustCompile("(?i)hos")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("hos")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("hos")).Find(&tag)

	err := db.Find(&hos_compute).Error
	if err != nil{
		logger.Error("Error: ",err)
		tx.Rollback()
	}

	db.Where("classifier = ?", hoscreds.ProjectName).Find(&hos_compute)

	if readstructconf.ReadStructConfigFile()!=0{
		standardresponse := []hosstruct.StandardizedHos{}

		for i:=0; i<len(hos_compute);i++{
			response := hosstruct.StandardizedHos{}
			response.Vm_Name = hos_compute[i].Vm_Name
			response.InstanceID = hos_compute[i].InstanceID
			response.FlavorName = hos_compute[i].FlavorName
			response.Disk = hos_compute[i].Disk
			response.Ram = hos_compute[i].Ram
			response.VCPU = hos_compute[i].VCPU
			//response.Tagname = hos_compute[i].Tagname
			response.Status = hos_compute[i].Status

			if len(tag) == 0 {
				response.Tagname = "Nil"
			}else {
				for _, el := range tag {
					if hos_compute[i].InstanceID != el.InstanceId {
						response.Tagname = "Nil"
					}else {
						response.Tagname = el.Tagname
						break
					}
				}
			}

			standardresponse = append(standardresponse, response)
		}

		_ = json.NewEncoder(w).Encode(&standardresponse)
	}else {
		for i:=0; i<len(hos_compute);i++ {
			response := hosstruct.HosResponse{}
			response.Vm_Name = hos_compute[i].Vm_Name
			response.InstanceID = hos_compute[i].InstanceID
			response.FlavorName = hos_compute[i].FlavorName
			response.Disk = hos_compute[i].Disk
			response.Ram = hos_compute[i].Ram
			response.VCPU = hos_compute[i].VCPU
			response.FlavorID = hos_compute[i].FlavorID
			response.Status = hos_compute[i].Status
			response.Image = hos_compute[i].Image
			response.AvailabilityZone = hos_compute[i].AvailabilityZone
			response.Classifier = hos_compute[i].Classifier
			response.Deleted = hos_compute[i].Deleted
			response.KeypairName = hos_compute[i].KeypairName
			response.SecurityGroups = hos_compute[i].SecurityGroups

			if len(tag) == 0 {
				response.Tagname = "Nil"
			}else {
				for _, el := range tag {
					if hos_compute[i].InstanceID != el.InstanceId {
						response.Tagname = "Nil"
					}else {
						response.Tagname = el.Tagname
						break
					}
				}
			}

			response_struct = append(response_struct, response)
		}
		_ = json.NewEncoder(w).Encode(&response_struct)
	}

	//_ = json.NewEncoder(w).Encode(db.Find(&hos_compute))

		if err != nil {
			logger.Error("Error: ",err)
			println(err)
		}
	logger.Info("Successful in Compute.")
	tx.Commit()
}

func (uc UserController) GetFlavorsDetails(w http.ResponseWriter, r *http.Request){

	res := compute.Flavors()
	_ = json.NewEncoder(w).Encode(&res)
        //fmt.Fprintf(w,res)
}

func (uc UserController) GetCeilometerStatitics(w http.ResponseWriter, r *http.Request){

	res := ceilometer.GetCpuUtilStatistics()
	logger.Info(res)
        fmt.Fprintf(w,res)

}

func (uc UserController) GetCeilometerDetails(w http.ResponseWriter, r *http.Request){

	//res := compute.GetCeilometerDetail()
	res := ceilometer.GetCeilometerDetail()
	logger.Info(res)
        fmt.Fprintf(w,res)

}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request){
	logger.Info("Hi You Just tested Server ping.")
	fmt.Fprintf(w, "Hi You Just tested Server ping.")
}


func (uc UserController) GetCompleteDetail(w http.ResponseWriter, r *http.Request){

	res := hosstaticdynamic.ComputeWithCPU()
	_ = json.NewEncoder(w).Encode(&res)

}

func (uc UserController) GetCompleteDynamicDetail(w http.ResponseWriter, r *http.Request){

	/*res, err := dynamicdetails.DynamicDetails()
	_ = json.NewEncoder(w).Encode(&res)

	if err != nil{
		return
	}*/


	tx := db.Begin()
	db.SingularTable(true)
	logger := Loggers.New()
	logger.Info("We are Fetching Static Data from Database.")
	hos_struct := []hosstruct.HosDynamicInstances{}

	errFind := db.Find(&hos_struct).Error

	if errFind != nil{
		logger.Error("Error: ", errcode.ErrFindDB)
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&hos_struct))

		if err != nil {
			logger.Error("Error :", err)
			println(err)
		}

	tx.Commit()
	logger.Info("Successful in Fetching Data from Database.")

}