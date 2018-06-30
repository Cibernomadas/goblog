package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cibernomadas/goblog/webapp/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/k0kubun/pp"
)

func Authenticated(c *gin.Context) {
	session := sessions.Default(c)
	u := session.Get("user")
	if u != nil {
		user := u.(models.User)
		if user.IsAuthenticated {
			dbc, exist := c.Get("db")
			if !exist {
				c.HTML(http.StatusOK, "login", gin.H{
					"title": "Hi! GoBlog.",
					"error": "We've got an internal error, please try later.",
				})
				return
			}
			db := dbc.(*gorm.DB)
			user.LastSeen = time.Now().UTC().Format("02-01-2006 15:04:05")
			db.Save(&user)
			c.Next()
			return
		}
	}
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Hi! GoBlog.",
		"error": "Restricted area, please login first.",
	})
	c.Abort()
}

func IndexFn(c *gin.Context) {
	user := models.User{}
	session := sessions.Default(c)
	u := session.Get("user")

	if u != nil {
		user = u.(models.User)
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hi! GoBlog.",
			"user":  user,
		})
		return
	}
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Hi! GoBlog.",
	})
}

func LoginFn(c *gin.Context) {
	user := models.User{}
	session := sessions.Default(c)
	u := session.Get("user")
	if u != nil {
		user = u.(models.User)
	}

	if u != nil && user.IsAuthenticated {
		// There is no need to see login page
		// Showing index
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hi! GoBlog.",
			"user":  user,
		})
		return
	}

	// The user is not authenticated
	if c.Request.Method == http.MethodGet { // Serve login page
		c.HTML(http.StatusOK, "login", gin.H{
			"title": "Hi! GoBlog.",
		})
	} else if c.Request.Method == http.MethodPost { // Process login
		dbc, exist := c.Get("db")
		if !exist {
			c.HTML(http.StatusOK, "login", gin.H{
				"title": "Hi! GoBlog.",
				"error": "We've got an internal error, please try later.",
			})
			return
		}

		var login models.LoginForm
		db := dbc.(*gorm.DB)
		if err := c.ShouldBind(&login); err == nil {
			db.Where(&models.User{Username: login.Username}).First(&user)
			if user.CheckPassword(login.Password) {
				user.IsAuthenticated = true
				user.LastSeen = time.Now().UTC().Format("02-01-2006 15:04:05")
				session.Set("user", user)
				if err = session.Save(); err != nil {
					c.HTML(http.StatusOK, "login", gin.H{
						"title": "Hi! GoBlog.",
						"error": "Internal error. Please try later",
					})
					return
				}
				c.HTML(http.StatusOK, "index", gin.H{
					"title": "Hi! GoBlog.",
					"user":  user,
				})
			} else {
				c.HTML(http.StatusOK, "login", gin.H{
					"title": "Hi! GoBlog.",
					"error": "Username or Password incorrect, please try again.",
				})
			}
		} else {
			c.HTML(http.StatusOK, "login", gin.H{
				"title": "Hi! GoBlog.",
				"error": "Required fields are not provided correctly.",
			})
		}
	}
}

func LogoutFn(c *gin.Context) {
	user := models.User{}
	session := sessions.Default(c)
	u := session.Get("user")
	if u != nil {
		user = u.(models.User)
		user.IsAuthenticated = false
	}
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Hi! GoBlog.",
	})
}

