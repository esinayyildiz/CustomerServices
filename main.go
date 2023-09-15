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

}

func main() {
	router := gin.Default()

	SetupRoutes(router)

	router.Run(":8080")
}
