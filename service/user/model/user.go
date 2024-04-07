package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"type:varchar(128);not null"`
	Gender   int8   `gorm:"type:tinyint(1);default:0"`
	Phone    string `gorm:"unique"`
}
