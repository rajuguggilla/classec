package overallcpuavg

import (

	//"gclassec/structs/azurestruct"
	"encoding/json"
	"net/http"
	"gclassec/structs/hosstruct"
	"gclassec/structs/vmwarestructs"
	"gclassec/structs/openstackInstance"
	//"regexp"
	//"fmt"
	"github.com/gorilla/mux"
	"strings"
	"gclassec/structs/azurestruct"
	"fmt"
)

var azure_avgcpu = []azurestruct.AzureCpu{}
var hos_avgcpu = []hosstruct.HOSCpu{}
var vmware_avgcpu = []vmwarestructs.VMwareCpu{}
var openstack_avgcpu = []openstackInstance.OpenstackCpu{}
func Getazureoverallcpu (w http.ResponseWriter, r *http.Request) {
	db1.SingularTable(true)
	db1.Find(&azure_avgcpu)
	_ = json.NewEncoder(w).Encode(db1.Find(&azure_avgcpu))

}

func Gethosoverallcpu(w http.ResponseWriter, r *http.Request){
	db1.SingularTable(true)
	db1.Find(&hos_avgcpu)
	_ = json.NewEncoder(w).Encode(db1.Find(&hos_avgcpu))
}

func Getvmwareoverallcpu(w http.ResponseWriter, r *http.Request) {
	db1.SingularTable(true)
	db1.Find(&vmware_avgcpu)
	_ = json.NewEncoder(w).Encode(db1.Find(&vmware_avgcpu))
}

func Getopenstackoverallcpu(w http.ResponseWriter, r *http.Request) {
	db1.SingularTable(true)
	db1.Find(&openstack_avgcpu)
	_ = json.NewEncoder(w).Encode(db1.Find(&openstack_avgcpu))
}

func Getoverallcpubyname(w http.ResponseWriter, r *http.Request){

	db1.SingularTable(true)
	vars := mux.Vars(r)
      	provider := vars["provider"]
	name := vars["name"]
	providername := strings.ToLower(provider)

	if providername == "azure" {
		_ = json.NewEncoder(w).Encode(db1.Where("name = ?",name).Find(&azure_avgcpu))
	}else if providername =="hos" {
		_ = json.NewEncoder(w).Encode(db1.Where("name = ?",name).Find(&hos_avgcpu))
	}else if providername =="vmware"{
		_ = json.NewEncoder(w).Encode(db1.Where("name = ?",name).Find(&vmware_avgcpu))
	}else if providername =="openstack"{
		_ = json.NewEncoder(w).Encode(db1.Where("name = ?",name).Find(&openstack_avgcpu))
	}else{
		fmt.Fprintf(w,"Not a valid provider it should be one of azure,hos,vmware or openstack")
	}

}
/*

func Getazureoverallcpu() {
}*/

