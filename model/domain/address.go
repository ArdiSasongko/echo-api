package domain

import (
	"time"
	//"gorm.io/gorm"
)

type Address struct {
	//gorm.Model        // karena sudah menggunakan gorm.Model maka CreatedAT dan UpdatedAT tidak diperluka
	AddressID  int    `gorm:"column:address_id;primaryKey;autoIncrement"` // primary key dan autoincrement
	UserIDFK   int    `gorm:"column:user_id_fk"`                          //should unique
	City       string `gorm:"column:city"`
	Province   string `gorm:"column:province"`
	PostalCode int    `gorm:"column:postal_code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (a *Address) TableName() string {
	return "address"
}
