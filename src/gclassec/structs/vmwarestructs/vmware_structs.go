package vmwarestructs


type VmwareInstances struct{

	Name string 		`gorm:"column:Name"`
	Uuid string 	        `gorm:"column:Uuid"`
	MemorySizeMB int32 	`gorm:"column:MemorySizeMB"`
	PowerState  string	`gorm:"column:PowerState"`
	NumofCPU int32           `gorm:"column:NumofCPU"`
	GuestFullName string 	`gorm:"column:GuestFullName"`
	IPaddress string        `gorm:"column:IPaddress"`

}


type DynamicValues struct{
	VMName			string
	OverallCpuUsage         int32
	GuestMemoryUsage	int32
	StorageCommitted	float32
}
type StaticDynamicValues struct{
	VMName              string
	Uuid                string
	MemorySizeMB        int32
	PowerState          string
	NumCpu              int32
	GuestFullName       string
	IpAddress           string
	OverallCpuUsage         int32
	GuestMemoryUsage	int32
	StorageCommitted	float32
	MemoryOverhead  		int64
	MaxCpuUsage 		 int32
	Uncommitted int64
	Unshared    int64
}

