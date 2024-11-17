package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/link33d/shorter/model"
	"github.com/link33d/shorter/service"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode() (string, error) {

	attempts := 0
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a unique code for the URL

	for {

		codeArray := make([]byte, 10)

		for i := 0; i < 10; i++ {
			codeArray[i] = charset[randGen.Intn(len(charset))]
		}

		code := string(codeArray)

		// Check if the generated code already exists in the database to avoid duplicates
		link, err := service.GetShortUrlByCode(code)
		if err != nil {
			return "", err
		}

		if link != nil {
			attempts++
			if attempts >= 10 {
				return "", fmt.Errorf("attempt limit exceeded")
			}
			continue
		}

		// Return code
		return code, nil
	}
}

func CreateShortUrl(ctx *gin.Context) {

	// Validate the format of the request body
	request := CreateShortUrlRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Validate if the provided URL is a valid URL format

	// Generate the code
	code, err := generateCode()
	if err != nil {
		fmt.Println(err)
		sendInternalServerError(ctx)
		return
	}

	// Try to insert into DB
	link := model.Link{
		Url:  request.Url,
		Code: code,
	}

	insertedLink, err := service.InsertShortUrl(link)
	if err != nil {
		fmt.Println(err)
		sendInternalServerError(ctx)
		return
	}

	// Send Success
	sendData(ctx, http.StatusCreated, "The link shortcut has been created successfully", &insertedLink)

}
