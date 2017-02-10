package vmwarecontroller

import (
	 _ "github.com/go-sql-driver/mysql"
	"gclassec/confmanagement/vmwareconf"
	"fmt"
	"net/url"
	"os"
	"gclassec/confmanagement/readazureconf"
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

	"github.com/vmware/govmomi/vim25/types"
)

//const (
//	EnvURL = "https://110.110.110.140:443/sdk"
//	EnvUserName = "administrator@vsphere.local"
//	EnvPassword = "Vcenter#1234"
//	EnvInsecure = "true"
//)

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


type FinalVmwareInstances struct {
	VMName              string
	Uuid                string
	MemorySizeMB        int32
	PowerState          string
	NumCpu              int32
	GuestFullName       string
	IpAddress           string

}
type DynamicValues struct{
	VMName			string
	OverallCpuUsage         int32
	GuestMemoryUsage	int32
	StorageCommitted	float32
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
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}



var dbcredentials1 = readazureconf.Configurtion()
var dbtype string = dbcredentials1.Dbtype
var dbname  string = dbcredentials1.Dbname
var dbusername string = dbcredentials1.Dbusername
var dbpassword string = dbcredentials1.Dbpassword
var dbhostname string = dbcredentials1.Dbhostname
var dbport string = dbcredentials1.Dbport

var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))

var db,err  = gorm.Open(dbtype, c)

func   (uc UserController) GetVcenterDetails(w http.ResponseWriter, r *http.Request)() {
	tx := db.Begin()
	db.SingularTable(true)
	vmware_struct := []vmwarestructs.VmwareInstances{}
	err := db.Find(&vmware_struct).Error
	if err != nil {

		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&vmware_struct))

	if err != nil {
		println(err)
	}

	tx.Commit()
}

func   (uc UserController) GetDynamicVcenterDetails(w http.ResponseWriter, r *http.Request)(){
       ctx, cancel := context.WithCancel(context.Background())
       defer cancel()
//       var insecureFlag = flag.Bool("insecure", true, insecureDescription)
       fmt.Println(*ENVinsecureFlag)
       //fmt.Println("Inside Vcenter get details !!!!!!!!! 1")

       flag.Parse()
//	var urlFlag = flag.String("url", EnvURL, urlDescription)
  //     // Parse URL from string
       u, err := url.Parse(*ENVurlFlag)
       if err != nil {
              exit(err)
       }

       // Override username and/or password as required
       ProcessOverride(u)

       // Connect and log in to ESX or vCenter
       c, err := govmomi.NewClient(ctx, u, *ENVinsecureFlag)
       if err != nil {
              exit(err)
       }

       f := find.NewFinder(c.Client, true)

       // Find one and only datacenter
       dc, err := f.DefaultDatacenter(ctx)
       if err != nil {
              exit(err)
       }

       // Make future calls local to this datacenterth
       f.SetDatacenter(dc)

       // Find virtual machines in datacenter
       vms, err := f.VirtualMachineList(ctx, "*")
       fmt.Println(vms)

       pc := property.DefaultCollector(c.Client)

       var refv []types.ManagedObjectReference
       for _, ds := range vms {
              refv = append(refv, ds.Reference())
       }

       // Retrieve name property for all vms
       var vmt []mo.VirtualMachine
       err = pc.Retrieve(ctx, refv, []string{"summary"}, &vmt)
       if err != nil {
              exit(err)
       }

       // Print summary
       tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

       fmt.Println("Virtual machines found:", len(vmt))
       fmt.Fprintf(w, "[")
       for _, vm := range vmt {
              //fmt.Fprintf(tw, "%s\n", vm.Name)
              fmt.Println("VM Name : ", vm.Summary.Config.Name)
              fmt.Println("Overall CPU : ", vm.Summary.QuickStats.OverallCpuUsage)
              fmt.Println("Guest memory : ", vm.Summary.QuickStats.GuestMemoryUsage)
              fmt.Println("Committed storage : ", units.ByteSize(vm.Summary.Storage.Committed))
              //_ = json.NewEncoder(os.Stdout).Encode(&vm)
              output := DynamicValues{VMName:vm.Summary.Config.Name,OverallCpuUsage:vm.Summary.QuickStats.OverallCpuUsage,GuestMemoryUsage:vm.Summary.QuickStats.GuestMemoryUsage,StorageCommitted:float32(vm.Summary.Storage.Committed)/float32(1024*1024*1024)}
              _ = json.NewEncoder(w).Encode(output)
             fmt.Fprintf(w, ",")
       }

      fmt.Fprintf(w, "{}]")

       tw.Flush()
}
//func VmwareInsert(){
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	fmt.Println(*ENVinsecureFlag)
//
//	flag.Parse()
//
//	// Parse URL from string
//	u, err := url.Parse(*ENVurlFlag)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// Override username and/or password as required
//	ProcessOverride(u)
//
//	// Connect and log in to ESX or vCenter
//	c, err := govmomi.NewClient(ctx, u, *ENVinsecureFlag)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	f := find.NewFinder(c.Client, true)
//
//	// Find one and only datacenter
//	dc, err := f.DefaultDatacenter(ctx)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// Make future calls local to this datacenter
//	f.SetDatacenter(dc)
//
//	// Find virtual machines in datacenter
//	vms, err := f.VirtualMachineList(ctx, "*")
//	fmt.Println(vms)
//
//	pc := property.DefaultCollector(c.Client)
//
//	var refv []types.ManagedObjectReference
//	for _, ds := range vms {
//		refv = append(refv, ds.Reference())
//	}
//
//	// Retrieve name property for all vms
//	var vmt []mo.VirtualMachine
//	err = pc.Retrieve(ctx, refv, []string{"summary"}, &vmt)
//	if err != nil {
//  		fmt.Println(err)
//	}
//
//
//	// Print summary
//	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
//
//	fmt.Println("Virtual machines found:", len(vmt))
//	for _, vm := range vmt {
//
//		output := vmwarestructs.VmwareInstances{Name:vm.Summary.Config.Name,Uuid:vm.Summary.Config.Uuid,MemorySizeMB:vm.Summary.Config.MemorySizeMB,PowerState:vm.Summary.Runtime.PowerState,NumofCPU:vm.Summary.Config.NumCpu,GuestFullName:vm.Summary.Guest.GuestFullName,IPaddress:vm.Summary.Guest.IpAddress}
//		//_ = json.NewEncoder(w).Encode(output)
//		db.Create(&output)
//	}
//
//	tw.Flush()
//}
//
//
//
////vm.Summary.Runtime.PowerState