func RegisterFn(c *gin.Context) {
	user := models.User{}
	session := sessions.Default(c)
	u := session.Get("user")
	if u != nil {
		user = u.(models.User)
	}

	if u != nil && user.IsAuthenticated {
		// There is no need to see login page
		// Showing index
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hi! GoBlog.",
			"user":  user,
		})
		return
	}
	// The user is not authenticated
	if c.Request.Method == http.MethodGet { // Serve registration page
		c.HTML(http.StatusOK, "register", gin.H{
			"title": "Hi! GoBlog.",
		})
	} else if c.Request.Method == http.MethodPost { // Process registration
		dbc, exist := c.Get("db")
		if !exist {
			c.HTML(http.StatusOK, "register", gin.H{
				"title": "Hi! GoBlog.",
				"error": "We've got an internal error, please try later.",
			})
			return
		}

		var register models.RegistrationForm
		db := dbc.(*gorm.DB)
		if err := c.ShouldBind(&register); err == nil {
			if register.CheckPasswords() {
				user.Username = register.Username
				user.Email = register.Email
				user.SetPassword(register.PasswordA)
				db.Create(&user)
				if db.NewRecord(user) { // New record means user has not a ID set
					c.HTML(http.StatusOK, "register", gin.H{
						"title": "Hi! GoBlog.",
						"error": "The username you have entered is already taken. Please choose another one.",
					})
					return
				}
				user.IsAuthenticated = true
				user.LastSeen = time.Now().UTC().Format("02-01-2006 15:04:05")
				session.Set("user", user)
				if err = session.Save(); err != nil {
					c.HTML(http.StatusOK, "index", gin.H{
						"title": "Hi! GoBlog.",
						"error": "Internal error. Please try later.",
					})
					return
				}
				c.HTML(http.StatusOK, "index", gin.H{
					"title": "Hi! GoBlog.",
					"user":  user,
				})
				return
			}
			c.HTML(http.StatusOK, "register", gin.H{
				"title": "Hi! GoBlog.",
				"error": "Passwords do not match.",
			})
			return
		}
		c.HTML(http.StatusOK, "register", gin.H{
			"title": "Hi! GoBlog.",
			"error": "Required fields are not provided correctly.",
		})
		return
	}
}

func ProfileFn(c *gin.Context) {
	currentUser := models.User{}
	session := sessions.Default(c)
	u := session.Get("user")
	if u != nil {
		currentUser = u.(models.User)
	}
	dbc, exist := c.Get("db")
	if !exist {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hi! GoBlog.",
			"error": "We've got an internal error, please try later.",
		})
		return
	}
	db := dbc.(*gorm.DB)
	username := c.Param("name")
	user := models.User{}
	db.Where(&models.User{Username: username}).First(&user)
	if user.Username != "" {
		posts := [...]struct {
			Author *models.User
			Body   string
		}{
			{&user, "Test post #1"},
			{&user, "Test post #2"},
		}
		c.HTML(http.StatusOK, "profile", gin.H{
			"title":        "Hi! GoBlog.",
			"user":         &user,
			"posts":        posts,
			"current_user": &currentUser,
		})
		return
	} else {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hi! GoBlog.",
			"error": "User not found",
		})
		return
	}
}

func EditProfileFn(c *gin.Context) {
	currentUser := models.User{}
	session := sessions.Default(c)
	u := session.Get("user")
	if u != nil {
		currentUser = u.(models.User)
	}
	dbc, exist := c.Get("db")
	if !exist {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hi! GoBlog.",
			"error": "We've got an internal error, please try later.",
		})
		return
	}
	if c.Request.Method == http.MethodGet { // Edit profile form
		c.HTML(http.StatusOK, "edit_profile", gin.H{
			"title": "Hi! GoBlog.",
			"user":  currentUser,
		})
	} else if c.Request.Method == http.MethodPost { // Process profile form
		var profileForm models.EditProfile
		if err := c.ShouldBind(&profileForm); err == nil {
			pp.Println(currentUser)
			pp.Println(profileForm)
			fmt.Println(currentUser.Username == profileForm.Username)
			if currentUser.Username == profileForm.Username {
				db := dbc.(*gorm.DB)
				currentUser.AboutMe = profileForm.AboutMe
				pp.Println(currentUser)
				db.Save(&currentUser)
				posts := [...]struct {
					Author *models.User
					Body   string
				}{
					{&currentUser, "Test post #1"},
					{&currentUser, "Test post #2"},
				}
				c.HTML(http.StatusOK, "profile", gin.H{
					"title": "Hi! GoBlog.",
					"user":  &currentUser,
					"posts": posts,
				})
			} else {
				c.HTML(http.StatusOK, "edit_profile", gin.H{
					"title": "Hi! GoBlog.",
					"error": "You do not own this profile.",
				})
				return
			}
		} else {
			c.HTML(http.StatusOK, "edit_profile", gin.H{
				"title": "Hi! GoBlog.",
				"error": "Required fields are not provided correctly.",
			})
			return
		}
	}
}
