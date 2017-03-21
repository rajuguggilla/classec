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
	//Volumes string 		`gorm:"column:volumes"`
	InsertionDate string 	`sql:"type:date" gorm:"column:insertion_date"`
	Tagname string		`gorm:"column:tagname"`
	Deleted bool     `sql:"type:varchar" gorm:"column:deleted"`
	Classifier	string	`gorm:"column:classifier"`

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
	Ram		int64		`json:"ram"`
	VCPUS		int64		`json:"vcpus"`
	Disk		int64		`json:"disk"`
	Links		[]SubLinks	`json:"links"`
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


//------------------------------------------------------New Structure to List Flavors --------------------------------------------------------//
type FlavorsListStruct struct {
 	Flavors 	[]FalvourSubStruct	`json:"flavors"`
}

type FalvourSubStruct struct {
	FlavorName		string		`json:"name"`
	Links			[]SubLinks	`json:"links`
	RAM			int64		`json:"ram"`
	OS_FLV_DISABLED		bool		`json:"OS-FLV-DISABLED:disabled"`
	VCPUS			int64		`json:"vcpus"`
	SWAP			string		`json:"swap"`
	OS_Flavor_Access	bool		`json:"os-flavor-access:is_public"`
	RXTX_Factor		float64		`json:"rxtx_factor"`
	OS_FLV_EXT_DATA		int64		`json:"OS-FLV-EXT-DATA:ephemeral"`
	Disk 			int64		`json:"disk"`
	FlavorId		string		`json:"id"`
}

type SubLinks struct {
	Href 		string                `json:"href"`
	Rel  		string                `json:"rel"`
}

//------------------------------------------------------New Structure to List Compute Instances--------------------------------------------------------//
type ComputeListStruct struct {
 	Servers 	[]ComputeSubStruct	`json:"servers"`
}

type ComputeSubStruct struct {
	TaskState		string		`json:"OS-EXT-STS:task_state"`
	Addresses		AddClassecNetwork	`json:"addresses`
	Links			[]SubLinks	`json:"links`
	Image			ImgStruct		`json:"image"`
	VmState			string		`json:"OS-EXT-STS:vm_state"`
	LaunchedAt		string		`json:"OS-SRV-USG:launched_at"`
	Flavor			FlavorsStruct		`json:"flavor"`
	ServerId		string		`json:"id"`
	SecurityGroups		[]SecurityGroupsSturct		`json:"security_groups"`
	UserId			string		`json:"user_id"`
	DiskConfig		string		`json:"OS-DCF:diskConfig"`
	AccessIPv4		string		`json:"accessIPv4"`
	AccessIPv6		string		`json:"accessIPv6"`
	Progress		int64		`json:"progress"`
	PowerState		int64		`json:"OS-EXT-STS:power_state"`
	AvailabilityZone	string		`json:"OS-EXT-AZ:availability_zone"`
	//Metadata		MetadataStruct	`json:"metadata"`
	Status			string		`json:"status"`
	Updated			string		`json:"updated"`
	HostId			string		`json:"hostId"`
	TerminatedAt 		string		`json:"OS-SRV-USG:terminated_at"`
	KeyName			string		`json:"key_name"`
	ServerName		string		`json:"name"`
	CreatedAt		string		`json:"created"`
	TenantId		string		`json:"tenant_id"`
	//VolumesAttached		[]VolumesStruct		`json:"os-extended-volumes:volumes_attached"`
	ConfigDrive 		string		`json:"config_drive"`

}

type AddClassecNetwork struct {
	ClassecNetwork 		[]SubAddressClassec                `json:"classec-network"`
}

type SubAddressClassec struct{
	MacAddress		string		`json:"OS-EXT-IPS-MAC:mac_addr"`
	Version			float64		`json:"version"`
	Addr			string		`json:"addr"`
	Type 			string		`json:"OS-EXT-IPS:type"`
}

type ImgStruct struct{
	ImageID		string		`json:"id"`
	Links		[]SubLinks	`json:"links"`
}

type SecurityGroupsSturct struct {
	SecurityGroupsName string `json:"name"`
}

type VolumesStruct struct{

}

type MetadataStruct struct{

}



//----------------------------------------------CEILOMETER STRUCTS--------------------------------------------//
type MeterStruct struct{
	Count 		int 		`json:"Count"`
        DurationStart 	string 		`json:"Duration_start"`
        Min 		float64 	`json:"Min"`
        DurationEnd 	string 		`json:"Duration_end"`
        Max 		float64 	`json:"Max"`
        Sum 		float64 	`json:"Sum"`
        Period 		int 		`json:"Period"`
        PeriodEnd 	string 		`json:"Period_end"`
        Duration 	float64 	`json:"Duration"`
        PeriodStart 	string 		`json:"Period_start"`
        Avg 		float64 	`json:"Avg"`
        Groupby 	string 		`json:"Groupby"`
        Unit 		string 		`json:"Unit"`
}
type CompleteDynamicResponse struct {
      Servers       []DynamicInstances_db    `json:"servers"`
}
type DynamicInstances_db struct {
	Vm_Name		string			`json:"Vm_Name"`
	InstanceID	string			`json:"InstanceId"`
	Count         int              		`json:"Count"`
	DurationStart string           		`json:"DurationStart"`
	Min           float64         		`json:"Min"`
	DurationEnd   string             	`json:"DurationEnd "`
	Max           float64       		`json:"Max"`
	Sum           float64        		`json:"Sum"`
	Period        int              		`json:"Period "`
	PeriodEnd     string             	`json:"PeriodEnd"`
	Duration      float64       		`json:"Duration"`
	PeriodStart   string              	`json:"PeriodStart"`
	Avg           float64       	 	`json:"Avg"`
	Groupby       string             	`json:"Groupby"`
	Unit          string                	`json:"Unit"`
}
type DynamicInstances struct {
	Vm_Name		string			`gorm:"column:Vm_Name"`
	InstanceID	string			`gorm:"column:InstanceID"`
	Count         int              		`gorm:"column:Count"`
	DurationStart string           		`gorm:"column:DurationStart"`
	Min           float64         		`gorm:"column:Min"`
	DurationEnd   string             	`gorm:"column:DurationEnd"`
	Max           float64       		`gorm:"column:Max"`
	Sum           float64        		`gorm:"column:Sum"`
	Period        int              		`gorm:"column:Period"`
	PeriodEnd     string             	`gorm:"column:PeriodEnd"`
	Duration      float64       		`gorm:"column:Duration"`
	PeriodStart   string              	`gorm:"column:PeriodStart"`
	Avg           float64       	 	`gorm:"column:Avg"`
	Groupby       string             	`gorm:"column:Groupby"`
	Unit          string                	`gorm:"column:Unit"`
}