package tagstruct

type Providers struct {
	InstanceId string `gorm:"column:InstanceId"`
	InstanceName	string	`gorm:"column:InstanceName"`
	Cloud string `gorm:"column:Cloud"`
	Tagname string `gorm:"column:Tagname"`
}