package router

import (
	"encoding/gob"
	"html/template"
	"math/rand"
	"path"
	"time"

	"github.com/cibernomadas/goblog/webapp/database"
	"github.com/cibernomadas/goblog/webapp/models"
	"github.com/cibernomadas/goblog/webapp/router/handlers"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	RegisterSerializers()

	srv := gin.Default()
	srv.Use(ResisterDatabase())
	srv.Use(RegisterSession())
	srv.HTMLRender = TemplateRender()
	return srv
}

func RegisterRoutes(srv *gin.Engine) {
	srv.GET("/", handlers.IndexFn)
	srv.GET("/login", handlers.LoginFn)
	srv.POST("/login", handlers.LoginFn)
	srv.GET("/register", handlers.RegisterFn)
	srv.POST("/register", handlers.RegisterFn)

	auth := srv.Group("", handlers.Authenticated)
	auth.GET("/logout", handlers.LogoutFn)
}

func ResisterDatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", database.DB)
		c.Next()
	}
}

func RegisterSerializers() {
	gob.Register(models.User{})
}

func RegisterSession() gin.HandlerFunc {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	store := cookie.NewStore([]byte("goblog"), b)
	return sessions.Sessions("sess", store)
}

func TemplateRender() *gintemplate.TemplateEngine {
	return gintemplate.New(gintemplate.TemplateConfig{
		Root:         path.Join("webapp", "template"),
		Extension:    ".tpl",
		Master:       path.Join("layouts", "base"),
		Partials:     TemplatePartials(),
		Funcs:        TemplateFuncs(),
		DisableCache: true,
	})
}

func TemplatePartials() []string {
	return []string{
		path.Join("partials", "menu"),
		path.Join("partials", "error"),
	}
}

func TemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"year": func() string {
			return time.Now().Format("2006")
		},
	}
}
