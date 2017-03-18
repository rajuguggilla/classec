package openstackInstance

type Instances struct{
	//Id int 			`gorm:"column:id"`
	Name string 		`gorm:"column:name"`
	InstanceID string 	`gorm:"column:instance_id"`
	Status string 		`gorm:"column:status"`
	AvailabilityZone string `gorm:"column:availability_zone"`
	Flavor string            `gorm:"column:flavor"`
	CreationTime string 	`gorm:"column:CreationTime"`
	FlavorID string 		`gorm:"column:flavor_id"`
	RAM int64 		`gorm:"column:ram"`
	VCPU int64 		`gorm:"column:vcpu"`
	Storage int64 		`gorm:"column:storage"`
	IPAddress string	`sql:"type:decimal" gorm:"column:ip_address"`
	SecurityGroup string 	`gorm:"column:security_group"`
	KeyPairName string 	`gorm:"column:keypair_name"`
	ImageName string 	`gorm:"column:image_name"`
	Volumes string 		`gorm:"column:volumes"`
	InsertionDate string 	`sql:"type:date" gorm:"column:insertion_date"`
	Tagname string		`gorm:"column:tagname"`
	Deleted bool     `sql:"type:varchar" gorm:"column:deleted"`

}

type ComputeResponse struct {
	Servers       []ServerResponse      `json:"servers"`
}

type ServerResponse struct {
	Id	string		`json:"id"`
	Name	string		`json:"name"`
	Image	ImageStruct	`json:"image"`
	Flavor	FlavorsStruct	`json:"flavor"`
	Status	string		`json:"status"`
        Updated	string		`json:"updated"`
        HostId	string		`json:"hostId"`
	Key_name	string	`json:"key_name"`
	Security_Groups	SubSecurityGroup `json:"security_groups"`
	Availability_Zone	string	`json:"OS-EXT-AZ:availability_zone"`
	Tenant_Id	string	`json:"tenant_id"`
	Addresses	SubAddr `json:"addresses"`
}

type SubAddress struct {
	//MacAddr		string	`json:"OS-EXT-IPS-MAC:mac_addr"`
	Version		int32	`json:"version"`
	IpAddress	string	`json:"addr"`
	//Type		string	`json:"OS-EXT-IPS:type"`
}

type SubAddr struct {
	Provider 	[]SubAddress	`json:"provider"`
}

type AddressesStruct struct{
	Addresses	SubAddr	`json:"addresses"`
	//Provider 	[]SubAddress	`json:"provider"`
}

type SubSecurityGroup struct {
	Name 	string		`json:"name"`
}

type ImageStruct struct {
	ImageID		string		`json:"id"`
}

type FlavorsStruct struct {
	FlavorID 	string 		`json:"id"`
	FlavorName 	string		`json:"name"`
	Ram		int32		`json:"ram"`
	VCPUS		int32		`json:"vcpus"`
	Disk		int32		`json:"disk"`
}

type FlvRespStruct struct {
	Flavors []FlavorsStruct		`json:"flavors"`
}

//---------------------------------Structure for Configuration File in OpenStack------------------------------------//
type Configuration struct {
    Host    string		`json:"Host"`
    Username   string		`json:"Username"`
    Password   string		`json:"Password"`
    ProjectID   string		`json:"ProjectID"`
    ProjectName   string	`json:"ProjectName"`
    Container   string		`json:"Container"`
    ImageRegion string		`json:"ImageRegion"`
    Controller string		`json:"Controller"`
}

//--------------------------Standardized struct------------------------------------//

type StandardizedOpenstack struct{
	Name string 		`json:"Name"`
	InstanceID string 	`json:"InstanceId"`
	Status string 		`json:"Status"`
	Flavor string            `json:"VmSize"`
	RAM int64 		`json:"RAM"`
	VCPU int64 		`json:"CPU"`
	Storage int64 		`json:"Storage"`
	Tagname string		`json:"Tagname"`
	//AvailabilityZone string `json:"AvailabilityZone"`
	//CreationTime string 	`json:"CreationTime"`
	//FlavorID string 		`json:"FlavorID"`
	//IPAddress string	`json:"IPAddress"`
	//SecurityGroup string 	`json:"SecurityGroup"`
	//KeyPairName string 	`json:"KeyPairName"`
	//ImageName string 	`json:"Image"`
	//Volumes string 		`json:"Volumes"`
	//InsertionDate string 	`json:"InsertionDate"`
	//Deleted bool     `json:"Deleted"`
}