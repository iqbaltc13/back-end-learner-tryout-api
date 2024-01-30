package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbaltc13/back-end-learner-tryout-api/helper"
	"github.com/iqbaltc13/back-end-learner-tryout-api/models"
)

func ListKelas(context *gin.Context) {
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	listPembayaran, err := models.FindPembayaranByUserId(user.ID)

	fmt.Println(listPembayaran)

	context.JSON(http.StatusOK, gin.H{
		"response_code": 200,
		"messages":      "Success",

		"data": listPembayaran,
	})
}
