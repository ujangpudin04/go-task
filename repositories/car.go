package repositories

import (
	"cakra/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	FindCar() ([]models.Car, error)
	GetCar(ID int) (models.Car, error)
	CreateCar(Car models.Car) (models.Car, error)
	UpdateCar(Car models.Car, ID int) (models.Car, error)
	DeleteCar(Car models.Car, ID int) (models.Car, error)
}

func RepositoryCar(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCar() ([]models.Car, error) {
	var cars []models.Car
	err := r.db.Find(&cars).Error

	return cars, err
}

func (r *repository) GetCar(ID int) (models.Car, error) {
	var car models.Car
	err := r.db.First(&car, ID).Error
	return car, err
}

// Write this code
func (r *repository) CreateCar(car models.Car) (models.Car, error) {
	err := r.db.Create(&car).Error

	return car, err
}

func (r *repository) UpdateCar(car models.Car, ID int) (models.Car, error) {
	err := r.db.Save(&car).Error

	return car, err
}

func (r *repository) DeleteCar(car models.Car, ID int) (models.Car, error) {
	err := r.db.Delete(&car).Error // Using Delete method

	return car, err
}
