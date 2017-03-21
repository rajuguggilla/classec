package vmwarecontroller

import (
	 _ "github.com/go-sql-driver/mysql"
	"gclassec/confmanagement/vmwareconf"
	"fmt"
	"net/url"
	"os"
	"strings"
	"github.com/jinzhu/gorm"
	"net/http"
	"gclassec/structs/vmwarestructs"
	"encoding/json"
	"context"
	"flag"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"text/tabwriter"
	"github.com/vmware/govmomi/units"
	"gclassec/loggers"
	"github.com/vmware/govmomi/vim25/types"

	"gclassec/errorcodes/errcode"
	"gclassec/structs/tagstruct"
	"regexp"
	"gclassec/dbmanagement"
	"gclassec/confmanagement/readstructconf"
)

//const (
//	EnvURL = "https://110.110.110.140:443/sdk"
//	EnvUserName = "administrator@vsphere.local"
//	EnvPassword = "Vcenter#1234"
//	EnvInsecure = "true"
//)
var logger = Loggers.New()
var vmwarecreds = vmwareconf.Configurtion()
var EnvURL string = vmwarecreds.EnvURL
var EnvUserName  string = vmwarecreds.EnvUserName
var EnvPassword string = vmwarecreds.EnvPassword
var EnvInsecure string = vmwarecreds.EnvInsecure

type (

    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}


var urlDescription = fmt.Sprintf("ESX or vCenter URL [%s]", EnvURL)
var ENVurlFlag = flag.String("url", EnvURL, urlDescription)

var insecureDescription = fmt.Sprintf("Don't verify the server's certificate chain [%s]", EnvInsecure)
var ENVinsecureFlag = flag.Bool("insecure", true, insecureDescription)

func ProcessOverride(u *url.URL) {
	//envUsername := os.Getenv(envUserName)
	//envPassword := os.Getenv(envPassword)

	// Override username if provided
	if EnvUserName != "" {
		var password string
		var ok bool

		if u.User != nil {
			password, ok = u.User.Password()
		}

		if ok {
			u.User = url.UserPassword(EnvUserName, password)
		} else {
			u.User = url.User(EnvUserName)
		}
	}

	// Override password if provided
	if EnvPassword != "" {
		var username string

		if u.User != nil {
			username = u.User.Username()
		}

		u.User = url.UserPassword(username, EnvPassword)
	}
}

