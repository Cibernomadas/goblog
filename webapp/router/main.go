package router

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	srv := gin.Default()
	srv.HTMLRender = templateRender()
	return srv
}

func RegisterRoutes(srv *gin.Engine) {
	srv.GET("/", holaMundo)
}

func holaMundo(c *gin.Context) {
	username := "Cibernómada"
	posts := [...]map[string]string{
		{"user": "Pedro",
			"body": "¡GoBlog mola mucho!"},
		{"user": "Juan",
			"body": "En cibernómadas publican cosas chulas."},
	}
	c.HTML(http.StatusOK, "index", gin.H{
		"title":    "Hi! GoBlog.",
		"username": username,
		"posts":    posts,
	})
}

func templateRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "webapp/template/base.html", "webapp/template/index.html")
	return r
}
