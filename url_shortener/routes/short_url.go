package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectUrl(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, "hello world")
}

func ShortUrl(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, "hello world")
}
