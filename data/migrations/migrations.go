package migrations

import (
	"database/sql"
	"fmt"
	"jw/data/sqlutils"
	"jw/utils"
	"log"
	"os"
	"regexp"
	"sort"
	"time"
)

func Add(n string) {
	timeStamp := time.Now().Unix()

	u := fmt.Sprintf("%d--%s_up.sql", timeStamp, n)
	d := fmt.Sprintf("%d--%s_down.sql", timeStamp, n)

	fmt.Println("...Creating migration files", u, d)
	os.Create("./" + u)
	os.Create("./" + d)
}

func Run(db *sql.DB) {
	fmt.Println("...run migrations")
	sqlutils.ExecFile(db, "./data/migrations/create-migrations-table_add.sql")

	files, err := utils.ReadDir("./data/migrations")
	if err != nil {
		log.Fatal(err)
		return
	}

	var migrations []string
	for _, f := range files {
		if matched, _ := regexp.MatchString(`^(\d+)--(.+)_up.sql$`, f); matched == true {
			migrations = append(migrations, f)
		}
	}
	sort.Strings(migrations)
	fmt.Println(migrations)

	for _, m := range migrations {
		sqlutils.ExecFile(db, "./data/migrations/"+m)
	}
}

func Drop(db *sql.DB) {
	fmt.Println("...drop migrations")
	// sqlutils.ExecFile(db, "./data/migrations/create-migrations-table_down.sql")

	files, err := utils.ReadDir("./data/migrations")
	if err != nil {
		log.Fatal(err)
		return
	}

	var migrations []string
	for _, f := range files {
		if matched, _ := regexp.MatchString(`^(\d+)--(.+)_down.sql$`, f); matched == true {
			migrations = append(migrations, f)
		}
	}
	sort.Strings(migrations)
	// sort.Reverse(sort.StringSlice(migrations))
	fmt.Println(migrations)

	// for _, m := range migrations {
	// 	sqlutils.ExecFile(db, "./data/migrations/"+m)
	// }
}

func Up() {
	fmt.Println("...run migrations up")
}

func Down() {
	fmt.Println("...run migrations down")
}
