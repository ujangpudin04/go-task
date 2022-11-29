package database

import (
	"cakra/models"
	"cakra/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	mysql.DB.AutoMigrate(&models.Car{}, &models.Ball{})
	fmt.Println("Migration Success")
}
