package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
