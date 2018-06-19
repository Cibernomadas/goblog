package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserSession struct {
	IsAuthenticated bool `gorm:"-"`
	IsActive        bool `gorm:"-"`
	IsAnonymous     bool `gorm:"-"`
}

type User struct {
	gorm.Model
	UserSession
	Username string `gorm:"type:varchar(32);UNIQUE"`
	Email    string `gorm:"type:varchar(256);UNIQUE"`
	Password string `gorm:"type:varchar(128)"`
}

func (u *User) SetPassword(p string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err == nil
}

type Post struct {
	gorm.Model
	UserID int
	User   User
	Body   string `gorm:"type:varchar(512)"`
}
