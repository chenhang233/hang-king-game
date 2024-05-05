package model

type SysSet struct {
	Id            int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	SysKey        string `gorm:"column:sys_key;type:varchar(255);NOT NULL" json:"sys_key"`
	SysValue      string `gorm:"column:sys_value;type:longtext;NOT NULL" json:"sys_value"`
	SysAnnotation string `gorm:"column:sys_annotation;type:longtext" json:"sys_annotation"`
}

func (m *SysSet) TableName() string {
	return "sys_set"
}
