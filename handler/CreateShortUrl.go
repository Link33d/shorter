package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/link33d/shorter/model"
	"github.com/link33d/shorter/service"
)

func CreateShortUrl(ctx *gin.Context) {

	// Validate the format of the request body
	request := CreateShortUrlRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Validate if the provided URL is a valid URL format
	// Generate a unique code for the URL
	// Check if the generated code already exists in the database to avoid duplicates

	code := "aaa"

	link := model.Link{
		Id:   0,
		Url:  request.Url,
		Code: code,
	}

	// Try to insert into DB
	insertedLink, err := service.CreateShortUrl(link)
	if err != nil {
		fmt.Println(err)
		sendMessage(ctx, http.StatusInternalServerError, "Something went wrong when communicating with the database")
		return
	}

	sendData(ctx, http.StatusCreated, "The link shortcut has been created successfully", insertedLink)

}
