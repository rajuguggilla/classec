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
