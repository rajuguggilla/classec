package main

import (
       "log"
       "gclassec/dbmanagement"
       "strings"
        "database/sql"
       _ "github.com/go-sql-driver/mysql"
       "fmt"
	"gclassec/structs/azurestruct"
	"github.com/jinzhu/gorm"
	"gclassec/structs/hosstruct"
	"gclassec/structs/vmwarestructs"
	"gclassec/structs/openstackInstance"
)

var dbtype string = dbmanagement.ENVdbtype
var dbname  string = dbmanagement.ENVdbnamegodb
var dbusername string = dbmanagement.ENVdbusername
var dbpassword string = dbmanagement.ENVdbpassword
var dbhostname string = dbmanagement.ENVdbhostname
var dbport string = dbmanagement.ENVdbport
var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}
var b1 []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))
var c1 string = (strings.Join(b1,""))
var db,err  = sql.Open(dbtype, c)
var db1,err1  = gorm.Open(dbtype, c1)



var (
name string
       minimum float64
       maximum float64
       average float64
)

func main(){
	Azurecpu()
	HOScpu()
	Openstackcpu()
	//VMwarecpu()
}
func Azurecpu(){

dynamic := []azurestruct.AzureCpu{}
	db1.SingularTable(true)
	db1.Find(&dynamic)
       rows, err := db.Query("select name,avg(minimum),avg(maximum) , avg(average) from azure_dynamic group by name;")
	if err != nil {
       log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
       err := rows.Scan(&name,&minimum,&maximum,&average)

       if err != nil {
              log.Fatal(err)
       }
       fmt.Println(minimum)
        fmt.Println(maximum)
	fmt.Println(name)
	fmt.Println(average)
	if (len(dynamic)== 0){
		dynamic := azurestruct.AzureCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
   			db1.Create(&dynamic)
	}else {
		for _, element := range dynamic {
				db1.Where("name = ?",name).Find(&dynamic)
			if(len(dynamic)==0){
				dynamic := azurestruct.AzureCpu{Name:name, Minimum:minimum, Maximum:maximum, Average:average}
				db1.Create(&dynamic)
			} else {
			dynamic := azurestruct.AzureCpu{Name:name, Minimum:minimum, Maximum:maximum, Average:average}
				db1.Model(&dynamic).Where("name =?", element.Name).Updates(dynamic)

			}
		}
	}
}
	err = rows.Err()
if err != nil {
       log.Fatal(err)
}
       if err != nil {
    log.Fatal(err)
	       db.Close()
	       db1.Close()
}


}
func HOScpu(){

dynamic := []hosstruct.HOSCpu{}
	db1.SingularTable(true)
	db1.Find(&dynamic)
       rows, err := db.Query("select Name,avg(Min),avg(Max) , avg(Avg) from hos_dynamic_instances group by Name;")
	if err != nil {
       log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
       err := rows.Scan(&name,&minimum,&maximum,&average)

       if err != nil {
              log.Fatal(err)
       }
       fmt.Println(minimum)
        fmt.Println(maximum)
	fmt.Println(name)
	fmt.Println(average)
	if (len(dynamic)== 0){
		dynamic := hosstruct.HOSCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
   			db1.Create(&dynamic)
	}else {
		for _, element := range dynamic {
			db1.Where("name = ?",name).Find(&dynamic)
			if(len(dynamic)==0){
				dynamic := hosstruct.HOSCpu{Name:name, Minimum:minimum, Maximum:maximum, Average:average}
				db1.Create(&dynamic)

			} else {
				dynamic := hosstruct.HOSCpu{Name:name, Minimum:minimum, Maximum:maximum, Average:average}
				db1.Model(&dynamic).Where("name =?", element.Name).Updates(dynamic)
			}
		}
	}
}
	err = rows.Err()
if err != nil {
       log.Fatal(err)
}
       if err != nil {
    log.Fatal(err)
	       db.Close()
	       db1.Close()
}


}
func VMwarecpu(){

dynamic := []vmwarestructs.VMwareCpu{}
	db1.SingularTable(true)
	db1.Find(&dynamic)
       rows, err := db.Query("select Name,avg(MinCpuUsage),avg(MaxCpuUsage) , avg(AvgCpuUsage) from vmware_dynamic_details group by name;")
	if err != nil {
       log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
       err := rows.Scan(&name,&minimum,&maximum,&average)

       if err != nil {
              log.Fatal(err)
       }
       fmt.Println(minimum)
        fmt.Println(maximum)
	fmt.Println(name)
	fmt.Println(average)
	if (len(dynamic)== 0){
		dynamic := vmwarestructs.VMwareCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
   			db1.Create(&dynamic)
	}else{
		for _,element := range dynamic{
			db1.Where("name = ?",name).Find(&dynamic)
			if(len(dynamic) ==0){
				dynamic := vmwarestructs.VMwareCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
				db1.Create(&dynamic)

			}else{
				dynamic := vmwarestructs.VMwareCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
				db1.Model(&dynamic).Where("name =?",element.Name).Updates(dynamic)
			}
		}
	}
}
	err = rows.Err()
if err != nil {
       log.Fatal(err)
}
       if err != nil {
    log.Fatal(err)
	       db.Close()
	       db1.Close()
}


}
func Openstackcpu(){

dynamic := []openstackInstance.OpenstackCpu{}
	db1.SingularTable(true)
	db1.Find(&dynamic)
       rows, err := db.Query("select Vm_Name,avg(Min),avg(Max) , avg(Avg) from dynamic_instances group by Vm_Name;")
	if err != nil {
       log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
       err := rows.Scan(&name,&minimum,&maximum,&average)

       if err != nil {
              log.Fatal(err)
       }
       fmt.Println(minimum)
        fmt.Println(maximum)
	fmt.Println(name)
	fmt.Println(average)
	if (len(dynamic)== 0){
		dynamic := openstackInstance.OpenstackCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
   			db1.Create(&dynamic)
	}
	for _,element := range dynamic{
		if name == element.Name{
			dynamic := openstackInstance.OpenstackCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
   			db1.Model(&dynamic).Where("name =?",element.Name).Updates(dynamic)
		}else{
			dynamic := openstackInstance.OpenstackCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
   			db1.Create(&dynamic)
		}
	}
}
	err = rows.Err()
if err != nil {
       log.Fatal(err)
}
       if err != nil {
    log.Fatal(err)

}
	db.Close()
	db1.Close()


}
