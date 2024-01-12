package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/iqbaltc13/back-end-learner-tryout-api/models"

	"github.com/iqbaltc13/back-end-learner-tryout-api/helper"

	"fmt"
	"time"

	"os/exec"

	"strings"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input models.RegistrationInput
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated UUID:")
	//fmt.Printf("%s", newUUID)
	currentTime := time.Now()

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
	idUser := strings.TrimSuffix(strings.ToLower(string(newUUID)), "\n")
	userEmailExist, err := models.FindUserByEmail(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	if userEmailExist.Email == input.Email {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         "Email is already taken",
		})
		return
	}

	user := models.User{
		ID:                    idUser,
		Name:                  input.Name,
		Username:              input.Username,
		Email:                 input.Email,
		Password:              input.Password,
		Phone:                 input.Phone,
		CurrentApkVersionName: input.CurrentApkVersionName,
		CurrentApkVersionCode: input.CurrentApkVersionCode,
		DeviceInfo:            input.DeviceInfo,
		CreatedAt:             currentTime.Format("2006-01-02 15:04:05"),
	}

	savedUser, err := user.Save()

	notifikasi := models.Notifikasi{
		ID:       strings.TrimSuffix(strings.ToLower(string(newUUID)), "\n"),
		Title:    "email_notification",
		Subtitle: "email_verification_after_regis",
		Action:   "redirect web page",

		ReceiverID: idUser,

		CreatedAt: currentTime.Format("2006-01-02 15:04:05"),
	}

	savedNotifikasi, err := notifikasi.Save()
	_ = savedNotifikasi
	var client = &http.Client{}

	request, err := http.NewRequest("GET", os.Getenv("BASE_URL_ADMIN_APP")+"/batch-send-verification-email", nil)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	response, err := client.Do(request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}
	defer response.Body.Close()

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
	//fmt.Println(user.Id)

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
