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