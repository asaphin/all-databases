package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/jmoiron/sqlx"
)

type SQLXAddressesRepository struct {
	db *sqlx.DB
}

func NewSQLXAddressesRepository() (*SQLXAddressesRepository, error) {
	db, err := NewSqlx(sqlxDatabaseName)
	if err != nil {
		return nil, err
	}

	return &SQLXAddressesRepository{
		db: db,
	}, nil
}

func (repo *SQLXAddressesRepository) Create(ctx context.Context, address *domain.Address) (string, error) {
	query := `
        INSERT INTO addresses 
            (type, in_care_of_name, street, street_number, apartment, locality, region, postal_code, country, additional_info, latitude, longitude) 
        VALUES 
            ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
        RETURNING id`

	additionalInfoJSON, err := json.Marshal(address.AdditionalInfo)
	if err != nil {
		return "", err
	}

	var id string
	err = repo.db.QueryRowContext(ctx, query,
		address.Type,
		address.InCareOfName,
		address.Street,
		address.StreetNumber,
		address.Apartment,
		address.Locality,
		address.Region,
		address.PostalCode,
		address.Country,
		additionalInfoJSON,
		address.Latitude,
		address.Longitude,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (repo *SQLXAddressesRepository) List(ctx context.Context, limit, offset int) ([]*domain.Address, error) {
	query := `SELECT * FROM addresses LIMIT $1 OFFSET $2`
	var addresses []*domain.Address
	err := repo.db.SelectContext(ctx, &addresses, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (repo *SQLXAddressesRepository) GetByID(ctx context.Context, addressID string) (*domain.Address, error) {
	query := `SELECT * FROM addresses WHERE id = $1`
	var address domain.Address
	err := repo.db.GetContext(ctx, &address, query, addressID)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (repo *SQLXAddressesRepository) Update(ctx context.Context, address *domain.Address) error {
	//query := `UPDATE addresses SET street=$1, city=$2, state=$3, zip=$4 WHERE id=$5`
	//_, err := repo.db.ExecContext(ctx, query, address.Street, address.City, address.State, address.Zip, address.ID)
	//if err != nil {
	//	return err
	//}
	//return nil
	return errors.New("unimplemented")
}

func (repo *SQLXAddressesRepository) Delete(ctx context.Context, addressID string) error {
	query := `DELETE FROM addresses WHERE id=$1`
	_, err := repo.db.ExecContext(ctx, query, addressID)
	if err != nil {
		return err
	}
	return nil
}
