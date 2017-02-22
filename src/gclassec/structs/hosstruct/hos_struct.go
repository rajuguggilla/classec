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
}


//---------------------------------Structure for ComputeVM in HOS------------------------------------//

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


//---------------------------------Structure for AuthToken in HOS------------------------------------//


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

//---------------------------------Structure for Configuration file in HOS------------------------------------//
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