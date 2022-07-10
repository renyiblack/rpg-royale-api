package main

import (
	"github.com/gin-gonic/gin"
	"rpg-royale-api/cmd/endpoints/account"
)

func main() {
	router := gin.Default()
	router.GET("/account", account.GetAccountByEmailAndPassword)
	router.POST("/account", account.PostAccount)
	//router.PUT("/account", account.PostAccount)
	//router.DELETE("/account", account.PostAccount)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
