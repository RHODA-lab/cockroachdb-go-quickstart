package main

import (
	"fmt"
	"github.com/RHEcosystemAppEng/sbo-go-library/pkg/binding/convert"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
)

func main() {
	migrateDB()
}

// Deprecated
func migrateDB() {
	connstring, err := convert.GetPostgreSQLConnectionString()
	if err != nil {
		log.Println(err)
	}
	log.Println(connstring)

	db, err := goose.OpenDBWithDriver("postgres", connstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()
	dir := "./sql"
	cmd := "status"
	var arguments []string
	if err := goose.Run(cmd, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", cmd, err)
	}

	fmt.Println("done")
}
