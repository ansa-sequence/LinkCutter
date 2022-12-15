package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(test *testing.T, databaseUrl string) (*Store, func(...string)) {
	test.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseUrl
	s := NewStore(config)
	if err := s.Open(); err != nil {
		test.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ","))); err != nil {
				test.Fatal(err)
			}
		}
		s.Close()
	}
}
