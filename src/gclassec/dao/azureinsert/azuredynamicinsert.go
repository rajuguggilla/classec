package azureinsert

import (
	"os"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
	"gclassec/goclientazure"
	"gclassec/readcredentials"
	"github.com/Azure/go-autorest/autorest/azure"
	"gclassec/errorcodes/errcode"
	"strings"
	"gclassec/structs/azurestruct"
)

func AzureDynamicInsert() error{
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

	ls, err := ac.ListAll()

	if err != nil{
		fmt.Println("Azure : ",errcode.ErrAuth)
		logger.Error("Azure :", errcode.ErrAuth)
		return err
	}

	for _, element := range *ls.Value {
		name := *(element.Name)
		rgroup := *(element.AvailabilitySet.ID)
		resourcegroupname := strings.Split(rgroup, "/")
		rsgroup := resourcegroupname[4]
		vmName := *(element.Name)
		fmt.Println(*(element.VMID))

		dc := goclientazure.NewDynamicUsageOperationsClient(c["AZURE_SUBSCRIPTION_ID"])
		dc.Authorizer = spt

		//Get current Status of instance
		instanceView, _ := ac.GetInstanceView(name, resourcegroupname[4])
		if *instanceView.Statuses[len(instanceView.Statuses) - 1].DisplayStatus != "VM running" {
			fmt.Println("Azure : ", errcode.ErrAzureDynamic)
			logger.Error("Azure :", errcode.ErrAzureDynamic)
			//return
		}else {
		dlist, er := dc.ListDynamic(vmName, rsgroup)
		if er != nil {
			fmt.Println("Azure : ", errcode.ErrAuth)
			logger.Error("Azure :", errcode.ErrAuth)
			return er
		}
		fmt.Println(dlist)
		logger.Info(dlist)
		for _, element1 := range *dlist.Value {
			if element1.Data[len(element1.Data) - 1].Average == nil{
				azure_dynamic := azurestruct.AzureDynamic{Name:*element.Name, Timestamp:*(element1.Data[len(element1.Data) - 1].TimeStamp), Minimum:0.0, Maximum:0.0, Average:0.0, Unit:*element1.Unit}
				db.Create(&azure_dynamic)
			}else{
				azure_dynamic := azurestruct.AzureDynamic{Name:*element.Name, Timestamp:*(element1.Data[len(element1.Data) - 1].TimeStamp), Minimum:*(element1.Data[len(element1.Data) - 1].Minimum), Maximum:*(element1.Data[len(element1.Data) - 1].Maximum), Average:*(element1.Data[len(element1.Data) - 1].Average), Unit:*element1.Unit}
				db.Create(&azure_dynamic)
			}
		}
		}

	}
	tx.Commit()
	return nil
}
