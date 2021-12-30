package main

import (
	"github.com/gin-gonic/gin"
	"gofinance/initializers"
	"gofinance/routes"
)

func main() {
	initializers.ConnectDatabase()

	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	routes.Routes(router)
	router.Run(":8000")
}
