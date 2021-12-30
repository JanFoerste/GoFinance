package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gofinance/actions/AccountActions"
	"gofinance/initializers"
	"gofinance/models"
	"io/ioutil"
	"net/http"
)

func AccountRoutes(router *gin.Engine) {
	router.GET("/accounts", getAccount)
	router.GET("/accounts/:id", findAccount)
	router.POST("/accounts", createAccount)
	router.DELETE("/accounts/:id", removeAccount)
	router.PATCH("/accounts/:id", updateAccount)
}

func getAccountById(id string, account *models.Account, context *gin.Context) {
	if err := initializers.DB.Where("id = ?", id).First(&account).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	}
}

func getAccount(context *gin.Context) {
	var accounts []models.Account
	initializers.DB.Find(&accounts)

	context.JSON(http.StatusOK, gin.H{"data": accounts})
}

func createAccount(context *gin.Context) {
	var input models.CreateAccountInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	color := input.Color
	if color == "" {
		color = AccountActions.GenerateUniqueColor()
	}

	balance := input.Balance

	account := models.Account{Name: input.Name, Balance: balance, Color: color}
	initializers.DB.Create(&account)

	context.JSON(http.StatusCreated, gin.H{"data": account})
}

func findAccount(context *gin.Context) {
	var account models.Account

	getAccountById(context.Param("id"), &account, context)

	context.JSON(http.StatusOK, gin.H{"data": account})
}

func removeAccount(context *gin.Context) {
	var account models.Account

	if err := initializers.DB.Where("id = ?", context.Param("id")).Delete(&account).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func updateAccount(context *gin.Context) {
	var account models.Account
	var input models.UpdateAccountInput

	getAccountById(context.Param("id"), &account, context)

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		return
	}

	if input.Name != "" {
		account.Name = input.Name
	}

	json.Unmarshal(body, &input)
	if input.Color.Set {
		if input.Color.Valid && input.Color.Value != "" {
			account.Color = input.Color.Value
		} else {
			account.Color = AccountActions.GenerateUniqueColor()
		}
	}

	context.JSON(http.StatusOK, gin.H{"data": account})
}
