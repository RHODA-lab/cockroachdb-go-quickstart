package main

import (
	"github.com/gorilla/mux"
	"github.com/myeung18/cockroachdb-go-quickstart/controller"
	"log"
	"net/http"
)

func main() {
	startWeb()
}

func startWeb() {
	r := mux.NewRouter().StrictSlash(false) //exact '/' match is needed
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/", fs)
	controller.NewController(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Fruit service is Listening..")
	server.ListenAndServe()
}
