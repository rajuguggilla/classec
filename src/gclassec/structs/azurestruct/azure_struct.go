package azurestruct

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
	Status string `gorm:"column:status"`
	Storage int32			`gorm:"column:storage"`
	RAM int32			`gorm:"column:ram"`
	NumCPU	int32			`gorm:"column:numcpu"`

}

type AzureDynamic struct {
	Name		string		`gorm:"column:name"`
	Timestamp	string		`gorm:"column:timestamp"`
	Minimum		float64		`sql:"type:varchar" gorm:"column:minimum"`
	Maximum		float64		`sql:"type:varchar" gorm:"column:maximum"`
	Average		float64		`sql:"type:varchar" gorm:"column:average"`
}


//---------------------------------Structure for Configuration File in Azure------------------------------------//
type Configuration struct {
    Clientid		string 		`json:"clientid"`
    Clientsecret	string		`json:"clientsecret"`
    Subscriptionid	string		`json:"subscriptionid"`
    Tenantid		string		`json:"tenantid"`
}


//--------------------------Standardized struct------------------------------------//
type AzureStandardResponse struct {
	Servers       []StandardizedAzure      `json:"servers"`
}

type StandardizedAzure struct{
	SubscriptionId	string		`json:"Classifier"`
	VmName string 			`json:"Name"`
	Type string 			`json:"Type"`
	Location string 		`json:"Location"`
	VmSize string           	`json:"VmSize"`
	Publisher string 		`json:"Publisher"`
	Offer string 			`json:"Offer"`
	SKU string 			`json:"SKU"`
	VmId string			`json:"InstanceId"`
	AvailabilitySetName string 	`json:"AvailabilitySetName"`
	Provisioningstate string	`json:"Provisioningstate"`
	ResourcegroupName string	`json:"ResourcegroupName"`
	Status string			`json:"Status"`
	Storage int32			`json:"Storage"`
	RAM int32			`json:"RAM"`
	NumCPU	int32			`json:"CPU"`
	Tagname string			`json:"Tagname"`
	Deleted bool     `json:"Deleted"`

}