package models

type Result struct {
	Id     uint8   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string  `gorm:"type:varchar(100);NOT_NULL" json:"name"`
	Gender string  `gorm:"type:varchar(100)" json:"gender"`
	Status string  `gorm:"type:varchar(100);NOT_NULL" json:"status"`
	Hobby  []Hobby `gorm:"foreignKey:Id"`
}