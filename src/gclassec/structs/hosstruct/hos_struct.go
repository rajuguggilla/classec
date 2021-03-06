package hosstruct

type HosInstances struct {
	Vm_Name	string		`gorm:"column:Name"`
	InstanceID	string	`gorm:"column:Instance_id"`
	FlavorID	string	`gorm:"column:Flavor_id"`
	FlavorName	string	`gorm:"column:Flavor_Name"`
	Status	string		`gorm:"column:Status"`
	Image	string		`gorm:"column:Image"`
	SecurityGroups	string	`gorm:"column:Security_Group"`
	AvailabilityZone	string	`gorm:"column:Availability_Zone"`
	//IPAddress	string	`gorm:"column:ip_address"`
	KeypairName	string	`gorm:"column:keypair_name"`
	Ram	int32	`gorm:"column:ram"`
	VCPU	int32	`gorm:"column:vcpu"`
	Disk	int32	`gorm:"column:disk"`
	//Volumes_Attached	string	`gorm:"column:volumes_attached"`
	Deleted bool     `sql:"type:varchar" gorm:"column:deleted"`
	Classifier string	`gorm:"column:classifier"`
}


type HosDynamicInstances struct {
        Vm_Name string          `gorm:"column:Name"`
        InstanceID      string  `gorm:"column:Instance_id"`
        Count int `gorm:"column:Count"`
        DurationStart string `sql:"type:timestamp" gorm:"column:Duration_start"`
        Min float64 `gorm:"column:Min"`
        DurationEnd string `sql:"type:timestamp" gorm:"column:Duration_end"`
        Max float64 `gorm:"column:Max"`
        Sum float64 `gorm:"column:Sum"`
        Period int `gorm:"column:Period"`
        PeriodEnd string `sql:"type:timestamp" gorm:"column:Period_end"`
        Duration float64 `gorm:"column:Duration"`
        PeriodStart string `sql:"type:timestamp" gorm:"column:Period_start"`
        Avg float64 `gorm:"column:Avg"`
        Unit string `gorm:"column:Unit"`

}

type HOSCpu struct{
	Name string `gorm:"column:name"`
	Vmid string `gorm:"column:vmid"`
	Minimum float64 `gorm:"column:min"`
	Maximum float64 `gorm:"column:max"`
	Average float64 `gorm:"column:avg"`
}


//-------------------------------- Structure to get AuthToken-----------------------------------------//

type HOSAutToken struct{
	Access 	AccessStruct	`json:"access"`
}


type  AccessStruct struct {
	Token  		TokenStruct		`json:"token"`
	ServiceCatalog	[]ServiceCatalogStruct	`json:"serviceCatalog"`
	User		UserStruct		`json:"user"`
	Metadata	Metadata		`json:"metadata"`

}

type TokenStruct struct{
	Issued_at	string		`json:"issued_at"`
	Expires		string		`json:"expires"`
	AuthToken	string		`json:"id"`
	Tenant		TenantStruct	`json:"tenant"`
	Audit_ids	[]string	`json:"audit_ids"`
}
type TenantStruct struct{
	Description	string		`json:"description"`
	Enabled		bool		`json:"enabled"`
	TenanatID	string		`json:"id"`
	TenantName	string		`json:"name"`
}

type ServiceCatalogStruct struct{
	Endpoints		[]EndpointsStruct	`json:"endpoints"`
	Endpoints_links		[]string		`json:"endpoints_links"`
	EndpointType		string			`json:"type"`
	EndpointName		string			`json:"name"`
}
type EndpointsStruct struct{
	AdminURL		string	`json:"adminURL"`
	Region			string	`json:"region"`
	EndpiontID		string	`json:"id"`
	InternalURL		string	`json:"internalURL"`
	PublicURL		string	`json:"publicURL"`
}

type UserStruct struct{
	UserName	string		`json:"username"`
	Roles_links	[]string	`json:"roles_links"`
	UserID		string		`json:"id"`
	Roles		[]RolesStruct	`json:"roles"`
	Name		string		`json:"name"`
}
type RolesStruct struct{
	RoleName 	string		`json:"name"`
}

type Metadata struct{
	Is_admin	int64		`json:"is_admin"`
	Roles		[]string	`json:"roles"`
}





