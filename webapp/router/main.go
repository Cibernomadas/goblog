package router

import (
	"html/template"
	"path"
	"time"

	"github.com/cibernomadas/goblog/webapp/database"
	"github.com/cibernomadas/goblog/webapp/router/handlers"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	srv := gin.Default()
	srv.Use(ResisterDatabase())
	srv.HTMLRender = TemplateRender()
	return srv
}

func RegisterRoutes(srv *gin.Engine) {
	srv.GET("/", handlers.IndexFn)
	srv.GET("/login", handlers.LoginFn)
	srv.POST("/login", handlers.LoginFn)
	srv.GET("/register", handlers.RegisterFn)
	srv.POST("/register", handlers.RegisterFn)
	srv.GET("/logout", handlers.LogoutFn)
}

func ResisterDatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", database.DB)
		c.Next()
	}
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
