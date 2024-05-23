package main

import (
	"github.com/alexcloudstar/go_12_apps/url_shortener/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    server := gin.Default()

    routes.RegisterRoutes(server)

    server.Run(":8000")
}
