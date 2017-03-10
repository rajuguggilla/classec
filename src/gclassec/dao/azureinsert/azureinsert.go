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
	"gclassec/structs/tagstruct"
	"regexp"
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

func AzureInsert() error{
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
		return err
	}
	spt, err := helpers.NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		logger.Error("Error: %v", err)
		fmt.Println("Error : ", err)
		return err
	}
	ac := goclientazure.NewVirtualMachinesClient(c["AZURE_SUBSCRIPTION_ID"])
	ac.Authorizer = spt

	tx := db.Begin()
	db.SingularTable(true)

	tag := []tagstruct.Providers{}
	azure_struct := []azurestruct.AzureInstances{}

	//create a regex `(?i)azure` will match string contains "azure" case insensitive
	reg := regexp.MustCompile("(?i)azure")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("Azure")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("Azure")).Find(&tag)

	er := db.Find(&azure_struct).Error

	if er != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Find(&azure_struct)

	ls, err := ac.ListAll()

	if err != nil{
		fmt.Println("Azure :", errcode.ErrAuth)
		logger.Error("Azure :", errcode.ErrAuth)
		return err
	}
	if (len(azure_struct)==0){
		for _, element := range *ls.Value {
		 rgroup := *(element.AvailabilitySet.ID)
                            resourcegroupname := strings.Split(rgroup, "/")
                            user := azurestruct.AzureInstances{SubscriptionId:subscriptionid,VmName:*element.Name, Type:*element.Type, Location:*element.Location, VmSize:element.VirtualMachineProperties.HardwareProfile.VMSize, VmId:*element.VMID, Publisher:*(element.StorageProfile.ImageReference.Publisher), Offer:*(element.StorageProfile.ImageReference.Offer), SKU:*(element.StorageProfile.ImageReference.Sku), AvailabilitySetName:*(element.AvailabilitySet.ID), Provisioningstate:*element.ProvisioningState, ResourcegroupName:resourcegroupname[4],Tagname:"Nil",Deleted:false}
                           db.Create(&user)
	}
	}else{
		for _, element := range *ls.Value {
		  	db.Where("name = ?",element.Name).Find(&azure_struct)
			     	if (len(azure_struct)==0){
				      rgroup := *(element.AvailabilitySet.ID)
					resourcegroupname := strings.Split(rgroup, "/")
                            user := azurestruct.AzureInstances{SubscriptionId:subscriptionid,VmName:*element.Name, Type:*element.Type, Location:*element.Location, VmSize:element.VirtualMachineProperties.HardwareProfile.VMSize, VmId:*element.VMID, Publisher:*(element.StorageProfile.ImageReference.Publisher), Offer:*(element.StorageProfile.ImageReference.Offer), SKU:*(element.StorageProfile.ImageReference.Sku), AvailabilitySetName:*(element.AvailabilitySet.ID), Provisioningstate:*element.ProvisioningState, ResourcegroupName:resourcegroupname[4],Tagname:"Nil",Deleted:false}
					db.Create(&user)
			     }else {
				      rgroup := *(element.AvailabilitySet.ID)
                            		resourcegroupname := strings.Split(rgroup, "/")
                            user := azurestruct.AzureInstances{SubscriptionId:subscriptionid,VmName:*element.Name, Type:*element.Type, Location:*element.Location, VmSize:element.VirtualMachineProperties.HardwareProfile.VMSize, VmId:*element.VMID, Publisher:*(element.StorageProfile.ImageReference.Publisher), Offer:*(element.StorageProfile.ImageReference.Offer), SKU:*(element.StorageProfile.ImageReference.Sku), AvailabilitySetName:*(element.AvailabilitySet.ID), Provisioningstate:*element.ProvisioningState, ResourcegroupName:resourcegroupname[4],Tagname:"Nil",Deleted:true}
                            db.Model(&user).Where("name =?",element.Name).Updates(user)
			     }
		}
	}
	_ = json.NewEncoder(os.Stdout).Encode(&ls)

	/*for _, element := range azure_struct {
              db.Table("azure_instances").Where("Name = ?",element.VmName).Update("deleted", true)
       }*/

	for _, i := range azure_struct {
		if len(tag) == 0 {
			fmt.Println("----Nothing in Tag----")
			db.Table("azure_instances").Where("vmid = ?",i.VmId).Update("tagname","Nil")
		} else {
			for _, el := range tag {
				if i.VmId != el.InstanceId {
					fmt.Println("----No Tag for this instance----")
					db.Table("azure_instances").Where("vmid = ?",i.VmId).Update("tagname","Nil")
				} else {
					fmt.Println("----Update Tag for this instance----")
					fmt.Println("Tagname : ", el.Tagname)
					db.Table("azure_instances").Where("vmid = ?",i.VmId).Update("tagname",el.Tagname)
				}
			}
		}
	}

	/*for _, element := range azure_struct {
                     for _, ele := range *ls.Value {
                            if element.VmName != *ele.Name {
                                   continue
                            }else{
                                   db.Table("azure_instances").Where("name = ?",element.VmName).Update("deleted", false)
                            }
                     }
                     }*/



	tx.Commit()
	return nil
}