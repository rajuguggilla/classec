package azurestruct

//import "github.com/Azure/go-autorest/autorest/date"

//import "github.com/Azure/go-autorest/autorest/date"

type AzureInstances struct{
	SubscriptionId	string		`gorm:"column:subscriptionid"`
	VmName string 			`gorm:"column:name"`
	Type string 			`gorm:"column:type"`
	Location string 		`gorm:"column:location"`
	VmSize string           	`sql:"type:varchar" gorm:"column:vmsize"`
	Publisher string 		`gorm:"column:publisher"`
	Offer string 			`gorm:"column:offer"`
	SKU string 			`gorm:"column:sku"`
	VmId string			`gorm:"column:vmid"`
	AvailabilitySetName string 	`gorm:"column:availabilitysetid"`
	Provisioningstate string	`sql:"type:decimal" gorm:"column:provisioningstate"`
	ResourcegroupName string	`gorm:"column:resourcegroupname"`
	Tagname string			`gorm:"column:tagname"`
	Deleted bool     `sql:"type:varchar" gorm:"column:deleted"`

}


//---------------------------------Structure for Configuration File in Azure------------------------------------//
type Configuration struct {
    Clientid		string 		`json:"clientid"`
    Clientsecret	string		`json:"clientsecret"`
    Subscriptionid	string		`json:"subscriptionid"`
    Tenantid		string		`json:"tenantid"`
}