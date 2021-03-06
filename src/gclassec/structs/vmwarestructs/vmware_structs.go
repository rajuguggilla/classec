package vmwarestructs


type VmwareInstances struct{

	Name string 		`gorm:"column:Name"`
	Uuid string 	        `gorm:"column:Uuid"`
	MemorySizeMB int32 	`gorm:"column:MemorySizeMB"`
	PowerState  string	`gorm:"column:PowerState"`
	NumofCPU int32          `gorm:"column:NumofCPU"`
	GuestFullName string 	`gorm:"column:GuestFullName"`
	IPaddress string        `gorm:"column:IPaddress"`
	StorageCommitted float32 `gorm:"column:StorageCommitted"`
	Deleted bool     	`sql:"type:varchar" gorm:"column:deleted"`
	Classifier string        `gorm:"column:classifier"`

}
type VMwareCpu struct{
	Name string `gorm:"column:name"`
	Vmid string  `gorm:"column:vmid"`
	Minimum float64 `gorm:"column:min"`
	Maximum float64 `gorm:"column:max"`
	Average float64 `gorm:"column:avg"`
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

type VmwareDynamicDetails struct{
	Name 			string			`gorm:"column:Name"`
	Uuid			string			`gorm:"column:Uuid"`
	Timestamp   		string		        `gorm:"column:Timestamp"`
	MaxCpuUsage 		int32			`gorm:"column:MaxCpuUsage"`
	MinCpuUsage 		int32			`gorm:"column:MinCpuUsage"`
	AvgCpuUsage 		int32			`gorm:"column:AvgCpuUsage"`

}

//---------------------------------Structure for Configuration File in OpenStack------------------------------------//
type Configuration struct {
    EnvURL		string		`json:"EnvURL"`
    EnvUserName		string		`json:"EnvUserName"`
    EnvPassword		string		`json:"EnvPassword"`
    EnvInsecure		string		`json:"EnvInsecure"`
}

//-------------------------------Response struct----------------------------------------//

type VmwareResponse struct{

	Name string 		`json:"Name"`
	Uuid string 	        `json:"Uuid"`
	MemorySizeMB int32 	`json:"MemorySizeMB"`
	PowerState  string	`json:"PowerState"`
	NumofCPU int32          `json:"NumofCPU"`
	GuestFullName string 	`json:"GuestFullName"`
	IPaddress string        `json:"IPaddress"`
	StorageCommitted float32 `json:"StorageCommitted"`
	Tagname	string		`json:"Tagname"`
	Deleted bool     	`json:"Deleted"`
	Classifier string        `json:"Classifier"`

}

//-------------------------------Standardized struct----------------------------------------//


type StandardizedVmware struct{

	Name string 		`json:"Name"`
	Uuid string 	        `json:"InstanceId"`
	PowerState  string	`json:"Status"`
	MemorySizeMB int32 	`json:"RAM"`
	NumofCPU int32          `json:"CPU"`
	StorageCommitted float32 `json:"Storage"`
	Tagname	string		`json:"Tagname"`
	//GuestFullName string 	`json:"GuestFullName"`
	//IPaddress string        `json:"IPAddress"`
	//Deleted bool     	`json:"Deleted"`
	//Classifier string       `json:"Classifier"`

}