package overallcpuavg

import (
     //  "log"
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
	"gclassec/errorcodes/errcode"
	"gclassec/loggers"
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

var logger = Loggers.New()

func Azurecpu() error{

dynamic := []azurestruct.AzureCpu{}
	db1.SingularTable(true)
	db1.Find(&dynamic)
       rows, err := db.Query("select name,avg(minimum),avg(maximum) , avg(average) from azure_dynamic group by name;")
	if err != nil {
           logger.Error(err)
	   logger.Error(errcode.ErrFindDB)
	   fmt.Println("Error:",errcode.ErrFindDB)
	   return err
        }


defer rows.Close()
for rows.Next() {
       err := rows.Scan(&name,&minimum,&maximum,&average)

       if err != nil {
              logger.Error(err)
	       logger.Error(err)
	       fmt.Println("Error:",errcode.ErrFindDB)
	       return err
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
       logger.Error(err)
	fmt.Println("error:",err)
	return err
}
       if err != nil {
       logger.Error(err)
	fmt.Println("error:",err)
	return err
}

	return nil
}



func HOScpu() error{

dynamic := []hosstruct.HOSCpu{}
	db1.SingularTable(true)
	db1.Find(&dynamic)
       rows, err := db.Query("select Name,avg(Min),avg(Max) , avg(Avg) from hos_dynamic_instances group by Name;")
	if err != nil {
		logger.Error(errcode.ErrFindDB)
		fmt.Println("Error:",errcode.ErrFindDB)
		//log.Println(err)
		return err
}

defer rows.Close()
for rows.Next() {
       err := rows.Scan(&name,&minimum,&maximum,&average)

       if err != nil {
              logger.Error("error:", errcode.ErrFindDB)
	       fmt.Println("error:", errcode.ErrFindDB)
	       return err
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
       logger.Error(err)
	fmt.Println("error:",err)
	return err
}
       if err != nil {
        logger.Error(err)
	fmt.Println("error:",err)
	return err

}
	return nil
}


func VMwarecpu() error{

dynamic := []vmwarestructs.VMwareCpu{}
	db1.SingularTable(true)
	db1.Find(&dynamic)
       rows, err := db.Query("select Name,avg(MinCpuUsage),avg(MaxCpuUsage) , avg(AvgCpuUsage) from vmware_dynamic_details group by Name;")
	if err != nil {
       		logger.Error(errcode.ErrFindDB)
		fmt.Println("Error:",errcode.ErrFindDB)
		//log.Println(err)
		return err
}
defer rows.Close()
for rows.Next() {
       err := rows.Scan(&name,&minimum,&maximum,&average)

       if err != nil {
	       logger.Error(errcode.ErrFindDB)
		fmt.Println("Error:",errcode.ErrFindDB)
		//log.Println(err)
		return err
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
			db1.Where("Name = ?",name).Find(&dynamic)
			if(len(dynamic) ==0){
				dynamic := vmwarestructs.VMwareCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
				db1.Create(&dynamic)

			}else{
				dynamic := vmwarestructs.VMwareCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
				db1.Model(&dynamic).Where("Name =?",element.Name).Updates(dynamic)
			}
		}
	}
}
	err = rows.Err()
if err != nil {
        logger.Error(err)
	fmt.Println("error:",err)
	return err
}
       if err != nil {
       logger.Error(err)
	fmt.Println("error:",err)
	return err

}
	return nil
}



func Openstackcpu() error{

dynamic := []openstackInstance.OpenstackCpu{}
	db1.SingularTable(true)
	db1.Find(&dynamic)
       rows, err := db.Query("select Vm_Name,avg(Min),avg(Max) , avg(Avg) from dynamic_instances group by Vm_Name;")
	if err != nil {
       		logger.Error(errcode.ErrFindDB)
		fmt.Println("Error:",errcode.ErrFindDB)
		//log.Println(err)
		return err
}
defer rows.Close()
for rows.Next() {
       err := rows.Scan(&name,&minimum,&maximum,&average)

       if err != nil {
	       logger.Error(errcode.ErrFindDB)
		fmt.Println("Error:",errcode.ErrFindDB)
		//log.Println(err)
		return err
       }
       fmt.Println(minimum)
        fmt.Println(maximum)
	fmt.Println(name)
	fmt.Println(average)
	if (len(dynamic)== 0){
		dynamic := openstackInstance.OpenstackCpu{Name:name,Minimum:minimum,Maximum:maximum,Average:average}
   			db1.Create(&dynamic)
	}else {
		for _, element := range dynamic {
			if name == element.Name {
				dynamic := openstackInstance.OpenstackCpu{Name:name, Minimum:minimum, Maximum:maximum, Average:average}
				db1.Model(&dynamic).Where("name =?", element.Name).Updates(dynamic)
			} else {
				dynamic := openstackInstance.OpenstackCpu{Name:name, Minimum:minimum, Maximum:maximum, Average:average}
				db1.Create(&dynamic)
			}
		}
	}
}
	err = rows.Err()
if err != nil {
        logger.Error(err)
	fmt.Println("error:",err)
	return err
}
       if err != nil {
       logger.Error(err)
	fmt.Println("error:",err)
	return err
}
	return nil
}
