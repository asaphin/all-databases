package app

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/domain"
)

type FilesRepository interface {
	Put(ctx context.Context, file *domain.File) error
	List(ctx context.Context) ([]*domain.FileListItem, error)
	Get(ctx context.Context, id string) (*domain.File, error)
	Delete(ctx context.Context, id string) error
}

type AddressesRepository interface {
	Create(ctx context.Context, address *domain.Address) (string, error)
	List(ctx context.Context, limit, offset int) ([]*domain.Address, error)
	GetByID(ctx context.Context, addressID string) (*domain.Address, error)
	Update(ctx context.Context, address *domain.Address) error
	Delete(ctx context.Context, addressID string) error
}

type BrandsRepository interface {
	Create(ctx context.Context, brand *domain.Brand) (string, error)
	List(ctx context.Context, limit, offset int) ([]*domain.BrandListItem, error)
	GetByID(ctx context.Context, brandID string) (*domain.Brand, error)
	Update(ctx context.Context, brand *domain.Brand) error
	Delete(ctx context.Context, brandID string) error
}
