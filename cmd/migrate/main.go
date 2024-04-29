package main

import (
	"github.com/asaphin/all-databases-go/internal/infrastructure/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := postgres.MigrateSQLX()
	if err != nil {
		log.WithError(err).Error("unable to make migrations for postgres sqlx database")
	}

	err = postgres.MigrateSQLXFiles()
	if err != nil {
		log.WithError(err).Error("unable to make migrations for postgres sqlx_files database")
	}

	postgres.Shutdown()
}
