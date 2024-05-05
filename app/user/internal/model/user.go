package model

import "time"

type User struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	UserId     string    `gorm:"column:user_id;type:varchar(255);NOT NULL" json:"user_id"`
	Nickname   string    `gorm:"column:nickname;type:varchar(255);NOT NULL" json:"nickname"`
	GameLevel  int       `gorm:"column:game_level;type:tinyint(4);NOT NULL" json:"game_level"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`
	Status     int       `gorm:"column:status;type:tinyint(4);comment:0 正常 1 封号;NOT NULL" json:"status"`
}

func (m *User) TableName() string {
	return "batt"
}

type Asset struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	UserId     int       `gorm:"column:user_id;type:int(11)" json:"user_id"`
	Code       int       `gorm:"column:code;type:int(11);comment:资产代码" json:"code"`
	Number     float64   `gorm:"column:number;type:decimal(10);comment:资产数量" json:"number"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
}

func (m *Asset) TableName() string {
	return "asset"
}

type AssetLog struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	UserId     int       `gorm:"column:user_id;type:int(11)" json:"user_id"`
	Event      string    `gorm:"column:event;type:longtext;comment:资产变更日志记录;NOT NULL" json:"event"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"create_time"`
}

func (m *AssetLog) TableName() string {
	return "asset_log"
}
