package models

type User struct {
	Id        uint8   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string  `gorm:"type:varchar(100);NOT_NULL" json:"name"`
	Gender    string  `gorm:"type:varchar(100)" json:"gender"`
	Status    string  `gorm:"type:varchar(100);NOT_NULL" json:"status"`
	UserHobby []Hobby `gorm:"foreignKey:Id" json:"hobbies"`
}
