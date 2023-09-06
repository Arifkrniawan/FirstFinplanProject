package models

type MapUserHobby struct {
	Id       uint8  `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user  string `gorm:"foreignKey;type:varchar(100)" json:"id_user"`
	Id_hobby string `gorm:"foreignKey;type:varchar(100);NOT_NULL" json:"id_hobby"`
	Status   string `gorm:"type:varchar(100);NOT_NULL" json:"status"`
}
