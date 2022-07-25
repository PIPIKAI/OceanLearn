package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username" gorm:"type:varchar(20);not null"`
	Password  string `json:"password" gorm:"size:255;not null"`
	Telephone string `json:"telephone" gorm:"type:varchar(110);not null;unique"`
}