func exit(err error) {
	logger.Error(os.Stderr, "Error: %s\n", err)
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	//os.Exit(1)
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

func   (uc UserController) GetStaticDynamicVcenterDetails(w http.ResponseWriter, r *http.Request)(){
       ctx, cancel := context.WithCancel(context.Background())
       defer cancel()
//       var insecureFlag = flag.Bool("insecure", true, insecureDescription)
       fmt.Println(*ENVinsecureFlag)
	logger.Debug(*ENVinsecureFlag)
       //fmt.Println("Inside Vcenter get details !!!!!!!!! 1")

       flag.Parse()
//	var urlFlag = flag.String("url", EnvURL, urlDescription)
  //     // Parse URL from string
       u, err := url.Parse(*ENVurlFlag)
       if err != nil {
	       logger.Error("Error: ",err)
              exit(err)
       }

       // Override username and/or password as required
       ProcessOverride(u)

       // Connect and log in to ESX or vCenter
       c, err := govmomi.NewClient(ctx, u, *ENVinsecureFlag)
       if err != nil {
	       logger.Error("VMWare : ",errcode.ErrAuth)
	       fmt.Println("VMWare : ", errcode.ErrAuth)
              exit(err)
       }

       f := find.NewFinder(c.Client, true)

       // Find one and only datacenter
       dc, err := f.DefaultDatacenter(ctx)
       if err != nil {
	       logger.Error("Error: ",err)
              exit(err)
       }

       // Make future calls local to this datacenterth
       f.SetDatacenter(dc)

       // Find virtual machines in datacenter
       vms, err := f.VirtualMachineList(ctx, "*")
       fmt.Println(vms)
	logger.Info(vms)

       pc := property.DefaultCollector(c.Client)

       var refv []types.ManagedObjectReference
       for _, ds := range vms {
              refv = append(refv, ds.Reference())
       }

       // Retrieve name property for all vms
       var vmt []mo.VirtualMachine
       err = pc.Retrieve(ctx, refv, []string{"summary"}, &vmt)
       if err != nil {
	       logger.Error("Error: ",err)
              exit(err)
       }

	tx := db.Begin()
	db.SingularTable(true)

	output := vmwarestructs.StaticDynamicValues{}
	tag := []tagstruct.Tags{}

	//create a regex `(?i)vmware` will match string contains "vmware" case insensitive
	reg := regexp.MustCompile("(?i)vmware")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("VMWARE")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("VMWARE")).Find(&tag)

       // Print summary
       tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

       logger.Info("Virtual machines found:", len(vmt))
       logger.Info(w,   "{\"Value\" : [")
	fmt.Fprintf(w, "{\"Value\" : [")
       //for _, vm := range vmt {
       //       //fmt.Fprintf(tw, "%s\n", vm.Name)
       //       logger.Info("VM Name : ", vm.Summary.Config.Name)
       //       logger.Info("Overall CPU : ", vm.Summary.QuickStats.OverallCpuUsage)
       //       logger.Info("Guest memory : ", vm.Summary.QuickStats.GuestMemoryUsage)
       //       logger.Info("Committed storage : ", units.ByteSize(vm.Summary.Storage.Committed))
       //       //_ = json.NewEncoder(os.Stdout).Encode(&vm)
       //       output := vmwarestructs.StaticDynamicValues{VMName:vm.Summary.Config.Name,Uuid:vm.Summary.Config.Uuid,MemorySizeMB:vm.Summary.Config.MemorySizeMB,PowerState:string(vm.Summary.Runtime.PowerState),NumCpu:vm.Summary.Config.NumCpu,GuestFullName:vm.Summary.Config.GuestFullName,IpAddress:vm.Summary.Guest.IpAddress,OverallCpuUsage:vm.Summary.QuickStats.OverallCpuUsage,GuestMemoryUsage:vm.Summary.QuickStats.GuestMemoryUsage,StorageCommitted:float32(vm.Summary.Storage.Committed)/float32(1024*1024*1024),MemoryOverhead  :vm.Summary.Runtime.MemoryOverhead ,MaxCpuUsage :vm.Summary.Runtime.MaxCpuUsage,Uncommitted:vm.Summary.Storage.Uncommitted,Unshared:vm.Summary.Storage.Unshared}
       //       _ = json.NewEncoder(w).Encode(output)
	//       logger.Info(",")
       //      fmt.Fprintf(w, ",")
       //}
	for i:=0; i<len(vmt); i++ {
              //fmt.Fprintf(tw, "%s\n", vm.Name)
              logger.Info("VM Name : ", vmt[i].Summary.Config.Name)
              logger.Info("Overall CPU : ", vmt[i].Summary.QuickStats.OverallCpuUsage)
              logger.Info("Guest memory : ", vmt[i].Summary.QuickStats.GuestMemoryUsage)
              logger.Info("Committed storage : ", units.ByteSize(vmt[i].Summary.Storage.Committed))
              //_ = json.NewEncoder(os.Stdout).Encode(&vm)
		fmt.Println("Tag : ", tag)
		if len(tag) == 0 {
			output = vmwarestructs.StaticDynamicValues{VMName:vmt[i].Summary.Config.Name, Uuid:vmt[i].Summary.Config.Uuid, MemorySizeMB:vmt[i].Summary.Config.MemorySizeMB, PowerState:string(vmt[i].Summary.Runtime.PowerState), NumCpu:vmt[i].Summary.Config.NumCpu, GuestFullName:vmt[i].Summary.Config.GuestFullName, IpAddress:vmt[i].Summary.Guest.IpAddress, OverallCpuUsage:vmt[i].Summary.QuickStats.OverallCpuUsage, GuestMemoryUsage:vmt[i].Summary.QuickStats.GuestMemoryUsage, StorageCommitted:float32(vmt[i].Summary.Storage.Committed) / float32(1024 * 1024 * 1024), MemoryOverhead:vmt[i].Summary.Runtime.MemoryOverhead, MaxCpuUsage:vmt[i].Summary.Runtime.MaxCpuUsage, Uncommitted:vmt[i].Summary.Storage.Uncommitted, Unshared:vmt[i].Summary.Storage.Unshared, Tagname:"Nil"}
			_ = json.NewEncoder(w).Encode(output)
		}else {
			for _, el := range tag {
				fmt.Println("In tag loop")
				if vmt[i].Summary.Config.Uuid != el.InstanceId {
					output = vmwarestructs.StaticDynamicValues{VMName:vmt[i].Summary.Config.Name, Uuid:vmt[i].Summary.Config.Uuid, MemorySizeMB:vmt[i].Summary.Config.MemorySizeMB, PowerState:string(vmt[i].Summary.Runtime.PowerState), NumCpu:vmt[i].Summary.Config.NumCpu, GuestFullName:vmt[i].Summary.Config.GuestFullName, IpAddress:vmt[i].Summary.Guest.IpAddress, OverallCpuUsage:vmt[i].Summary.QuickStats.OverallCpuUsage, GuestMemoryUsage:vmt[i].Summary.QuickStats.GuestMemoryUsage, StorageCommitted:float32(vmt[i].Summary.Storage.Committed) / float32(1024 * 1024 * 1024), MemoryOverhead:vmt[i].Summary.Runtime.MemoryOverhead, MaxCpuUsage:vmt[i].Summary.Runtime.MaxCpuUsage, Uncommitted:vmt[i].Summary.Storage.Uncommitted, Unshared:vmt[i].Summary.Storage.Unshared, Tagname:"Nil"}
				}else {
					output = vmwarestructs.StaticDynamicValues{VMName:vmt[i].Summary.Config.Name, Uuid:vmt[i].Summary.Config.Uuid, MemorySizeMB:vmt[i].Summary.Config.MemorySizeMB, PowerState:string(vmt[i].Summary.Runtime.PowerState), NumCpu:vmt[i].Summary.Config.NumCpu, GuestFullName:vmt[i].Summary.Config.GuestFullName, IpAddress:vmt[i].Summary.Guest.IpAddress, OverallCpuUsage:vmt[i].Summary.QuickStats.OverallCpuUsage, GuestMemoryUsage:vmt[i].Summary.QuickStats.GuestMemoryUsage, StorageCommitted:float32(vmt[i].Summary.Storage.Committed) / float32(1024 * 1024 * 1024), MemoryOverhead:vmt[i].Summary.Runtime.MemoryOverhead, MaxCpuUsage:vmt[i].Summary.Runtime.MaxCpuUsage, Uncommitted:vmt[i].Summary.Storage.Uncommitted, Unshared:vmt[i].Summary.Storage.Unshared, Tagname:el.Tagname}
				}
				_ = json.NewEncoder(w).Encode(output)
			}
		}
	      if i< (len(vmt)-1){
		     logger.Info(",")
		     fmt.Fprintf(w, ",")
	     }
       }
	logger.Info("]}")
      fmt.Fprintf(w, "]}")
//,PowerState:vm.Summary.Runtime.PowerState
       tw.Flush()
	tx.Commit()
}

