package postgres

import "github.com/jmoiron/sqlx"

type SQLXFilesRepository struct {
	db *sqlx.DB
}

func NewSQLXFilesRepository() (*SQLXFilesRepository, error) {
	db, err := NewSqlx(sqlxFilesDatabaseName)
	if err != nil {
		return nil, err
	}

	return &SQLXFilesRepository{
		db: db,
	}, nil
}
