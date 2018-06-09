package main

import (
	"github.com/cibernomadas/goblog/webapp/database"
	"github.com/cibernomadas/goblog/webapp/router"
)

func main() {
	database.DB = database.Init()
	defer database.DB.Close()

	srv := router.NewServer()
	router.RegisterRoutes(srv)

	srv.Run("127.0.0.1:8000")
}