type Configuration struct {
	IdentityEndpoint	string	`json:"IdentityEndpoint"`
    	UserName		string	`json:"userName"`
	Password		string	`json:"password"`
    	TenantName 		string	`json:"tenantName"`
    	TenantId 		string	`json:"tenantID"`
	ProjectId		string	`json:"projectID"`
	ProjectName		string	`json:"projectName"`
    	Container 		string	`json:"container"`
    	Region	 		string	`json:"region"`
}




//---------------------------------Structure for CcmputeVM in HOS------------------------------------//

type ComputeResponse struct {
      Servers       []ServersResponse      `json:"servers"`
}

type ServersResponse struct {
       Status     	      string `json:"status"`
       Updated                string `json:"updated"`
       HostId                 string `json:"hostId"`
       HostName		      string `json:"OS-EXT-SRV-ATTR:host"`
       //Addresses	      AddressesStruct `json:"addresses"`
       //Links		      []SubLinks	`json:"links"`
       Image 		      ImageStruct	`json:"image"`
       Key_name               string `json:"key_name"`
       Task_State	      string	`json:"OS-EXT-STS:task_state"`
       Vm_State		string		`json:"OS-EXT-STS:vm_state"`
       //InstanceName           string `json:"OS-EXT-SRV-ATTR:instance_name"`
       //Launched_At           string `json:"OS-SRV-USG:launched_at"`
       //Hypervisor_Hostname           string `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
       InstanceID             string `json:"id"`
       Flavor                 FlavorsStruct `json:"flavor"`
       Security_Groups       SubSecurityGroup `json:"security_groups"`
       //Terminated_At               string `json:"OS-SRV-USG:terminated_at"`
       Availability_Zone               string `json:"OS-EXT-AZ:availability_zone"`
       User_Id               string `json:"user_id"`
       Vm_Name               string `json:"name"`
       //Created_At               string `json:"created"`
       Tenant_Id               string `json:"tenant_id"`
       //DiskConfig               string `json:"OS-DCF:diskConfig"`
       //volumes_attached               string `json:"os-extended-volumes:volumes_attached"`
       //AccessIPv4               string `json:"accessIPv4"`
       //AccessIPv6               string `json:"accessIPv6"`
       //Progress               int32 `json:"progress"`
       Power_State               int32 `json:"OS-EXT-STS:power_state"`
       //Config_Drive               string `json:"config_drive"`
       //Metadata               string `json:"metadata"`


}

type SubAddress struct {
	MacAddr		string	`json:"OS-EXT-IPS-MAC:mac_addr"`
	Version		string	`json:"version"`
	IpAddress	string	`json:"addr"`
	Type		string	`json:"OS-EXT-IPS:type"`
}

type AddressesStruct struct{
	Lbpvtnet 	[]SubAddress	`json:"lbpvtnet"`
}
type SubLinks struct {
	Href 	string	`json:"href"`
	Rel 	string 	`json:"rel"`
}

type ImageStruct struct {
	ImageID		string		`json:"id"`
	//ImageLinks	SubLinks	`json:"links"`
}

type SubSecurityGroup struct {
	Name 	string		`json:"name"`
}


//---------------------------------Structure for Flavors in HOS------------------------------------//


type FlavorsStruct struct{

	FlavorID 	string 		`json:"id"`
	FlavorName 	string		`json:"name"`
	Ram		int32		`json:"ram"`
	VCPUS		int32		`json:"vcpus"`
	Disk		int32		`json:"disk"`
	//Links 		SubLinks	`json:"links"`

}
type FlvRespStruct struct {
	Flavors []FlavorsStruct		`json:"flavors"`
}




//-------------------------------- Structure for Dynamic Details in HOS --------------------------//

type DynamicData []struct {
        Count int `json:"Count"`
        DurationStart string `json:"Duration_start"`
        Min float64 `json:"Min"`
        DurationEnd string `json:"Duration_end"`
        Max float64 `json:"Max"`
        Sum float64 `json:"Sum"`
        Period int `json:"Period"`
        PeriodEnd string `json:"Period_end"`
        Duration float64 `json:"Duration"`
        PeriodStart string `json:"Period_start"`
        Avg float64 `json:"Avg"`
        Groupby Group `json:"Groupby"`
        Unit string `json:"Unit"`
 }


type Group struct{
        GroupBy string `json:"Groupby"`
}

type DynamicDataResponse struct {
        Data []DynamicData
}


//-------------------------------- Structure for Complete Dynamic Data --------------------------


type CompleteDynamicResponse struct {
      Servers       []CompleteDynamicData      `json:"servers"`
}

