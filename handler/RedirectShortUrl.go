package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/link33d/shorter/service"
)

func RedirectShortUrl(ctx *gin.Context) {
	code := ctx.Param("code")

	if code == "" {
		sendHtmlError(ctx, http.StatusBadRequest)
		return
	}

	link, err := service.GetShortUrlByCode(code)
	if err != nil {
		sendHtmlError(ctx, http.StatusInternalServerError)
		return
	}

	if link == nil {
		sendHtmlError(ctx, http.StatusNotFound)
		return
	}

	ctx.HTML(http.StatusOK, "redirect.html", gin.H{
		"url": link.Url,
	})

}
