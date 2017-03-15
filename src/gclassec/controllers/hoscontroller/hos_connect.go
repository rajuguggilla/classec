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

func (uc UserController) CpuUtilDetails(w http.ResponseWriter, r *http.Request){
        vars := mux.Vars(r)
        id := vars["id"]
        res := util.GetCpuUtilDetails(id)
	logger.Info(res)
        fmt.Fprintf(w,res)

}
func (uc UserController) GetComputeDetails(w http.ResponseWriter, r *http.Request){

	res := compute.Compute()
        _ = json.NewEncoder(w).Encode(&res)
	//fmt.Fprintf(w,res)

}

func (uc UserController) Compute(w http.ResponseWriter, r *http.Request){
	tx := db.Begin()
	db.SingularTable(true)

	hos_compute := []hosstruct.HosInstances{}

	err := db.Find(&hos_compute).Error

	if err != nil{
		logger.Error("Error: ",err)
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&hos_compute))

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