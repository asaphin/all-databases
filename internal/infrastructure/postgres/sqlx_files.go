package postgres

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/jmoiron/sqlx"
)

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

func (repo *SQLXFilesRepository) Put(ctx context.Context, file *domain.File) error {
	query := `
		INSERT INTO files (id, name, type, created_at, updated_at, data) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE 
		SET name = EXCLUDED.name, type = EXCLUDED.type, 
			created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at, 
			data = EXCLUDED.data
	`
	_, err := repo.db.ExecContext(ctx, query, file.ID, file.Name, file.Type, file.CreatedAt, file.UpdatedAt, file.Data)
	return err
}

func (repo *SQLXFilesRepository) List(ctx context.Context) ([]*domain.FileListItem, error) {
	var files []*domain.FileListItem
	query := `SELECT id, name, type, created_at, updated_at FROM files`
	err := repo.db.SelectContext(ctx, &files, query)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (repo *SQLXFilesRepository) Get(ctx context.Context, id string) (*domain.File, error) {
	var file domain.File
	query := `SELECT id, name, type, created_at, updated_at, data FROM files WHERE id = $1`
	err := repo.db.GetContext(ctx, &file, query, id)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (repo *SQLXFilesRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM files WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, id)
	return err
}
