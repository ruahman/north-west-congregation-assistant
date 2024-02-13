package database

import (
	"database/sql"
	"fmt"
	"jw/utils"
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
func CreateDB(db *sql.DB, database string) {
	_, _ = Exec(db, "CREATE DATABASE "+database)
}

/***
 * Drop the database
 * @param db *sql.DB - database connection
 */
func DropDB(db *sql.DB) {
	_, _ = ExecFile(db, DATABASE_PATH+"/database_down.sql")
}

/***
 * Execute a sql Query
 * @param db *sql.DB - database connection
 * @param q string - sql Query
 * @param params ...interface{} - Query parameters
 * @return sql.Result - result of the sql execution
 * @return error - error if any
 */
func Exec(db *sql.DB, q string, params ...interface{}) (sql.Result, error) {
	res, err := db.Exec(q, params...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

/***
 * Execute a sql file
 * @param db *sql.DB - database connection
 * @param p string - path to the sql file
 * @return sql.Result - result of the sql execution
 * @return error - error if any
 */
func ExecFile(db *sql.DB, p string) (sql.Result, error) {
	sql, err := utils.ReadFile(p)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	res, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

func Query(db *sql.DB, q string, params ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(q, params...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return rows, nil
}
