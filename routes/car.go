package routes

import (
	"cakra/handlers"
	"cakra/pkg/mysql"
	"cakra/repositories"

	"github.com/gorilla/mux"
)

func CarRoutes(r *mux.Router) {
	CarRepository := repositories.RepositoryCar(mysql.DB)
	h := handlers.HandlerCar(CarRepository)

	r.HandleFunc("/car", h.FindCar).Methods("GET")
	r.HandleFunc("/car/{id}", h.GetCar).Methods("GET")
	// r.HandleFunc("/film", h.CreateFilm).Methods("POST")
	r.HandleFunc("/car", h.CreateCar).Methods("POST")
	r.HandleFunc("/car/{id}", h.UpdateCar).Methods("PATCH")
	r.HandleFunc("/car/{id}", h.DeleteCar).Methods("DELETE")
}
