package postgres

import (
	"database/sql"
	"fmt"
	"github.com/asaphin/all-databases-go/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"sync"
)

var dbInstances = make(map[string]*sql.DB)
var shutdownSync = sync.Once{}

func New(dbName string) (*sql.DB, error) {
	if _, ok := dbInstances[dbName]; !ok {
		cfg := config.Get()

		log.WithField("dbName", dbName).Debug("connecting to database")

		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.Username,
			cfg.Postgres.Password,
			dbName,
			cfg.Postgres.SSLMode))
		if err != nil {
			return nil, err
		}

		log.WithField("dbName", dbName).Debug("database connected")

		err = db.Ping()
		if err != nil {
			closeErr := db.Close()
			if closeErr != nil {
				log.WithField("dbName", dbName).WithError(closeErr).Warn("failed to close database connection")
			}

			return nil, err
		}

		log.WithField("dbName", dbName).Debug("ping successfull")

		dbInstances[dbName] = db
	}

	return dbInstances[dbName], nil
}

func Shutdown() {
	log.Trace("shutting down postgres database instances")

	shutdownSync.Do(func() {
		for name, db := range dbInstances {
			log.WithField("dbName", name).Trace("shutting down postgres database instance")

			if err := db.Close(); err == nil {
				log.WithField("name", name).Debug("connection closed")
			} else {
				log.WithField("name", name).WithError(err).Warn("unable to close connection")
			}
		}
	})
}

func NewSqlx(dbName string) (*sqlx.DB, error) {
	innerDB, err := New(dbName)
	if err != nil {
		return nil, err
	}

	log.WithField("dbName", dbName).Debug("sqlx instance created")

	return sqlx.NewDb(innerDB, "postgres"), nil
}
