package azureinsert


import (
	"os"
	"github.com/Azure/go-autorest/autorest/azure"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"gclassec/structs/azurestruct"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"gclassec/readcredentials"
	"gclassec/goclientazure"
	"gclassec/loggers"
	"gclassec/errorcodes/errcode"
	"gclassec/dbmanagement"
)

type ls struct {

}

var logger = Loggers.New()
var dbtype string = dbmanagement.ENVdbtype
var dbname  string = dbmanagement.ENVdbnamegodb
var dbusername string = dbmanagement.ENVdbusername
var dbpassword string = dbmanagement.ENVdbpassword
var dbhostname string = dbmanagement.ENVdbhostname
var dbport string = dbmanagement.ENVdbport


var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))
var db,err  = gorm.Open(dbtype, c)

var azure_details = readazurecreds.Configurtion()
var subscriptionid = azure_details.SubscriptionId


func checkEnvVar(envVars *map[string]string) error {
	var missingVars []string
	for varName, value := range *envVars {
		if value == "" {
			missingVars = append(missingVars, varName)
		}
	}
	if len(missingVars) > 0 {
		return fmt.Errorf("Missing environment variables %v", missingVars)
	}
	return nil
}

func AzureInsert() (error,int,int){
	var storage int32
	var ram int32
	var numCPU int32
	var azureCreds = readazurecreds.Configurtion()
	os.Setenv("AZURE_CLIENT_ID", azureCreds.ClientId)
	os.Setenv("AZURE_CLIENT_SECRET", azureCreds.ClientSecret)
	os.Setenv("AZURE_SUBSCRIPTION_ID", azureCreds.SubscriptionId)
	os.Setenv("AZURE_TENANT_ID", azureCreds.TenantId)
	println("------------AZURE CLIENT ID--------------")
	println(azureCreds.ClientId)
	logger.Debug("------------AZURE CLIENT ID--------------")
	logger.Debug(azureCreds.ClientId)
	c := map[string]string{
		"AZURE_CLIENT_ID":       os.Getenv("AZURE_CLIENT_ID"),
		"AZURE_CLIENT_SECRET":   os.Getenv("AZURE_CLIENT_SECRET"),
		"AZURE_SUBSCRIPTION_ID": os.Getenv("AZURE_SUBSCRIPTION_ID"),
		"AZURE_TENANT_ID":       os.Getenv("AZURE_TENANT_ID")}
	if err := checkEnvVar(&c); err != nil {
		logger.Error("Error: %v", err)
		fmt.Println("Error : ", err)
		return err,0,0
	}
	spt, err := helpers.NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		logger.Error("Error: %v", err)
		fmt.Println("Error : ", err)
		return err,0,0
	}
	ac := goclientazure.NewVirtualMachinesClient(c["AZURE_SUBSCRIPTION_ID"])
	ac.Authorizer = spt

	tx := db.Begin()
	db.SingularTable(true)

	azure_struct := []azurestruct.AzureInstances{}

	er := db.Find(&azure_struct).Error

	if er != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		//tx.Rollback()
		return er,0,0
	}
	db.Find(&azure_struct)


	ls, err := ac.ListAll()

	poweredoncount := 0
	poweredoffcount := 0

	for _,element1:=range *ls.Value{
		name := *(element1.Name)
		rgroup := *(element1.AvailabilitySet.ID)
                        resourcegroupname := strings.Split(rgroup, "/")
			//Get current Status of instance
			instanceView, _ := ac.GetInstanceView(name,resourcegroupname[4])
		if *instanceView.Statuses[len(instanceView.Statuses) - 1].DisplayStatus == "VM running"{
			poweredoncount++
		}else{
			poweredoffcount++
		}
	}

	if err != nil{
		fmt.Println("Azure :", errcode.ErrAuth)
		logger.Error("Azure :", errcode.ErrAuth)
		return err,0,0
	}
	for _, element := range *ls.Value{
		name := *(element.Name)
		rgroup := *(element.AvailabilitySet.ID)
                resourcegroupname := strings.Split(rgroup, "/")
		//Get Storage, RAM and CPU count for instance
		size, errSize := ac.ListAvailableSizes(resourcegroupname[4], name)
		if errSize != nil{
			fmt.Println("Error : ", errSize)
			return errSize, 0, 0
		}
		fmt.Println("Size : ", size.Value)
		for _, ele := range *size.Value {
			if *ele.Name == element.VirtualMachineProperties.HardwareProfile.VMSize {
				storage = (*ele.ResourceDiskSizeInMB)/1024
				ram = *ele.MemoryInMB
				numCPU = *ele.NumberOfCores
			}
		}
	}
	fmt.Println("storage : ", storage)

	if (len(azure_struct)==0){
		for _, element := range *ls.Value {
			name := *(element.Name)
		 	rgroup := *(element.AvailabilitySet.ID)
                        resourcegroupname := strings.Split(rgroup, "/")
			//Get current Status of instance
			instanceView, _ := ac.GetInstanceView(name,resourcegroupname[4])
			user := azurestruct.AzureInstances{SubscriptionId:subscriptionid,VmName:*element.Name, Type:*element.Type, Location:*element.Location, VmSize:element.VirtualMachineProperties.HardwareProfile.VMSize, VmId:*element.VMID, Publisher:*(element.StorageProfile.ImageReference.Publisher), Offer:*(element.StorageProfile.ImageReference.Offer), SKU:*(element.StorageProfile.ImageReference.Sku), AvailabilitySetName:*(element.AvailabilitySet.ID), Provisioningstate:*element.ProvisioningState, ResourcegroupName:resourcegroupname[4], Status:*instanceView.Statuses[len(instanceView.Statuses) - 1].DisplayStatus, Storage:storage, RAM:ram, NumCPU:numCPU, Deleted:false}
                        db.Create(&user)
		}
	}else{
		for _, element := range *ls.Value {
			name := *(element.Name)
		 	rgroup := *(element.AvailabilitySet.ID)
                        resourcegroupname := strings.Split(rgroup, "/")
			//Get current Status of instance
			instanceView, _ := ac.GetInstanceView(name,resourcegroupname[4])
		  	db.Where("name = ?",element.Name).Find(&azure_struct)
			if (len(azure_struct)==0){
			        rgroup := *(element.AvailabilitySet.ID)
				resourcegroupname := strings.Split(rgroup, "/")
                        	user := azurestruct.AzureInstances{SubscriptionId:subscriptionid,VmName:*element.Name, Type:*element.Type, Location:*element.Location, VmSize:element.VirtualMachineProperties.HardwareProfile.VMSize, VmId:*element.VMID, Publisher:*(element.StorageProfile.ImageReference.Publisher), Offer:*(element.StorageProfile.ImageReference.Offer), SKU:*(element.StorageProfile.ImageReference.Sku), AvailabilitySetName:*(element.AvailabilitySet.ID), Provisioningstate:*element.ProvisioningState, ResourcegroupName:resourcegroupname[4],Status:*instanceView.Statuses[len(instanceView.Statuses) - 1].DisplayStatus, Storage:storage, RAM:ram, NumCPU:numCPU, Deleted:false}
				db.Create(&user)
			}else {
				rgroup := *(element.AvailabilitySet.ID)
                        	resourcegroupname := strings.Split(rgroup, "/")
                                user := azurestruct.AzureInstances{SubscriptionId:subscriptionid,VmName:*element.Name, Type:*element.Type, Location:*element.Location, VmSize:element.VirtualMachineProperties.HardwareProfile.VMSize, VmId:*element.VMID, Publisher:*(element.StorageProfile.ImageReference.Publisher), Offer:*(element.StorageProfile.ImageReference.Offer), SKU:*(element.StorageProfile.ImageReference.Sku), AvailabilitySetName:*(element.AvailabilitySet.ID), Provisioningstate:*element.ProvisioningState, ResourcegroupName:resourcegroupname[4],Status:*instanceView.Statuses[len(instanceView.Statuses) - 1].DisplayStatus, Storage:storage, RAM:ram, NumCPU:numCPU, Deleted:true}
				db.Model(&user).Where("name =?",element.Name).Updates(user)
			}
		}
	}

	_ = json.NewEncoder(os.Stdout).Encode(&ls)

	/*for _, element := range azure_struct {
              db.Table("azure_instances").Where("Name = ?",element.VmName).Update("deleted", true)
       }*/
	db.Find(&azure_struct)
	for _, element := range azure_struct {
                     for _, ele := range *ls.Value {
                            if element.VmName != *ele.Name {
                                   continue
                            }else{
                                   db.Table("azure_instances").Where("name = ?",element.VmName).Update("deleted", false)
                            }
                     }
                     }

	tx.Commit()
	return nil,poweredoncount,poweredoffcount
}