package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse request because %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	err = user.Save()
	if err != nil {
		message := fmt.Sprintf("Couldn't save user because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse request because %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		message := fmt.Sprintf("Couldn't login because %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": message})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		message := fmt.Sprintf("Couldn't login because %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})
}
