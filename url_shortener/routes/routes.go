package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
    server.GET("/:shortUrl", RedirectUrl)
    server.POST("/:shortUrl", ShortUrl)

    server.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
}
