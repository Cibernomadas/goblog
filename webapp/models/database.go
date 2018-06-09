package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);UNIQUE"`
	Email    string `gorm:"type:varchar(256);UNIQUE"`
	Password string `gorm:"type:varchar(128)"`
}

type Post struct {
	gorm.Model
	UserID int
	User   User
	Body   string `gorm:"type:varchar(512)"`
}
