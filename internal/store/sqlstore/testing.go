package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestDB(test *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	test.Helper()

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		test.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		test.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}
		db.Close()
	}
}
