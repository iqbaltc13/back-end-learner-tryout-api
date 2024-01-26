package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbaltc13/back-end-learner-tryout-api/helper"
)

func Home(context *gin.Context) {
	user, err := helper.CurrentUser(context)
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

		"data": user,
	})
}
