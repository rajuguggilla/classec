package vmwareinsert

import (
	"fmt"
	"flag"
	"net/url"
	"github.com/vmware/govmomi"
	"gclassec/errorcodes/errcode"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"text/tabwriter"
	"os"
	"gclassec/structs/vmwarestructs"
	"context"
	"github.com/vmware/govmomi/vim25/types"
	"gclassec/structs/tagstruct"
	"regexp"
	"time"
	"gclassec/controllers/vmwarecontroller"
)

func VmwareDynamicInsert() error{
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	CurrentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("--------",CurrentTime)



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
		return err
	}

	// Override username and/or password as required
	ProcessOverride(u)

	// Connect and log in to ESX or vCenter
	c, err := govmomi.NewClient(ctx, u, *insecureFlag)
	if err != nil {
		logger.Error("VMWare : ", errcode.ErrAuth)
		fmt.Println("VMWare : ", errcode.ErrAuth)
		return err
	}

	f := find.NewFinder(c.Client, true)

	// Find one and only datacenter
	dc, err := f.DefaultDatacenter(ctx)
	if err != nil {
		logger.Error("Error: ",err)
		fmt.Println("Error : ", err)
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

	tx := db.Begin()
	db.SingularTable(true)

	tag := []tagstruct.Providers{}

	//create a regex `(?i)vmware` will match string contains "vmware" case insensitive
	reg := regexp.MustCompile("(?i)vmware")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("VMWARE")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("VMWARE")).Find(&tag)

	fmt.Println("Tag : ", tag)

	vmware_struct := []vmwarestructs.VmwareInstances{}
	er := db.Find(&vmware_struct).Error
	if er != nil {
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Find(&vmware_struct)

	// Print summary
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	logger.Info("Virtual machines found:", len(vmt))

	for _, element := range vmware_struct {
       		db.Table("vmware_instances").Where("Name = ?",element.Name).Update("deleted", true)
	}

	for _, vm := range vmt {
       		for _, ele := range vmware_struct {
              if vm.Summary.Config.Name != ele.Name {
                     continue
              }else {

		      fmt.Println("Vm naame :",vm.Summary.Config.Name)
		      fmt.Println("Uuid  :",vm.Summary.Config.Uuid)
		      fmt.Println("timestamp :",vm.Summary.Storage.Timestamp)
		      fmt.Println("Max CPU  :",vm.Summary.Runtime.MaxCpuUsage)
		      //fmt.Println("Min CPU :",vm.Summary.Runtime.MinCpuUsage)
		      //fmt.Println("Avg CPU  :",vm.Summary.Runtime.AvgCpuUsage)
		      user := vmwarestructs.VmwareDynamicDetails{Name:vm.Summary.Config.Name, Uuid:vm.Summary.Config.Uuid,Timestamp:CurrentTime, MaxCpuUsage:vm.Summary.Runtime.MaxCpuUsage}//, MinCpuUsage:vm.Summary.Runtime.MinCpuUsage, AvgCpuUsage:vm.Summary.Runtime.AvgCpuUsage}//,Tagname:"Nil", Deleted:true}
                   //  db.Model(&user).Where("Name =?",vm.Summary.Config.Name).Updates(user)
		      db.Create(&user)
			//db.Model(&user).Updates(&user)
              }
       }
}


	tw.Flush()
//	tx.Commit()
	return nil
}

