package tagstruct

type Providers struct {
	InstanceId string `gorm:"column:InstanceId"`
	Cloud string `gorm:"column:Cloud"`
	Tagname string `gorm:"column:Tagname"`
}