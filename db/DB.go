package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v4"
)

var lock = &sync.Mutex{}

var instance *pgx.Conn

func GetInstance() *pgx.Conn {

	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			conn, err := pgx.Connect(context.Background(), "user=postgres host=localhost port=5432 database=dp_final sslmode=disable")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to connect %v\n", err)
				os.Exit(1)
			}
			instance = conn
		}
	}
	return instance
}