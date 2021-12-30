package routes

import (
	"github.com/gin-gonic/gin"
	"gofinance/controllers"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", test)

	controllers.CategoryRoutes(router)
}

func test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "API Test",
	})

	return
}
