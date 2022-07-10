package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpg-royale-api/cmd/entities"
)

var accounts = []entities.Account{
	{ID: 1, Name: "victor", Email: "victor@hotmail.com", Password: "123"},
}

func GetAccountByEmailAndPassword(c *gin.Context) {
	var account entities.Account

	if err := c.BindJSON(&account); err != nil {
		return
	}

	for _, a := range accounts {
		if a.Email == account.Email && a.Password == account.Password {
			c.IndentedJSON(http.StatusOK, gin.H{"name": a.Name})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "account not found"})
}

func PostAccount(c *gin.Context) {
	var newAccount entities.Account

	if err := c.BindJSON(&newAccount); err != nil {
		return
	}

	for _, a := range accounts {
		if a.Email == newAccount.Email {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "email already in use"})
			return
		}
	}

	newAccount.ID = len(accounts)

	accounts = append(accounts, newAccount)
	newAccount.Password = ""
	c.IndentedJSON(http.StatusCreated, newAccount)
}
