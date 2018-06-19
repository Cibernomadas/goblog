package database

import (
	"log"

	"github.com/cibernomadas/goblog/webapp/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data/goblog.db")
	if err != nil {
		log.Fatal(err)

	}

	db.AutoMigrate(&models.User{}, &models.Post{})
	return db
}
