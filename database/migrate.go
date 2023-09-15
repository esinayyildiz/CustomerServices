package database

import "github.com/esinayyildiz/CustomerServices/models"

func Migrate() {
	Conn.AutoMigrate(&models.Customer{})
}
