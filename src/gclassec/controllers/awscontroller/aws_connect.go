package awscontroller

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gclassec/structs/awsstructs"
	"gclassec/loggers"
	"gclassec/dbmanagement"
)

type (
    // UserController represents the controller for operating on the User resource
    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}


var dbtype string = dbmanagement.ENVdbtype
var dbname  string = dbmanagement.ENVdbnameaws
var dbusername string = dbmanagement.ENVdbusername
var dbpassword string = dbmanagement.ENVdbpassword
var dbhostname string = dbmanagement.ENVdbhostname
var dbport string = dbmanagement.ENVdbport
var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))

var db,err  = gorm.Open(dbtype, c)
var logger = Loggers.New()
func (uc UserController) GetDetailsById(w http.ResponseWriter, r *http.Request) {
	dbObj := []awsstructs.Rds_dynamic{}

	//id := p.ByName("id")

	vars := mux.Vars(r)
	id := vars["id"]

	db.SingularTable(true)

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj, "DBInstanceIdentifier = ?", id))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
		logger.Error(err)
	}
	//db.Close()
}

func (uc UserController) GetDB(w http.ResponseWriter, r *http.Request) {
	tx := db.Begin()
	dbObj := []awsstructs.Rds_dynamic{}

	queryValue1 := r.URL.Query().Get("CPUUtilization_max")

	queryValue2 := r.URL.Query().Get("DatabaseConnections_max")

	println(queryValue1)
	logger.Info(queryValue1)
	println(queryValue2)
	logger.Info(queryValue2)
	db.SingularTable(true)

	err := db.Where("CPUUtilization_max < ? AND DatabaseConnections_max = ?", queryValue1, queryValue2).Find(&dbObj).Error
	if err != nil{
		logger.Error("Error: ",err)
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Where("CPUUtilization_max < ? AND DatabaseConnections_max = ?", queryValue1, queryValue2).Find(&dbObj))

	// Ping function checks database connectivity
	err = db.DB().Ping()
	if err != nil {
		logger.Error("Error: ",err)
		panic(err)
	}
	logger.Info("Successful in GetDB_AWS")
	tx.Commit()

}

//Get pricing depending on instance type



func (uc UserController) GetPrice(w http.ResponseWriter, r *http.Request) {
	tx := db.Begin()
	dbObj := []awsstructs.Vw_rds{}


	db.SingularTable(true)
	err := db.Find(&dbObj).Error
	if err != nil{
		logger.Error("Error: ",err)
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&dbObj))

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		logger.Error("Error: ",err)
		panic(err)
	}
	logger.Info("Successful in GetPrice_AWS")
	tx.Commit()
}

func (uc UserController) GetDetails(w http.ResponseWriter, r *http.Request){

	tx := db.Begin()
	db.SingularTable(true)

	rds_dynamic := []awsstructs.Rds_dynamic{}

	err := db.Find(&rds_dynamic).Error

	if err != nil{
		logger.Error("Error: ",err)
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&rds_dynamic))

		if err != nil {
			logger.Error("Error: ",err)
			println(err)
		}
	logger.Info("Successful in GetDetails_AWS")
	tx.Commit()
}

