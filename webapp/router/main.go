package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	srv := gin.Default()
	return srv
}

func RegisterRoutes(srv *gin.Engine) {
	srv.GET("/", holaMundo)
}

func holaMundo(c *gin.Context) {
	c.String(http.StatusOK, "¡Hola Mundo!")
}
