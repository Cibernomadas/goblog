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
	srv.GET("/", indexFn)

	srv.GET("/login", loginFn)
	srv.POST("/login", loginFn)
}

func indexFn(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Hi! GoBlog.",
	})
}

func loginFn(c *gin.Context) {
	if c.Request.Method == http.MethodGet { // Serve login page
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hi! GoBlog.",
		})
	} else if c.Request.Method == http.MethodPost { // Process login
		var login loginForm
		if err := c.ShouldBind(&login); err == nil {
			// TODO:
		} else {
			c.HTML(http.StatusOK, "index", gin.H{
				"title": "Hi! GoBlog.",
				"error": "Required fields not provided.",
			})
		}
	}
}

type loginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func templateRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "webapp/template/base.html", "webapp/template/index.html")
	return r
}
