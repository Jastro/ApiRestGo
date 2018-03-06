package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Car struct {
	ID         string   `json:"id,omitempty`
	Color      string   `json:"color,omitempty`
	Brand      string   `json:"brand,omitempty"`
	Model      string   `json:"model,omitempty"`
	Enrollment string   `json:"enrollment,omitempty"`
	Data   	   *Data    `json:"data,omitempty"`
}

type Data struct {
	Type  string `json:"type,omitempty"`
	Tare  int 	 `json:"tare,omitempty"`
	Seats int    `json:"seats,omitempty"`
}

var cars []Car

func GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Car{})
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cars)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = params["id"]
	cars = append(cars, car)
	json.NewEncoder(w).Encode(cars)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(cars)
	}
}

func generateData() {
	cars = append(cars, Car{ID: "1", Brand: "Citroen", Model: "Saxo", Color: "Rojo", Enrollment: "1565 JHY", Data: &Data{Type: "26-42HX", Tare: 1985, Seats: 6}})
	cars = append(cars, Car{ID: "2", Brand: "Citroen", Model: "Saxo", Color: "Azul", Enrollment: "1565 JHY", Data: &Data{Type: "26-42HX", Tare: 1985, Seats: 6}})
	cars = append(cars, Car{ID: "3", Brand: "Citroen", Model: "Saxo", Color: "Amarillo", Enrollment: "1565 JHY", Data: &Data{Type: "26-42HX", Tare: 1985, Seats: 6}})
}

func main() {
	router := mux.NewRouter()
	generateData();
	router.HandleFunc("/cars", GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	router.HandleFunc("/cars/{id}", CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}