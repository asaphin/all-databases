package postgres

import (
	"context"
	"errors"
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/asaphin/all-databases-go/internal/utils"
	"github.com/jmoiron/sqlx"
)

type SQLXBrandsRepository struct {
	db *sqlx.DB
}

func NewSQLXBrandsRepository() (*SQLXBrandsRepository, error) {
	db, err := NewSqlx(sqlxDatabaseName)
	if err != nil {
		return nil, err
	}

	return &SQLXBrandsRepository{
		db: db,
	}, nil
}

func (r *SQLXBrandsRepository) Create(ctx context.Context, brand *domain.Brand) (string, error) {
	query := "INSERT INTO brands (name, slogan) VALUES ($1, $2) RETURNING id"

	var id string

	err := r.db.QueryRowContext(ctx, query, brand.Name, brand.Slogan).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *SQLXBrandsRepository) List(ctx context.Context, limit, offset int) ([]*domain.BrandListItem, error) {
	query := "SELECT id, name FROM brands"
	var args []interface{}

	if limit > 0 {
		query += " LIMIT $1"
		args = append(args, limit)
	}

	if offset > 0 {
		if limit <= 0 {
			return nil, errors.New("offset specified without limit")
		}
		query += " OFFSET $2"
		args = append(args, offset)
	}

	rows, err := r.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer utils.LogAsWarningIfReturnsError(func() error { return rows.Close() }, "unable to close rows")

	var brands []*domain.BrandListItem

	for rows.Next() {
		var brand domain.BrandListItem
		err := rows.StructScan(&brand)
		if err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return brands, nil
}

func (r *SQLXBrandsRepository) GetByID(ctx context.Context, brandID string) (*domain.Brand, error) {
	query := "SELECT id, name, slogan FROM brands WHERE id = $1"

	var brand domain.Brand

	err := r.db.GetContext(ctx, &brand, query, brandID)
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

func (r *SQLXBrandsRepository) Update(ctx context.Context, brand *domain.Brand) error {
	query := "UPDATE brands SET name=$1, slogan=$2 WHERE id=$3"

	result, err := r.db.ExecContext(ctx, query, brand.Name, brand.Slogan, brand.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *SQLXBrandsRepository) Delete(ctx context.Context, brandID string) error {
	query := "DELETE FROM brands WHERE id=$1"

	result, err := r.db.ExecContext(ctx, query, brandID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
