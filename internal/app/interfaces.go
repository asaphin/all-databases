package app

import (
	"context"
	"github.com/asaphin/all-databases-go/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) (string, error)
	List(ctx context.Context) ([]*domain.User, error)
	GetByID(ctx context.Context, userID string) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, userID string) error
}