type CompleteDynamicData struct {
	InstanceID             string `json:"id"`
	Vm_Name               string `json:"name"`
        Count int `json:"Count"`
        DurationStart string `json:"Duration_start"`
        Min float64 `json:"Min"`
        DurationEnd string `json:"Duration_end"`
        Max float64 `json:"Max"`
        Sum float64 `json:"Sum"`
        Period int `json:"Period"`
        PeriodEnd string `json:"Period_end"`
        Duration float64 `json:"Duration"`
        PeriodStart string `json:"Period_start"`
        Avg float64 `json:"Avg"`
        Groupby Group `json:"Groupby"`
        Unit string `json:"Unit"`
 }






///----------------------------------Structure for Complete Data(static data with avg cpu_util--------//


type CompleteComputeResponse struct {
      Servers       []CompleteServersResponse      `json:"servers"`
}


type CompleteServersResponse struct {
       Status     	      string `json:"status"`
       Updated                string `json:"updated"`
       HostId                 string `json:"hostId"`
       HostName		      string `json:"OS-EXT-SRV-ATTR:host"`
       Image 		      ImageStruct	`json:"image"`
       Key_name               string `json:"key_name"`
       Task_State	      string	`json:"OS-EXT-STS:task_state"`
       Vm_State		string		`json:"OS-EXT-STS:vm_state"`
       InstanceID             string `json:"id"`
       Flavor                FlavorsStruct `json:"flavor"`
       Security_Groups       SubSecurityGroup `json:"security_groups"`
       Availability_Zone               string `json:"OS-EXT-AZ:availability_zone"`
       User_Id               string `json:"user_id"`
       Vm_Name               string `json:"name"`
       Tenant_Id               string `json:"tenant_id"`
       Power_State               int32 `json:"OS-EXT-STS:power_state"`
	Cpu_Util	float64		`json:"Cpu_util"`
	Tagname	string	`json:"Tagname"`

}



//--------------- Structure to get Avg of each instacne---------------------//

type LatestDynamicData struct {
        Count int `json:"Count"`
        DurationStart string `json:"Duration_start"`
        Min float64 `json:"Min"`
        DurationEnd string `json:"Duration_end"`
        Max float64 `json:"Max"`
        Sum float64 `json:"Sum"`
        Period int `json:"Period"`
        PeriodEnd string `json:"Period_end"`
        Duration float64 `json:"Duration"`
        PeriodStart string `json:"Period_start"`
        Avg float64 `json:"Avg"`
        Groupby Group `json:"Groupby"`
        Unit string `json:"Unit"`
 }

//--------------------------Response struct------------------------------------//
type HosResponse struct {
	Vm_Name	string		`json:"Vm_Name"`
	InstanceID	string	`json:"InstanceID"`
	FlavorID	string	`json:"FlavorID"`
	FlavorName	string	`json:"FlavorName"`
	Status	string		`json:"Status"`
	Image	string		`json:"Image"`
	SecurityGroups	string	`json:"SecurityGroups"`
	AvailabilityZone	string	`json:"AvailabilityZone"`
	//IPAddress	string	`gorm:"column:ip_address"`
	KeypairName	string	`json:"KeypairName"`
	Ram	int32	`json:"Ram"`
	VCPU	int32	`json:"VCPU"`
	Disk	int32	`json:"Disk"`
	//Volumes_Attached	string	`gorm:"column:volumes_attached"`
	Tagname	string		`json:"Tagname"`
	Deleted bool     `json:"Deleted"`
	Classifier string	`json:"Classifier"`
}



//--------------------------Standardized struct------------------------------------//

type StandardizedHos struct {
	Vm_Name	string		`json:"Name"`
	InstanceID	string	`json:"InstanceId"`
	Status	string		`json:"Status"`
	FlavorName	string	`json:"VmSize"`
	Ram	int32	`json:"RAM"`
	VCPU	int32	`json:"CPU"`
	Disk	int32	`json:"Storage"`
	Tagname	string	`json:"Tagname"`
	//FlavorID	string	`json:"FlavorID"`
	//Image	string		`json:"Image"`
	//SecurityGroups	string	`json:"SecurityGroup"`
	//AvailabilityZone	string	`json:"AvailabilityZone"`
	//IPAddress	string	`json:"IPAddress"`
	//KeypairName	string	`json:"KeyPairName"`
	//Volumes_Attached	string	`json:"Volumes"`
	//Deleted bool     `json:"Deleted"`
}