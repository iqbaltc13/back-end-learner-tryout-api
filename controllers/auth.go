package controllers

import (
	"net/http"

	"github.com/iqbaltc13/back-end-learner-tryout-api/models"

	"github.com/iqbaltc13/back-end-learner-tryout-api/helper"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input models.RegistrationInput

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}
	if input.Password != input.ConfirmPassword {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         "Password and confirm password do not match",
		})
		return
	}

	err, userEmailExist := models.isEmailTaken(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	if userEmailExist == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         "Email is already taken",
		})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"response_code": 200,
		"data":          savedUser,
	})
}

func Login(context *gin.Context) {
	var input models.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	//user, err := model.FindUserByUsername(input.Username)
	user, err := models.FindUserByEmail(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"response_code": 200,
		"messages":      "Login Success",
		"token_jwt":     jwt,
		"data":          user,
	})
}
