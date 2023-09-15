package main

import (
	"github.com/esinayyildiz/CustomerServices/api"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/customer", api.SaveCustomerApi)
	router.DELETE("/customer", api.DeleteCustomerApi)
	router.GET("/customer/:id", api.GetCustomerApi)
	router.PUT("/customer/:id", api.UpdateCustomerApi)

	// Diğer API uçları da burada tanımlanabilir
}

func main() {
	router := gin.Default()

	// API uçlarını yükle
	SetupRoutes(router)

	// Server'ı başlat
	router.Run(":8080")
}
