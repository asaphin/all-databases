package postgres

import (
	"fmt"
	"github.com/asaphin/all-databases-go/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var dbInstances = make(map[string]*sqlx.DB)
var shutdownSync = sync.Once{}

func NewDB(dbName string) *sqlx.DB {
	if _, ok := dbInstances[dbName]; !ok {
		cfg := config.Get()

		db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.Username,
			cfg.Postgres.Password,
			dbName,
			cfg.Postgres.SSLMode))
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		dbInstances[dbName] = db

		shutdownSync.Do(func() {
			go handleShutdown()
		})
	}

	return dbInstances[dbName]
}

func handleShutdown() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	sig := <-sigCh
	fmt.Printf("got system signal %s, shutting down...\n", sig.String())

	for name, db := range dbInstances {
		if err := db.Close(); err == nil {
			fmt.Printf("postgres %s connection closed\n", name)
		} else {
			fmt.Printf("unable to close connection to postgres %s: %s\n", name, err.Error())
		}
	}

	os.Exit(0)
}
