package api

import (
	"net/http"
	"strconv"

	"github.com/esinayyildiz/CustomerServices/data"
	"github.com/esinayyildiz/CustomerServices/models"
	"github.com/gin-gonic/gin"
)

func SaveCustomerApi(c *gin.Context) {
	var customer models.Customer

	// control the json data
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save the customer object
	err := data.SaveCustomer(customer)
	//if any error exists
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The customer could not be saved"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The customer has been successfully saved"})
}

func DeleteCustomerApi(c *gin.Context) {
	// id
	id, _ := strconv.Atoi(c.Query("id"))

	// delete customer
	err := data.DeleteCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The customer could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The customer has been successfully deleted."})
}

func GetCustomerApi(c *gin.Context) {
	// id from user
	id, _ := strconv.Atoi(c.Query("id"))

	//read the customer from database
	customer, err := data.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The customer could not be read"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func UpdateCustomerApi(c *gin.Context) {
	// id from the user
	id, _ := strconv.Atoi(c.Param("id"))

	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.UpdateCustomer(id, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The customer could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The customer has been successflly updated"})
}
