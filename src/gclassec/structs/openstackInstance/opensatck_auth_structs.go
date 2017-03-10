package openstackInstance


//------------------------------------------------------Responsebody Structure For V2 authToken Request--------------------------------------------------------//
type OpenStackAutToken_v2 struct{
	Access 	AccessStruct_v2	`json:"access"`
}


type  AccessStruct_v2 struct {
	Token  		TokenStruct_v2		`json:"token"`
	ServiceCatalog	[]ServiceStructure_v2	`json:"serviceCatalog"`
	User		UserStruct_v2		`json:"user"`
	Metadata	Metadata_v2		`json:"metadata"`

}

type TokenStruct_v2 struct{
	Issued_at	string		`json:"issued_at"`
	Expires		string		`json:"expires"`
	AuthToken	string		`json:"id"`
	Tenant		TenantStruct_v2	`json:"tenant"`
	Audit_ids	[]string	`json:"audit_ids"`
}
type TenantStruct_v2 struct{
	Description	string		`json:"description"`
	Enabled		bool		`json:"enabled"`
	TenanatID	string		`json:"id"`
	TenantName	string		`json:"name"`
}

type ServiceStructure_v2 struct{
	Endpoints		[]EndpointsStruct_v2	`json:"endpoints"`
	Endpoints_links		[]string		`json:"endpoints_links"`
	EndpointType		string			`json:"type"`
	EndpointName		string			`json:"name"`
}
type EndpointsStruct_v2 struct{
	AdminURL		string	`json:"adminURL"`
	Region			string	`json:"region"`
	EndpiontID		string	`json:"id"`
	InternalURL		string	`json:"internalURL"`
	PublicURL		string	`json:"publicURL"`
}

type UserStruct_v2 struct{
	UserName	string		`json:"username"`
	Roles_links	[]string	`json:"roles_links"`
	UserID		string		`json:"id"`
	Roles		[]Roles_v2		`json:"roles"`
	Name		string		`json:"name"`
}
type Roles_v2 struct{
	RoleName 	string		`json:"name"`
}

type Metadata_v2 struct{
	Is_admin	int64		`json:"is_admin"`
	Roles		[]string	`json:"roles"`
}



//------------------------------------------------------Responsebody Structure For V3 authToken Request--------------------------------------------------------//

type OpenStackAutToken_v3 struct{
	Token 	AccessStruct_v3	`json:"token"`
}

type  AccessStruct_v3 struct {
	Is_Domain  		bool			`json:"is_domain"`
	Methods			[]string		`json:"methods"`
	Roles			[]DemoStruct1_v3		`json:"roles"`
	Expires_At		string			`json:"expires_at"`
	Project  		DemoStruct2_v3		`json:"project"`
	Catalog			[]SingleCatalogStruct_v3	`json:"catalog"`
	User			DemoStruct2_v3		`json:"user"`
	Audit_Ids		[]string		`json:"audit_ids"`
	Issued_At		string			`json:"issued_at"`
}

type DemoStruct1_v3 struct{
	Id	string		`json:"id"`
	Name	string		`json:"name"`
}

type DemoStruct2_v3 struct{
	Domain 	DemoStruct1_v3	`json:"domain"`
	Id	string		`json:"id"`
	Name	string		`json:"name"`
}

type SingleEndpointStruct_v3 struct{
	Region_Id	string		`json:"region_id"`
	URL		string		`json:"url"`
	Interface	string		`json:"interface"`
	Id		string		`json:"id"`
	Region		string		`json:"region"`
}

type SingleCatalogStruct_v3 struct{
	Endpoints 	[]SingleEndpointStruct_v3	`json:"endpoints"`
	Id		string			`json:"id"`
	Type		string			`json:"type"`
	Name		string			`json:"name"`
}


//------------------------------------------------------Structure to List All Api Endpoints --------------------------------------------------------//
type OpenStackEndpoints struct {
 	ApiEndpoints 	[]EndpointStruct
}

type EndpointStruct struct {
	EndpointName	string
	EndpointURL	string
	EndpointType	string
}

//------------------------------------------------------Structure to read Configuration file --------------------------------------------------------//

type OpenStackUserConfig struct {
	IdentityEndpoint	string	`json:"identityEndpoint"`
    	UserName		string	`json:"userName"`
	Password		string	`json:"password"`
	Domain			string	`json:"domain"`
    	TenantName 		string	`json:"tenantName"`
    	TenantId 		string	`json:"tenantID"`
	ProjectId		string	`json:"projectID"`
	ProjectName		string	`json:"projectName"`
    	Container 		string	`json:"container"`
    	Region	 		string	`json:"region"`
	Controller 		string	`json:"controller"`
}
