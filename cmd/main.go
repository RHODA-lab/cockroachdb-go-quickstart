package main

import (
	"github.com/gorilla/mux"
	"github.com/myeung18/cockroachdb-go-quickstart/pkg/controller"
	"github.com/myeung18/cockroachdb-go-quickstart/pkg/database"
	"log"
	"net/http"
)

func main() {
	startWeb()
}

func init() {
	database.MigrateWithEmbed()
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
