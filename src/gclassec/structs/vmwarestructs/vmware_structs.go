package vmwarestructs


type VmwareInstances struct{

	Name string 		`gorm:"column:Name"`
	Uuid string 	        `gorm:"column:Uuid"`
	MemorySizeMB int32 	`gorm:"column:MemorySizeMB"`
	PowerState  string	`gorm:"column:PowerState"`
	NumofCPU int32          `gorm:"column:NumofCPU"`
	GuestFullName string 	`gorm:"column:GuestFullName"`
	IPaddress string        `gorm:"column:IPaddress"`
	Tagname	string		`gorm:"column:tagname"`
	Deleted bool     	`sql:"type:varchar" gorm:"column:deleted"`
	Classifier string        `gorm:"column:classifier"`

}


type DynamicValues struct{
	VMName			string
	OverallCpuUsage         int32
	GuestMemoryUsage	int32
	StorageCommitted	float32
}
type StaticDynamicValues struct{
	VMName              	string	//`json:"Name,omitempty"`
	Uuid                	string	//`json:"Uuid,omitempty"`
	MemorySizeMB        	int32	//`json:"MemorySizeMB,omitempty"`
	PowerState          	string	//`json:"PowerState,omitempty"`
	NumCpu              	int32	//`json:"NumofCPU,omitempty"`
	GuestFullName       	string	//`json:"GuestFullName,omitempty"`
	IpAddress           	string	//`json:"IPaddress,omitempty"`
	OverallCpuUsage         int32	//`json:"OverallCpuUsage,omitempty"`
	GuestMemoryUsage	int32	//`json:"GuestMemoryUsage,omitempty"`
	StorageCommitted	float32	//`json:"StorageCommitted,omitempty"`
	MemoryOverhead  	int64	//`json:"MemoryOverhead,omitempty"`
	MaxCpuUsage 		int32	//`json:"MaxCpuUsage,omitempty"`
	Uncommitted 		int64	//`json:"Uncommitted,omitempty"`
	Unshared    		int64	//`json:"Unshared,omitempty"`
	Tagname	string	//`json:"Tagname,omitempty"`
}


//---------------------------------Structure for Configuration File in OpenStack------------------------------------//
type Configuration struct {
    EnvURL		string		`json:"EnvURL"`
    EnvUserName		string		`json:"EnvUserName"`
    EnvPassword		string		`json:"EnvPassword"`
    EnvInsecure		string		`json:"EnvInsecure"`
}

