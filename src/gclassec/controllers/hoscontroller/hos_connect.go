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
	"gclassec/confmanagement/readazureconf"
	"gclassec/goclienthos/util"
	"gclassec/goclienthos/ceilometer"
)

type (

    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}

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

func (uc UserController) CpuUtilDetails(w http.ResponseWriter, r *http.Request){
        vars := mux.Vars(r)
        id := vars["id"]
        res := util.GetCpuUtilDetails(id)
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

		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&hos_compute))

		if err != nil {
			println(err)
		}

	tx.Commit()
}

func (uc UserController) GetFlavorsDetails(w http.ResponseWriter, r *http.Request){

	res := compute.Flavors()
	_ = json.NewEncoder(w).Encode(&res)
        //fmt.Fprintf(w,res)
}

func (uc UserController) GetCeilometerStatitics(w http.ResponseWriter, r *http.Request){

	res := ceilometer.GetCpuUtilStatistics()
        fmt.Fprintf(w,res)

}

func (uc UserController) GetCeilometerDetails(w http.ResponseWriter, r *http.Request){

	//res := compute.GetCeilometerDetail()
	res := ceilometer.GetCeilometerDetail()
        fmt.Fprintf(w,res)

}

func (uc UserController) Index(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Hi You Just tested Server ping.")
}