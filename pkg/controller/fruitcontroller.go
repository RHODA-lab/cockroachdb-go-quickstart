package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RHEcosystemAppEng/sbo-go-library/pkg/binding/convert"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/myeung18/cockroachdb-go-quickstart/pkg/fruit"
	"log"
	"net/http"
)

type Env struct {
	service *fruit.Service
}

func NewController(r *mux.Router) {
	connstring, err := convert.GetPostgreSQLConnectionString()
	if err != nil {
		log.Println(err)
	}
	//log.Println(connstring)
	config, err1 := pgx.ParseConfig(connstring)
	if err1 != nil {
		log.Println("error configuring the database: ", err1)
	}
	conn, err2 := pgx.ConnectConfig(context.Background(), config)
	if err2 != nil {
		log.Println("error connecting to the database: ", err2)
	}

	env := &Env{service: fruit.NewFruitService(conn)}

	r.HandleFunc("/fruits", env.listFruits).Methods("GET")
	r.HandleFunc("/fruits/{id}", env.getByID).Methods("GET")
	r.HandleFunc("/fruits", env.createFruit).Methods("POST")
	r.HandleFunc("/fruits/{id}", env.updateFruit).Methods("PUT")
	r.HandleFunc("/fruits/{id}", env.deleteFruit).Methods("DELETE")
}

func (env *Env) deleteFruit(w http.ResponseWriter, r *http.Request) {
	//call model / db
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := env.service.DeleteByID(id)
	if err != nil {
		log.Println(fmt.Sprintf("failed to delete fruit %v", id), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	setHeader(w)
	w.WriteHeader(http.StatusOK)
}

func (env *Env) updateFruit(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var fruit fruit.Fruit
	err := json.NewDecoder(r.Body).Decode(&fruit)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("put err exit", fruit)
		return
	}

	fruit.Id = id
	err = env.service.Update(fruit)
	if err != nil {
		log.Println(fmt.Sprintf("failed to update fruit %v", fruit.Id), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	setHeader(w)
	w.WriteHeader(http.StatusOK)
}

func (env *Env) createFruit(w http.ResponseWriter, r *http.Request) {
	var fruit fruit.Fruit
	err := json.NewDecoder(r.Body).Decode(&fruit)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("post err exit", fruit)
		return
	}
	fruit.Id = uuid.New().String()
	env.service.Create(fruit)

	setHeader(w)
	w.WriteHeader(http.StatusCreated)
	fmt.Println("post ", fruit)
}

func (env *Env) getByID(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fruit, err := env.service.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	setHeader(w)
	enc := json.NewEncoder(w)
	enc.Encode(fruit)
}

func (env *Env) listFruits(w http.ResponseWriter, r *http.Request) {
	fruits := env.service.ListFruits()
	setHeader(w)
	enc := json.NewEncoder(w)
	enc.Encode(fruits)
}

func setHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}
