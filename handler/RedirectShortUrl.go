package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/link33d/shorter/service"
)

func RedirectShortUrl(ctx *gin.Context) {
	code := ctx.Param("code")

	// Check if code is empty
	if code == "" {
		sendHtmlError(ctx, http.StatusBadRequest)
		return
	}

	// Try getting the URL from the database
	link, err := service.GetShortUrlByCode(code)
	if err != nil {
		sendHtmlError(ctx, http.StatusInternalServerError)
		return
	}

	// If you don't find, send 404
	if link == nil {
		sendHtmlError(ctx, http.StatusNotFound)
		return
	}

	// If found, redirect
	ctx.HTML(http.StatusOK, "redirect.html", gin.H{
		"url": link.Url,
	})

}
