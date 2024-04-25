package postgres

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func MigrateSQLX() error {
	db, err := New("sqlx")
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/infrastructure/postgres/migrations/sqlx",
		"sqlx", driver)
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("no new migrations found")

			return nil
		}

		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) { // or m.Step(2) if you want to explicitly set the number of migrations to run
		return err
	}

	return nil
}
