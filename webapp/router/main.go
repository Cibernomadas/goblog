package router

import (
	"github.com/cibernomadas/goblog/webapp/database"
	"github.com/cibernomadas/goblog/webapp/router/handlers"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	srv := gin.Default()
	srv.Use(ResisterDatabase())
	srv.HTMLRender = templateRender()
	return srv
}

func RegisterRoutes(srv *gin.Engine) {
	srv.GET("/", handlers.IndexFn)

	srv.GET("/login", handlers.LoginFn)
	srv.POST("/login", handlers.LoginFn)
}

func ResisterDatabase() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", database.DB)
		c.Next()
	}
}

func templateRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "webapp/template/base.html", "webapp/template/index.html")
	return r
}