func   (uc UserController) GetVcenterDetails(w http.ResponseWriter, r *http.Request)() {
	tx := db.Begin()
	db.SingularTable(true)
	vmware_struct := []vmwarestructs.VmwareInstances{}
	errFind := db.Find(&vmware_struct).Error
	if errFind != nil {
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}

	db.Where("classifier = ?",vmwarecreds.EnvUserName).Find(&vmware_struct)

	//_ = json.NewEncoder(w).Encode(db.Where("classifier = ?",vmwarecreds.EnvUserName).Find(&vmware_struct))

	if readstructconf.ReadStructConfigFile()!=0{
		standardresponse := []vmwarestructs.StandardizedVmware{}

		for i:=0; i<len(vmware_struct);i++{
			response := vmwarestructs.StandardizedVmware{}
			response.Name = vmware_struct[i].Name
			response.Uuid = vmware_struct[i].Uuid
			response.PowerState = vmware_struct[i].PowerState
			response.MemorySizeMB = vmware_struct[i].MemorySizeMB
			response.Tagname = vmware_struct[i].Tagname
			response.NumofCPU = vmware_struct[i].NumofCPU
			response.StorageCommitted = vmware_struct[i].StorageCommitted

			standardresponse = append(standardresponse, response)
		}
		_ = json.NewEncoder(w).Encode(&standardresponse)
	}else {
		_ = json.NewEncoder(w).Encode(&vmware_struct)
	}

	if err != nil {
		logger.Error("Error: ",err)
		println(err)
	}
	logger.Info("Successful in GetVcenterDetails.")
	tx.Commit()
}

