package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iqbaltc13/back-end-learner-tryout-api/helper"
)

func Home(context *gin.Context) {
	token := helper.ValidateJWT(context)

	fmt.Println(token)
}
