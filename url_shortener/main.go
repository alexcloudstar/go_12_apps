package main

import (
	"github.com/alexcloudstar/go_12_apps/url_shortener/db"
	"github.com/alexcloudstar/go_12_apps/url_shortener/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

    db.Init()

	routes.RegisterRoutes(server)


	server.Run(":8000")
}
