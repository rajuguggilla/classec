package vmwareinsert

import (
	"gclassec/confmanagement/readazureconf"
	"strings"
	"context"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/jinzhu/gorm"
	"fmt"
	"flag"
	"net/url"
	"text/tabwriter"
	"os"
	"gclassec/structs/vmwarestructs"
	_ "github.com/go-sql-driver/mysql"
	"gclassec/confmanagement/vmwareconf"
	"gclassec/controllers/vmwarecontroller"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"gclassec/loggers"
	
)
var vmwarecreds = vmwareconf.Configurtion()
var EnvURL string = vmwarecreds.EnvURL
var EnvUserName  string = vmwarecreds.EnvUserName
var EnvPassword string = vmwarecreds.EnvPassword
var EnvInsecure string = vmwarecreds.EnvInsecure
var logger = Loggers.New()
//var urlDescription = fmt.Sprintf("ESX or vCenter URL [%s]", EnvURL)
////var urlFlag = flag.String("url", EnvURL, urlDescription)
//
//var insecureDescription = fmt.Sprintf("Don't verify the server's certificate chain [%s]", EnvInsecure)
////var insecureFlag = flag.Bool("insecure", true, insecureDescription)

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
	//fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	//os.Exit(1)

}


/*func VmWareInsertDB(){
	err := VmwareInsert()
	if err != nil{

	}
}*/

func VmwareInsert() error{
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()


	fmt.Println("dbtype string =", dbcredentials.Dbtype)
	fmt.Println(" dbname  string =", dbcredentials.Dbname)
	fmt.Println(" dbusername string =", dbcredentials.Dbusername)
	fmt.Println(" dbpassword string =", dbcredentials.Dbpassword)
	fmt.Println(" dbhostname string =", dbcredentials.Dbhostname)
	fmt.Println("dbport string = ",dbcredentials.Dbport)
	fmt.Println(" EnvURL string = ",vmwarecreds.EnvURL)
	fmt.Println(" EnvUserName  string =", vmwarecreds.EnvUserName)
	fmt.Println(" EnvPassword string =", vmwarecreds.EnvPassword)
	fmt.Println(" EnvInsecure string =", vmwarecreds.EnvInsecure)





	var insecureFlag =  vmwarecontroller.ENVinsecureFlag/*flag.Bool("insecure", true, insecureDescription)*/
	logger.Info(*insecureFlag)

	flag.Parse()
	var urlFlag =vmwarecontroller.ENVurlFlag
	// Parse URL from string
	u, err := url.Parse(*urlFlag)
	if err != nil {
		logger.Error("Error: ",err)
		//fmt.Println(err)
		//exit(err)
		return err
	}

	// Override username and/or password as required
	ProcessOverride(u)

	// Connect and log in to ESX or vCenter
	c, err := govmomi.NewClient(ctx, u, *insecureFlag)
	if err != nil {
		logger.Error("\n\n Failed to connect VMWare  ")
		fmt.Println("\n\n Failed to connect VMWare ")
		logger.Error("Error: ",err)
		//fmt.Println(err)
		//exit(err)
		return err
	}

	f := find.NewFinder(c.Client, true)

	// Find one and only datacenter
	dc, err := f.DefaultDatacenter(ctx)
	if err != nil {
		logger.Error("Error: ",err)
		//fmt.Println(err)
		//exit(err)
		return err
	}

	// Make future calls local to this datacenter
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
  		fmt.Println(err)
		return err
	}


	// Print summary
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	logger.Info("Virtual machines found:", len(vmt))
	for _, vm := range vmt {

		output := vmwarestructs.VmwareInstances{Name:vm.Summary.Config.Name,Uuid:vm.Summary.Config.Uuid,MemorySizeMB:vm.Summary.Config.MemorySizeMB,PowerState:string(vm.Summary.Runtime.PowerState),NumofCPU:vm.Summary.Config.NumCpu,GuestFullName:vm.Summary.Guest.GuestFullName,IPaddress:vm.Summary.Guest.IpAddress}
		//_ = json.NewEncoder(w).Encode(output)
		db.Create(&output)
		db.Model(&output).Updates(&output)
	}
	logger.Info("Successful in VmWareInsert.")
	tw.Flush()

	return nil
}
