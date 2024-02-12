package database

import (
	"database/sql"
	"fmt"
	"jw/data/sqlutils"
	"log"

	_ "github.com/lib/pq"
)

const (
	DATABASE_PATH = "data/database"
)

/***
 * Connect to the database
 * @param u string - username
 * @param p string - password
 * @param h string - host
 * @param d string - database
 * @return *sql.DB - database connection
 */
func Connect(u string, p string, h string, d string) (*sql.DB, error) {
	// connection string to database
	c := fmt.Sprintf(
		"postgres://%s:%s@%s:5432/%s?sslmode=disable",
		u,
		p,
		h,
		d,
	)

	// connect to a postgres database
	db, err := sql.Open("postgres", c)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// ping the database
	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

/***
 * Create the database
 * @param db *sql.DB - database connection
 */
func CreateDB(db *sql.DB) {
	sqlutils.ExecFile(db, DATABASE_PATH+"/create-database_up.sql")
}

/***
 * Drop the database
 * @param db *sql.DB - database connection
 */
func DropDB(db *sql.DB) {
	sqlutils.ExecFile(db, DATABASE_PATH+"/create-database_down.sql")
}
