package database

import (
	"embed"
	"fmt"
	"github.com/RHEcosystemAppEng/sbo-go-library/pkg/binding/convert"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"strings"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func MigrateWithEmbed() {
	connString, err := convert.GetPostgreSQLConnectionString()
	if err != nil {
		log.Println(err)
	}
	printsConnString(connString)

	goose.SetBaseFS(embedMigrations)

	db, err := goose.OpenDBWithDriver("postgres", connString)
	if err != nil {
		log.Println(fmt.Sprintf("goose: failed to open DB: %v\n", err))
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(fmt.Sprintf("goose: failed to close DB: %v\n", err))
		}
	}()
	if err := goose.Up(db, "migrations"); err != nil {
		log.Println(fmt.Sprintf("goose: failed to goose up : %v\n", err))
	}
}

func printsConnString(connString string) {

	parts := strings.Split(connString, ":")
	if len(parts) == 4 {
		atIdx := strings.Index(parts[2], "@")
		if atIdx >= 0 {
			parts[2] = "<pwd_not_shown>" + parts[2][atIdx:]
		} else {
			parts[2] = ""
		}
		log.Println(strings.Join(parts, ":"))
	}

}
