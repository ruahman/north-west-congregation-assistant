package database

import (
	"fmt"
	"jw/utils"
	"testing"
)

func TestDatabaseConnect(t *testing.T) {
	// Connect to the database
	db, err := Connect("postgres", "password", "postgres", "postgres")
	if err != nil {
		t.Error("Connect failed")
	}

	fmt.Println("database stats")
	utils.PrettyJSON(db.Stats())
}
