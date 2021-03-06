package azurecontroller

import(
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"gclassec/structs/azurestruct"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"os"
//	"log"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"gclassec/goclientazure"
	"gclassec/readcredentials"
	"gclassec/loggers"
	"gclassec/structs/tagstruct"
	"gclassec/errorcodes/errcode"
	"regexp"
	"gclassec/dbmanagement"
	"gclassec/confmanagement/readstructconf"
)
type (

    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}
var logger = Loggers.New()
//var counter = 0
var dbtype string = dbmanagement.ENVdbtype
var dbname  string = dbmanagement.ENVdbnamegodb
var dbusername string = dbmanagement.ENVdbusername
var dbpassword string = dbmanagement.ENVdbpassword
var dbhostname string = dbmanagement.ENVdbhostname
var dbport string = dbmanagement.ENVdbport
var b []string = []string{dbusername,":",dbpassword,"@tcp","(",dbhostname,":",dbport,")","/",dbname}

var c string = (strings.Join(b,""))

var db,err  = gorm.Open(dbtype, c)

var azurecreds = readazurecreds.Configurtion()

func (uc UserController) GetAzureStaticDynamic(w http.ResponseWriter, r *http.Request)(){
	var azureCreds = readazurecreds.Configurtion()
	os.Setenv("AZURE_CLIENT_ID", azureCreds.ClientId)
	os.Setenv("AZURE_CLIENT_SECRET", azureCreds.ClientSecret)
	os.Setenv("AZURE_SUBSCRIPTION_ID", azureCreds.SubscriptionId)
	os.Setenv("AZURE_TENANT_ID", azureCreds.TenantId)
	println("------------AZURE CLIENT ID--------------")
	println(azureCreds.ClientId)
	logger.Debug("AZURE_CLIENT_ID", azureCreds.ClientId)
	logger.Debug("AZURE_CLIENT_SECRET", azureCreds.ClientSecret)
	logger.Debug("AZURE_SUBSCRIPTION_ID", azureCreds.SubscriptionId)
	logger.Debug("AZURE_TENANT_ID", azureCreds.TenantId)
	logger.Info("------------AZURE CLIENT ID--------------")
	logger.Info(azureCreds.ClientId)
	c := map[string]string{
		"AZURE_CLIENT_ID":       os.Getenv("AZURE_CLIENT_ID"),
		"AZURE_CLIENT_SECRET":   os.Getenv("AZURE_CLIENT_SECRET"),
		"AZURE_SUBSCRIPTION_ID": os.Getenv("AZURE_SUBSCRIPTION_ID"),
		"AZURE_TENANT_ID":       os.Getenv("AZURE_TENANT_ID")}
	if err := checkEnvVar(&c); err != nil {
		//log.Fatalf("Error: %v", err)
		fmt.Println("Error: ", err)
		logger.Error("Error: %v", err)
		return
	}
	spt, err := helpers.NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		//log.Fatalf("Error: %v", errcode.ErrAuth)
		fmt.Println("Azure : ",errcode.ErrAuth)
		logger.Error("Azure : ", errcode.ErrAuth)
		return
	}
	ac := goclientazure.NewVirtualMachinesClient(c["AZURE_SUBSCRIPTION_ID"])
	ac.Authorizer = spt

	ls, err := ac.ListAll()

		if err != nil{
		fmt.Println("Azure : ",errcode.ErrAuth)
		logger.Error("Azure :", errcode.ErrAuth)
//		return
	}

	tx := db.Begin()
	db.SingularTable(true)

	//obj := &azurestruct.VirtualMachineStaticDynamic{}
	obj := []azurestruct.VirtualMachineStaticDynamic{}
	tag := []tagstruct.Tags{}

	//create a regex `(?i)azure` will match string contains "azure" case insensitive
	reg := regexp.MustCompile("(?i)azure")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("Azure")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("Azure")).Find(&tag)

	fmt.Println(obj)
	logger.Info(obj)

	for _, element := range *ls.Value {
		rgroup := *(element.AvailabilitySet.ID)
		resourcegroupname := strings.Split(rgroup, "/")
		rsgroup := resourcegroupname[4]
		vmName := *(element.Name)
		vmId := *(element.VMID)

		dc := goclientazure.NewDynamicUsageOperationsClient(c["AZURE_SUBSCRIPTION_ID"])
		dc.Authorizer = spt

		dlist, _ := dc.ListDynamic(vmName, rsgroup)
		for _, element1 := range *dlist.Value {
			res := azurestruct.VirtualMachineStaticDynamic{}
			if len(tag) == 0 {
				res.Tagname = "Nil"
			}else {
				for _, el := range tag {
					if vmId != el.InstanceId {
						res.Tagname = "Nil"
					}else {
						res.Tagname = el.Tagname
					}
				}
			}
			res.VmName = *element.Name
			res.VmId = *element.VMID
			res.Location = *element.Location
			res.Offer = *element.StorageProfile.ImageReference.Offer
			res.Publisher = *element.StorageProfile.ImageReference.Publisher
			res.SKU = *element.StorageProfile.ImageReference.Sku
			res.AvailabilitySetName = *element.AvailabilitySet.ID
			res.Provisioningstate = *element.ProvisioningState
			res.ResourcegroupName = rsgroup
			res.Type = *element.Type
			res.VmSize = element.VirtualMachineProperties.HardwareProfile.VMSize
			res.TimeStamp = *(element1.Data[len(element1.Data) - 2].TimeStamp)
			res.Unit = *element1.Unit
			if element1.Data[len(element1.Data) - 1].Average == nil{
				res.Average = 0.0
			}else {
				res.Average = *(element1.Data[len(element1.Data) - 2].Average)
			}
			obj = append(obj, res)
		}
	}
	_ = json.NewEncoder(w).Encode(&obj)
	tx.Commit()
}

