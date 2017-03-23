package overallcpuavg

import (
	"encoding/json"
	"net/http"
	"gclassec/structs/hosstruct"
	"gclassec/structs/vmwarestructs"
	"gclassec/structs/openstackInstance"
	"gclassec/structs/azurestruct"
)

var azure_avgcpu = []azurestruct.AzureCpu{}
var hos_avgcpu = []hosstruct.HOSCpu{}
var vmware_avgcpu = []vmwarestructs.VMwareCpu{}
var openstack_avgcpu = []openstackInstance.OpenstackCpu{}
func Getazureoverallcpu (w http.ResponseWriter, r *http.Request) {
	db1.SingularTable(true)
	instanceid := r.URL.Query().Get("instanceid")
	if instanceid == ""{
		_ = json.NewEncoder(w).Encode(db1.Find(&azure_avgcpu))
	}else{
		_ = json.NewEncoder(w).Encode(db1.Where("vmid = ?",instanceid).Find(&azure_avgcpu))
	}}


func Gethosoverallcpu(w http.ResponseWriter, r *http.Request){
	db1.SingularTable(true)
	instanceid := r.URL.Query().Get("instanceid")
	if instanceid == ""{
		_ = json.NewEncoder(w).Encode(db1.Find(&hos_avgcpu))
	}else{
		_ = json.NewEncoder(w).Encode(db1.Where("vmid = ?",instanceid).Find(&hos_avgcpu))
	}}

func Getvmwareoverallcpu(w http.ResponseWriter, r *http.Request) {
	db1.SingularTable(true)
	instanceid := r.URL.Query().Get("instanceid")
	if instanceid == ""{
		_ = json.NewEncoder(w).Encode(db1.Find(&vmware_avgcpu))
	}else{
		_ = json.NewEncoder(w).Encode(db1.Where("vmid = ?",instanceid).Find(&vmware_avgcpu))
	}}

func Getopenstackoverallcpu(w http.ResponseWriter, r *http.Request) {
	db1.SingularTable(true)
	instanceid := r.URL.Query().Get("instanceid")
	if instanceid == ""{
		_ = json.NewEncoder(w).Encode(db1.Find(&openstack_avgcpu))
	}else{
		_ = json.NewEncoder(w).Encode(db1.Where("vmid = ?",instanceid).Find(&openstack_avgcpu))
	}


}



