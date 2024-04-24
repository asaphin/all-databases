package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
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
		return err
	}

	// /home/anton/GoProjects/personal/all-databases-go/internal/infrastructure/postgres/migrations/sqlx/0001_create_addresses_table.down.sql

	return m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
}
