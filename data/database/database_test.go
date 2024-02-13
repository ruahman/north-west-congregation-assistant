package database

import (
	"database/sql"
	"fmt"
	"jw/utils"
	"log"
	"testing"
)

const (
	USERNAME = "postgres"
	PASSWORD = "password"
	HOST     = "postgres"
	DATABASE = "postgres"
)

func setup(optional ...string) (*sql.DB, func(db *sql.DB)) {
	var database string
	if len(optional) > 0 {
		database = optional[0]
	} else {
		database = DATABASE
	}

	db, err := Connect(
		USERNAME,
		PASSWORD,
		HOST,
		database,
	)
	if err != nil {
		log.Fatal("Connect failed")
	}

	return db, func(db *sql.DB) {
		db.Close()
	}
}

func Teardown(db *sql.DB) {
	db.Close()
}

func TestDatabaseConnect(t *testing.T) {
	// Connect to the database
	db, err := Connect("postgres", "password", "postgres", "postgres")
	if err != nil {
		t.Error("Connect failed")
	}

	fmt.Println("database stats")
	utils.PrettyJSON(db.Stats())
}

func TestExecFile(t *testing.T) {
	db, teardown := setup("postgres")
	defer teardown(db)

	res, err := ExecFile(db, "fixtures/get_number_of_databases.sql")
	if err != nil {
		t.Error("ExecFile failed")
	}
	fmt.Println("ExecFile result")
	utils.PrettyJSON(res)
}

func TestCreateDB(t *testing.T) {
	db, teardown := setup("postgres")
	defer teardown(db)

	CreateDB(db, "test")
	defer (func() {
		_, _ = Exec(db, "DROP DATABASE test")
	})()

	rows, err := Query(db, "SELECT datname FROM pg_database")
	if err != nil {
		t.Error("Query failed")
	}

	defer rows.Close()
	var datnames []string
	for rows.Next() {
		var datname string
		if err := rows.Scan(&datname); err != nil {
			t.Error("Scan failed")
		}
		fmt.Println(datname)
		datnames = append(datnames, datname)
	}
	if err := rows.Err(); err != nil {
		t.Error("Rows error")
	}

	fmt.Println(datnames)
	fmt.Println(len(datnames))

	if len(datnames) != 4 {
		t.Error("Database not created")
	}
}
