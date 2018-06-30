package models

import (
	"crypto/md5"
	"fmt"

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
	AboutMe  string `gorm:"type:varchar(512)"`
	LastSeen string `gorm:"type:varchar(20)"`
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

func (u *User) Avatar(s int) string {
	if s < 1 || s > 2048 {
		s = 128
	}
	email := []byte(u.Email)
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?d=identicon&s=%d", md5.Sum(email), s)
}

type Post struct {
	gorm.Model
	UserID int
	User   User
	Body   string `gorm:"type:varchar(512)"`
}
