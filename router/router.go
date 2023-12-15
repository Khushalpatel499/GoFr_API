package router

import (
	"github.com/gorilla/mux"
	"github.com/khushalpatel499/gofr_api/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/cars", controller.GetAllCars).Methods("GET")
	router.HandleFunc("/api/car", controller.InsertOneCar).Methods("POST")
	router.HandleFunc("/api/cars/{id}", controller.UpdateOneCar).Methods("PUT")
	router.HandleFunc("/api/cars/{id}", controller.DeleteACars).Methods("DELETE")
	router.HandleFunc("/api/deleteallcars", controller.DeleteAllCars).Methods("DELETE")

	return router
}
