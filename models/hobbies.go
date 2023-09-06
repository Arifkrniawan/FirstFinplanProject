package models

type Hobby struct {
	Id    uint8  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:varchar(100);NOT_NULL" json:"name"`
	Level int    `gorm:"type:int" json:"level"`
}
