package azurecontroller

import(
	"strings"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"gclassec/structs/azurestruct"
	"gclassec/confmanagement/readazureconf"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"os"
	"log"
	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
	"github.com/Azure/go-autorest/autorest/azure"
	"gclassec/goclientazure"
	"gclassec/readcredentials"
)
type (

    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
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

func (uc UserController) GetAzureStaticDynamic(w http.ResponseWriter, r *http.Request)(){
	var azureCreds = readazurecreds.Configurtion()
	os.Setenv("AZURE_CLIENT_ID", azureCreds.ClientId)
	os.Setenv("AZURE_CLIENT_SECRET", azureCreds.ClientSecret)
	os.Setenv("AZURE_SUBSCRIPTION_ID", azureCreds.SubscriptionId)
	os.Setenv("AZURE_TENANT_ID", azureCreds.TenantId)
	println("------------AZURE CLIENT ID--------------")
	println(azureCreds.ClientId)
	c := map[string]string{
		"AZURE_CLIENT_ID":       os.Getenv("AZURE_CLIENT_ID"),
		"AZURE_CLIENT_SECRET":   os.Getenv("AZURE_CLIENT_SECRET"),
		"AZURE_SUBSCRIPTION_ID": os.Getenv("AZURE_SUBSCRIPTION_ID"),
		"AZURE_TENANT_ID":       os.Getenv("AZURE_TENANT_ID")}
	if err := checkEnvVar(&c); err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	spt, err := helpers.NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	ac := goclientazure.NewVirtualMachinesClient(c["AZURE_SUBSCRIPTION_ID"])
	ac.Authorizer = spt

	ls, _ := ac.ListAll()

	obj := &azurestruct.VirtualMachineStaticDynamic{}
	for _, element := range *ls.Value {
		rgroup := *(element.AvailabilitySet.ID)
		resourcegroupname := strings.Split(rgroup, "/")
		rsgroup := resourcegroupname[4]
		vmName := *(element.Name)

		dc := goclientazure.NewDynamicUsageOperationsClient(c["AZURE_SUBSCRIPTION_ID"])
		dc.Authorizer = spt

		dlist, _ := dc.ListDynamic(vmName, rsgroup)
		fmt.Println(dlist)
		for _, element1 := range *dlist.Value {
			obj = &azurestruct.VirtualMachineStaticDynamic{VmName:element.Name, Type:element.Type, Location:element.Location, VmSize:element.VirtualMachineProperties.HardwareProfile.VMSize, VmId:element.VMID, Publisher:element.StorageProfile.ImageReference.Publisher, Offer:element.StorageProfile.ImageReference.Offer, SKU:element.StorageProfile.ImageReference.Sku, AvailabilitySetName:element.AvailabilitySet.ID, Provisioningstate:element.ProvisioningState, ResourcegroupName:rsgroup,TimeStamp:element1.Data[len(element1.Data)-2].TimeStamp,Average:element1.Data[len(element1.Data)-2].Average}
		}
	}
	fmt.Println(obj)
	_ = json.NewEncoder(w).Encode(&obj)
}

func   (uc UserController) GetAzureDetails(w http.ResponseWriter, r *http.Request)(){

	tx := db.Begin()
	db.SingularTable(true)

	azure_struct := []azurestruct.AzureInstances{}

	err := db.Find(&azure_struct).Error

	if err != nil{

		tx.Rollback()
	}

	_ = json.NewEncoder(w).Encode(db.Find(&azure_struct))

		if err != nil {
			println(err)
		}

	tx.Commit()
}

func   (uc UserController) GetDynamicAzureDetails(w http.ResponseWriter, r *http.Request)(){
	vars := mux.Vars(r)
	name := vars["name"]
	resourceGrp := vars["resourceGroup"]


	//var drggroup string

	c := map[string]string{
		"AZURE_CLIENT_ID":       os.Getenv("AZURE_CLIENT_ID"),
		"AZURE_CLIENT_SECRET":   os.Getenv("AZURE_CLIENT_SECRET"),
		"AZURE_SUBSCRIPTION_ID": os.Getenv("AZURE_SUBSCRIPTION_ID"),
		"AZURE_TENANT_ID":       os.Getenv("AZURE_TENANT_ID")}
	if err := checkEnvVar(&c); err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	spt, err := helpers.NewServicePrincipalTokenFromCredentials(c, azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	dc := goclientazure.NewDynamicUsageOperationsClient(c["AZURE_SUBSCRIPTION_ID"])
	dc.Authorizer = spt

	dlist, _ := dc.ListDynamic(name,resourceGrp)
	fmt.Println(dlist)

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
		return fmt.Errorf("Missing environment variables %v", missingVars)
	}
	return nil
}