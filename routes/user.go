package routes

import (
	"net/http"
	"strconv"

	"restapi/models"
	"restapi/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
    
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func getuserById(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param returns the value of the URL param
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user ID."})
		return
	}
	user, err := models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id.."})
		return
	}
	context.JSON(http.StatusOK, user)
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCreddentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})

}
