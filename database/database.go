package database

import (
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

var lock = &sync.Mutex{}

var instance *sqlx.DB

func GetInstance() *sqlx.DB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			conn, err := sqlx.Connect("postgres", "user=postgres host=localhost port=5432 database=dp_final sslmode=disable")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to connect %v\n", err)
				os.Exit(1)
			}
			instance = conn
		}
	}

	return instance
}
