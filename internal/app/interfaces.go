package app

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/domain"
)

type AddressesRepository interface {
	Create(ctx context.Context, address *domain.Address) (string, error)
	List(ctx context.Context, limit, offset int) ([]*domain.Address, error)
	GetByID(ctx context.Context, addressID string) (*domain.Address, error)
	Update(ctx context.Context, address *domain.Address) error
	Delete(ctx context.Context, addressID string) error
}
