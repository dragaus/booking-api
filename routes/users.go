package routes

import (
	"net/http"

	"example.com/booking/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "User cannot be saved"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created succesfully"})
}