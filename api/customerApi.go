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

	// JSON verisini al ve customer nesnesine çöz
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// customer nesnesini veritabanına kaydet
	err := data.SaveCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The customer could not be saved"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The customer has been successfully saved"})
}

func DeleteCustomerApi(c *gin.Context) {
	// URL'den müşteri kimliğini al
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
	// URL'den müşteri kimliğini al
	id, _ := strconv.Atoi(c.Query("id"))

	// Müşteriyi veritabanından oku
	customer, err := data.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The customer could not be read"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func UpdateCustomerApi(c *gin.Context) {
	// URL'den müşteri kimliğini al
	id, _ := strconv.Atoi(c.Param("id"))

	// JSON verisini al ve customer nesnesine çöz
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Müşteriyi veritabanında güncelle
	err := data.UpdateCustomer(id, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Müşteri güncellenemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Müşteri başarıyla güncellendi"})
}
