package handlers

import (
	"net/http"

	"github.com/cibernomadas/goblog/webapp/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func IndexFn(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Hi! GoBlog.",
	})
}

func LoginFn(c *gin.Context) {
	if c.Request.Method == http.MethodGet { // Serve login page
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hi! GoBlog.",
		})
	} else if c.Request.Method == http.MethodPost { // Process login
		db, exist := c.Get("db")
		if !exist {
			c.HTML(http.StatusOK, "index", gin.H{
				"title": "Hi! GoBlog.",
				"error": "We've got an internal error, please try later.",
			})
		}
		var login models.LoginForm
		if err := c.ShouldBind(&login); err == nil {
			// TODO:
			u := models.User{
				Username: login.Username,
				Password: login.Password,
				Email:    "asd@asd.com",
			}
			db.(*gorm.DB).Create(&u)
			p := models.Post{
				Body: "asd",
				User: u,
			}
			db.(*gorm.DB).Create(&p)
		} else {
			c.HTML(http.StatusOK, "index", gin.H{
				"title": "Hi! GoBlog.",
				"error": "Required fields not provided.",
			})
		}
	}
}
