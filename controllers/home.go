package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Home(context *gin.Context) string {

	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""

}
