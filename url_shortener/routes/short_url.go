package routes

import (
	"fmt"
	"net/http"

	"github.com/alexcloudstar/go_12_apps/url_shortener/utils"
	"github.com/gin-gonic/gin"
)

type Url struct {
    Short string `json:"short"`
    Long string `json:"long"`
}

type Body struct {
    Url string `json:"url"`
}

var urls []Url

func RedirectUrl(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
        "urls": urls,
    })
}

func ShortUrl(ctx *gin.Context) {
    var body Body

    err := ctx.BindJSON(&body)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    hashedUrl := utils.GetHash(5)

    urls = append(urls, Url{
        Short: hashedUrl,
        Long: body.Url,
    })

    ctx.JSON(http.StatusOK, gin.H{
        "url": hashedUrl,
    })
}
