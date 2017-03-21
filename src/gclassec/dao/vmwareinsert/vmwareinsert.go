package vmwareinsert

import (
       //"gclassec/confmanagement/readazureconf"
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

       "gclassec/errorcodes/errcode"
       "gclassec/structs/tagstruct"
       "regexp"
       "gclassec/dbmanagement"
       //"github.com/vmware/govmomi/govc/vm"
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

//var dbcredentials = dbman.Configurtion()
var dbtype string = dbmanagement.ENVdbtype
var dbname  string = dbmanagement.ENVdbnamegodb
var dbusername string = dbmanagement.ENVdbusername
var dbpassword string = dbmanagement.ENVdbpassword
var dbhostname string = dbmanagement.ENVdbhostname
var dbport string = dbmanagement.ENVdbport

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




func VmwareInsert() (error,int,int){
       ctx, cancel := context.WithCancel(context.Background())
       defer cancel()


       /*fmt.Println("dbtype string =", dbcredentials.Dbtype)
       fmt.Println(" dbname  string =", dbcredentials.Dbname)
       fmt.Println(" dbusername string =", dbcredentials.Dbusername)
       fmt.Println(" dbpassword string =", dbcredentials.Dbpassword)
       fmt.Println(" dbhostname string =", dbcredentials.Dbhostname)
       fmt.Println("dbport string = ",dbcredentials.Dbport)
       fmt.Println(" EnvURL string = ",vmwarecreds.EnvURL)
       fmt.Println(" EnvUserName  string =", vmwarecreds.EnvUserName)
       fmt.Println(" EnvPassword string =", vmwarecreds.EnvPassword)
       fmt.Println(" EnvInsecure string =", vmwarecreds.EnvInsecure)
*/




       var insecureFlag =  vmwarecontroller.ENVinsecureFlag/*flag.Bool("insecure", true, insecureDescription)*/
       logger.Info(*insecureFlag)

       flag.Parse()
       var urlFlag =vmwarecontroller.ENVurlFlag
       // Parse URL from string
       u, err := url.Parse(*urlFlag)
       if err != nil {
              logger.Error("Error: ",err)
              fmt.Println("Error :", err)
              //exit(err)
              return err,0,0
       }

       // Override username and/or password as required
       ProcessOverride(u)

       // Connect and log in to ESX or vCenter
       c, err := govmomi.NewClient(ctx, u, *insecureFlag)
       if err != nil {
              logger.Error("VMWare : ", errcode.ErrAuth)
              fmt.Println("VMWare : ", errcode.ErrAuth)
              return err,0,0
       }

       f := find.NewFinder(c.Client, true)

       // Find one and only datacenter
       dc, err := f.DefaultDatacenter(ctx)
       if err != nil {
              logger.Error("Error: ",err)
              fmt.Println("Error : ", err)
              //exit(err)
              return err,0,0
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
              return err,0,0
       }

       tx := db.Begin()
       db.SingularTable(true)

       tag := []tagstruct.Tags{}

       //create a regex `(?i)vmware` will match string contains "vmware" case insensitive
       reg := regexp.MustCompile("(?i)vmware")

       //Do the match operation using FindString() function
       er1 := db.Where("Cloud = ?", reg.FindString("VMWARE")).Find(&tag).Error
       if er1 != nil{
              logger.Error("Error: ",errcode.ErrFindDB)
              //tx.Rollback()
              return er1,0,0
       }
       db.Where("Cloud = ?", reg.FindString("VMWARE")).Find(&tag)

       fmt.Println("Tag : ", tag)
	count:=0
	count2:=0
	for _,element1:=range vmt{
		if element1.Summary.Runtime.PowerState == "poweredOn"{
			count++
		}else{
			count2++
		}
	}

       vmware_struct := []vmwarestructs.VmwareInstances{}
       er := db.Find(&vmware_struct).Error
       if er != nil {
              logger.Error("Error: ",errcode.ErrFindDB)
              //tx.Rollback()
              return err,0,0
       }
       /*for _, element := range vmware_struct {
                      db.Table("vmware_instances").Where("Name = ?",element.Name).Update("deleted", true)
       }*/
       db.Find(&vmware_struct)
       if (len(vmware_struct) == 0) {
              for _, vm := range vmt {
                     user := vmwarestructs.VmwareInstances{Name:vm.Summary.Config.Name, Uuid:vm.Summary.Config.Uuid, MemorySizeMB:vm.Summary.Config.MemorySizeMB, PowerState:string(vm.Summary.Runtime.PowerState), NumofCPU:vm.Summary.Config.NumCpu, GuestFullName:vm.Summary.Guest.GuestFullName, IPaddress:vm.Summary.Guest.IpAddress,StorageCommitted:float32(vm.Summary.Storage.Committed)/float32(1024*1024*1024), Tagname:"Nil", Deleted:false, Classifier:vmwarecreds.EnvUserName}
                     db.Create(&user)
              }
       }else{
              for _, vm := range vmt {
              db.Where("Name = ?",vm.Summary.Config.Name).Find(&vmware_struct)
              if (len(vmware_struct)==0) {
                     user := vmwarestructs.VmwareInstances{Name:vm.Summary.Config.Name, Uuid:vm.Summary.Config.Uuid, MemorySizeMB:vm.Summary.Config.MemorySizeMB, PowerState:string(vm.Summary.Runtime.PowerState), NumofCPU:vm.Summary.Config.NumCpu, GuestFullName:vm.Summary.Guest.GuestFullName, IPaddress:vm.Summary.Guest.IpAddress,StorageCommitted:float32(vm.Summary.Storage.Committed)/float32(1024*1024*1024), Tagname:"Nil", Deleted:false, Classifier:vmwarecreds.EnvUserName}
                     db.Create(&user)
              }else{
                     user := vmwarestructs.VmwareInstances{Name:vm.Summary.Config.Name, Uuid:vm.Summary.Config.Uuid, MemorySizeMB:vm.Summary.Config.MemorySizeMB, PowerState:string(vm.Summary.Runtime.PowerState), NumofCPU:vm.Summary.Config.NumCpu, GuestFullName:vm.Summary.Guest.GuestFullName, IPaddress:vm.Summary.Guest.IpAddress,StorageCommitted:float32(vm.Summary.Storage.Committed)/float32(1024*1024*1024), Tagname:"Nil", Deleted:false, Classifier:vmwarecreds.EnvUserName}
                     db.Model(&user).Where("Name =?", vm.Summary.Config.Name).Updates(user)
              }
       }
       }
       // Print summary
       tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

       logger.Info("Virtual machines found:", len(vmt))
       db.Find(&vmware_struct)
       for _, i := range vmware_struct {
              if len(tag) != 0 {
                     for _, el := range tag {
                            if i.Uuid == el.InstanceId{
                                   fmt.Println("----Update Tag for this instance----")
                                   fmt.Println("el.InstanceId: ", el.InstanceId)
                                   db.Model(vmwarestructs.VmwareInstances{}).Where("Name = ?", i.Name).Update("tagname",el.Tagname)
                                   //db.Table("vmware_instances").Where("Name = ?", i.Name).Update("tagname",el.Tagname)
                            }
                     }

              }
       }
       for _, element := range vmware_struct {
              fmt.Println("inside delete")
              for _, ele := range vmt{
                     if element.Name != ele.Summary.Config.Name {
                            fmt.Println("insdie  continue")
                          continue

                     }else{
                            db.Table("vmware_instances").Where("Name = ?",element.Name ).Update("deleted", false)
              }
             }
       }
       logger.Info("Successful in VmWareInsert.")
       tw.Flush()
       tx.Commit()
       return nil,count,count2
}