func   (uc UserController) GetDynamicVcenterDetails(w http.ResponseWriter, r *http.Request)(){
       ctx, cancel := context.WithCancel(context.Background())
       defer cancel()
//       var insecureFlag = flag.Bool("insecure", true, insecureDescription)
       fmt.Println(*ENVinsecureFlag)
	logger.Debug(*ENVinsecureFlag)
       //fmt.Println("Inside Vcenter get details !!!!!!!!! 1")

       flag.Parse()
//	var urlFlag = flag.String("url", EnvURL, urlDescription)
  //     // Parse URL from string
       u, err := url.Parse(*ENVurlFlag)
       if err != nil {
	       logger.Error("Error: ",err)
              exit(err)
       }

       // Override username and/or password as required
       ProcessOverride(u)

       // Connect and log in to ESX or vCenter
       c, err := govmomi.NewClient(ctx, u, *ENVinsecureFlag)
       if err != nil {
	       logger.Error("VMWare: ", errcode.ErrAuth)
	       fmt.Println("VMWare : ", errcode.ErrAuth)
              exit(err)
       }

       f := find.NewFinder(c.Client, true)

       // Find one and only datacenter
       dc, err := f.DefaultDatacenter(ctx)
       if err != nil {
	       logger.Error("Error: ",err)
              exit(err)
       }

       // Make future calls local to this datacenterth
       f.SetDatacenter(dc)

       // Find virtual machines in datacenter
       vms, err := f.VirtualMachineList(ctx, "*")
       fmt.Println(vms)
	logger.Info(vms)

       pc := property.DefaultCollector(c.Client)

       var refv []types.ManagedObjectReference
       for _, ds := range vms {
              refv = append(refv, ds.Reference())
       }

       // Retrieve name property for all vms
       var vmt []mo.VirtualMachine
       err = pc.Retrieve(ctx, refv, []string{"summary"}, &vmt)
       if err != nil {
	       logger.Error("Error: ",err)
              exit(err)
       }

       // Print summary
       tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

       logger.Info("Virtual machines found:", len(vmt))
	logger.Info(w,"{\"Value\" : [")

       fmt.Fprintf(w, "{\"Value\" : [")
       //for _, vm := range vmt {
       //       //fmt.Fprintf(tw, "%s\n", vm.Name)
       //       logger.Info("VM Name : ", vm.Summary.Config.Name)
       //       logger.Info("Overall CPU : ", vm.Summary.QuickStats.OverallCpuUsage)
       //       logger.Info("Guest memory : ", vm.Summary.QuickStats.GuestMemoryUsage)
       //       logger.Info("Committed storage : ", units.ByteSize(vm.Summary.Storage.Committed))
       //       //_ = json.NewEncoder(os.Stdout).Encode(&vm)
       //       output := vmwarestructs.DynamicValues{VMName:vm.Summary.Config.Name,OverallCpuUsage:vm.Summary.QuickStats.OverallCpuUsage,GuestMemoryUsage:vm.Summary.QuickStats.GuestMemoryUsage,StorageCommitted:float32(vm.Summary.Storage.Committed)/float32(1024*1024*1024)}
       //       _ = json.NewEncoder(w).Encode(output)
	//       logger.Info(",")
	//
       //      fmt.Fprintf(w, ",")
       //}
	for i:=0; i<len(vmt); i++ {
              //fmt.Fprintf(tw, "%s\n", vm.Name)
              logger.Info("VM Name : ", vmt[i].Summary.Config.Name)
              logger.Info("Overall CPU : ", vmt[i].Summary.QuickStats.OverallCpuUsage)
              logger.Info("Guest memory : ", vmt[i].Summary.QuickStats.GuestMemoryUsage)
              logger.Info("Committed storage : ", units.ByteSize(vmt[i].Summary.Storage.Committed))
              //_ = json.NewEncoder(os.Stdout).Encode(&vm)
              output := vmwarestructs.DynamicValues{VMName:vmt[i].Summary.Config.Name,OverallCpuUsage:vmt[i].Summary.QuickStats.OverallCpuUsage,GuestMemoryUsage:vmt[i].Summary.QuickStats.GuestMemoryUsage,StorageCommitted:float32(vmt[i].Summary.Storage.Committed)/float32(1024*1024*1024)}
              _ = json.NewEncoder(w).Encode(output)

	     if i< (len(vmt)-1){
		     logger.Info(",")
		     fmt.Fprintf(w, ",")
	     }

       }
	logger.Info("]}")
      fmt.Fprintf(w, "]}")

       tw.Flush()
}
func   (uc UserController) GetDynamicVcenterUpdateDetails(w http.ResponseWriter, r *http.Request)() {
	tx := db.Begin()
	db.SingularTable(true)
	vmware_struct := []vmwarestructs.VmwareDynamicDetails{}
	errFind := db.Find(&vmware_struct).Error
	if errFind != nil {
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&vmware_struct))

	if err != nil {
		logger.Error("Error: ",err)
		println(err)
	}
	logger.Info("Successful in GetVcenterDetails.")
	tx.Commit()
}
