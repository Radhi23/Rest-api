package middleware

import (
	"fmt"
	"net/http"
	"userapp/rest-api/models"
	"userapp/rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	fmt.Println(err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get the event"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