func   (uc UserController) GetAzureDetails(w http.ResponseWriter, r *http.Request)(){

	tx := db.Begin()
	db.SingularTable(true)

	azure_struct := []azurestruct.AzureInstances{}
	response_struct := []azurestruct.AzureResponse{}
	tag := []tagstruct.Tags{}

	//create a regex `(?i)azure` will match string contains "azure" case insensitive
	reg := regexp.MustCompile("(?i)azure")

	//Do the match operation using FindString() function
	er1 := db.Where("Cloud = ?", reg.FindString("Azure")).Find(&tag).Error
	if er1 != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}
	db.Where("Cloud = ?", reg.FindString("Azure")).Find(&tag)

	errFind := db.Find(&azure_struct).Error

	if errFind != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}

	db.Where("subscriptionid =?",azurecreds.SubscriptionId).Find(&azure_struct)

	if readstructconf.ReadStructConfigFile()!=0{
		standardresponse := []azurestruct.StandardizedAzure{}
		for i:=0; i<len(azure_struct);i++{
			response := azurestruct.StandardizedAzure{}
			response.VmName = azure_struct[i].VmName
			response.VmId = azure_struct[i].VmId
			response.Status = azure_struct[i].Status
			response.RAM = azure_struct[i].RAM
			response.NumCPU = azure_struct[i].NumCPU
			response.Storage = azure_struct[i].Storage
			//response.Tagname = azure_struct[i].Tagname
			response.VmSize = azure_struct[i].VmSize
			if len(tag) == 0 {
				response.Tagname = "Nil"
			}else {
				for _, el := range tag {
					if azure_struct[i].VmId != el.InstanceId {
						response.Tagname = "Nil"
					}else {
						response.Tagname = el.Tagname
						break
					}
				}
			}
			standardresponse = append(standardresponse, response)
		}
		_ = json.NewEncoder(w).Encode(&standardresponse)
	}else {
		for i:=0; i<len(azure_struct);i++ {
			response := azurestruct.AzureResponse{}
			response.VmName = azure_struct[i].VmName
			response.VmId = azure_struct[i].VmId
			response.Status = azure_struct[i].Status
			response.RAM = azure_struct[i].RAM
			response.NumCPU = azure_struct[i].NumCPU
			response.Storage = azure_struct[i].Storage
			response.VmSize = azure_struct[i].VmSize
			response.Deleted = azure_struct[i].Deleted
			response.AvailabilitySetName = azure_struct[i].AvailabilitySetName
			response.Location = azure_struct[i].Location
			response.Offer = azure_struct[i].Offer
			response.Provisioningstate = azure_struct[i].Provisioningstate
			response.Publisher = azure_struct[i].Publisher
			response.ResourcegroupName = azure_struct[i].ResourcegroupName
			response.SKU = azure_struct[i].SKU
			response.SubscriptionId = azure_struct[i].SubscriptionId
			response.Type = azure_struct[i].Type

			if len(tag) == 0 {
				response.Tagname = "Nil"
			}else {
				for _, el := range tag {
					if azure_struct[i].VmId != el.InstanceId {
						response.Tagname = "Nil"
					}else {
						response.Tagname = el.Tagname
						break
					}
				}
			}
			response_struct = append(response_struct, response)
		}
		_ = json.NewEncoder(w).Encode(&response_struct)
	}

	//_ = json.NewEncoder(w).Encode(db.Where("subscriptionid =?",azurecreds.SubscriptionId).Find(&azure_struct))

		if err != nil {
			logger.Error("Error: ",err)
			println(err)
		}
	logger.Info("Successful in GetAzureDetails.")
	tx.Commit()
}

