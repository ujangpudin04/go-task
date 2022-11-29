package handlers

import (
	carsdto "cakra/dto/cars"
	dto "cakra/dto/result"
	"cakra/models"
	"cakra/repositories"

	// "context"
	"encoding/json"
	"net/http"

	// "os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// Declare handler struct here ...
type handlerCar struct {
	CarRepository repositories.CarRepository
}

// Declare Handler function here ...
func HandlerCar(CarRepository repositories.CarRepository) *handlerCar {
	return &handlerCar{CarRepository}
}

// Declare FindUsers method here ...
func (h *handlerCar) FindCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cars, err := h.CarRepository.FindCar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// for i, p := range films {
	// 	films[i].Thumbnailfilm = os.Getenv("PATH_FILE") + p.Thumbnailfilm
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: cars}
	json.NewEncoder(w).Encode(response)
}

// Declare GetUser method here ...
func (h *handlerCar) GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	car, err := h.CarRepository.GetCar(id)
	// film.Thumbnailfilm = os.Getenv("PATH_FILE") + film.Thumbnailfilm

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCar(car)}
	json.NewEncoder(w).Encode(response)
}

// Write this code
func (h *handlerCar) CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(carsdto.CarCreateRequest)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// data form pattern submit to pattern entity db user
	car := models.Car{
		Price: request.Price,
		Brand: request.Brand,
		Type:  request.Type,
	}

	data, err := h.CarRepository.CreateCar(car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCar(data)}
	json.NewEncoder(w).Encode(response)

}

// Write this code
func (h *handlerCar) UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(carsdto.CarUpdateRequest) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	car, _ := h.CarRepository.GetCar(int(id))

	if request.Price != "" {
		car.Price = request.Price
	}

	if request.Brand != "" {
		car.Brand = request.Brand
	}

	if request.Type != "" {
		car.Type = request.Type
	}

	data, err := h.CarRepository.UpdateCar(car, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCar(data)}
	json.NewEncoder(w).Encode(response)
}

// Write this code
func (h *handlerCar) DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	car, err := h.CarRepository.GetCar(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CarRepository.DeleteCar(car, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCar(data)}
	json.NewEncoder(w).Encode(response)
}

// Declare convertResponse function here ...
func convertResponseCar(u models.Car) carsdto.CarResponse {
	return carsdto.CarResponse{
		ID:    u.ID,
		Price: u.Price,
		Brand: u.Brand,
		Type:  u.Type,
	}
}
