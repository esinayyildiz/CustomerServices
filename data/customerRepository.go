package data

import (
	"github.com/esinayyildiz/CustomerServices/database"
	"github.com/esinayyildiz/CustomerServices/models"
)

func SaveCustomer(customer models.Customer) error {
	result := database.Conn.Create(&customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteCustomer(id int) error {
	result := database.Conn.Delete(&models.Customer{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCustomer(id int) (*models.Customer, error) {
	var customer models.Customer
	result := database.Conn.First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func UpdateCustomer(id int, customer models.Customer) error {
	result := database.Conn.Model(&models.Customer{}).Where("id = ?", id).Updates(customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
