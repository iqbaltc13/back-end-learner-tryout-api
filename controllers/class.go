package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iqbaltc13/back-end-learner-tryout-api/helper"
	"github.com/iqbaltc13/back-end-learner-tryout-api/models"
)

func ListClass(context *gin.Context) {
	user, err := helper.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"response_code": 500,
			"error":         err.Error(),
		})
		return
	}

	listPembayaran, err := models.FindPembayaranByUserId(user.ID)
	var ClassIds []string
	for _, element := range listPembayaran {

		ClassIds = append(ClassIds, element.Classid)

	}
	listClass, err := models.FindClassByIds(ClassIds)

	context.JSON(http.StatusOK, gin.H{
		"response_code": 200,
		"messages":      "Success",

		"data": listClass,
	})
}
