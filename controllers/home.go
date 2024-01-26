package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

func extractTokenFromHeader(header string) (string, error) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid or missing bearer token")
	}
	return parts[1], nil
}

func Home(w http.ResponseWriter, r *http.Request) {

	tokenString, err := extractTokenFromHeader(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Println(tokenString)

}
