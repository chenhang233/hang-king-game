package model

type Product struct {
	Id    int     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name  string  `gorm:"column:name;type:varchar(20)" json:"name"`
	Price float64 `gorm:"column:price;type:decimal(10,2)" json:"price"`
}

func (m *Product) TableName() string {
	return "product"
}
