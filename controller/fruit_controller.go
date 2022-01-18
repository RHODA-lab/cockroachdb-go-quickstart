package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewController(r *mux.Router) {
	r.HandleFunc("/fruits", listFruits).Methods("GET")
	r.HandleFunc("/fruits/{id}", getByID).Methods("GET")
	r.HandleFunc("/fruits", createFruit).Methods("POST")
	r.HandleFunc("/fruits/{id}", updateFruit).Methods("PUT")
	r.HandleFunc("/fruits/{id}", deleteFruit).Methods("DELETE")
}

func deleteFruit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete", mux.Vars(r)["id"])
}

func updateFruit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("put", mux.Vars(r)["id"])
}

func createFruit(w http.ResponseWriter, r *http.Request) {
	var fruit Fruit
	err := json.NewDecoder(r.Body).Decode(&fruit)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("post ", fruit)
}

func getByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get", mux.Vars(r)["id"])
}

func listFruits(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all")
}

type Fruit struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
