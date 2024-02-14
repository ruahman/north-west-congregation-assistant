package database

import (
	"database/sql"
	"fmt"
	"jw/data/models"
	"jw/utils"
	"log"
	"testing"
)

const (
	USERNAME  = "postgres"
	PASSWORD  = "password"
	HOST      = "postgres"
	DATABASE  = "postgres"
	NOT_FOUND = -1
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

func getDatabases(db *sql.DB) ([]string, error) {
	rows, err := Query(db, "SELECT datname FROM pg_database")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	databases, err := models.Map[string](rows, func(rows *sql.Rows) (string, error) {
		var datname string
		if err := rows.Scan(&datname); err != nil {
			return "", err
		}
		return datname, nil
	})
	if err != nil {
		return nil, err
	}

	return databases, nil
}

func TestCreateDB(t *testing.T) {
	db, teardown := setup("postgres")
	defer teardown(db)

	CreateDB(db, "test")
	defer (func() {
		_, _ = Exec(db, "DROP DATABASE test")
	})()

	databases, err := getDatabases(db)
	if err != nil {
		t.Error("getDatabases failed")
	}

	fmt.Println("databases", databases)

	if len(databases) != 4 {
		t.Error("Database test was not created")
	}

	if x := utils.Search(databases, "test"); x == NOT_FOUND {
		t.Error("Database test was not created")
	}

	fmt.Println("Database test was created")
}

func TestDropDB(t *testing.T) {
	db, teardown := setup("postgres")
	defer teardown(db)

	CreateDB(db, "test")

	databases, err := getDatabases(db)
	if err != nil {
		t.Error("getDatabases failed")
	}

	fmt.Println("databases", databases)

	if len(databases) != 4 {
		t.Error("Database not created")
	}

	if x := utils.Search(databases, "test"); x == NOT_FOUND {
		t.Error("Database test was not created")
	}

	DropDB(db, "test")

	databases, err = getDatabases(db)
	if err != nil {
		t.Error("getDatabases failed")
	}

	if len(databases) != 3 {
		t.Error("Database not created")
	}

	if x := utils.Search(databases, "test"); x != NOT_FOUND {
		t.Error("Database was not deleted")
	}
	println("test db was deleted")
}
