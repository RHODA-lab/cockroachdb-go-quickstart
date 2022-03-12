package main

import (
	"github.com/golang-migrate/migrate"
	"github.com/gorilla/mux"
	"github.com/myeung18/cockroachdb-go-quickstart/pkg/controller"
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


func migrateDB() {
	m, err := migrate.New(
		"file://db/migrations",
		"cockroachdb://cockroach:@localhost:26257/example?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}