package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gofinance/actions/CategoryActions"
	"gofinance/initializers"
	"gofinance/models"
	"io/ioutil"
	"net/http"
)

func CategoryRoutes(router *gin.Engine) {
	router.GET("/categories", getCategory)
	router.GET("/categories/:id", findCategory)
	router.POST("/categories", createCategory)
	router.DELETE("/categories/:id", removeCategory)
	router.PATCH("/categories/:id", updateCategory)
}

func getCategoryById(id string, category *models.Category, context *gin.Context) {
	if err := initializers.DB.Where("id = ?", id).First(&category).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	}
}

func getCategory(context *gin.Context) {
	var categories []models.Category
	initializers.DB.Find(&categories)

	context.JSON(http.StatusOK, gin.H{"data": categories})
}

func createCategory(context *gin.Context) {
	var input models.CreateCategoryInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Color == "" {
		input.Color = CategoryActions.GenerateUniqueColor()
	}

	category := models.Category{Name: input.Name, Color: input.Color}
	initializers.DB.Create(&category)

	context.JSON(http.StatusCreated, gin.H{"data": category})
}

func findCategory(context *gin.Context) {
	var category models.Category

	getCategoryById(context.Param("id"), &category, context)

	context.JSON(http.StatusOK, gin.H{"data": category})
}

func removeCategory(context *gin.Context) {
	var category models.Category

	if err := initializers.DB.Where("id = ?", context.Param("id")).Delete(&category).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func updateCategory(context *gin.Context) {
	var category models.Category
	var input models.UpdateCategoryInput

	getCategoryById(context.Param("id"), &category, context)

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		return
	}

	if input.Name != "" {
		category.Name = input.Name
	}

	json.Unmarshal(body, &input)
	if input.Color.Set {
		if input.Color.Valid && input.Color.Value != "" {
			category.Color = input.Color.Value
		} else {
			category.Color = CategoryActions.GenerateUniqueColor()
		}
	}

	initializers.DB.Save(&category)

	context.JSON(http.StatusOK, gin.H{"data": category})
}