func   (uc UserController) GetDynamicAzureDetails(w http.ResponseWriter, r *http.Request)(){
	vars := mux.Vars(r)
	name := vars["name"]
	resourceGrp := vars["resourceGroup"]

	var azureCreds = readazurecreds.Configurtion()
	os.Setenv("AZURE_CLIENT_ID", azureCreds.ClientId)
	os.Setenv("AZURE_CLIENT_SECRET", azureCreds.ClientSecret)
	os.Setenv("AZURE_SUBSCRIPTION_ID", azureCreds.SubscriptionId)
	os.Setenv("AZURE_TENANT_ID", azureCreds.TenantId)
	println("------------AZURE CLIENT ID--------------")
	println(azureCreds.ClientId)
	logger.Debug("AZURE_CLIENT_ID", azureCreds.ClientId)
	logger.Debug("AZURE_CLIENT_SECRET", azureCreds.ClientSecret)
	logger.Debug("AZURE_SUBSCRIPTION_ID", azureCreds.SubscriptionId)
	logger.Debug("AZURE_TENANT_ID", azureCreds.TenantId)
	logger.Info("------------AZURE CLIENT ID--------------")
	logger.Info(azureCreds.ClientId)
	//var drggroup string

	c := map[string]string{
		"AZURE_CLIENT_ID":       os.Getenv("AZURE_CLIENT_ID"),
		"AZURE_CLIENT_SECRET":   os.Getenv("AZURE_CLIENT_SECRET"),
		"AZURE_SUBSCRIPTION_ID": os.Getenv("AZURE_SUBSCRIPTION_ID"),
		"AZURE_TENANT_ID":       os.Getenv("AZURE_TENANT_ID")}
	if err := checkEnvVar(&c); err != nil {
		logger.Error("Error: %v", err)
		fmt.Println("Error: ", err)
		//log.Fatalf("Error: %v", err)
		return
	}
	spt, err := helpers.NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		logger.Error("Azure : ", errcode.ErrAuth)
		fmt.Println("Azure : ", errcode.ErrAuth)
		return
	}

	dc := goclientazure.NewDynamicUsageOperationsClient(c["AZURE_SUBSCRIPTION_ID"])
	dc.Authorizer = spt

	dlist, _ := dc.ListDynamic(name,resourceGrp)
	logger.Info(dlist)

	_ = json.NewEncoder(w).Encode(&dlist)

}


func checkEnvVar(envVars *map[string]string) error {
	var missingVars []string
	for varName, value := range *envVars {
		if value == "" {
			missingVars = append(missingVars, varName)
		}
	}
	if len(missingVars) > 0 {
		logger.Error("Missing environment variables %v", missingVars)
		return fmt.Errorf("Missing environment variables %v", missingVars)
	}
	return nil
}

func   (uc UserController) GetAzureDynamic(w http.ResponseWriter, r *http.Request)(){

	tx := db.Begin()
	db.SingularTable(true)

	azure_struct := []azurestruct.AzureDynamic{}

	errFind := db.Find(&azure_struct).Error

	if errFind != nil{
		logger.Error("Error: ",errcode.ErrFindDB)
		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&azure_struct))

		if err != nil {
			logger.Error("Error: ",err)
			println(err)
		}
	logger.Info("Successful in GetAzureDynamic.")
	tx.Commit()
}