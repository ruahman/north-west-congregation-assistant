package database

import (
	"database/sql"
	"fmt"
	"jw/data/sqlutils"
	"log"
)

const (
	DATABASE_PATH = "data/database"
)

func ConnectDB(u string, p string, h string, d string) *sql.DB {
	c := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		u,
		p,
		h,
		d,
	)

	db, err := sql.Open("postgres", c)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}

func CreateDB(db *sql.DB) {
	sqlutils.ExecFile(db, "./create-database_up.sql")
}

func DropDB(db *sql.DB) {
	sqlutils.ExecFile(db, "./create-database_down.sql")
}
